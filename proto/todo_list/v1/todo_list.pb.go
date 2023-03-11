// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        (unknown)
// source: todo_list/v1/todo_list.proto

package todo_listv1

import (
	_ "github.com/aalindg/boilerplate/google/api"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PRIORITY int32

const (
	PRIORITY_PRIORITY_P0_UNSPECIFIED PRIORITY = 0
	PRIORITY_PRIORITY_P1             PRIORITY = 1
	PRIORITY_PRIORITY_P2             PRIORITY = 2
	PRIORITY_PRIORITY_P3             PRIORITY = 3
)

// Enum value maps for PRIORITY.
var (
	PRIORITY_name = map[int32]string{
		0: "PRIORITY_P0_UNSPECIFIED",
		1: "PRIORITY_P1",
		2: "PRIORITY_P2",
		3: "PRIORITY_P3",
	}
	PRIORITY_value = map[string]int32{
		"PRIORITY_P0_UNSPECIFIED": 0,
		"PRIORITY_P1":             1,
		"PRIORITY_P2":             2,
		"PRIORITY_P3":             3,
	}
)

func (x PRIORITY) Enum() *PRIORITY {
	p := new(PRIORITY)
	*p = x
	return p
}

func (x PRIORITY) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PRIORITY) Descriptor() protoreflect.EnumDescriptor {
	return file_todo_list_v1_todo_list_proto_enumTypes[0].Descriptor()
}

func (PRIORITY) Type() protoreflect.EnumType {
	return &file_todo_list_v1_todo_list_proto_enumTypes[0]
}

func (x PRIORITY) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PRIORITY.Descriptor instead.
func (PRIORITY) EnumDescriptor() ([]byte, []int) {
	return file_todo_list_v1_todo_list_proto_rawDescGZIP(), []int{0}
}

type CreateListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreateListRequest) Reset() {
	*x = CreateListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_list_v1_todo_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateListRequest) ProtoMessage() {}

