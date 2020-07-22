// Code generated by protoc-gen-go. DO NOT EDIT.
// source: raft.proto

package raft

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Status int32

const (
	Status_Leader    Status = 0
	Status_Follower  Status = 1
	Status_Candidate Status = 2
)

var Status_name = map[int32]string{
	0: "Leader",
	1: "Follower",
	2: "Candidate",
}

var Status_value = map[string]int32{
	"Leader":    0,
	"Follower":  1,
	"Candidate": 2,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{0}
}

type Peer struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Addresses            []string `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{0}
}

func (m *Peer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Peer.Unmarshal(m, b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
}
func (m *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(m, src)
}
func (m *Peer) XXX_Size() int {
	return xxx_messageInfo_Peer.Size(m)
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

func (m *Peer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Peer) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

// Cluster config message
type Cluster struct {
	Peers                []*Peer  `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{1}
}

func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster.Unmarshal(m, b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
}
func (m *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(m, src)
}
func (m *Cluster) XXX_Size() int {
	return xxx_messageInfo_Cluster.Size(m)
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

// Receiver implementation:
// 1. Reply false if term < currentTerm (§5.1)
// 2. If votedFor is null or candidateId, and candidate’s log is at
//    least as up-to-date as receiver’s log, grant vote (§5.2, §5.4)
type RequestVote struct {
	Term                 int64    `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	CandidateId          string   `protobuf:"bytes,2,opt,name=candidateId,proto3" json:"candidateId,omitempty"`
	LastLogIndex         int64    `protobuf:"varint,3,opt,name=lastLogIndex,proto3" json:"lastLogIndex,omitempty"`
	LastLogTerm          int64    `protobuf:"varint,4,opt,name=lastLogTerm,proto3" json:"lastLogTerm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestVote) Reset()         { *m = RequestVote{} }
func (m *RequestVote) String() string { return proto.CompactTextString(m) }
func (*RequestVote) ProtoMessage()    {}
func (*RequestVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{2}
}

func (m *RequestVote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestVote.Unmarshal(m, b)
}
func (m *RequestVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestVote.Marshal(b, m, deterministic)
}
func (m *RequestVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestVote.Merge(m, src)
}
func (m *RequestVote) XXX_Size() int {
	return xxx_messageInfo_RequestVote.Size(m)
}
func (m *RequestVote) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestVote.DiscardUnknown(m)
}

var xxx_messageInfo_RequestVote proto.InternalMessageInfo

func (m *RequestVote) GetTerm() int64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RequestVote) GetCandidateId() string {
	if m != nil {
		return m.CandidateId
	}
	return ""
}

func (m *RequestVote) GetLastLogIndex() int64 {
	if m != nil {
		return m.LastLogIndex
	}
	return 0
}

func (m *RequestVote) GetLastLogTerm() int64 {
	if m != nil {
		return m.LastLogTerm
	}
	return 0
}

type ResponseVote struct {
	Term                 int64    `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	VoteGranted          bool     `protobuf:"varint,2,opt,name=voteGranted,proto3" json:"voteGranted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseVote) Reset()         { *m = ResponseVote{} }
func (m *ResponseVote) String() string { return proto.CompactTextString(m) }
func (*ResponseVote) ProtoMessage()    {}
func (*ResponseVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{3}
}

func (m *ResponseVote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseVote.Unmarshal(m, b)
}
func (m *ResponseVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseVote.Marshal(b, m, deterministic)
}
func (m *ResponseVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseVote.Merge(m, src)
}
func (m *ResponseVote) XXX_Size() int {
	return xxx_messageInfo_ResponseVote.Size(m)
}
func (m *ResponseVote) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseVote.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseVote proto.InternalMessageInfo

func (m *ResponseVote) GetTerm() int64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *ResponseVote) GetVoteGranted() bool {
	if m != nil {
		return m.VoteGranted
	}
	return false
}

