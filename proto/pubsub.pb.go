// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: pubsub.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

type PublishRespond_PublishStatus int32

const (
	PublishRespond_SUCCESS PublishRespond_PublishStatus = 0
	PublishRespond_FAIL    PublishRespond_PublishStatus = 1
)

// Enum value maps for PublishRespond_PublishStatus.
var (
	PublishRespond_PublishStatus_name = map[int32]string{
		0: "SUCCESS",
		1: "FAIL",
	}
	PublishRespond_PublishStatus_value = map[string]int32{
		"SUCCESS": 0,
		"FAIL":    1,
	}
)

func (x PublishRespond_PublishStatus) Enum() *PublishRespond_PublishStatus {
	p := new(PublishRespond_PublishStatus)
	*p = x
	return p
}

func (x PublishRespond_PublishStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PublishRespond_PublishStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_pubsub_proto_enumTypes[0].Descriptor()
}

func (PublishRespond_PublishStatus) Type() protoreflect.EnumType {
	return &file_pubsub_proto_enumTypes[0]
}

func (x PublishRespond_PublishStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PublishRespond_PublishStatus.Descriptor instead.
func (PublishRespond_PublishStatus) EnumDescriptor() ([]byte, []int) {
	return file_pubsub_proto_rawDescGZIP(), []int{1, 0}
}

// The request message for publish
type PublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicMessagePairs []*TopicMessagesPair `protobuf:"bytes,1,rep,name=topic_message_pairs,json=topicMessagePairs,proto3" json:"topic_message_pairs,omitempty"`
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pubsub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pubsub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_pubsub_proto_rawDescGZIP(), []int{0}
}

func (x *PublishRequest) GetTopicMessagePairs() []*TopicMessagesPair {
	if x != nil {
		return x.TopicMessagePairs
	}
	return nil
}

// The response message for publish
type PublishRespond struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status PublishRespond_PublishStatus `protobuf:"varint,1,opt,name=status,proto3,enum=pubsub.PublishRespond_PublishStatus" json:"status,omitempty"`
}

func (x *PublishRespond) Reset() {
	*x = PublishRespond{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pubsub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRespond) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRespond) ProtoMessage() {}