func (x *CreateListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_list_v1_todo_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateListRequest.ProtoReflect.Descriptor instead.
func (*CreateListRequest) Descriptor() ([]byte, []int) {
	return file_todo_list_v1_todo_list_proto_rawDescGZIP(), []int{0}
}

func (x *CreateListRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateListRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CreateListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List *List `protobuf:"bytes,1,opt,name=list,proto3" json:"list,omitempty"`
}

func (x *CreateListResponse) Reset() {
	*x = CreateListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_list_v1_todo_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateListResponse) ProtoMessage() {}

func (x *CreateListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_list_v1_todo_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateListResponse.ProtoReflect.Descriptor instead.
func (*CreateListResponse) Descriptor() ([]byte, []int) {
	return file_todo_list_v1_todo_list_proto_rawDescGZIP(), []int{1}
}

func (x *CreateListResponse) GetList() *List {
	if x != nil {
		return x.List
	}
	return nil
}

type AddListElementRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListId  string   `protobuf:"bytes,1,opt,name=list_id,json=listId,proto3" json:"list_id,omitempty"`
	Element *Element `protobuf:"bytes,2,opt,name=element,proto3" json:"element,omitempty"`
}

func (x *AddListElementRequest) Reset() {
	*x = AddListElementRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_list_v1_todo_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddListElementRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddListElementRequest) ProtoMessage() {}

func (x *AddListElementRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_list_v1_todo_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddListElementRequest.ProtoReflect.Descriptor instead.
func (*AddListElementRequest) Descriptor() ([]byte, []int) {
	return file_todo_list_v1_todo_list_proto_rawDescGZIP(), []int{2}
}

func (x *AddListElementRequest) GetListId() string {
	if x != nil {
		return x.ListId
	}
	return ""
}

func (x *AddListElementRequest) GetElement() *Element {
	if x != nil {
		return x.Element
	}
	return nil
}

type AddListElementResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List *List `protobuf:"bytes,1,opt,name=list,proto3" json:"list,omitempty"`
}

func (x *AddListElementResponse) Reset() {
	*x = AddListElementResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_list_v1_todo_list_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddListElementResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddListElementResponse) ProtoMessage() {}

func (x *AddListElementResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_list_v1_todo_list_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddListElementResponse.ProtoReflect.Descriptor instead.
func (*AddListElementResponse) Descriptor() ([]byte, []int) {
	return file_todo_list_v1_todo_list_proto_rawDescGZIP(), []int{3}
}

func (x *AddListElementResponse) GetList() *List {
	if x != nil {
		return x.List
	}
	return nil
}

type GetListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListId string `protobuf:"bytes,1,opt,name=list_id,json=listId,proto3" json:"list_id,omitempty"`
}

func (x *GetListRequest) Reset() {
	*x = GetListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_list_v1_todo_list_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListRequest) ProtoMessage() {}

func (x *GetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_list_v1_todo_list_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListRequest.ProtoReflect.Descriptor instead.
func (*GetListRequest) Descriptor() ([]byte, []int) {
	return file_todo_list_v1_todo_list_proto_rawDescGZIP(), []int{4}
}

func (x *GetListRequest) GetListId() string {
	if x != nil {
		return x.ListId
	}
	return ""
}

type GetListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List *List `protobuf:"bytes,1,opt,name=list,proto3" json:"list,omitempty"`
}

func (x *GetListResponse) Reset() {
	*x = GetListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_list_v1_todo_list_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListResponse) ProtoMessage() {}

func (x *GetListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_list_v1_todo_list_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListResponse.ProtoReflect.Descriptor instead.
func (*GetListResponse) Descriptor() ([]byte, []int) {
	return file_todo_list_v1_todo_list_proto_rawDescGZIP(), []int{5}
}

func (x *GetListResponse) GetList() *List {
	if x != nil {
		return x.List
	}
	return nil
}

type FetchAllListsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FetchAllListsRequest) Reset() {
	*x = FetchAllListsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_list_v1_todo_list_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchAllListsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchAllListsRequest) ProtoMessage() {}

func (x *FetchAllListsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_list_v1_todo_list_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchAllListsRequest.ProtoReflect.Descriptor instead.
func (*FetchAllListsRequest) Descriptor() ([]byte, []int) {
	return file_todo_list_v1_todo_list_proto_rawDescGZIP(), []int{6}
}

type FetchAllListsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List *List `protobuf:"bytes,1,opt,name=list,proto3" json:"list,omitempty"`
}

func (x *FetchAllListsResponse) Reset() {
	*x = FetchAllListsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_list_v1_todo_list_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchAllListsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchAllListsResponse) ProtoMessage() {}

func (x *FetchAllListsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_list_v1_todo_list_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchAllListsResponse.ProtoReflect.Descriptor instead.
func (*FetchAllListsResponse) Descriptor() ([]byte, []int) {
	return file_todo_list_v1_todo_list_proto_rawDescGZIP(), []int{7}
}

func (x *FetchAllListsResponse) GetList() *List {
	if x != nil {
		return x.List
	}
	return nil
}

type List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListId    string     `protobuf:"bytes,1,opt,name=list_id,json=listId,proto3" json:"list_id,omitempty"`
	Title     string     `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Elements  []*Element `protobuf:"bytes,3,rep,name=elements,proto3" json:"elements,omitempty"`
	CreatedBy string     `protobuf:"bytes,4,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
}

func (x *List) Reset() {
	*x = List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_list_v1_todo_list_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*List) ProtoMessage() {}

func (x *List) ProtoReflect() protoreflect.Message {
	mi := &file_todo_list_v1_todo_list_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use List.ProtoReflect.Descriptor instead.
func (*List) Descriptor() ([]byte, []int) {
	return file_todo_list_v1_todo_list_proto_rawDescGZIP(), []int{8}
}

func (x *List) GetListId() string {
	if x != nil {
		return x.ListId
	}
	return ""
}

func (x *List) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *List) GetElements() []*Element {
	if x != nil {
		return x.Elements
	}
	return nil
}

func (x *List) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

type Element struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text      string                 `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	IsDone    bool                   `protobuf:"varint,2,opt,name=is_done,json=isDone,proto3" json:"is_done,omitempty"`
	Priority  PRIORITY               `protobuf:"varint,3,opt,name=priority,proto3,enum=todo_list.v1.PRIORITY" json:"priority,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Element) Reset() {
	*x = Element{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_list_v1_todo_list_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Element) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Element) ProtoMessage() {}

