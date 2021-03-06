// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.2
// source: proto/raft.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Role int32

const (
	Role_Leader    Role = 0
	Role_Follower  Role = 1
	Role_Candidate Role = 2
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "Leader",
		1: "Follower",
		2: "Candidate",
	}
	Role_value = map[string]int32{
		"Leader":    0,
		"Follower":  1,
		"Candidate": 2,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_raft_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_proto_raft_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{0}
}

type AppendEntriesResponse_Status int32

const (
	AppendEntriesResponse_SUCCESS AppendEntriesResponse_Status = 0
	AppendEntriesResponse_FAILURE AppendEntriesResponse_Status = 1
)

// Enum value maps for AppendEntriesResponse_Status.
var (
	AppendEntriesResponse_Status_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILURE",
	}
	AppendEntriesResponse_Status_value = map[string]int32{
		"SUCCESS": 0,
		"FAILURE": 1,
	}
)

func (x AppendEntriesResponse_Status) Enum() *AppendEntriesResponse_Status {
	p := new(AppendEntriesResponse_Status)
	*p = x
	return p
}

func (x AppendEntriesResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppendEntriesResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_raft_proto_enumTypes[1].Descriptor()
}

func (AppendEntriesResponse_Status) Type() protoreflect.EnumType {
	return &file_proto_raft_proto_enumTypes[1]
}

func (x AppendEntriesResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppendEntriesResponse_Status.Descriptor instead.
func (AppendEntriesResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{1, 0}
}

type AppendEntriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Leader's term
	Term uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	// Leader's id in form of addr:port. So that the follower can redirect the client
	LeaderId string `protobuf:"bytes,2,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
	// Index of log entry immediately preceding new ones
	PrevLogIndex uint64 `protobuf:"varint,3,opt,name=prev_log_index,json=prevLogIndex,proto3" json:"prev_log_index,omitempty"`
	// Term of prevLogIndex
	PrevLogTerm uint64 `protobuf:"varint,4,opt,name=prev_log_term,json=prevLogTerm,proto3" json:"prev_log_term,omitempty"`
	// Log entries to store (empty for heartbeat)
	Entries []*LogEntry `protobuf:"bytes,5,rep,name=entries,proto3" json:"entries,omitempty"`
	// Leader's commit index
	Commit uint64 `protobuf:"varint,6,opt,name=commit,proto3" json:"commit,omitempty"`
}

func (x *AppendEntriesRequest) Reset() {
	*x = AppendEntriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raft_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntriesRequest) ProtoMessage() {}

func (x *AppendEntriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntriesRequest.ProtoReflect.Descriptor instead.
func (*AppendEntriesRequest) Descriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{0}
}

func (x *AppendEntriesRequest) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntriesRequest) GetLeaderId() string {
	if x != nil {
		return x.LeaderId
	}
	return ""
}

func (x *AppendEntriesRequest) GetPrevLogIndex() uint64 {
	if x != nil {
		return x.PrevLogIndex
	}
	return 0
}

func (x *AppendEntriesRequest) GetPrevLogTerm() uint64 {
	if x != nil {
		return x.PrevLogTerm
	}
	return 0
}

func (x *AppendEntriesRequest) GetEntries() []*LogEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *AppendEntriesRequest) GetCommit() uint64 {
	if x != nil {
		return x.Commit
	}
	return 0
}

type AppendEntriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Current term, for leader to update itself
	Term uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	// Status to indicate whether the RPC is success
	// SUCCESS if the follower contained entry matching prevLogIndex and prevLogTerm
	Status AppendEntriesResponse_Status `protobuf:"varint,2,opt,name=status,proto3,enum=pubsub.AppendEntriesResponse_Status" json:"status,omitempty"`
	// Additional information
	Info string `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *AppendEntriesResponse) Reset() {
	*x = AppendEntriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raft_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntriesResponse) ProtoMessage() {}

func (x *AppendEntriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntriesResponse.ProtoReflect.Descriptor instead.
func (*AppendEntriesResponse) Descriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{1}
}

func (x *AppendEntriesResponse) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntriesResponse) GetStatus() AppendEntriesResponse_Status {
	if x != nil {
		return x.Status
	}
	return AppendEntriesResponse_SUCCESS
}