// Receiver implementation:
// 1. Reply false if term < currentTerm (§5.1)
// 2. Reply false if log doesn’t contain an entry at prevLogIndex
//    whose term matches prevLogTerm (§5.3)
// 3. If an existing entry conflicts with a new one (same index
//    but different terms), delete the existing entry and all that
//    follow it (§5.3)
// 4. Append any new entries not already in the log
// 5. If leaderCommit > commitIndex, set commitIndex =
//    min(leaderCommit, index of last new entry)
type RequestAppendEntries struct {
	Term         int64    `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	LeaderId     string   `protobuf:"bytes,2,opt,name=leaderId,proto3" json:"leaderId,omitempty"`
	PrevLogIndex int64    `protobuf:"varint,3,opt,name=prevLogIndex,proto3" json:"prevLogIndex,omitempty"`
	PrevLogTerm  int64    `protobuf:"varint,4,opt,name=prevLogTerm,proto3" json:"prevLogTerm,omitempty"`
	Entries      [][]byte `protobuf:"bytes,5,rep,name=entries,proto3" json:"entries,omitempty"`
	// send more than one for efficiency)
	LeaderCommit         int64    `protobuf:"varint,6,opt,name=leaderCommit,proto3" json:"leaderCommit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestAppendEntries) Reset()         { *m = RequestAppendEntries{} }
func (m *RequestAppendEntries) String() string { return proto.CompactTextString(m) }
func (*RequestAppendEntries) ProtoMessage()    {}
func (*RequestAppendEntries) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{4}
}

func (m *RequestAppendEntries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestAppendEntries.Unmarshal(m, b)
}
func (m *RequestAppendEntries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestAppendEntries.Marshal(b, m, deterministic)
}
func (m *RequestAppendEntries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestAppendEntries.Merge(m, src)
}
func (m *RequestAppendEntries) XXX_Size() int {
	return xxx_messageInfo_RequestAppendEntries.Size(m)
}
func (m *RequestAppendEntries) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestAppendEntries.DiscardUnknown(m)
}

var xxx_messageInfo_RequestAppendEntries proto.InternalMessageInfo

func (m *RequestAppendEntries) GetTerm() int64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RequestAppendEntries) GetLeaderId() string {
	if m != nil {
		return m.LeaderId
	}
	return ""
}

func (m *RequestAppendEntries) GetPrevLogIndex() int64 {
	if m != nil {
		return m.PrevLogIndex
	}
	return 0
}

func (m *RequestAppendEntries) GetPrevLogTerm() int64 {
	if m != nil {
		return m.PrevLogTerm
	}
	return 0
}

func (m *RequestAppendEntries) GetEntries() [][]byte {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *RequestAppendEntries) GetLeaderCommit() int64 {
	if m != nil {
		return m.LeaderCommit
	}
	return 0
}

