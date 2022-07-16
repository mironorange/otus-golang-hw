// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.14.0
// source: EventService.proto

package pb

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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UUID - уникальный идентификатор события.
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Заголовок - короткий текст.
	Summary string `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	// Unix timestamp даты и времени начала события.
	StartedAt int32 `protobuf:"varint,3,opt,name=startedAt,proto3" json:"startedAt,omitempty"`
	// Unix timestamp даты и времени завершения события.
	FinishedAt int32 `protobuf:"varint,4,opt,name=finishedAt,proto3" json:"finishedAt,omitempty"`
	// Описание события - длинный текст.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// UUID пользователя, владельца события.
	UserUuid string `protobuf:"bytes,6,opt,name=userUuid,proto3" json:"userUuid,omitempty"`
	// Unix timestamp даты и времени уведомления о событии.
	NotificationAt int32 `protobuf:"varint,7,opt,name=notificationAt,proto3" json:"notificationAt,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[0]
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
	return file_EventService_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Event) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Event) GetStartedAt() int32 {
	if x != nil {
		return x.StartedAt
	}
	return 0
}

func (x *Event) GetFinishedAt() int32 {
	if x != nil {
		return x.FinishedAt
	}
	return 0
}

func (x *Event) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Event) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *Event) GetNotificationAt() int32 {
	if x != nil {
		return x.NotificationAt
	}
	return 0
}

type UpdateEventAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Заголовок - короткий текст.
	Summary string `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
	// Unix timestamp даты и времени начала события.
	StartedAt int32 `protobuf:"varint,2,opt,name=startedAt,proto3" json:"startedAt,omitempty"`
	// Unix timestamp даты и времени завершения события.
	FinishedAt int32 `protobuf:"varint,3,opt,name=finishedAt,proto3" json:"finishedAt,omitempty"`
	// Описание события - длинный текст.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// UUID пользователя, владельца события.
	UserUuid string `protobuf:"bytes,5,opt,name=userUuid,proto3" json:"userUuid,omitempty"`
	// Unix timestamp даты и времени уведомления о событии.
	NotificationAt int32 `protobuf:"varint,6,opt,name=notificationAt,proto3" json:"notificationAt,omitempty"`
}

func (x *UpdateEventAttributes) Reset() {
	*x = UpdateEventAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEventAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventAttributes) ProtoMessage() {}

func (x *UpdateEventAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEventAttributes.ProtoReflect.Descriptor instead.
func (*UpdateEventAttributes) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateEventAttributes) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *UpdateEventAttributes) GetStartedAt() int32 {
	if x != nil {
		return x.StartedAt
	}
	return 0
}

func (x *UpdateEventAttributes) GetFinishedAt() int32 {
	if x != nil {
		return x.FinishedAt
	}
	return 0
}

func (x *UpdateEventAttributes) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateEventAttributes) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *UpdateEventAttributes) GetNotificationAt() int32 {
	if x != nil {
		return x.NotificationAt
	}
	return 0
}

type CreateEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UUID - уникальный идентификатор события.
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Заголовок - короткий текст.
	Summary string `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	// Unix timestamp даты и времени начала события.
	StartedAt int32 `protobuf:"varint,3,opt,name=startedAt,proto3" json:"startedAt,omitempty"`
	// Unix timestamp даты и времени завершения события.
	FinishedAt int32 `protobuf:"varint,4,opt,name=finishedAt,proto3" json:"finishedAt,omitempty"`
	// Описание события - длинный текст.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// UUID пользователя, владельца события.
	UserUuid string `protobuf:"bytes,6,opt,name=userUuid,proto3" json:"userUuid,omitempty"`
	// Unix timestamp даты и времени уведомления о событии.
	NotificationAt int32 `protobuf:"varint,7,opt,name=notificationAt,proto3" json:"notificationAt,omitempty"`
}

func (x *CreateEventRequest) Reset() {
	*x = CreateEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventRequest) ProtoMessage() {}

