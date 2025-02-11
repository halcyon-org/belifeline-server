// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: belifeline/models/v1/koyo.proto

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
// Koyos Information
// KoyoInformation is basic information about Koyo
type KoyoInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KoyoId          *ULID                  `protobuf:"bytes,1,opt,name=koyo_id,json=koyoId,proto3,oneof" json:"koyo_id,omitempty"`
	KoyoName        *string                `protobuf:"bytes,2,opt,name=koyo_name,json=koyoName,proto3,oneof" json:"koyo_name,omitempty"`
	KoyoDescription *string                `protobuf:"bytes,3,opt,name=koyo_description,json=koyoDescription,proto3,oneof" json:"koyo_description,omitempty"`                                                                    // Koyo's developers explain what this Koyo is all about
	NeedExternal    []*ULID                `protobuf:"bytes,4,rep,name=need_external,json=needExternal,proto3" json:"need_external,omitempty"`                                                                                   // List of ExternalInformation required for Koyo to work
	KoyoParams      map[string]string      `protobuf:"bytes,5,rep,name=koyo_params,json=koyoParams,proto3" json:"koyo_params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Koyo parameters and their default values.
	KoyoScales      []float32              `protobuf:"fixed32,6,rep,packed,name=koyo_scales,json=koyoScales,proto3" json:"koyo_scales,omitempty"`                                                                                // List of Koyo scale (data resolution) (Area where data exists)
	KoyoDataIds     []*ULID                `protobuf:"bytes,7,rep,name=koyo_data_ids,json=koyoDataIds,proto3" json:"koyo_data_ids,omitempty"`                                                                                    // Koyo's data list
	Version         *Version               `protobuf:"bytes,8,opt,name=version,proto3,oneof" json:"version,omitempty"`
	License         *string                `protobuf:"bytes,10,opt,name=license,proto3,oneof" json:"license,omitempty"`                                                       // Koyo's license type
	ExtLicenses     []string               `protobuf:"bytes,11,rep,name=ext_licenses,json=extLicenses,proto3" json:"ext_licenses,omitempty"`                                  // List of licenses on which Koyo depends
	DataType        *DataType              `protobuf:"varint,12,opt,name=data_type,json=dataType,proto3,enum=belifeline.models.v1.DataType,oneof" json:"data_type,omitempty"` // Format of data returned by Koyo (in KoyoData)
	ApiKey          *APIKey                `protobuf:"bytes,13,opt,name=api_key,json=apiKey,proto3,oneof" json:"api_key,omitempty"`
	FirstEntryAt    *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=first_entry_at,json=firstEntryAt,proto3,oneof" json:"first_entry_at,omitempty"`    // Immutable at the date and time Koyo is registered in the belifeline
	LastEntryAt     *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=last_entry_at,json=lastEntryAt,proto3,oneof" json:"last_entry_at,omitempty"`       // Time koyo last added data
	LastUpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=last_updated_at,json=lastUpdatedAt,proto3,oneof" json:"last_updated_at,omitempty"` // Time koyo information last updated
}

func (x *KoyoInformation) Reset() {
	*x = KoyoInformation{}
	mi := &file_belifeline_models_v1_koyo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KoyoInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KoyoInformation) ProtoMessage() {}

