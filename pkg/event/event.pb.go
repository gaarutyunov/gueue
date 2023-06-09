// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: pkg/event/event.proto

package event

import (
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

// Event to send to a specific topic
type EventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Topic to send event into
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	// Event correlation id
	CorrelationId string `protobuf:"bytes,2,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	// Event body
	Message []byte `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EventRequest) Reset() {
	*x = EventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_event_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventRequest) ProtoMessage() {}

func (x *EventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_event_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventRequest.ProtoReflect.Descriptor instead.
func (*EventRequest) Descriptor() ([]byte, []int) {
	return file_pkg_event_event_proto_rawDescGZIP(), []int{0}
}

func (x *EventRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *EventRequest) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *EventRequest) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

// Event response with event id
type EventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Event unique id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EventResponse) Reset() {
	*x = EventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_event_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventResponse) ProtoMessage() {}

func (x *EventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_event_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventResponse.ProtoReflect.Descriptor instead.
func (*EventResponse) Descriptor() ([]byte, []int) {
	return file_pkg_event_event_proto_rawDescGZIP(), []int{1}
}

func (x *EventResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Event sent over stream
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Event unique id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Topic to send event into
	Topic string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	// Event correlation id
	CorrelationId string `protobuf:"bytes,3,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	// Event body
	Message []byte `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_event_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_event_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_pkg_event_event_proto_rawDescGZIP(), []int{2}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Event) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *Event) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

// Request for topic subscription
type TopicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Topic names
	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	// The subscribing consumer id
	ConsumerId string `protobuf:"bytes,2,opt,name=consumer_id,json=consumerId,proto3" json:"consumer_id,omitempty"`
	// Size of buffer
	Buffer uint32 `protobuf:"varint,3,opt,name=buffer,proto3" json:"buffer,omitempty"`
}

func (x *TopicRequest) Reset() {
	*x = TopicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_event_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicRequest) ProtoMessage() {}

func (x *TopicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_event_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicRequest.ProtoReflect.Descriptor instead.
func (*TopicRequest) Descriptor() ([]byte, []int) {
	return file_pkg_event_event_proto_rawDescGZIP(), []int{3}
}

func (x *TopicRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *TopicRequest) GetConsumerId() string {
	if x != nil {
		return x.ConsumerId
	}
	return ""
}

func (x *TopicRequest) GetBuffer() uint32 {
	if x != nil {
		return x.Buffer
	}
	return 0
}

// Unbind consumer from topic request
type UnbindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Consumer id to unbind
	ConsumerId string `protobuf:"bytes,1,opt,name=consumer_id,json=consumerId,proto3" json:"consumer_id,omitempty"`
}

func (x *UnbindRequest) Reset() {
	*x = UnbindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_event_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnbindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnbindRequest) ProtoMessage() {}

func (x *UnbindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_event_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnbindRequest.ProtoReflect.Descriptor instead.
func (*UnbindRequest) Descriptor() ([]byte, []int) {
	return file_pkg_event_event_proto_rawDescGZIP(), []int{4}
}

func (x *UnbindRequest) GetConsumerId() string {
	if x != nil {
		return x.ConsumerId
	}
	return ""
}

// Response with operation status
type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Operation status
	Status uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_event_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_event_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_pkg_event_event_proto_rawDescGZIP(), []int{5}
}

func (x *StatusResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_pkg_event_event_proto protoreflect.FileDescriptor

var file_pkg_event_event_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x65,
	0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f,
	0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6e, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5d, 0x0a, 0x0c, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x62,
	0x75, 0x66, 0x66, 0x65, 0x72, 0x22, 0x30, 0x0a, 0x0d, 0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x32, 0x42, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x12, 0x36, 0x0a,
	0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x13, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x75, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x12, 0x30, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x13, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22,
	0x00, 0x30, 0x01, 0x12, 0x37, 0x0a, 0x06, 0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x12, 0x14, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x61, 0x72, 0x75,
	0x74, 0x79, 0x75, 0x6e, 0x6f, 0x76, 0x2f, 0x67, 0x75, 0x65, 0x75, 0x65, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_event_event_proto_rawDescOnce sync.Once
	file_pkg_event_event_proto_rawDescData = file_pkg_event_event_proto_rawDesc
)

func file_pkg_event_event_proto_rawDescGZIP() []byte {
	file_pkg_event_event_proto_rawDescOnce.Do(func() {
		file_pkg_event_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_event_event_proto_rawDescData)
	})
	return file_pkg_event_event_proto_rawDescData
}

var file_pkg_event_event_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_event_event_proto_goTypes = []interface{}{
	(*EventRequest)(nil),   // 0: event.EventRequest
	(*EventResponse)(nil),  // 1: event.EventResponse
	(*Event)(nil),          // 2: event.Event
	(*TopicRequest)(nil),   // 3: event.TopicRequest
	(*UnbindRequest)(nil),  // 4: event.UnbindRequest
	(*StatusResponse)(nil), // 5: event.StatusResponse
}
var file_pkg_event_event_proto_depIdxs = []int32{
	0, // 0: event.Producer.Publish:input_type -> event.EventRequest
	3, // 1: event.Consumer.Consume:input_type -> event.TopicRequest
	4, // 2: event.Consumer.Unbind:input_type -> event.UnbindRequest
	1, // 3: event.Producer.Publish:output_type -> event.EventResponse
	2, // 4: event.Consumer.Consume:output_type -> event.Event
	5, // 5: event.Consumer.Unbind:output_type -> event.StatusResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_event_event_proto_init() }
func file_pkg_event_event_proto_init() {
	if File_pkg_event_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_event_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventRequest); i {
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
		file_pkg_event_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventResponse); i {
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
		file_pkg_event_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_pkg_event_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicRequest); i {
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
		file_pkg_event_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnbindRequest); i {
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
		file_pkg_event_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
			RawDescriptor: file_pkg_event_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_pkg_event_event_proto_goTypes,
		DependencyIndexes: file_pkg_event_event_proto_depIdxs,
		MessageInfos:      file_pkg_event_event_proto_msgTypes,
	}.Build()
	File_pkg_event_event_proto = out.File
	file_pkg_event_event_proto_rawDesc = nil
	file_pkg_event_event_proto_goTypes = nil
	file_pkg_event_event_proto_depIdxs = nil
}
