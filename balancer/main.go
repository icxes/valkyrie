package balancer

import (
	"context"
	"net/http"
	"time"

	"github.com/R-a-dio/valkyrie/balancer/current"
	"github.com/R-a-dio/valkyrie/config"
	"github.com/R-a-dio/valkyrie/errors"
	"github.com/R-a-dio/valkyrie/storage"
)

// Execute executes the balancer with the context ctx and config cfg.
// Execution of the balancer can be halted by cancelling ctx.
func Execute(ctx context.Context, cfg config.Config) error {
	const op errors.Op = "balancer/Execute"

	br, err := NewBalancer(ctx, cfg)
	if err != nil {
		return errors.E(op, err)
	}
	err = br.start(ctx)
	if err != nil {
		return errors.E(op, err)
	}
	return nil
}

// NewBalancer returns an initialized Balancer.
func NewBalancer(ctx context.Context, cfg config.Config) (*Balancer, error) {
	const op errors.Op = "balancer/NewBalancer"

	c := cfg.Conf()
	ss, err := storage.Open(cfg)
	if err != nil {
		return nil, errors.E(op, err)
	}

	br := &Balancer{
		Config:  cfg,
		storage: ss,
		manager: c.Manager.Client(),
	}

	br.c = current.NewCurrent(c.Balancer.Fallback)
	mux := http.NewServeMux()
	mux.HandleFunc("/", br.getIndex)
	mux.HandleFunc("/status", br.getStatus)
	mux.HandleFunc("/main", br.getMain)

	br.serv = &http.Server{
		Handler:      mux,
		Addr:         c.Balancer.Addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return br, nil
}