func (x *KoyoInformation) ProtoReflect() protoreflect.Message {
	mi := &file_belifeline_models_v1_koyo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KoyoInformation.ProtoReflect.Descriptor instead.
func (*KoyoInformation) Descriptor() ([]byte, []int) {
	return file_belifeline_models_v1_koyo_proto_rawDescGZIP(), []int{0}
}

func (x *KoyoInformation) GetKoyoId() *ULID {
	if x != nil {
		return x.KoyoId
	}
	return nil
}

func (x *KoyoInformation) GetKoyoName() string {
	if x != nil && x.KoyoName != nil {
		return *x.KoyoName
	}
	return ""
}

func (x *KoyoInformation) GetKoyoDescription() string {
	if x != nil && x.KoyoDescription != nil {
		return *x.KoyoDescription
	}
	return ""
}

func (x *KoyoInformation) GetNeedExternal() []*ULID {
	if x != nil {
		return x.NeedExternal
	}
	return nil
}

func (x *KoyoInformation) GetKoyoParams() map[string]string {
	if x != nil {
		return x.KoyoParams
	}
	return nil
}

func (x *KoyoInformation) GetKoyoScales() []float32 {
	if x != nil {
		return x.KoyoScales
	}
	return nil
}

func (x *KoyoInformation) GetKoyoDataIds() []*ULID {
	if x != nil {
		return x.KoyoDataIds
	}
	return nil
}

func (x *KoyoInformation) GetVersion() *Version {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *KoyoInformation) GetLicense() string {
	if x != nil && x.License != nil {
		return *x.License
	}
	return ""
}

func (x *KoyoInformation) GetExtLicenses() []string {
	if x != nil {
		return x.ExtLicenses
	}
	return nil
}

func (x *KoyoInformation) GetDataType() DataType {
	if x != nil && x.DataType != nil {
		return *x.DataType
	}
	return DataType_DATA_TYPE_UNSPECIFIED
}

func (x *KoyoInformation) GetApiKey() *APIKey {
	if x != nil {
		return x.ApiKey
	}
	return nil
}

func (x *KoyoInformation) GetFirstEntryAt() *timestamppb.Timestamp {
	if x != nil {
		return x.FirstEntryAt
	}
	return nil
}

func (x *KoyoInformation) GetLastEntryAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastEntryAt
	}
	return nil
}

func (x *KoyoInformation) GetLastUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdatedAt
	}
	return nil
}

// *
// Koyo Data
// KoyoData is the data returned by Koyo. That is, data generated (calculated) by Koyo's algorithm.
type KoyoData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KoyoDataId     *ULID                  `protobuf:"bytes,1,opt,name=koyo_data_id,json=koyoDataId,proto3,oneof" json:"koyo_data_id,omitempty"`
	KoyoId         *ULID                  `protobuf:"bytes,2,opt,name=koyo_id,json=koyoId,proto3,oneof" json:"koyo_id,omitempty"`
	KoyoScale      *float32               `protobuf:"fixed32,3,opt,name=koyo_scale,json=koyoScale,proto3,oneof" json:"koyo_scale,omitempty"`                                                                                                  // The scale used when generating data
	KoyoDataParams map[string]string      `protobuf:"bytes,4,rep,name=koyo_data_params,json=koyoDataParams,proto3" json:"koyo_data_params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Parameters used when generating data
	Content        []*Content             `protobuf:"bytes,5,rep,name=content,proto3" json:"content,omitempty"`                                                                                                                               // Data content
	EntryAt        *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=entry_at,json=entryAt,proto3,oneof" json:"entry_at,omitempty"`                                                                                                          // Time data was generated
	TargetAt       *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=target_at,json=targetAt,proto3,oneof" json:"target_at,omitempty"`                                                                                                       // Time data is targeted (external information daat time)
}

func (x *KoyoData) Reset() {
	*x = KoyoData{}
	mi := &file_belifeline_models_v1_koyo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KoyoData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KoyoData) ProtoMessage() {}

func (x *KoyoData) ProtoReflect() protoreflect.Message {
	mi := &file_belifeline_models_v1_koyo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KoyoData.ProtoReflect.Descriptor instead.
func (*KoyoData) Descriptor() ([]byte, []int) {
	return file_belifeline_models_v1_koyo_proto_rawDescGZIP(), []int{1}
}

func (x *KoyoData) GetKoyoDataId() *ULID {
	if x != nil {
		return x.KoyoDataId
	}
	return nil
}

func (x *KoyoData) GetKoyoId() *ULID {
	if x != nil {
		return x.KoyoId
	}
	return nil
}

func (x *KoyoData) GetKoyoScale() float32 {
	if x != nil && x.KoyoScale != nil {
		return *x.KoyoScale
	}
	return 0
}

func (x *KoyoData) GetKoyoDataParams() map[string]string {
	if x != nil {
		return x.KoyoDataParams
	}
	return nil
}

func (x *KoyoData) GetContent() []*Content {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *KoyoData) GetEntryAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EntryAt
	}
	return nil
}

func (x *KoyoData) GetTargetAt() *timestamppb.Timestamp {
	if x != nil {
		return x.TargetAt
	}
	return nil
}

var File_belifeline_models_v1_koyo_proto protoreflect.FileDescriptor

