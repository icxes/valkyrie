package manager

import (
	"context"
	"log"
	"time"

	radio "github.com/R-a-dio/valkyrie"
	"github.com/R-a-dio/valkyrie/errors"
	"github.com/R-a-dio/valkyrie/rpc"
	"github.com/R-a-dio/valkyrie/util/eventstream"
	"google.golang.org/grpc"
)

// NewHTTPServer sets up a net/http server ready to serve RPC requests
func NewHTTPServer(m *Manager) (*grpc.Server, error) {
	gs := grpc.NewServer()
	rpc.RegisterManagerServer(gs, rpc.NewManager(m))

	return gs, nil
}

func (m *Manager) CurrentUser(ctx context.Context) (eventstream.Stream[radio.User], error) {
	return m.userStream.SubStream(ctx), nil
}

func (m *Manager) CurrentThread(ctx context.Context) (eventstream.Stream[radio.Thread], error) {
	return m.threadStream.SubStream(ctx), nil
}

func (m *Manager) CurrentSong(ctx context.Context) (eventstream.Stream[*radio.SongUpdate], error) {
	return m.songStream.SubStream(ctx), nil
}

func (m *Manager) CurrentListeners(ctx context.Context) (eventstream.Stream[radio.Listeners], error) {
	return m.listenerStream.SubStream(ctx), nil
}

// Status returns the current status of the radio
func (m *Manager) Status(ctx context.Context) (*radio.Status, error) {
	m.mu.Lock()
	status := m.status.Copy()
	m.mu.Unlock()
	return &status, nil
}

// UpdateUser sets information about the current streamer
func (m *Manager) UpdateUser(ctx context.Context, u radio.User) error {
	defer m.updateStreamStatus()
	m.userStream.Send(u)

	m.mu.Lock()

	m.status.StreamerName = u.DJ.Name
	m.status.User = u

	isRobot := u.Username == "AFK"
	if isRobot && m.status.SongInfo.IsFallback {
		// since we're setting the DJ and are already on a fallback with our listener, we
		// try and just start the streamer straight away
		m.tryStartStreamer(time.Second * 0)
	}
	if !isRobot {
		m.stopStartStreamer()
	}

	m.mu.Unlock()
	log.Printf("manager: updating user to: %s (%s)", u.DJ.Name, u.Username)
	return nil
}

// UpdateSong sets information about the currently playing song
func (m *Manager) UpdateSong(ctx context.Context, update *radio.SongUpdate) error {
	const op errors.Op = "manager/Manager.UpdateSong"

	new := update
	info := update.Info

	// first we check if this is the same song as the previous one we received to
	// avoid double announcement or drifting start/end timings
	m.mu.Lock()
	if m.status.Song.Metadata == new.Metadata && !info.IsFallback {
		m.mu.Unlock()
		return nil
	}

	// check if a robot is streaming
	// TODO: don't hardcode this
	//isRobot := m.status.User.Username == "AFK"
	isRobot := true

	// check if we're on a fallback stream
	if info.IsFallback {
		log.Printf("manager: fallback engaged: %s", new.Metadata)
		// if we have a robot user we want to start the automated streamer, but only if
		// there isn't already a timer running
		if isRobot {
			// TODO: don't hardcode this
			timeout := time.Second * 15
			m.tryStartStreamer(timeout)
		}
		m.status.SongInfo.IsFallback = info.IsFallback
		m.mu.Unlock()
		return nil
	}
	// if we're not on a fallback we want to stop the timer for the automated streamer
	m.stopStartStreamer()
	m.mu.Unlock()

	// otherwise continue like it's a new song
	defer m.updateStreamStatus()

	ss, tx, err := m.Storage.SongTx(ctx, nil)
	if err != nil {
		return errors.E(op, err)
	}
	defer tx.Rollback()

	// we assume that the song we received has very little or no data except for the
	// Metadata field. So we try and find more info from that
	song, err := ss.FromMetadata(new.Metadata)
	if err != nil && !errors.Is(errors.SongUnknown, err) {
		return errors.E(op, err)
	}

	// if we don't have this song in the database create a new entry for it
	if song == nil {
		song, err = ss.Create(new.Metadata)
		if err != nil {
			return errors.E(op, err)
		}
	}

	// calculate start and end time only if they're zero
	if info.Start.IsZero() {
		// we assume the song just started if it wasn't set
		info.Start = time.Now()
	}
	if info.End.IsZero() {
		// set end to start if we got passed a zero time
		info.End = info.Start
	}
	if song.Length > 0 {
		// add the song length if we have one
		info.End = info.End.Add(song.Length)
	}

	var prev radio.Song
	var prevInfo radio.SongInfo
	var listenerCountDiff *int

	// critical section to swap our new song with the previous one
	m.mu.Lock()

	prev, m.status.Song = m.status.Song, *song
	prevInfo, m.status.SongInfo = m.status.SongInfo, info

	// record listener count and calculate the difference between start/end of song
	currentListenerCount := m.status.Listeners
	// update and retrieve listener count of start of song
	var startListenerCount int
	startListenerCount, m.songStartListenerCount = m.songStartListenerCount, currentListenerCount

	m.mu.Unlock()

	// only calculate a diff if we have more than 10 listeners
	if currentListenerCount > 10 && startListenerCount > 10 {
		diff := currentListenerCount - startListenerCount
		listenerCountDiff = &diff
	}

	log.Printf("manager: set song: \"%s\" (%s)\n", song.Metadata, song.Length)

	// send an event out
	m.songStream.Send(&radio.SongUpdate{Song: *song, Info: info})

	// =============================================
	// finish up database work for the previous song
	//
	// after this point, any reference to the `song` variable is an error, so we
	// make it nil so it will panic if done by mistake
	song = nil
	if prev.ID == 0 { // protect against a zero'd song
		return nil
	}

	// insert a played entry
	err = ss.AddPlay(prev, listenerCountDiff)
	if err != nil {
		return errors.E(op, err)
	}

	// update lastplayed if the streamer is a robot and the song has a track
	if prev.HasTrack() && isRobot {
		ts, _, err := m.Storage.TrackTx(ctx, tx)
		if err != nil {
			return errors.E(op, err)
		}

		err = ts.UpdateLastPlayed(prev.TrackID)
		if err != nil {
			return errors.E(op, err, prev)
		}
	}

	// update song length only if it didn't already have one
	if prev.Length == 0 {
		err = ss.UpdateLength(prev, time.Since(prevInfo.Start))
		if err != nil {
			return errors.E(op, err, prev)
		}
	}

	err = tx.Commit()
	if err != nil {
		return errors.E(op, errors.TransactionCommit, err, prev)
	}
	return nil
}

// UpdateThread sets the current thread information on the front page and chats
func (m *Manager) UpdateThread(ctx context.Context, thread radio.Thread) error {
	defer m.updateStreamStatus()
	m.threadStream.Send(thread)

	m.mu.Lock()
	m.status.Thread = thread
	m.mu.Unlock()
	return nil
}

// UpdateListeners sets the listener count
func (m *Manager) UpdateListeners(ctx context.Context, listeners radio.Listeners) error {
	defer m.updateStreamStatus()
	m.listenerStream.Send(listeners)

	m.mu.Lock()
	m.status.Listeners = int(listeners)
	m.mu.Unlock()
	return nil
}
