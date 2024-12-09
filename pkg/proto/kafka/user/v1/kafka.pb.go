// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: user/v1/kafka.proto

package v1

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

type ChangeType int32

const (
	ChangeType_CHANGE_TYPE_UNSPECIFIED ChangeType = 0
	ChangeType_CHANGE_TYPE_CREATED     ChangeType = 1
	ChangeType_CHANGE_TYPE_DELETED     ChangeType = 2
)

// Enum value maps for ChangeType.
var (
	ChangeType_name = map[int32]string{
		0: "CHANGE_TYPE_UNSPECIFIED",
		1: "CHANGE_TYPE_CREATED",
		2: "CHANGE_TYPE_DELETED",
	}
	ChangeType_value = map[string]int32{
		"CHANGE_TYPE_UNSPECIFIED": 0,
		"CHANGE_TYPE_CREATED":     1,
		"CHANGE_TYPE_DELETED":     2,
	}
)

func (x ChangeType) Enum() *ChangeType {
	p := new(ChangeType)
	*p = x
	return p
}

func (x ChangeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangeType) Descriptor() protoreflect.EnumDescriptor {
	return file_user_v1_kafka_proto_enumTypes[0].Descriptor()
}

func (ChangeType) Type() protoreflect.EnumType {
	return &file_user_v1_kafka_proto_enumTypes[0]
}

func (x ChangeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangeType.Descriptor instead.
func (ChangeType) EnumDescriptor() ([]byte, []int) {
	return file_user_v1_kafka_proto_rawDescGZIP(), []int{0}
}

type ChangedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is unique and sortable user identifier
	Id         string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ChangeType ChangeType `protobuf:"varint,2,opt,name=change_type,json=changeType,proto3,enum=user.v1.ChangeType" json:"change_type,omitempty"`
}

func (x *ChangedEvent) Reset() {
	*x = ChangedEvent{}
	mi := &file_user_v1_kafka_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangedEvent) ProtoMessage() {}

func (x *ChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_kafka_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangedEvent.ProtoReflect.Descriptor instead.
func (*ChangedEvent) Descriptor() ([]byte, []int) {
	return file_user_v1_kafka_proto_rawDescGZIP(), []int{0}
}

func (x *ChangedEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChangedEvent) GetChangeType() ChangeType {
	if x != nil {
		return x.ChangeType
	}
	return ChangeType_CHANGE_TYPE_UNSPECIFIED
}

var File_user_v1_kafka_proto protoreflect.FileDescriptor

var file_user_v1_kafka_proto_rawDesc = []byte{
	0x0a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x54,
	0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x34,
	0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x2a, 0x5b, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x17, 0x0a, 0x13, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43,
	0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x48, 0x41, 0x4e,
	0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10,
	0x02, 0x42, 0x09, 0x5a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_v1_kafka_proto_rawDescOnce sync.Once
	file_user_v1_kafka_proto_rawDescData = file_user_v1_kafka_proto_rawDesc
)

func file_user_v1_kafka_proto_rawDescGZIP() []byte {
	file_user_v1_kafka_proto_rawDescOnce.Do(func() {
		file_user_v1_kafka_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_v1_kafka_proto_rawDescData)
	})
	return file_user_v1_kafka_proto_rawDescData
}

var file_user_v1_kafka_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_v1_kafka_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_user_v1_kafka_proto_goTypes = []any{
	(ChangeType)(0),      // 0: user.v1.ChangeType
	(*ChangedEvent)(nil), // 1: user.v1.ChangedEvent
}
var file_user_v1_kafka_proto_depIdxs = []int32{
	0, // 0: user.v1.ChangedEvent.change_type:type_name -> user.v1.ChangeType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_v1_kafka_proto_init() }
func file_user_v1_kafka_proto_init() {
	if File_user_v1_kafka_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_v1_kafka_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_v1_kafka_proto_goTypes,
		DependencyIndexes: file_user_v1_kafka_proto_depIdxs,
		EnumInfos:         file_user_v1_kafka_proto_enumTypes,
		MessageInfos:      file_user_v1_kafka_proto_msgTypes,
	}.Build()
	File_user_v1_kafka_proto = out.File
	file_user_v1_kafka_proto_rawDesc = nil
	file_user_v1_kafka_proto_goTypes = nil
	file_user_v1_kafka_proto_depIdxs = nil
}
