// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/streamer/streamer.proto

package streamer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Null struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Null) Reset()         { *m = Null{} }
func (m *Null) String() string { return proto.CompactTextString(m) }
func (*Null) ProtoMessage()    {}
func (*Null) Descriptor() ([]byte, []int) {
	return fileDescriptor_streamer_00cca327fd5fdd43, []int{0}
}
func (m *Null) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Null.Unmarshal(m, b)
}
func (m *Null) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Null.Marshal(b, m, deterministic)
}
func (dst *Null) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Null.Merge(dst, src)
}
func (m *Null) XXX_Size() int {
	return xxx_messageInfo_Null.Size(m)
}
func (m *Null) XXX_DiscardUnknown() {
	xxx_messageInfo_Null.DiscardUnknown(m)
}

var xxx_messageInfo_Null proto.InternalMessageInfo

type StopRequest struct {
	ForceStop            bool     `protobuf:"varint,1,opt,name=force_stop,json=forceStop" json:"force_stop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopRequest) Reset()         { *m = StopRequest{} }
func (m *StopRequest) String() string { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()    {}
func (*StopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_streamer_00cca327fd5fdd43, []int{1}
}
func (m *StopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopRequest.Unmarshal(m, b)
}
func (m *StopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopRequest.Marshal(b, m, deterministic)
}
func (dst *StopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopRequest.Merge(dst, src)
}
func (m *StopRequest) XXX_Size() int {
	return xxx_messageInfo_StopRequest.Size(m)
}
func (m *StopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopRequest proto.InternalMessageInfo

func (m *StopRequest) GetForceStop() bool {
	if m != nil {
		return m.ForceStop
	}
	return false
}

type QueueEntry struct {
	// is_request indicates if this was a request made by a human
	IsRequest bool `protobuf:"varint,1,opt,name=is_request,json=isRequest" json:"is_request,omitempty"`
	// user_identifier is the way we identify the user that added this to the
	// queue; This can be anything that uniquely identifies a user
	UserIdentifier string `protobuf:"bytes,2,opt,name=user_identifier,json=userIdentifier" json:"user_identifier,omitempty"`
	// estimated_play_time is the estimated time for this track to be played
	EstimatedPlayTime string `protobuf:"bytes,3,opt,name=estimated_play_time,json=estimatedPlayTime" json:"estimated_play_time,omitempty"`
	// track_id is the id used in the database table
	TrackId int64 `protobuf:"varint,4,opt,name=track_id,json=trackId" json:"track_id,omitempty"`
	// track_tags is a short `artist - title` field
	TrackTags            string   `protobuf:"bytes,5,opt,name=track_tags,json=trackTags" json:"track_tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueueEntry) Reset()         { *m = QueueEntry{} }
func (m *QueueEntry) String() string { return proto.CompactTextString(m) }
func (*QueueEntry) ProtoMessage()    {}
func (*QueueEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_streamer_00cca327fd5fdd43, []int{2}
}
func (m *QueueEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueEntry.Unmarshal(m, b)
}
func (m *QueueEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueEntry.Marshal(b, m, deterministic)
}
func (dst *QueueEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueEntry.Merge(dst, src)
}
func (m *QueueEntry) XXX_Size() int {
	return xxx_messageInfo_QueueEntry.Size(m)
}
func (m *QueueEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueEntry.DiscardUnknown(m)
}

var xxx_messageInfo_QueueEntry proto.InternalMessageInfo

func (m *QueueEntry) GetIsRequest() bool {
	if m != nil {
		return m.IsRequest
	}
	return false
}

func (m *QueueEntry) GetUserIdentifier() string {
	if m != nil {
		return m.UserIdentifier
	}
	return ""
}

func (m *QueueEntry) GetEstimatedPlayTime() string {
	if m != nil {
		return m.EstimatedPlayTime
	}
	return ""
}

func (m *QueueEntry) GetTrackId() int64 {
	if m != nil {
		return m.TrackId
	}
	return 0
}

func (m *QueueEntry) GetTrackTags() string {
	if m != nil {
		return m.TrackTags
	}
	return ""
}

type StatusResponse struct {
	Running              bool          `protobuf:"varint,1,opt,name=running" json:"running,omitempty"`
	Queue                []*QueueEntry `protobuf:"bytes,2,rep,name=queue" json:"queue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_streamer_00cca327fd5fdd43, []int{3}
}
func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse.Unmarshal(m, b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
}
func (dst *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(dst, src)
}
func (m *StatusResponse) XXX_Size() int {
	return xxx_messageInfo_StatusResponse.Size(m)
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetRunning() bool {
	if m != nil {
		return m.Running
	}
	return false
}

func (m *StatusResponse) GetQueue() []*QueueEntry {
	if m != nil {
		return m.Queue
	}
	return nil
}

type TrackRequest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	Track                int64    `protobuf:"varint,2,opt,name=track" json:"track,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrackRequest) Reset()         { *m = TrackRequest{} }
func (m *TrackRequest) String() string { return proto.CompactTextString(m) }
func (*TrackRequest) ProtoMessage()    {}
func (*TrackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_streamer_00cca327fd5fdd43, []int{4}
}
func (m *TrackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackRequest.Unmarshal(m, b)
}
func (m *TrackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackRequest.Marshal(b, m, deterministic)
}
func (dst *TrackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackRequest.Merge(dst, src)
}
func (m *TrackRequest) XXX_Size() int {
	return xxx_messageInfo_TrackRequest.Size(m)
}
func (m *TrackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TrackRequest proto.InternalMessageInfo

func (m *TrackRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *TrackRequest) GetTrack() int64 {
	if m != nil {
		return m.Track
	}
	return 0
}

type RequestResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestResponse) Reset()         { *m = RequestResponse{} }
func (m *RequestResponse) String() string { return proto.CompactTextString(m) }
func (*RequestResponse) ProtoMessage()    {}
func (*RequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_streamer_00cca327fd5fdd43, []int{5}
}
func (m *RequestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestResponse.Unmarshal(m, b)
}
func (m *RequestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestResponse.Marshal(b, m, deterministic)
}
func (dst *RequestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestResponse.Merge(dst, src)
}
func (m *RequestResponse) XXX_Size() int {
	return xxx_messageInfo_RequestResponse.Size(m)
}
func (m *RequestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RequestResponse proto.InternalMessageInfo

func (m *RequestResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RequestResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Null)(nil), "radio.internal.streamer.Null")
	proto.RegisterType((*StopRequest)(nil), "radio.internal.streamer.StopRequest")
	proto.RegisterType((*QueueEntry)(nil), "radio.internal.streamer.QueueEntry")
	proto.RegisterType((*StatusResponse)(nil), "radio.internal.streamer.StatusResponse")
	proto.RegisterType((*TrackRequest)(nil), "radio.internal.streamer.TrackRequest")
	proto.RegisterType((*RequestResponse)(nil), "radio.internal.streamer.RequestResponse")
}

func init() {
	proto.RegisterFile("rpc/streamer/streamer.proto", fileDescriptor_streamer_00cca327fd5fdd43)
}

var fileDescriptor_streamer_00cca327fd5fdd43 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xe3, 0x24, 0x4d, 0xa6, 0x55, 0x0b, 0x0b, 0x12, 0xa6, 0xa8, 0x28, 0x32, 0xa0, 0xfa,
	0x80, 0x8c, 0x54, 0x4e, 0x1c, 0xb8, 0x20, 0x7a, 0xa8, 0x90, 0x50, 0x71, 0x72, 0xe2, 0x62, 0x2d,
	0xf6, 0x34, 0x5a, 0xe1, 0xbf, 0xee, 0xcc, 0x1e, 0xf2, 0x72, 0xbc, 0x03, 0x6f, 0x84, 0x76, 0x6d,
	0x07, 0x83, 0xe4, 0xe6, 0xb6, 0x33, 0xf3, 0xcd, 0xcc, 0x37, 0xdf, 0xcc, 0xc2, 0x0b, 0xdd, 0x64,
	0xef, 0x88, 0x35, 0xca, 0x12, 0xf5, 0xfe, 0x11, 0x37, 0xba, 0xe6, 0x5a, 0x3c, 0xd3, 0x32, 0x57,
	0x75, 0xac, 0x2a, 0x46, 0x5d, 0xc9, 0x22, 0xee, 0xc3, 0xe1, 0x1c, 0xa6, 0x5f, 0x4d, 0x51, 0x84,
	0x6f, 0xe1, 0x78, 0xcd, 0x75, 0x93, 0xe0, 0xbd, 0x41, 0x62, 0x71, 0x01, 0x70, 0x57, 0xeb, 0x0c,
	0x53, 0xe2, 0xba, 0x09, 0xbc, 0x95, 0x17, 0x2d, 0x92, 0xa5, 0xf3, 0x58, 0x54, 0xf8, 0xcb, 0x03,
	0xf8, 0x66, 0xd0, 0xe0, 0x75, 0xc5, 0x7a, 0x67, 0xd1, 0x8a, 0x52, 0xdd, 0xe6, 0xf6, 0x68, 0x45,
	0x7d, 0xb1, 0x4b, 0x38, 0x33, 0x84, 0x3a, 0x55, 0x39, 0x56, 0xac, 0xee, 0x14, 0xea, 0x60, 0xb2,
	0xf2, 0xa2, 0x65, 0x72, 0x6a, 0xdd, 0x37, 0x7b, 0xaf, 0x88, 0xe1, 0x09, 0x12, 0xab, 0x52, 0x32,
	0xe6, 0x69, 0x53, 0xc8, 0x5d, 0xca, 0xaa, 0xc4, 0xc0, 0x77, 0xe0, 0xc7, 0xfb, 0xd0, 0x6d, 0x21,
	0x77, 0x1b, 0x55, 0xa2, 0x78, 0x0e, 0x0b, 0xd6, 0x32, 0xfb, 0x99, 0xaa, 0x3c, 0x98, 0xae, 0xbc,
	0xc8, 0x4f, 0x8e, 0x9c, 0x7d, 0x93, 0x5b, 0x4a, 0x6d, 0x88, 0xe5, 0x96, 0x82, 0x99, 0xab, 0xb0,
	0x74, 0x9e, 0x8d, 0xdc, 0x52, 0x88, 0x70, 0xba, 0x66, 0xc9, 0x86, 0x12, 0xa4, 0xa6, 0xae, 0x08,
	0x45, 0x00, 0x47, 0xda, 0x54, 0x95, 0xaa, 0xb6, 0xdd, 0x00, 0xbd, 0x29, 0x3e, 0xc0, 0xec, 0xde,
	0xce, 0x1a, 0x4c, 0x56, 0x7e, 0x74, 0x7c, 0xf5, 0x2a, 0x1e, 0xd1, 0x32, 0xfe, 0xab, 0x48, 0xd2,
	0x66, 0x84, 0x9f, 0xe1, 0x64, 0x63, 0x7b, 0xf6, 0x4a, 0xbc, 0x04, 0x18, 0x88, 0xe0, 0x39, 0x56,
	0x03, 0x8f, 0x78, 0x0a, 0x33, 0xc7, 0xd1, 0xe9, 0xe3, 0x27, 0xad, 0x11, 0x7e, 0x84, 0xb3, 0xae,
	0xc0, 0x90, 0x2d, 0x99, 0x2c, 0x43, 0xa2, 0x9e, 0x6d, 0x67, 0x8a, 0x47, 0xe0, 0x97, 0xb4, 0xed,
	0x04, 0xb6, 0xcf, 0xab, 0xdf, 0x13, 0x58, 0xac, 0x3b, 0x8e, 0xe2, 0x1a, 0x66, 0x6b, 0x96, 0x9a,
	0xc5, 0xc5, 0xe8, 0x18, 0xf6, 0x1e, 0xce, 0x1f, 0x0e, 0x8b, 0x2f, 0x30, 0xb5, 0x87, 0x20, 0x5e,
	0x8f, 0xc2, 0x06, 0xd7, 0x74, 0xa8, 0xd8, 0x2d, 0xcc, 0xdb, 0x65, 0x1c, 0x22, 0x75, 0xf9, 0x40,
	0xb7, 0x7f, 0x96, 0x99, 0xc2, 0x49, 0xd7, 0xdb, 0xc9, 0x2f, 0xde, 0x8c, 0x26, 0x0e, 0xd7, 0x73,
	0x1e, 0x8d, 0xc2, 0xfe, 0xd3, 0xff, 0x13, 0x7c, 0x5f, 0xf4, 0xb1, 0x1f, 0x73, 0xf7, 0xc5, 0xde,
	0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x83, 0x6d, 0x7c, 0x81, 0x03, 0x00, 0x00,
}
