// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: belifeline/models/v1/provider.proto

package modelsv1

import (
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

// *
// Client Information
// ClientInformation is basic information about Client
type ClientInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId      *ULID                  `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3,oneof" json:"client_id,omitempty"`
	Username      *string                `protobuf:"bytes,2,opt,name=username,proto3,oneof" json:"username,omitempty"`
	ApiKey        *ApiKey                `protobuf:"bytes,3,opt,name=api_key,json=apiKey,proto3,oneof" json:"api_key,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`               // Immutable at the date and time the client is registered in the belifeline
	LastUsedAt    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=last_used_at,json=lastUsedAt,proto3,oneof" json:"last_used_at,omitempty"`          // Time the client last used the belifeline
	LastUpdatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=last_updated_at,json=lastUpdatedAt,proto3,oneof" json:"last_updated_at,omitempty"` // Time the client information last updated
}

func (x *ClientInformation) Reset() {
	*x = ClientInformation{}
	mi := &file_belifeline_models_v1_provider_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInformation) ProtoMessage() {}

func (x *ClientInformation) ProtoReflect() protoreflect.Message {
	mi := &file_belifeline_models_v1_provider_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientInformation.ProtoReflect.Descriptor instead.
func (*ClientInformation) Descriptor() ([]byte, []int) {
	return file_belifeline_models_v1_provider_proto_rawDescGZIP(), []int{0}
}

func (x *ClientInformation) GetClientId() *ULID {
	if x != nil {
		return x.ClientId
	}
	return nil
}

func (x *ClientInformation) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *ClientInformation) GetApiKey() *ApiKey {
	if x != nil {
		return x.ApiKey
	}
	return nil
}

func (x *ClientInformation) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ClientInformation) GetLastUsedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUsedAt
	}
	return nil
}

func (x *ClientInformation) GetLastUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdatedAt
	}
	return nil
}

var File_belifeline_models_v1_provider_proto protoreflect.FileDescriptor

var file_belifeline_models_v1_provider_proto_rawDesc = []byte{
	0x0a, 0x23, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x62, 0x65, 0x6c,
	0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5,
	0x03, 0x0a, 0x11, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x4c, 0x49, 0x44, 0x48, 0x00, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x69, 0x4b,
	0x65, 0x79, 0x48, 0x02, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x12,
	0x3e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48,
	0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x41, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x48, 0x04, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x73, 0x65, 0x64, 0x41, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x47, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x05, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x6b,
	0x65, 0x79, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x6c, 0x63, 0x79, 0x6f, 0x6e, 0x2d, 0x6f, 0x72, 0x67,
	0x2f, 0x6b, 0x69, 0x7a, 0x75, 0x6e, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x65, 0x6c, 0x69,
	0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_belifeline_models_v1_provider_proto_rawDescOnce sync.Once
	file_belifeline_models_v1_provider_proto_rawDescData = file_belifeline_models_v1_provider_proto_rawDesc
)

func file_belifeline_models_v1_provider_proto_rawDescGZIP() []byte {
	file_belifeline_models_v1_provider_proto_rawDescOnce.Do(func() {
		file_belifeline_models_v1_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_belifeline_models_v1_provider_proto_rawDescData)
	})
	return file_belifeline_models_v1_provider_proto_rawDescData
}

var file_belifeline_models_v1_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_belifeline_models_v1_provider_proto_goTypes = []any{
	(*ClientInformation)(nil),     // 0: belifeline.models.v1.ClientInformation
	(*ULID)(nil),                  // 1: belifeline.models.v1.ULID
	(*ApiKey)(nil),                // 2: belifeline.models.v1.ApiKey
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_belifeline_models_v1_provider_proto_depIdxs = []int32{
	1, // 0: belifeline.models.v1.ClientInformation.client_id:type_name -> belifeline.models.v1.ULID
	2, // 1: belifeline.models.v1.ClientInformation.api_key:type_name -> belifeline.models.v1.ApiKey
	3, // 2: belifeline.models.v1.ClientInformation.created_at:type_name -> google.protobuf.Timestamp
	3, // 3: belifeline.models.v1.ClientInformation.last_used_at:type_name -> google.protobuf.Timestamp
	3, // 4: belifeline.models.v1.ClientInformation.last_updated_at:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_belifeline_models_v1_provider_proto_init() }
func file_belifeline_models_v1_provider_proto_init() {
	if File_belifeline_models_v1_provider_proto != nil {
		return
	}
	file_belifeline_models_v1_types_proto_init()
	file_belifeline_models_v1_provider_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_belifeline_models_v1_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_belifeline_models_v1_provider_proto_goTypes,
		DependencyIndexes: file_belifeline_models_v1_provider_proto_depIdxs,
		MessageInfos:      file_belifeline_models_v1_provider_proto_msgTypes,
	}.Build()
	File_belifeline_models_v1_provider_proto = out.File
	file_belifeline_models_v1_provider_proto_rawDesc = nil
	file_belifeline_models_v1_provider_proto_goTypes = nil
	file_belifeline_models_v1_provider_proto_depIdxs = nil
}
