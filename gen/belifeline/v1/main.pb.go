// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: belifeline/v1/main.proto

package mainv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_belifeline_v1_main_proto protoreflect.FileDescriptor

var file_belifeline_v1_main_proto_rawDesc = []byte{
	0x0a, 0x18, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x62, 0x65, 0x6c, 0x69,
	0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x62, 0x65, 0x6c, 0x69, 0x66,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x99, 0x01, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1b, 0x2e,
	0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x65, 0x6c,
	0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x05, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x1b, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x32, 0xef,
	0x06, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x50, 0x0a, 0x09, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x12, 0x1f, 0x2e, 0x62,
	0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x53, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x20, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x62, 0x65, 0x6c,
	0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x59, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x76, 0x6f, 0x6b,
	0x65, 0x12, 0x22, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x76, 0x6f,
	0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x77, 0x0a, 0x16,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x2c, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x80, 0x01, 0x0a, 0x19, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x2f, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0a, 0x4b, 0x6f, 0x79, 0x6f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6f, 0x79, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6f, 0x79, 0x6f, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a,
	0x0a, 0x4b, 0x6f, 0x79, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x62, 0x65,
	0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6f, 0x79, 0x6f,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6f,
	0x79, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0d, 0x4b, 0x6f, 0x79, 0x6f, 0x41, 0x70, 0x69, 0x52, 0x65, 0x76,
	0x6f, 0x6b, 0x65, 0x12, 0x23, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6f, 0x79, 0x6f, 0x41, 0x70, 0x69, 0x52, 0x65, 0x76, 0x6f, 0x6b,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6f, 0x79, 0x6f, 0x41, 0x70, 0x69,
	0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x32, 0xaf, 0x03, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x7a, 0x0a, 0x17, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x2d, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e,
	0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4f, 0x0a, 0x08, 0x4b, 0x6f, 0x79, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x62,
	0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6f, 0x79,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x62,
	0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6f, 0x79,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x77, 0x0a, 0x16, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x12, 0x2c, 0x2e, 0x62, 0x65,
	0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x62, 0x65, 0x6c, 0x69,
	0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0b, 0x4b, 0x6f,
	0x79, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x47, 0x65, 0x74, 0x12, 0x21, 0x2e, 0x62, 0x65, 0x6c, 0x69,
	0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6f, 0x79, 0x6f, 0x44, 0x61,
	0x74, 0x61, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x62,
	0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6f, 0x79,
	0x6f, 0x44, 0x61, 0x74, 0x61, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x32, 0xc3, 0x01, 0x0a, 0x1a, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xa4, 0x01, 0x0a, 0x25, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x2e, 0x62, 0x65,
	0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0xba, 0x01, 0x0a, 0x0b, 0x4b, 0x6f, 0x79,
	0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0a, 0x4b, 0x6f, 0x79, 0x6f,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6f, 0x79, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6f, 0x79, 0x6f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a,
	0x0b, 0x4b, 0x6f, 0x79, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x41, 0x64, 0x64, 0x12, 0x21, 0x2e, 0x62,
	0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6f, 0x79,
	0x6f, 0x44, 0x61, 0x74, 0x61, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x62, 0x65, 0x6c, 0x69, 0x66, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4b, 0x6f, 0x79, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x6c, 0x63, 0x79, 0x6f, 0x6e, 0x2d, 0x6f, 0x72, 0x67, 0x2f,
	0x6b, 0x69, 0x7a, 0x75, 0x6e, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x65, 0x6c, 0x69, 0x66,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x61, 0x69, 0x6e, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_belifeline_v1_main_proto_goTypes = []any{
	(*CheckRequest)(nil),                                  // 0: belifeline.v1.CheckRequest
	(*WatchRequest)(nil),                                  // 1: belifeline.v1.WatchRequest
	(*ClientSetRequest)(nil),                              // 2: belifeline.v1.ClientSetRequest
	(*ClientListRequest)(nil),                             // 3: belifeline.v1.ClientListRequest
	(*ClientDeleteRequest)(nil),                           // 4: belifeline.v1.ClientDeleteRequest
	(*ClientRevokeRequest)(nil),                           // 5: belifeline.v1.ClientRevokeRequest
	(*ExternalInformationSetRequest)(nil),                 // 6: belifeline.v1.ExternalInformationSetRequest
	(*ExternalInformationDeleteRequest)(nil),              // 7: belifeline.v1.ExternalInformationDeleteRequest
	(*KoyoCreateRequest)(nil),                             // 8: belifeline.v1.KoyoCreateRequest
	(*KoyoDeleteRequest)(nil),                             // 9: belifeline.v1.KoyoDeleteRequest
	(*KoyoApiRevokeRequest)(nil),                          // 10: belifeline.v1.KoyoApiRevokeRequest
	(*ExternalInformationListRequest)(nil),                // 11: belifeline.v1.ExternalInformationListRequest
	(*KoyoListRequest)(nil),                               // 12: belifeline.v1.KoyoListRequest
	(*ExternalInformationGetRequest)(nil),                 // 13: belifeline.v1.ExternalInformationGetRequest
	(*KoyoDataGetRequest)(nil),                            // 14: belifeline.v1.KoyoDataGetRequest
	(*ExternalInformationUpdateNotificationRequest)(nil),  // 15: belifeline.v1.ExternalInformationUpdateNotificationRequest
	(*KoyoUpdateRequest)(nil),                             // 16: belifeline.v1.KoyoUpdateRequest
	(*KoyoDataAddRequest)(nil),                            // 17: belifeline.v1.KoyoDataAddRequest
	(*CheckResponse)(nil),                                 // 18: belifeline.v1.CheckResponse
	(*WatchResponse)(nil),                                 // 19: belifeline.v1.WatchResponse
	(*ClientSetResponse)(nil),                             // 20: belifeline.v1.ClientSetResponse
	(*ClientListResponse)(nil),                            // 21: belifeline.v1.ClientListResponse
	(*ClientDeleteResponse)(nil),                          // 22: belifeline.v1.ClientDeleteResponse
	(*ClientRevokeResponse)(nil),                          // 23: belifeline.v1.ClientRevokeResponse
	(*ExternalInformationSetResponse)(nil),                // 24: belifeline.v1.ExternalInformationSetResponse
	(*ExternalInformationDeleteResponse)(nil),             // 25: belifeline.v1.ExternalInformationDeleteResponse
	(*KoyoCreateResponse)(nil),                            // 26: belifeline.v1.KoyoCreateResponse
	(*KoyoDeleteResponse)(nil),                            // 27: belifeline.v1.KoyoDeleteResponse
	(*KoyoApiRevokeResponse)(nil),                         // 28: belifeline.v1.KoyoApiRevokeResponse
	(*ExternalInformationListResponse)(nil),               // 29: belifeline.v1.ExternalInformationListResponse
	(*KoyoListResponse)(nil),                              // 30: belifeline.v1.KoyoListResponse
	(*ExternalInformationGetResponse)(nil),                // 31: belifeline.v1.ExternalInformationGetResponse
	(*KoyoDataGetResponse)(nil),                           // 32: belifeline.v1.KoyoDataGetResponse
	(*ExternalInformationUpdateNotificationResponse)(nil), // 33: belifeline.v1.ExternalInformationUpdateNotificationResponse
	(*KoyoUpdateResponse)(nil),                            // 34: belifeline.v1.KoyoUpdateResponse
	(*KoyoDataAddResponse)(nil),                           // 35: belifeline.v1.KoyoDataAddResponse
}
var file_belifeline_v1_main_proto_depIdxs = []int32{
	0,  // 0: belifeline.v1.HealthService.Check:input_type -> belifeline.v1.CheckRequest
	1,  // 1: belifeline.v1.HealthService.Watch:input_type -> belifeline.v1.WatchRequest
	2,  // 2: belifeline.v1.AdminService.ClientSet:input_type -> belifeline.v1.ClientSetRequest
	3,  // 3: belifeline.v1.AdminService.ClientList:input_type -> belifeline.v1.ClientListRequest
	4,  // 4: belifeline.v1.AdminService.ClientDelete:input_type -> belifeline.v1.ClientDeleteRequest
	5,  // 5: belifeline.v1.AdminService.ClientRevoke:input_type -> belifeline.v1.ClientRevokeRequest
	6,  // 6: belifeline.v1.AdminService.ExternalInformationSet:input_type -> belifeline.v1.ExternalInformationSetRequest
	7,  // 7: belifeline.v1.AdminService.ExternalInformationDelete:input_type -> belifeline.v1.ExternalInformationDeleteRequest
	8,  // 8: belifeline.v1.AdminService.KoyoCreate:input_type -> belifeline.v1.KoyoCreateRequest
	9,  // 9: belifeline.v1.AdminService.KoyoDelete:input_type -> belifeline.v1.KoyoDeleteRequest
	10, // 10: belifeline.v1.AdminService.KoyoApiRevoke:input_type -> belifeline.v1.KoyoApiRevokeRequest
	11, // 11: belifeline.v1.ProviderService.ExternalInformationList:input_type -> belifeline.v1.ExternalInformationListRequest
	12, // 12: belifeline.v1.ProviderService.KoyoList:input_type -> belifeline.v1.KoyoListRequest
	13, // 13: belifeline.v1.ProviderService.ExternalInformationGet:input_type -> belifeline.v1.ExternalInformationGetRequest
	14, // 14: belifeline.v1.ProviderService.KoyoDataGet:input_type -> belifeline.v1.KoyoDataGetRequest
	15, // 15: belifeline.v1.ExternalInformationService.ExternalInformationUpdateNotification:input_type -> belifeline.v1.ExternalInformationUpdateNotificationRequest
	16, // 16: belifeline.v1.KoyoService.KoyoUpdate:input_type -> belifeline.v1.KoyoUpdateRequest
	17, // 17: belifeline.v1.KoyoService.KoyoDataAdd:input_type -> belifeline.v1.KoyoDataAddRequest
	18, // 18: belifeline.v1.HealthService.Check:output_type -> belifeline.v1.CheckResponse
	19, // 19: belifeline.v1.HealthService.Watch:output_type -> belifeline.v1.WatchResponse
	20, // 20: belifeline.v1.AdminService.ClientSet:output_type -> belifeline.v1.ClientSetResponse
	21, // 21: belifeline.v1.AdminService.ClientList:output_type -> belifeline.v1.ClientListResponse
	22, // 22: belifeline.v1.AdminService.ClientDelete:output_type -> belifeline.v1.ClientDeleteResponse
	23, // 23: belifeline.v1.AdminService.ClientRevoke:output_type -> belifeline.v1.ClientRevokeResponse
	24, // 24: belifeline.v1.AdminService.ExternalInformationSet:output_type -> belifeline.v1.ExternalInformationSetResponse
	25, // 25: belifeline.v1.AdminService.ExternalInformationDelete:output_type -> belifeline.v1.ExternalInformationDeleteResponse
	26, // 26: belifeline.v1.AdminService.KoyoCreate:output_type -> belifeline.v1.KoyoCreateResponse
	27, // 27: belifeline.v1.AdminService.KoyoDelete:output_type -> belifeline.v1.KoyoDeleteResponse
	28, // 28: belifeline.v1.AdminService.KoyoApiRevoke:output_type -> belifeline.v1.KoyoApiRevokeResponse
	29, // 29: belifeline.v1.ProviderService.ExternalInformationList:output_type -> belifeline.v1.ExternalInformationListResponse
	30, // 30: belifeline.v1.ProviderService.KoyoList:output_type -> belifeline.v1.KoyoListResponse
	31, // 31: belifeline.v1.ProviderService.ExternalInformationGet:output_type -> belifeline.v1.ExternalInformationGetResponse
	32, // 32: belifeline.v1.ProviderService.KoyoDataGet:output_type -> belifeline.v1.KoyoDataGetResponse
	33, // 33: belifeline.v1.ExternalInformationService.ExternalInformationUpdateNotification:output_type -> belifeline.v1.ExternalInformationUpdateNotificationResponse
	34, // 34: belifeline.v1.KoyoService.KoyoUpdate:output_type -> belifeline.v1.KoyoUpdateResponse
	35, // 35: belifeline.v1.KoyoService.KoyoDataAdd:output_type -> belifeline.v1.KoyoDataAddResponse
	18, // [18:36] is the sub-list for method output_type
	0,  // [0:18] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_belifeline_v1_main_proto_init() }
func file_belifeline_v1_main_proto_init() {
	if File_belifeline_v1_main_proto != nil {
		return
	}
	file_belifeline_v1_api_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_belifeline_v1_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   5,
		},
		GoTypes:           file_belifeline_v1_main_proto_goTypes,
		DependencyIndexes: file_belifeline_v1_main_proto_depIdxs,
	}.Build()
	File_belifeline_v1_main_proto = out.File
	file_belifeline_v1_main_proto_rawDesc = nil
	file_belifeline_v1_main_proto_goTypes = nil
	file_belifeline_v1_main_proto_depIdxs = nil
}