func (x *Element) ProtoReflect() protoreflect.Message {
	mi := &file_todo_list_v1_todo_list_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Element.ProtoReflect.Descriptor instead.
func (*Element) Descriptor() ([]byte, []int) {
	return file_todo_list_v1_todo_list_proto_rawDescGZIP(), []int{9}
}

func (x *Element) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Element) GetIsDone() bool {
	if x != nil {
		return x.IsDone
	}
	return false
}

func (x *Element) GetPriority() PRIORITY {
	if x != nil {
		return x.Priority
	}
	return PRIORITY_PRIORITY_P0_UNSPECIFIED
}

func (x *Element) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_todo_list_v1_todo_list_proto protoreflect.FileDescriptor

var file_todo_list_v1_todo_list_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x6f, 0x64, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x3c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x61, 0x0a,
	0x15, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x2f, 0x0a, 0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x22, 0x40, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x22, 0x39, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x41, 0x6c, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x3f, 0x0a, 0x15, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x6c, 0x6c, 0x4c, 0x69, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x22, 0x87, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x69,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x69, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x6f,
	0x64, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0xa5, 0x01, 0x0a, 0x07,
	0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x73, 0x5f, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73,
	0x44, 0x6f, 0x6e, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x52, 0x08,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x2a, 0x5a, 0x0a, 0x08, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x12,
	0x1b, 0x0a, 0x17, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x50, 0x30, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b,
	0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x50, 0x31, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x50, 0x32, 0x10, 0x02, 0x12, 0x0f,
	0x0a, 0x0b, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x50, 0x33, 0x10, 0x03, 0x32,
	0xcf, 0x03, 0x0a, 0x0f, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x1f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22,
	0x0d, 0x2f, 0x76, 0x69, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x12, 0x7d,
	0x0a, 0x0e, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x23, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76, 0x69, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x73, 0x2f, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x64, 0x64, 0x12, 0x63, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f,
	0x76, 0x69, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x2f, 0x7b, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x7d, 0x12, 0x6d, 0x0a, 0x0d, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x6c, 0x6c, 0x4c, 0x69,
	0x73, 0x74, 0x73, 0x12, 0x22, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x6c, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x6c, 0x6c, 0x4c,
	0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x69, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x30,
	0x01, 0x42, 0xa7, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x61, 0x6c, 0x69, 0x6e, 0x64, 0x67, 0x2f, 0x62, 0x6f, 0x69, 0x6c,
	0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x17, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x54,
	0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_todo_list_v1_todo_list_proto_rawDescOnce sync.Once
	file_todo_list_v1_todo_list_proto_rawDescData = file_todo_list_v1_todo_list_proto_rawDesc
)

func file_todo_list_v1_todo_list_proto_rawDescGZIP() []byte {
	file_todo_list_v1_todo_list_proto_rawDescOnce.Do(func() {
		file_todo_list_v1_todo_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_todo_list_v1_todo_list_proto_rawDescData)
	})
	return file_todo_list_v1_todo_list_proto_rawDescData
}

