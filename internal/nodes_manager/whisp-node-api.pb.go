// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.1
// source: api/whisp-node-api.proto

package nodes_manager

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

type Blank struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Blank) Reset() {
	*x = Blank{}
	mi := &file_api_whisp_node_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Blank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blank) ProtoMessage() {}

func (x *Blank) ProtoReflect() protoreflect.Message {
	mi := &file_api_whisp_node_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blank.ProtoReflect.Descriptor instead.
func (*Blank) Descriptor() ([]byte, []int) {
	return file_api_whisp_node_api_proto_rawDescGZIP(), []int{0}
}

type HostChatRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChatId        string                 `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	OwnerUsername string                 `protobuf:"bytes,2,opt,name=owner_username,json=ownerUsername,proto3" json:"owner_username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HostChatRequest) Reset() {
	*x = HostChatRequest{}
	mi := &file_api_whisp_node_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HostChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostChatRequest) ProtoMessage() {}

func (x *HostChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_whisp_node_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostChatRequest.ProtoReflect.Descriptor instead.
func (*HostChatRequest) Descriptor() ([]byte, []int) {
	return file_api_whisp_node_api_proto_rawDescGZIP(), []int{1}
}

func (x *HostChatRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *HostChatRequest) GetOwnerUsername() string {
	if x != nil {
		return x.OwnerUsername
	}
	return ""
}

type DropChatRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChatId        string                 `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DropChatRequest) Reset() {
	*x = DropChatRequest{}
	mi := &file_api_whisp_node_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DropChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropChatRequest) ProtoMessage() {}

func (x *DropChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_whisp_node_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropChatRequest.ProtoReflect.Descriptor instead.
func (*DropChatRequest) Descriptor() ([]byte, []int) {
	return file_api_whisp_node_api_proto_rawDescGZIP(), []int{2}
}

func (x *DropChatRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

type NodeInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Chats         []*Chat                `protobuf:"bytes,1,rep,name=chats,proto3" json:"chats,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeInfoResponse) Reset() {
	*x = NodeInfoResponse{}
	mi := &file_api_whisp_node_api_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfoResponse) ProtoMessage() {}

func (x *NodeInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_whisp_node_api_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfoResponse.ProtoReflect.Descriptor instead.
func (*NodeInfoResponse) Descriptor() ([]byte, []int) {
	return file_api_whisp_node_api_proto_rawDescGZIP(), []int{3}
}

func (x *NodeInfoResponse) GetChats() []*Chat {
	if x != nil {
		return x.Chats
	}
	return nil
}

type Chat struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NumberOfClients uint32                 `protobuf:"varint,2,opt,name=number_of_clients,json=numberOfClients,proto3" json:"number_of_clients,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Chat) Reset() {
	*x = Chat{}
	mi := &file_api_whisp_node_api_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_api_whisp_node_api_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat.ProtoReflect.Descriptor instead.
func (*Chat) Descriptor() ([]byte, []int) {
	return file_api_whisp_node_api_proto_rawDescGZIP(), []int{4}
}

func (x *Chat) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Chat) GetNumberOfClients() uint32 {
	if x != nil {
		return x.NumberOfClients
	}
	return 0
}

var File_api_whisp_node_api_proto protoreflect.FileDescriptor

var file_api_whisp_node_api_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x68, 0x69, 0x73, 0x70, 0x2d, 0x6e, 0x6f, 0x64, 0x65,
	0x2d, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x42, 0x6c,
	0x61, 0x6e, 0x6b, 0x22, 0x51, 0x0a, 0x0f, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12,
	0x25, 0x0a, 0x0e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x0f, 0x44, 0x72, 0x6f, 0x70, 0x43, 0x68,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74,
	0x49, 0x64, 0x22, 0x2f, 0x0a, 0x10, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x05, 0x63, 0x68,
	0x61, 0x74, 0x73, 0x22, 0x42, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x84, 0x01, 0x0a, 0x09, 0x57, 0x68, 0x69, 0x73,
	0x70, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x68, 0x61,
	0x74, 0x12, 0x10, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x42, 0x6c, 0x61, 0x6e, 0x6b, 0x22, 0x00, 0x12, 0x26, 0x0a,
	0x08, 0x44, 0x72, 0x6f, 0x70, 0x43, 0x68, 0x61, 0x74, 0x12, 0x10, 0x2e, 0x44, 0x72, 0x6f, 0x70,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x42, 0x6c,
	0x61, 0x6e, 0x6b, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x06, 0x2e, 0x42, 0x6c, 0x61, 0x6e, 0x6b, 0x1a, 0x11, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x40,
	0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x75, 0x73,
	0x74, 0x44, 0x65, 0x61, 0x6e, 0x2f, 0x77, 0x68, 0x69, 0x73, 0x70, 0x2d, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_whisp_node_api_proto_rawDescOnce sync.Once
	file_api_whisp_node_api_proto_rawDescData = file_api_whisp_node_api_proto_rawDesc
)

func file_api_whisp_node_api_proto_rawDescGZIP() []byte {
	file_api_whisp_node_api_proto_rawDescOnce.Do(func() {
		file_api_whisp_node_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_whisp_node_api_proto_rawDescData)
	})
	return file_api_whisp_node_api_proto_rawDescData
}

var file_api_whisp_node_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_whisp_node_api_proto_goTypes = []any{
	(*Blank)(nil),            // 0: Blank
	(*HostChatRequest)(nil),  // 1: HostChatRequest
	(*DropChatRequest)(nil),  // 2: DropChatRequest
	(*NodeInfoResponse)(nil), // 3: NodeInfoResponse
	(*Chat)(nil),             // 4: Chat
}
var file_api_whisp_node_api_proto_depIdxs = []int32{
	4, // 0: NodeInfoResponse.chats:type_name -> Chat
	1, // 1: WhispNode.HostChat:input_type -> HostChatRequest
	2, // 2: WhispNode.DropChat:input_type -> DropChatRequest
	0, // 3: WhispNode.NodeInfo:input_type -> Blank
	0, // 4: WhispNode.HostChat:output_type -> Blank
	0, // 5: WhispNode.DropChat:output_type -> Blank
	3, // 6: WhispNode.NodeInfo:output_type -> NodeInfoResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_whisp_node_api_proto_init() }
func file_api_whisp_node_api_proto_init() {
	if File_api_whisp_node_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_whisp_node_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_whisp_node_api_proto_goTypes,
		DependencyIndexes: file_api_whisp_node_api_proto_depIdxs,
		MessageInfos:      file_api_whisp_node_api_proto_msgTypes,
	}.Build()
	File_api_whisp_node_api_proto = out.File
	file_api_whisp_node_api_proto_rawDesc = nil
	file_api_whisp_node_api_proto_goTypes = nil
	file_api_whisp_node_api_proto_depIdxs = nil
}