func (x *CreateEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventRequest.ProtoReflect.Descriptor instead.
func (*CreateEventRequest) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{2}
}

func (x *CreateEventRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *CreateEventRequest) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *CreateEventRequest) GetStartedAt() int32 {
	if x != nil {
		return x.StartedAt
	}
	return 0
}

func (x *CreateEventRequest) GetFinishedAt() int32 {
	if x != nil {
		return x.FinishedAt
	}
	return 0
}

func (x *CreateEventRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateEventRequest) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *CreateEventRequest) GetNotificationAt() int32 {
	if x != nil {
		return x.NotificationAt
	}
	return 0
}

type CreateEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateEventResponse) Reset() {
	*x = CreateEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventResponse) ProtoMessage() {}

func (x *CreateEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventResponse.ProtoReflect.Descriptor instead.
func (*CreateEventResponse) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{3}
}

type UpdateEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid       string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Attributes *UpdateEventAttributes `protobuf:"bytes,2,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *UpdateEventRequest) Reset() {
	*x = UpdateEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventRequest) ProtoMessage() {}

func (x *UpdateEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEventRequest.ProtoReflect.Descriptor instead.
func (*UpdateEventRequest) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateEventRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *UpdateEventRequest) GetAttributes() *UpdateEventAttributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type UpdateEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateEventResponse) Reset() {
	*x = UpdateEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventResponse) ProtoMessage() {}

func (x *UpdateEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEventResponse.ProtoReflect.Descriptor instead.
func (*UpdateEventResponse) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{5}
}

type GetEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SinceNotificationAt int32 `protobuf:"varint,1,opt,name=sinceNotificationAt,proto3" json:"sinceNotificationAt,omitempty"`
}

func (x *GetEventsRequest) Reset() {
	*x = GetEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsRequest) ProtoMessage() {}

func (x *GetEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsRequest.ProtoReflect.Descriptor instead.
func (*GetEventsRequest) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{6}
}

func (x *GetEventsRequest) GetSinceNotificationAt() int32 {
	if x != nil {
		return x.SinceNotificationAt
	}
	return 0
}

type GetEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Event `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetEventsResponse) Reset() {
	*x = GetEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsResponse) ProtoMessage() {}

func (x *GetEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsResponse.ProtoReflect.Descriptor instead.
func (*GetEventsResponse) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{7}
}

func (x *GetEventsResponse) GetItems() []*Event {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetEventRequest) Reset() {
	*x = GetEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventRequest) ProtoMessage() {}

func (x *GetEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventRequest.ProtoReflect.Descriptor instead.
func (*GetEventRequest) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{8}
}