var file_todo_list_v1_todo_list_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_todo_list_v1_todo_list_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_todo_list_v1_todo_list_proto_goTypes = []interface{}{
	(PRIORITY)(0),                  // 0: todo_list.v1.PRIORITY
	(*CreateListRequest)(nil),      // 1: todo_list.v1.CreateListRequest
	(*CreateListResponse)(nil),     // 2: todo_list.v1.CreateListResponse
	(*AddListElementRequest)(nil),  // 3: todo_list.v1.AddListElementRequest
	(*AddListElementResponse)(nil), // 4: todo_list.v1.AddListElementResponse
	(*GetListRequest)(nil),         // 5: todo_list.v1.GetListRequest
	(*GetListResponse)(nil),        // 6: todo_list.v1.GetListResponse
	(*FetchAllListsRequest)(nil),   // 7: todo_list.v1.FetchAllListsRequest
	(*FetchAllListsResponse)(nil),  // 8: todo_list.v1.FetchAllListsResponse
	(*List)(nil),                   // 9: todo_list.v1.List
	(*Element)(nil),                // 10: todo_list.v1.Element
	(*timestamppb.Timestamp)(nil),  // 11: google.protobuf.Timestamp
}
var file_todo_list_v1_todo_list_proto_depIdxs = []int32{
	9,  // 0: todo_list.v1.CreateListResponse.list:type_name -> todo_list.v1.List
	10, // 1: todo_list.v1.AddListElementRequest.element:type_name -> todo_list.v1.Element
	9,  // 2: todo_list.v1.AddListElementResponse.list:type_name -> todo_list.v1.List
	9,  // 3: todo_list.v1.GetListResponse.list:type_name -> todo_list.v1.List
	9,  // 4: todo_list.v1.FetchAllListsResponse.list:type_name -> todo_list.v1.List
	10, // 5: todo_list.v1.List.elements:type_name -> todo_list.v1.Element
	0,  // 6: todo_list.v1.Element.priority:type_name -> todo_list.v1.PRIORITY
	11, // 7: todo_list.v1.Element.created_at:type_name -> google.protobuf.Timestamp
	1,  // 8: todo_list.v1.TodoListService.CreateList:input_type -> todo_list.v1.CreateListRequest
	3,  // 9: todo_list.v1.TodoListService.AddListElement:input_type -> todo_list.v1.AddListElementRequest
	5,  // 10: todo_list.v1.TodoListService.GetList:input_type -> todo_list.v1.GetListRequest
	7,  // 11: todo_list.v1.TodoListService.FetchAllLists:input_type -> todo_list.v1.FetchAllListsRequest
	2,  // 12: todo_list.v1.TodoListService.CreateList:output_type -> todo_list.v1.CreateListResponse
	4,  // 13: todo_list.v1.TodoListService.AddListElement:output_type -> todo_list.v1.AddListElementResponse
	6,  // 14: todo_list.v1.TodoListService.GetList:output_type -> todo_list.v1.GetListResponse
	8,  // 15: todo_list.v1.TodoListService.FetchAllLists:output_type -> todo_list.v1.FetchAllListsResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_todo_list_v1_todo_list_proto_init() }
func file_todo_list_v1_todo_list_proto_init() {
	if File_todo_list_v1_todo_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todo_list_v1_todo_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateListRequest); i {
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
		file_todo_list_v1_todo_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateListResponse); i {
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
		file_todo_list_v1_todo_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddListElementRequest); i {
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
		file_todo_list_v1_todo_list_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddListElementResponse); i {
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
		file_todo_list_v1_todo_list_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListRequest); i {
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
		file_todo_list_v1_todo_list_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListResponse); i {
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
		file_todo_list_v1_todo_list_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchAllListsRequest); i {
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
		file_todo_list_v1_todo_list_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchAllListsResponse); i {
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
		file_todo_list_v1_todo_list_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*List); i {
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
		file_todo_list_v1_todo_list_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Element); i {
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
			RawDescriptor: file_todo_list_v1_todo_list_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_list_v1_todo_list_proto_goTypes,
		DependencyIndexes: file_todo_list_v1_todo_list_proto_depIdxs,
		EnumInfos:         file_todo_list_v1_todo_list_proto_enumTypes,
		MessageInfos:      file_todo_list_v1_todo_list_proto_msgTypes,
	}.Build()
	File_todo_list_v1_todo_list_proto = out.File
	file_todo_list_v1_todo_list_proto_rawDesc = nil
	file_todo_list_v1_todo_list_proto_goTypes = nil
	file_todo_list_v1_todo_list_proto_depIdxs = nil
}