func (x *PublishRespond) ProtoReflect() protoreflect.Message {
	mi := &file_pubsub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRespond.ProtoReflect.Descriptor instead.
func (*PublishRespond) Descriptor() ([]byte, []int) {
	return file_pubsub_proto_rawDescGZIP(), []int{1}
}

func (x *PublishRespond) GetStatus() PublishRespond_PublishStatus {
	if x != nil {
		return x.Status
	}
	return PublishRespond_SUCCESS
}

// The request message for subscribe
type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topics []*Topic `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pubsub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pubsub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_pubsub_proto_rawDescGZIP(), []int{2}
}

func (x *SubscribeRequest) GetTopics() []*Topic {
	if x != nil {
		return x.Topics
	}
	return nil
}

// The response message for subscribe
type SubscribeRespond struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicMessagePairs []*TopicMessagesPair `protobuf:"bytes,1,rep,name=topic_message_pairs,json=topicMessagePairs,proto3" json:"topic_message_pairs,omitempty"`
}

func (x *SubscribeRespond) Reset() {
	*x = SubscribeRespond{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pubsub_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRespond) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRespond) ProtoMessage() {}

func (x *SubscribeRespond) ProtoReflect() protoreflect.Message {
	mi := &file_pubsub_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRespond.ProtoReflect.Descriptor instead.
func (*SubscribeRespond) Descriptor() ([]byte, []int) {
	return file_pubsub_proto_rawDescGZIP(), []int{3}
}

func (x *SubscribeRespond) GetTopicMessagePairs() []*TopicMessagesPair {
	if x != nil {
		return x.TopicMessagePairs
	}
	return nil
}

// Topic to messages pair
type TopicMessagesPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic    *Topic     `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Messages []*Message `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *TopicMessagesPair) Reset() {
	*x = TopicMessagesPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pubsub_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicMessagesPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicMessagesPair) ProtoMessage() {}

func (x *TopicMessagesPair) ProtoReflect() protoreflect.Message {
	mi := &file_pubsub_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicMessagesPair.ProtoReflect.Descriptor instead.
func (*TopicMessagesPair) Descriptor() ([]byte, []int) {
	return file_pubsub_proto_rawDescGZIP(), []int{4}
}

func (x *TopicMessagesPair) GetTopic() *Topic {
	if x != nil {
		return x.Topic
	}
	return nil
}

func (x *TopicMessagesPair) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

// Defination for a single message
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pubsub_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_pubsub_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_pubsub_proto_rawDescGZIP(), []int{5}
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// Defination for a topic
type Topic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *Topic) Reset() {
	*x = Topic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pubsub_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topic) ProtoMessage() {}

func (x *Topic) ProtoReflect() protoreflect.Message {
	mi := &file_pubsub_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topic.ProtoReflect.Descriptor instead.
func (*Topic) Descriptor() ([]byte, []int) {
	return file_pubsub_proto_rawDescGZIP(), []int{6}
}

func (x *Topic) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

var File_pubsub_proto protoreflect.FileDescriptor

var file_pubsub_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x22, 0x5b, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x13, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x50, 0x61, 0x69, 0x72,
	0x52, 0x11, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61,
	0x69, 0x72, 0x73, 0x22, 0x76, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x64, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x2e, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x26, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x22, 0x39, 0x0a, 0x10, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x06,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x22, 0x5d, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x12, 0x49, 0x0a, 0x13, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x69, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62,
	0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x50, 0x61,
	0x69, 0x72, 0x52, 0x11, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x50, 0x61, 0x69, 0x72, 0x73, 0x22, 0x65, 0x0a, 0x11, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x50, 0x61, 0x69, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x75, 0x62, 0x73,
	0x75, 0x62, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12,
	0x2b, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x23, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x1d, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x32, 0x8a, 0x01, 0x0a, 0x06, 0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x12, 0x3b, 0x0a, 0x07, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x16, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x22, 0x00, 0x30, 0x01, 0x42, 0x28, 0x0a,
	0x09, 0x69, 0x6f, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x42, 0x0b, 0x50, 0x75, 0x62, 0x53,
	0x75, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0c, 0x70, 0x75, 0x62, 0x73, 0x75,
	0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pubsub_proto_rawDescOnce sync.Once
	file_pubsub_proto_rawDescData = file_pubsub_proto_rawDesc
)

func file_pubsub_proto_rawDescGZIP() []byte {
	file_pubsub_proto_rawDescOnce.Do(func() {
		file_pubsub_proto_rawDescData = protoimpl.X.CompressGZIP(file_pubsub_proto_rawDescData)
	})
	return file_pubsub_proto_rawDescData
}

var file_pubsub_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pubsub_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pubsub_proto_goTypes = []interface{}{
	(PublishRespond_PublishStatus)(0), // 0: pubsub.PublishRespond.PublishStatus
	(*PublishRequest)(nil),            // 1: pubsub.PublishRequest
	(*PublishRespond)(nil),            // 2: pubsub.PublishRespond
	(*SubscribeRequest)(nil),          // 3: pubsub.SubscribeRequest
	(*SubscribeRespond)(nil),          // 4: pubsub.SubscribeRespond
	(*TopicMessagesPair)(nil),         // 5: pubsub.TopicMessagesPair
	(*Message)(nil),                   // 6: pubsub.Message
	(*Topic)(nil),                     // 7: pubsub.Topic
}
var file_pubsub_proto_depIdxs = []int32{
	5, // 0: pubsub.PublishRequest.topic_message_pairs:type_name -> pubsub.TopicMessagesPair
	0, // 1: pubsub.PublishRespond.status:type_name -> pubsub.PublishRespond.PublishStatus
	7, // 2: pubsub.SubscribeRequest.topics:type_name -> pubsub.Topic
	5, // 3: pubsub.SubscribeRespond.topic_message_pairs:type_name -> pubsub.TopicMessagesPair
	7, // 4: pubsub.TopicMessagesPair.topic:type_name -> pubsub.Topic
	6, // 5: pubsub.TopicMessagesPair.messages:type_name -> pubsub.Message
	1, // 6: pubsub.PubSub.publish:input_type -> pubsub.PublishRequest
	3, // 7: pubsub.PubSub.subscribe:input_type -> pubsub.SubscribeRequest
	2, // 8: pubsub.PubSub.publish:output_type -> pubsub.PublishRespond
	4, // 9: pubsub.PubSub.subscribe:output_type -> pubsub.SubscribeRespond
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pubsub_proto_init() }
func file_pubsub_proto_init() {
	if File_pubsub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pubsub_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRequest); i {
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
		file_pubsub_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRespond); i {
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
		file_pubsub_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_pubsub_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRespond); i {
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
		file_pubsub_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicMessagesPair); i {
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
		file_pubsub_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_pubsub_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Topic); i {
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
			RawDescriptor: file_pubsub_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pubsub_proto_goTypes,
		DependencyIndexes: file_pubsub_proto_depIdxs,
		EnumInfos:         file_pubsub_proto_enumTypes,
		MessageInfos:      file_pubsub_proto_msgTypes,
	}.Build()
	File_pubsub_proto = out.File
	file_pubsub_proto_rawDesc = nil
	file_pubsub_proto_goTypes = nil
	file_pubsub_proto_depIdxs = nil
}