func (x *GetEventRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *Event `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *GetEventResponse) Reset() {
	*x = GetEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventResponse) ProtoMessage() {}

func (x *GetEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventResponse.ProtoReflect.Descriptor instead.
func (*GetEventResponse) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{9}
}

func (x *GetEventResponse) GetItem() *Event {
	if x != nil {
		return x.Item
	}
	return nil
}

type DeleteEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *DeleteEventRequest) Reset() {
	*x = DeleteEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventRequest) ProtoMessage() {}

func (x *DeleteEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventRequest.ProtoReflect.Descriptor instead.
func (*DeleteEventRequest) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteEventRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type DeleteEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteEventResponse) Reset() {
	*x = DeleteEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventResponse) ProtoMessage() {}

func (x *DeleteEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventResponse.ProtoReflect.Descriptor instead.
func (*DeleteEventResponse) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{11}
}

var File_EventService_proto protoreflect.FileDescriptor

var file_EventService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0xd9, 0x01, 0x0a, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12,
	0x26, 0x0a, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x22, 0xd5, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x22,
	0xe6, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64,
	0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x66, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x3c, 0x0a, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x44,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x13, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x74, 0x22, 0x37, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x25, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x22, 0x34, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x28, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xe3, 0x02, 0x0a, 0x08,
	0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x12, 0x46, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x46, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x19, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x17, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_EventService_proto_rawDescOnce sync.Once
	file_EventService_proto_rawDescData = file_EventService_proto_rawDesc
)

func file_EventService_proto_rawDescGZIP() []byte {
	file_EventService_proto_rawDescOnce.Do(func() {
		file_EventService_proto_rawDescData = protoimpl.X.CompressGZIP(file_EventService_proto_rawDescData)
	})
	return file_EventService_proto_rawDescData
}

var file_EventService_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_EventService_proto_goTypes = []interface{}{
	(*Event)(nil),                 // 0: event.Event
	(*UpdateEventAttributes)(nil), // 1: event.UpdateEventAttributes
	(*CreateEventRequest)(nil),    // 2: event.CreateEventRequest
	(*CreateEventResponse)(nil),   // 3: event.CreateEventResponse
	(*UpdateEventRequest)(nil),    // 4: event.UpdateEventRequest
	(*UpdateEventResponse)(nil),   // 5: event.UpdateEventResponse
	(*GetEventsRequest)(nil),      // 6: event.GetEventsRequest
	(*GetEventsResponse)(nil),     // 7: event.GetEventsResponse
	(*GetEventRequest)(nil),       // 8: event.GetEventRequest
	(*GetEventResponse)(nil),      // 9: event.GetEventResponse
	(*DeleteEventRequest)(nil),    // 10: event.DeleteEventRequest
	(*DeleteEventResponse)(nil),   // 11: event.DeleteEventResponse
}
var file_EventService_proto_depIdxs = []int32{
	1,  // 0: event.UpdateEventRequest.attributes:type_name -> event.UpdateEventAttributes
	0,  // 1: event.GetEventsResponse.items:type_name -> event.Event
	0,  // 2: event.GetEventResponse.item:type_name -> event.Event
	2,  // 3: event.Calendar.CreateEvent:input_type -> event.CreateEventRequest
	4,  // 4: event.Calendar.UpdateEvent:input_type -> event.UpdateEventRequest
	6,  // 5: event.Calendar.GetEvents:input_type -> event.GetEventsRequest
	8,  // 6: event.Calendar.GetEvent:input_type -> event.GetEventRequest
	10, // 7: event.Calendar.DeleteEvent:input_type -> event.DeleteEventRequest
	3,  // 8: event.Calendar.CreateEvent:output_type -> event.CreateEventResponse
	5,  // 9: event.Calendar.UpdateEvent:output_type -> event.UpdateEventResponse
	7,  // 10: event.Calendar.GetEvents:output_type -> event.GetEventsResponse
	9,  // 11: event.Calendar.GetEvent:output_type -> event.GetEventResponse
	11, // 12: event.Calendar.DeleteEvent:output_type -> event.DeleteEventResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_EventService_proto_init() }
func file_EventService_proto_init() {
	if File_EventService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_EventService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_EventService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEventAttributes); i {
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
		file_EventService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventRequest); i {
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
		file_EventService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventResponse); i {
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
		file_EventService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEventRequest); i {
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
		file_EventService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEventResponse); i {
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
		file_EventService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsRequest); i {
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
		file_EventService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsResponse); i {
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
		file_EventService_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventRequest); i {
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
		file_EventService_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventResponse); i {
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
		file_EventService_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEventRequest); i {
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
		file_EventService_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEventResponse); i {
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
			RawDescriptor: file_EventService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_EventService_proto_goTypes,
		DependencyIndexes: file_EventService_proto_depIdxs,
		MessageInfos:      file_EventService_proto_msgTypes,
	}.Build()
	File_EventService_proto = out.File
	file_EventService_proto_rawDesc = nil
	file_EventService_proto_goTypes = nil
	file_EventService_proto_depIdxs = nil
}