func (x *AppendEntriesResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

type RequestVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Candidate's term
	Term uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	// The id of the candidate that is requesting the vote (addr:port)
	CandidateId string `protobuf:"bytes,2,opt,name=candidate_id,json=candidateId,proto3" json:"candidate_id,omitempty"`
	// Index of the candidate's last log entry
	LastLogIndex uint64 `protobuf:"varint,3,opt,name=last_log_index,json=lastLogIndex,proto3" json:"last_log_index,omitempty"`
	// Term of the candidate's last log entry
	LastLogTerm uint64 `protobuf:"varint,4,opt,name=last_log_term,json=lastLogTerm,proto3" json:"last_log_term,omitempty"`
}

func (x *RequestVoteRequest) Reset() {
	*x = RequestVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raft_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteRequest) ProtoMessage() {}

func (x *RequestVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteRequest.ProtoReflect.Descriptor instead.
func (*RequestVoteRequest) Descriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{2}
}

func (x *RequestVoteRequest) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteRequest) GetCandidateId() string {
	if x != nil {
		return x.CandidateId
	}
	return ""
}

func (x *RequestVoteRequest) GetLastLogIndex() uint64 {
	if x != nil {
		return x.LastLogIndex
	}
	return 0
}

func (x *RequestVoteRequest) GetLastLogTerm() uint64 {
	if x != nil {
		return x.LastLogTerm
	}
	return 0
}

type RequestVoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Current term, for candidate to update itself
	Term uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	// True if the candidate received the vote
	VoteGranted bool `protobuf:"varint,2,opt,name=vote_granted,json=voteGranted,proto3" json:"vote_granted,omitempty"`
}

func (x *RequestVoteResponse) Reset() {
	*x = RequestVoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raft_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteResponse) ProtoMessage() {}

func (x *RequestVoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteResponse.ProtoReflect.Descriptor instead.
func (*RequestVoteResponse) Descriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{3}
}

func (x *RequestVoteResponse) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteResponse) GetVoteGranted() bool {
	if x != nil {
		return x.VoteGranted
	}
	return false
}

type LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term  uint64   `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Topic *Topic   `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Msg   *Message `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *LogEntry) Reset() {
	*x = LogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raft_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntry) ProtoMessage() {}

func (x *LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntry.ProtoReflect.Descriptor instead.
func (*LogEntry) Descriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{4}
}

func (x *LogEntry) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *LogEntry) GetTopic() *Topic {
	if x != nil {
		return x.Topic
	}
	return nil
}

func (x *LogEntry) GetMsg() *Message {
	if x != nil {
		return x.Msg
	}
	return nil
}

var File_proto_raft_proto protoreflect.FileDescriptor

var file_proto_raft_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5,
	0x01, 0x0a, 0x14, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x76,
	0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x22,
	0x0a, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65,
	0x72, 0x6d, 0x12, 0x2a, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x4c, 0x6f, 0x67,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x22, 0xa1, 0x01, 0x0a, 0x15, 0x41, 0x70, 0x70, 0x65, 0x6e,
	0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x74, 0x65, 0x72, 0x6d, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x41, 0x70,
	0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x22, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x01, 0x22, 0x95, 0x01, 0x0a, 0x12, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x22,
	0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x54, 0x65,
	0x72, 0x6d, 0x22, 0x4c, 0x0a, 0x13, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x21, 0x0a,
	0x0c, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64,
	0x22, 0x66, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d,
	0x12, 0x23, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x21, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x2a, 0x2f, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65,
	0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x61,
	0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x10, 0x02, 0x32, 0xa7, 0x01, 0x0a, 0x0b, 0x52, 0x61,
	0x66, 0x74, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x12, 0x4e, 0x0a, 0x0d, 0x41, 0x70, 0x70,
	0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x70, 0x75, 0x62,
	0x73, 0x75, 0x62, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75,
	0x62, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75,
	0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_raft_proto_rawDescOnce sync.Once
	file_proto_raft_proto_rawDescData = file_proto_raft_proto_rawDesc
)

func file_proto_raft_proto_rawDescGZIP() []byte {
	file_proto_raft_proto_rawDescOnce.Do(func() {
		file_proto_raft_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_raft_proto_rawDescData)
	})
	return file_proto_raft_proto_rawDescData
}

var file_proto_raft_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_raft_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_raft_proto_goTypes = []interface{}{
	(Role)(0),                         // 0: pubsub.Role
	(AppendEntriesResponse_Status)(0), // 1: pubsub.AppendEntriesResponse.Status
	(*AppendEntriesRequest)(nil),      // 2: pubsub.AppendEntriesRequest
	(*AppendEntriesResponse)(nil),     // 3: pubsub.AppendEntriesResponse
	(*RequestVoteRequest)(nil),        // 4: pubsub.RequestVoteRequest
	(*RequestVoteResponse)(nil),       // 5: pubsub.RequestVoteResponse
	(*LogEntry)(nil),                  // 6: pubsub.LogEntry
	(*Topic)(nil),                     // 7: pubsub.Topic
	(*Message)(nil),                   // 8: pubsub.Message
}
var file_proto_raft_proto_depIdxs = []int32{
	6, // 0: pubsub.AppendEntriesRequest.entries:type_name -> pubsub.LogEntry
	1, // 1: pubsub.AppendEntriesResponse.status:type_name -> pubsub.AppendEntriesResponse.Status
	7, // 2: pubsub.LogEntry.topic:type_name -> pubsub.Topic
	8, // 3: pubsub.LogEntry.msg:type_name -> pubsub.Message
	2, // 4: pubsub.RaftSidecar.AppendEntries:input_type -> pubsub.AppendEntriesRequest
	4, // 5: pubsub.RaftSidecar.RequestVote:input_type -> pubsub.RequestVoteRequest
	3, // 6: pubsub.RaftSidecar.AppendEntries:output_type -> pubsub.AppendEntriesResponse
	5, // 7: pubsub.RaftSidecar.RequestVote:output_type -> pubsub.RequestVoteResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_raft_proto_init() }
func file_proto_raft_proto_init() {
	if File_proto_raft_proto != nil {
		return
	}
	file_proto_pubsub_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_raft_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntriesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_raft_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntriesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_raft_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestVoteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_raft_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestVoteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_raft_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_raft_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_raft_proto_goTypes,
		DependencyIndexes: file_proto_raft_proto_depIdxs,
		EnumInfos:         file_proto_raft_proto_enumTypes,
		MessageInfos:      file_proto_raft_proto_msgTypes,
	}.Build()
	File_proto_raft_proto = out.File
	file_proto_raft_proto_rawDesc = nil
	file_proto_raft_proto_goTypes = nil
	file_proto_raft_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RaftSidecarClient is the client API for RaftSidecar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RaftSidecarClient interface {
	// Invoked by leader to replicate log entries. Also used as heart beat.
	AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error)
	// Invoked by candidates to gather votes
	RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error)
}

type raftSidecarClient struct {
	cc grpc.ClientConnInterface
}

func NewRaftSidecarClient(cc grpc.ClientConnInterface) RaftSidecarClient {
	return &raftSidecarClient{cc}
}

func (c *raftSidecarClient) AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error) {
	out := new(AppendEntriesResponse)
	err := c.cc.Invoke(ctx, "/pubsub.RaftSidecar/AppendEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftSidecarClient) RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error) {
	out := new(RequestVoteResponse)
	err := c.cc.Invoke(ctx, "/pubsub.RaftSidecar/RequestVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RaftSidecarServer is the server API for RaftSidecar service.
type RaftSidecarServer interface {
	// Invoked by leader to replicate log entries. Also used as heart beat.
	AppendEntries(context.Context, *AppendEntriesRequest) (*AppendEntriesResponse, error)
	// Invoked by candidates to gather votes
	RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteResponse, error)
}

// UnimplementedRaftSidecarServer can be embedded to have forward compatible implementations.
type UnimplementedRaftSidecarServer struct {
}

func (*UnimplementedRaftSidecarServer) AppendEntries(context.Context, *AppendEntriesRequest) (*AppendEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendEntries not implemented")
}
func (*UnimplementedRaftSidecarServer) RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestVote not implemented")
}

func RegisterRaftSidecarServer(s *grpc.Server, srv RaftSidecarServer) {
	s.RegisterService(&_RaftSidecar_serviceDesc, srv)
}

func _RaftSidecar_AppendEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftSidecarServer).AppendEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pubsub.RaftSidecar/AppendEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftSidecarServer).AppendEntries(ctx, req.(*AppendEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftSidecar_RequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftSidecarServer).RequestVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pubsub.RaftSidecar/RequestVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftSidecarServer).RequestVote(ctx, req.(*RequestVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RaftSidecar_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pubsub.RaftSidecar",
	HandlerType: (*RaftSidecarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppendEntries",
			Handler:    _RaftSidecar_AppendEntries_Handler,
		},
		{
			MethodName: "RequestVote",
			Handler:    _RaftSidecar_RequestVote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/raft.proto",
}