var file_belifeline_models_v1_koyo_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x6f, 0x79, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x08, 0x0a, 0x0f, 0x4b,
	0x6f, 0x79, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38,
	0x0a, 0x07, 0x6b, 0x6f, 0x79, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x4c, 0x49, 0x44, 0x48, 0x00, 0x52, 0x06, 0x6b,
	0x6f, 0x79, 0x6f, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x6b, 0x6f, 0x79, 0x6f,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x6b,
	0x6f, 0x79, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x10, 0x6b, 0x6f,
	0x79, 0x6f, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0f, 0x6b, 0x6f, 0x79, 0x6f, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x0d, 0x6e, 0x65,
	0x65, 0x64, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x4c, 0x49, 0x44, 0x52, 0x0c, 0x6e,
	0x65, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x56, 0x0a, 0x0b, 0x6b,
	0x6f, 0x79, 0x6f, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x35, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6f, 0x79, 0x6f, 0x49, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4b, 0x6f, 0x79, 0x6f, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x6b, 0x6f, 0x79, 0x6f, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6b, 0x6f, 0x79, 0x6f, 0x5f, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x02, 0x52, 0x0a, 0x6b, 0x6f, 0x79, 0x6f, 0x53, 0x63,
	0x61, 0x6c, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x0d, 0x6b, 0x6f, 0x79, 0x6f, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x65,
	0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x4c, 0x49, 0x44, 0x52, 0x0b, 0x6b, 0x6f, 0x79, 0x6f, 0x44, 0x61, 0x74,
	0x61, 0x49, 0x64, 0x73, 0x12, 0x3c, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x48, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88,
	0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x74, 0x5f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x74, 0x4c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x48, 0x05, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x50, 0x49, 0x4b, 0x65, 0x79, 0x48, 0x06, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x88,
	0x01, 0x01, 0x12, 0x45, 0x0a, 0x0e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x07, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x43, 0x0a, 0x0d, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x08, 0x52, 0x0b,
	0x6c, 0x61, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x47,
	0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x48, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x1a, 0x3d, 0x0a, 0x0f, 0x4b, 0x6f, 0x79, 0x6f, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6b, 0x6f, 0x79, 0x6f, 0x5f,
	0x69, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6b, 0x6f, 0x79, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x42, 0x13, 0x0a, 0x11, 0x5f, 0x6b, 0x6f, 0x79, 0x6f, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x61, 0x74, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x61, 0x74, 0x42, 0x12, 0x0a, 0x10,
	0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x22, 0xc6, 0x04, 0x0a, 0x08, 0x4b, 0x6f, 0x79, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x12, 0x41, 0x0a,
	0x0c, 0x6b, 0x6f, 0x79, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x4c, 0x49, 0x44, 0x48,
	0x00, 0x52, 0x0a, 0x6b, 0x6f, 0x79, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x38, 0x0a, 0x07, 0x6b, 0x6f, 0x79, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x4c, 0x49, 0x44, 0x48, 0x01, 0x52,
	0x06, 0x6b, 0x6f, 0x79, 0x6f, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x6b, 0x6f,
	0x79, 0x6f, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x48, 0x02,
	0x52, 0x09, 0x6b, 0x6f, 0x79, 0x6f, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x5c,
	0x0a, 0x10, 0x6b, 0x6f, 0x79, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4b, 0x6f, 0x79, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4b, 0x6f, 0x79, 0x6f, 0x44, 0x61, 0x74,
	0x61, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x6b, 0x6f,
	0x79, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x37, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x48, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x41, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x3c, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x48, 0x04, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x74, 0x88, 0x01, 0x01, 0x1a,
	0x41, 0x0a, 0x13, 0x4b, 0x6f, 0x79, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6b, 0x6f, 0x79, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6b, 0x6f, 0x79, 0x6f, 0x5f, 0x69, 0x64, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x6b, 0x6f, 0x79, 0x6f, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x61, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x74, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x6c, 0x63, 0x79, 0x6f, 0x6e, 0x2d,
	0x6f, 0x72, 0x67, 0x2f, 0x6b, 0x69, 0x7a, 0x75, 0x6e, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62,
	0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_belifeline_models_v1_koyo_proto_rawDescOnce sync.Once
	file_belifeline_models_v1_koyo_proto_rawDescData = file_belifeline_models_v1_koyo_proto_rawDesc
)