type ResponseAppendEntries struct {
	Term                 int64    `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseAppendEntries) Reset()         { *m = ResponseAppendEntries{} }
func (m *ResponseAppendEntries) String() string { return proto.CompactTextString(m) }
func (*ResponseAppendEntries) ProtoMessage()    {}
func (*ResponseAppendEntries) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{5}
}

func (m *ResponseAppendEntries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseAppendEntries.Unmarshal(m, b)
}
func (m *ResponseAppendEntries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseAppendEntries.Marshal(b, m, deterministic)
}
func (m *ResponseAppendEntries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseAppendEntries.Merge(m, src)
}
func (m *ResponseAppendEntries) XXX_Size() int {
	return xxx_messageInfo_ResponseAppendEntries.Size(m)
}
func (m *ResponseAppendEntries) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseAppendEntries.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseAppendEntries proto.InternalMessageInfo

func (m *ResponseAppendEntries) GetTerm() int64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *ResponseAppendEntries) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

// Invoked by leader to send chunks of a snapshot to a follower. Leaders always
// send chunks in order.
// Receiver implementation:
// 1. Reply immediately if term < currentTerm
// 2. Create new snapshot file if first chunk (offset is 0)
// 3. Write data into snapshot file at given offset
// 4. Reply and wait for more data chunks if done is false
// 5. Save snapshot file, discard any existing or partial snapshot
//    with a smaller index
// 6. If existing log entry has same index and term as snapshot’s
//    last included entry, retain log entries following it and reply
// 7. Discard the entire log
// 8. Reset state machine using snapshot contents (and load
//    snapshot’s cluster configuration)
type RequestInstallSnapshot struct {
	Term              int64  `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	LeaderId          string `protobuf:"bytes,2,opt,name=leaderId,proto3" json:"leaderId,omitempty"`
	LastIncludedIndex int64  `protobuf:"varint,3,opt,name=lastIncludedIndex,proto3" json:"lastIncludedIndex,omitempty"`
	// and including this index
	LastIncludedTerm int64 `protobuf:"varint,4,opt,name=lastIncludedTerm,proto3" json:"lastIncludedTerm,omitempty"`
	Offset           int64 `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	// where chunk is positioned in the snapshot file
	Data                 []byte   `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	Done                 bool     `protobuf:"varint,7,opt,name=done,proto3" json:"done,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestInstallSnapshot) Reset()         { *m = RequestInstallSnapshot{} }
func (m *RequestInstallSnapshot) String() string { return proto.CompactTextString(m) }
func (*RequestInstallSnapshot) ProtoMessage()    {}
func (*RequestInstallSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{6}
}

func (m *RequestInstallSnapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestInstallSnapshot.Unmarshal(m, b)
}
func (m *RequestInstallSnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestInstallSnapshot.Marshal(b, m, deterministic)
}
func (m *RequestInstallSnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestInstallSnapshot.Merge(m, src)
}
func (m *RequestInstallSnapshot) XXX_Size() int {
	return xxx_messageInfo_RequestInstallSnapshot.Size(m)
}
func (m *RequestInstallSnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestInstallSnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_RequestInstallSnapshot proto.InternalMessageInfo

func (m *RequestInstallSnapshot) GetTerm() int64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RequestInstallSnapshot) GetLeaderId() string {
	if m != nil {
		return m.LeaderId
	}
	return ""
}

func (m *RequestInstallSnapshot) GetLastIncludedIndex() int64 {
	if m != nil {
		return m.LastIncludedIndex
	}
	return 0
}

func (m *RequestInstallSnapshot) GetLastIncludedTerm() int64 {
	if m != nil {
		return m.LastIncludedTerm
	}
	return 0
}

func (m *RequestInstallSnapshot) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *RequestInstallSnapshot) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *RequestInstallSnapshot) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

type ResponseInstallSnapshot struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseInstallSnapshot) Reset()         { *m = ResponseInstallSnapshot{} }
func (m *ResponseInstallSnapshot) String() string { return proto.CompactTextString(m) }
func (*ResponseInstallSnapshot) ProtoMessage()    {}
func (*ResponseInstallSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{7}
}

func (m *ResponseInstallSnapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseInstallSnapshot.Unmarshal(m, b)
}
func (m *ResponseInstallSnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseInstallSnapshot.Marshal(b, m, deterministic)
}
func (m *ResponseInstallSnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseInstallSnapshot.Merge(m, src)
}
func (m *ResponseInstallSnapshot) XXX_Size() int {
	return xxx_messageInfo_ResponseInstallSnapshot.Size(m)
}
func (m *ResponseInstallSnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseInstallSnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseInstallSnapshot proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("raft.Status", Status_name, Status_value)
	proto.RegisterType((*Peer)(nil), "raft.Peer")
	proto.RegisterType((*Cluster)(nil), "raft.Cluster")
	proto.RegisterType((*RequestVote)(nil), "raft.RequestVote")
	proto.RegisterType((*ResponseVote)(nil), "raft.ResponseVote")
	proto.RegisterType((*RequestAppendEntries)(nil), "raft.RequestAppendEntries")
	proto.RegisterType((*ResponseAppendEntries)(nil), "raft.ResponseAppendEntries")
	proto.RegisterType((*RequestInstallSnapshot)(nil), "raft.RequestInstallSnapshot")
	proto.RegisterType((*ResponseInstallSnapshot)(nil), "raft.ResponseInstallSnapshot")
}

func init() { proto.RegisterFile("raft.proto", fileDescriptor_b042552c306ae59b) }

var fileDescriptor_b042552c306ae59b = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5d, 0x6e, 0xd4, 0x30,
	0x10, 0x26, 0xfb, 0xbf, 0xb3, 0x29, 0xa4, 0x16, 0x14, 0x13, 0x8a, 0x14, 0xe5, 0x69, 0x55, 0xa0,
	0x12, 0x85, 0x0b, 0xa0, 0xa5, 0xc0, 0x4a, 0x05, 0x21, 0x17, 0xf1, 0x6e, 0xe2, 0x59, 0x1a, 0x29,
	0x6b, 0x07, 0xdb, 0x29, 0x5c, 0x81, 0xab, 0x71, 0x05, 0xde, 0x38, 0x09, 0x8a, 0x37, 0x29, 0x4e,
	0x8a, 0x40, 0xbc, 0xcd, 0x7c, 0x93, 0x99, 0xf9, 0xbe, 0xcf, 0x76, 0x00, 0x34, 0xdf, 0xd8, 0xe3,
	0x52, 0x2b, 0xab, 0xc8, 0xa8, 0x8e, 0xd3, 0x67, 0x30, 0x7a, 0x87, 0xa8, 0xc9, 0x4d, 0x18, 0xe4,
	0x82, 0x06, 0x49, 0xb0, 0x9c, 0xb3, 0x41, 0x2e, 0xc8, 0x21, 0xcc, 0xb9, 0x10, 0x1a, 0x8d, 0x41,
	0x43, 0x07, 0xc9, 0x70, 0x39, 0x67, 0xbf, 0x81, 0xf4, 0x21, 0x4c, 0x57, 0x45, 0x65, 0x2c, 0x6a,
	0x92, 0xc0, 0xb8, 0x44, 0xd4, 0x86, 0x06, 0xc9, 0x70, 0xb9, 0x38, 0x81, 0x63, 0xb7, 0xa2, 0x9e,
	0xc9, 0x76, 0x85, 0xf4, 0x5b, 0x00, 0x0b, 0x86, 0x9f, 0x2b, 0x34, 0xf6, 0x83, 0xb2, 0x48, 0x08,
	0x8c, 0x2c, 0xea, 0xad, 0x5b, 0x36, 0x64, 0x2e, 0x26, 0x09, 0x2c, 0x32, 0x2e, 0x45, 0x2e, 0xb8,
	0xc5, 0xb5, 0xa0, 0x03, 0xc7, 0xc3, 0x87, 0x48, 0x0a, 0x61, 0xc1, 0x8d, 0x3d, 0x53, 0x9f, 0xd6,
	0x52, 0xe0, 0x57, 0x3a, 0x74, 0xdd, 0x1d, 0xac, 0x9e, 0xd2, 0xe4, 0xef, 0xeb, 0x05, 0x23, 0xf7,
	0x89, 0x0f, 0xa5, 0x2f, 0x20, 0x64, 0x68, 0x4a, 0x25, 0x0d, 0xfe, 0x8d, 0xcb, 0xa5, 0xb2, 0xf8,
	0x4a, 0x73, 0x69, 0x71, 0xc7, 0x65, 0xc6, 0x7c, 0x28, 0xfd, 0x1e, 0xc0, 0xed, 0x46, 0xd1, 0xf3,
	0xb2, 0x44, 0x29, 0x4e, 0xa5, 0xd5, 0x39, 0x9a, 0x3f, 0x8e, 0x8b, 0x61, 0x56, 0x20, 0x17, 0xa8,
	0xaf, 0x74, 0x5d, 0xe5, 0xb5, 0xa8, 0x52, 0xe3, 0x65, 0x5f, 0x94, 0x8f, 0xd5, 0x74, 0x9a, 0xdc,
	0x17, 0xe5, 0x41, 0x84, 0xc2, 0x14, 0x77, 0x04, 0xe8, 0x38, 0x19, 0x2e, 0x43, 0xd6, 0xa6, 0xce,
	0x34, 0xb7, 0x6b, 0xa5, 0xb6, 0xdb, 0xdc, 0xd2, 0x49, 0x63, 0x9a, 0x87, 0xa5, 0xa7, 0x70, 0xa7,
	0xb5, 0xe4, 0xdf, 0x62, 0x28, 0x4c, 0x4d, 0x95, 0x65, 0x68, 0x4c, 0xe3, 0x4b, 0x9b, 0xa6, 0x3f,
	0x03, 0x38, 0x68, 0x3c, 0x59, 0x4b, 0x63, 0x79, 0x51, 0x9c, 0x4b, 0x5e, 0x9a, 0x0b, 0x65, 0xff,
	0xdb, 0x95, 0x47, 0xb0, 0x5f, 0x9f, 0xd9, 0x5a, 0x66, 0x45, 0x25, 0x50, 0xf8, 0xd6, 0x5c, 0x2f,
	0x90, 0x23, 0x88, 0x7c, 0xd0, 0x33, 0xe9, 0x1a, 0x4e, 0x0e, 0x60, 0xa2, 0x36, 0x1b, 0x83, 0x96,
	0x8e, 0xdd, 0x17, 0x4d, 0x56, 0x33, 0x14, 0xdc, 0x72, 0xe7, 0x4f, 0xc8, 0x5c, 0xec, 0x30, 0x25,
	0x91, 0x4e, 0x9d, 0x4e, 0x17, 0xa7, 0xf7, 0xe0, 0x6e, 0xeb, 0x55, 0x4f, 0xe4, 0xd1, 0x13, 0x98,
	0x9c, 0x5b, 0x6e, 0x2b, 0x43, 0x00, 0x26, 0x67, 0x4e, 0x4a, 0x74, 0x83, 0x84, 0x30, 0x7b, 0xa9,
	0x8a, 0x42, 0x7d, 0x41, 0x1d, 0x05, 0x64, 0x0f, 0xe6, 0xab, 0xf6, 0x4a, 0x47, 0x83, 0x93, 0x1f,
	0x01, 0x44, 0x8c, 0x6f, 0x6c, 0xdd, 0x87, 0x6f, 0x78, 0x76, 0x91, 0x4b, 0x24, 0x8f, 0x61, 0xe4,
	0x6e, 0xe6, 0xfe, 0xee, 0x21, 0x79, 0x0f, 0x27, 0x26, 0x2d, 0xe4, 0x5d, 0xe0, 0xd7, 0xb0, 0xd7,
	0x3d, 0xb5, 0xb8, 0xd3, 0xd7, 0xa9, 0xc5, 0xf7, 0xbb, 0x03, 0xba, 0x8d, 0x6f, 0xe1, 0x56, 0xff,
	0xe0, 0x0e, 0x3b, 0xb3, 0x7a, 0xd5, 0xf8, 0x41, 0x77, 0x5a, 0xaf, 0xfc, 0x71, 0xe2, 0x7e, 0x33,
	0x4f, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x3b, 0xdc, 0xf6, 0x74, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RaftStateMachineClient is the client API for RaftStateMachine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RaftStateMachineClient interface {
	Vote(ctx context.Context, in *RequestVote, opts ...grpc.CallOption) (*ResponseVote, error)
	AppendEntries(ctx context.Context, in *RequestAppendEntries, opts ...grpc.CallOption) (*ResponseAppendEntries, error)
	InstallSnapshot(ctx context.Context, in *RequestInstallSnapshot, opts ...grpc.CallOption) (*ResponseInstallSnapshot, error)
}

type raftStateMachineClient struct {
	cc *grpc.ClientConn
}

func NewRaftStateMachineClient(cc *grpc.ClientConn) RaftStateMachineClient {
	return &raftStateMachineClient{cc}
}

func (c *raftStateMachineClient) Vote(ctx context.Context, in *RequestVote, opts ...grpc.CallOption) (*ResponseVote, error) {
	out := new(ResponseVote)
	err := c.cc.Invoke(ctx, "/raft.RaftStateMachine/Vote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftStateMachineClient) AppendEntries(ctx context.Context, in *RequestAppendEntries, opts ...grpc.CallOption) (*ResponseAppendEntries, error) {
	out := new(ResponseAppendEntries)
	err := c.cc.Invoke(ctx, "/raft.RaftStateMachine/AppendEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftStateMachineClient) InstallSnapshot(ctx context.Context, in *RequestInstallSnapshot, opts ...grpc.CallOption) (*ResponseInstallSnapshot, error) {
	out := new(ResponseInstallSnapshot)
	err := c.cc.Invoke(ctx, "/raft.RaftStateMachine/InstallSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RaftStateMachineServer is the server API for RaftStateMachine service.
type RaftStateMachineServer interface {
	Vote(context.Context, *RequestVote) (*ResponseVote, error)
	AppendEntries(context.Context, *RequestAppendEntries) (*ResponseAppendEntries, error)
	InstallSnapshot(context.Context, *RequestInstallSnapshot) (*ResponseInstallSnapshot, error)
}

func RegisterRaftStateMachineServer(s *grpc.Server, srv RaftStateMachineServer) {
	s.RegisterService(&_RaftStateMachine_serviceDesc, srv)
}

func _RaftStateMachine_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftStateMachineServer).Vote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft.RaftStateMachine/Vote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftStateMachineServer).Vote(ctx, req.(*RequestVote))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftStateMachine_AppendEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAppendEntries)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftStateMachineServer).AppendEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft.RaftStateMachine/AppendEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftStateMachineServer).AppendEntries(ctx, req.(*RequestAppendEntries))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftStateMachine_InstallSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestInstallSnapshot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftStateMachineServer).InstallSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft.RaftStateMachine/InstallSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftStateMachineServer).InstallSnapshot(ctx, req.(*RequestInstallSnapshot))
	}
	return interceptor(ctx, in, info, handler)
}

var _RaftStateMachine_serviceDesc = grpc.ServiceDesc{
	ServiceName: "raft.RaftStateMachine",
	HandlerType: (*RaftStateMachineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Vote",
			Handler:    _RaftStateMachine_Vote_Handler,
		},
		{
			MethodName: "AppendEntries",
			Handler:    _RaftStateMachine_AppendEntries_Handler,
		},
		{
			MethodName: "InstallSnapshot",
			Handler:    _RaftStateMachine_InstallSnapshot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "raft.proto",
}