func file_belifeline_models_v1_koyo_proto_rawDescGZIP() []byte {
	file_belifeline_models_v1_koyo_proto_rawDescOnce.Do(func() {
		file_belifeline_models_v1_koyo_proto_rawDescData = protoimpl.X.CompressGZIP(file_belifeline_models_v1_koyo_proto_rawDescData)
	})
	return file_belifeline_models_v1_koyo_proto_rawDescData
}

var file_belifeline_models_v1_koyo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_belifeline_models_v1_koyo_proto_goTypes = []any{
	(*KoyoInformation)(nil),       // 0: belifeline.models.v1.KoyoInformation
	(*KoyoData)(nil),              // 1: belifeline.models.v1.KoyoData
	nil,                           // 2: belifeline.models.v1.KoyoInformation.KoyoParamsEntry
	nil,                           // 3: belifeline.models.v1.KoyoData.KoyoDataParamsEntry
	(*ULID)(nil),                  // 4: belifeline.models.v1.ULID
	(*Version)(nil),               // 5: belifeline.models.v1.Version
	(DataType)(0),                 // 6: belifeline.models.v1.DataType
	(*APIKey)(nil),                // 7: belifeline.models.v1.APIKey
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*Content)(nil),               // 9: belifeline.models.v1.Content
}
var file_belifeline_models_v1_koyo_proto_depIdxs = []int32{
	4,  // 0: belifeline.models.v1.KoyoInformation.koyo_id:type_name -> belifeline.models.v1.ULID
	4,  // 1: belifeline.models.v1.KoyoInformation.need_external:type_name -> belifeline.models.v1.ULID
	2,  // 2: belifeline.models.v1.KoyoInformation.koyo_params:type_name -> belifeline.models.v1.KoyoInformation.KoyoParamsEntry
	4,  // 3: belifeline.models.v1.KoyoInformation.koyo_data_ids:type_name -> belifeline.models.v1.ULID
	5,  // 4: belifeline.models.v1.KoyoInformation.version:type_name -> belifeline.models.v1.Version
	6,  // 5: belifeline.models.v1.KoyoInformation.data_type:type_name -> belifeline.models.v1.DataType
	7,  // 6: belifeline.models.v1.KoyoInformation.api_key:type_name -> belifeline.models.v1.APIKey
	8,  // 7: belifeline.models.v1.KoyoInformation.first_entry_at:type_name -> google.protobuf.Timestamp
	8,  // 8: belifeline.models.v1.KoyoInformation.last_entry_at:type_name -> google.protobuf.Timestamp
	8,  // 9: belifeline.models.v1.KoyoInformation.last_updated_at:type_name -> google.protobuf.Timestamp
	4,  // 10: belifeline.models.v1.KoyoData.koyo_data_id:type_name -> belifeline.models.v1.ULID
	4,  // 11: belifeline.models.v1.KoyoData.koyo_id:type_name -> belifeline.models.v1.ULID
	3,  // 12: belifeline.models.v1.KoyoData.koyo_data_params:type_name -> belifeline.models.v1.KoyoData.KoyoDataParamsEntry
	9,  // 13: belifeline.models.v1.KoyoData.content:type_name -> belifeline.models.v1.Content
	8,  // 14: belifeline.models.v1.KoyoData.entry_at:type_name -> google.protobuf.Timestamp
	8,  // 15: belifeline.models.v1.KoyoData.target_at:type_name -> google.protobuf.Timestamp
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_belifeline_models_v1_koyo_proto_init() }
func file_belifeline_models_v1_koyo_proto_init() {
	if File_belifeline_models_v1_koyo_proto != nil {
		return
	}
	file_belifeline_models_v1_types_proto_init()
	file_belifeline_models_v1_koyo_proto_msgTypes[0].OneofWrappers = []any{}
	file_belifeline_models_v1_koyo_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_belifeline_models_v1_koyo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_belifeline_models_v1_koyo_proto_goTypes,
		DependencyIndexes: file_belifeline_models_v1_koyo_proto_depIdxs,
		MessageInfos:      file_belifeline_models_v1_koyo_proto_msgTypes,
	}.Build()
	File_belifeline_models_v1_koyo_proto = out.File
	file_belifeline_models_v1_koyo_proto_rawDesc = nil
	file_belifeline_models_v1_koyo_proto_goTypes = nil
	file_belifeline_models_v1_koyo_proto_depIdxs = nil
}
