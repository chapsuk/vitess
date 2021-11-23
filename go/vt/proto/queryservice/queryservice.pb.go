//
//Copyright 2019 The Vitess Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// This file contains the service VtTablet exposes for queries.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: queryservice.proto

package queryservice

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	binlogdata "vitess.io/vitess/go/vt/proto/binlogdata"
	query "vitess.io/vitess/go/vt/proto/query"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_queryservice_proto protoreflect.FileDescriptor

var file_queryservice_proto_rawDesc = []byte{
	0x0a, 0x12, 0x71, 0x75, 0x65, 0x72, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x71, 0x75, 0x65, 0x72, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x10, 0x62, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xb3, 0x11, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3a, 0x0a, 0x07, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0c, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x34, 0x0a, 0x05, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x12, 0x14, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x16, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x6f,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3a, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x12, 0x15, 0x2e, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0e,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x64, 0x12, 0x1c,
	0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x50, 0x72, 0x65,
	0x70, 0x61, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x50, 0x72, 0x65, 0x70, 0x61,
	0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a,
	0x10, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65,
	0x64, 0x12, 0x1e, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46,
	0x0a, 0x0b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x19, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x19, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x65,
	0x74, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e,
	0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x6f,
	0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52,
	0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1d, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x49, 0x0a, 0x0c, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x12, 0x1a, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a,
	0x11, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x1f, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x42, 0x65, 0x67, 0x69,
	0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x12, 0x42, 0x65, 0x67, 0x69, 0x6e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x20, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4e, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1b, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x43, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x41, 0x63, 0x6b, 0x12, 0x18, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0e, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x1c, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x13,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x14,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x72, 0x0a, 0x19, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x42, 0x65, 0x67, 0x69,
	0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x27,
	0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x42, 0x65,
	0x67, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3a, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x12, 0x15, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x12, 0x1a, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x46,
	0x0a, 0x07, 0x56, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1a, 0x2e, 0x62, 0x69, 0x6e, 0x6c,
	0x6f, 0x67, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x56, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x56, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x52, 0x0a, 0x0b, 0x56, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x1e, 0x2e, 0x62, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x56, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x6f, 0x77, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x62, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x56, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x6f, 0x77, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x5b, 0x0a, 0x0e, 0x56, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x21, 0x2e, 0x62,
	0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x56, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x62, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x56, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x2b, 0x5a, 0x29, 0x76, 0x69, 0x74, 0x65, 0x73,
	0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x76,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_queryservice_proto_goTypes = []interface{}{
	(*query.ExecuteRequest)(nil),                    // 0: query.ExecuteRequest
	(*query.ExecuteBatchRequest)(nil),               // 1: query.ExecuteBatchRequest
	(*query.StreamExecuteRequest)(nil),              // 2: query.StreamExecuteRequest
	(*query.BeginRequest)(nil),                      // 3: query.BeginRequest
	(*query.CommitRequest)(nil),                     // 4: query.CommitRequest
	(*query.RollbackRequest)(nil),                   // 5: query.RollbackRequest
	(*query.PrepareRequest)(nil),                    // 6: query.PrepareRequest
	(*query.CommitPreparedRequest)(nil),             // 7: query.CommitPreparedRequest
	(*query.RollbackPreparedRequest)(nil),           // 8: query.RollbackPreparedRequest
	(*query.CreateTransactionRequest)(nil),          // 9: query.CreateTransactionRequest
	(*query.StartCommitRequest)(nil),                // 10: query.StartCommitRequest
	(*query.SetRollbackRequest)(nil),                // 11: query.SetRollbackRequest
	(*query.ConcludeTransactionRequest)(nil),        // 12: query.ConcludeTransactionRequest
	(*query.ReadTransactionRequest)(nil),            // 13: query.ReadTransactionRequest
	(*query.BeginExecuteRequest)(nil),               // 14: query.BeginExecuteRequest
	(*query.BeginExecuteBatchRequest)(nil),          // 15: query.BeginExecuteBatchRequest
	(*query.BeginStreamExecuteRequest)(nil),         // 16: query.BeginStreamExecuteRequest
	(*query.MessageStreamRequest)(nil),              // 17: query.MessageStreamRequest
	(*query.MessageAckRequest)(nil),                 // 18: query.MessageAckRequest
	(*query.ReserveExecuteRequest)(nil),             // 19: query.ReserveExecuteRequest
	(*query.ReserveBeginExecuteRequest)(nil),        // 20: query.ReserveBeginExecuteRequest
	(*query.ReserveStreamExecuteRequest)(nil),       // 21: query.ReserveStreamExecuteRequest
	(*query.ReserveBeginStreamExecuteRequest)(nil),  // 22: query.ReserveBeginStreamExecuteRequest
	(*query.ReleaseRequest)(nil),                    // 23: query.ReleaseRequest
	(*query.StreamHealthRequest)(nil),               // 24: query.StreamHealthRequest
	(*binlogdata.VStreamRequest)(nil),               // 25: binlogdata.VStreamRequest
	(*binlogdata.VStreamRowsRequest)(nil),           // 26: binlogdata.VStreamRowsRequest
	(*binlogdata.VStreamResultsRequest)(nil),        // 27: binlogdata.VStreamResultsRequest
	(*query.ExecuteResponse)(nil),                   // 28: query.ExecuteResponse
	(*query.ExecuteBatchResponse)(nil),              // 29: query.ExecuteBatchResponse
	(*query.StreamExecuteResponse)(nil),             // 30: query.StreamExecuteResponse
	(*query.BeginResponse)(nil),                     // 31: query.BeginResponse
	(*query.CommitResponse)(nil),                    // 32: query.CommitResponse
	(*query.RollbackResponse)(nil),                  // 33: query.RollbackResponse
	(*query.PrepareResponse)(nil),                   // 34: query.PrepareResponse
	(*query.CommitPreparedResponse)(nil),            // 35: query.CommitPreparedResponse
	(*query.RollbackPreparedResponse)(nil),          // 36: query.RollbackPreparedResponse
	(*query.CreateTransactionResponse)(nil),         // 37: query.CreateTransactionResponse
	(*query.StartCommitResponse)(nil),               // 38: query.StartCommitResponse
	(*query.SetRollbackResponse)(nil),               // 39: query.SetRollbackResponse
	(*query.ConcludeTransactionResponse)(nil),       // 40: query.ConcludeTransactionResponse
	(*query.ReadTransactionResponse)(nil),           // 41: query.ReadTransactionResponse
	(*query.BeginExecuteResponse)(nil),              // 42: query.BeginExecuteResponse
	(*query.BeginExecuteBatchResponse)(nil),         // 43: query.BeginExecuteBatchResponse
	(*query.BeginStreamExecuteResponse)(nil),        // 44: query.BeginStreamExecuteResponse
	(*query.MessageStreamResponse)(nil),             // 45: query.MessageStreamResponse
	(*query.MessageAckResponse)(nil),                // 46: query.MessageAckResponse
	(*query.ReserveExecuteResponse)(nil),            // 47: query.ReserveExecuteResponse
	(*query.ReserveBeginExecuteResponse)(nil),       // 48: query.ReserveBeginExecuteResponse
	(*query.ReserveStreamExecuteResponse)(nil),      // 49: query.ReserveStreamExecuteResponse
	(*query.ReserveBeginStreamExecuteResponse)(nil), // 50: query.ReserveBeginStreamExecuteResponse
	(*query.ReleaseResponse)(nil),                   // 51: query.ReleaseResponse
	(*query.StreamHealthResponse)(nil),              // 52: query.StreamHealthResponse
	(*binlogdata.VStreamResponse)(nil),              // 53: binlogdata.VStreamResponse
	(*binlogdata.VStreamRowsResponse)(nil),          // 54: binlogdata.VStreamRowsResponse
	(*binlogdata.VStreamResultsResponse)(nil),       // 55: binlogdata.VStreamResultsResponse
}
var file_queryservice_proto_depIdxs = []int32{
	0,  // 0: queryservice.Query.Execute:input_type -> query.ExecuteRequest
	1,  // 1: queryservice.Query.ExecuteBatch:input_type -> query.ExecuteBatchRequest
	2,  // 2: queryservice.Query.StreamExecute:input_type -> query.StreamExecuteRequest
	3,  // 3: queryservice.Query.Begin:input_type -> query.BeginRequest
	4,  // 4: queryservice.Query.Commit:input_type -> query.CommitRequest
	5,  // 5: queryservice.Query.Rollback:input_type -> query.RollbackRequest
	6,  // 6: queryservice.Query.Prepare:input_type -> query.PrepareRequest
	7,  // 7: queryservice.Query.CommitPrepared:input_type -> query.CommitPreparedRequest
	8,  // 8: queryservice.Query.RollbackPrepared:input_type -> query.RollbackPreparedRequest
	9,  // 9: queryservice.Query.CreateTransaction:input_type -> query.CreateTransactionRequest
	10, // 10: queryservice.Query.StartCommit:input_type -> query.StartCommitRequest
	11, // 11: queryservice.Query.SetRollback:input_type -> query.SetRollbackRequest
	12, // 12: queryservice.Query.ConcludeTransaction:input_type -> query.ConcludeTransactionRequest
	13, // 13: queryservice.Query.ReadTransaction:input_type -> query.ReadTransactionRequest
	14, // 14: queryservice.Query.BeginExecute:input_type -> query.BeginExecuteRequest
	15, // 15: queryservice.Query.BeginExecuteBatch:input_type -> query.BeginExecuteBatchRequest
	16, // 16: queryservice.Query.BeginStreamExecute:input_type -> query.BeginStreamExecuteRequest
	17, // 17: queryservice.Query.MessageStream:input_type -> query.MessageStreamRequest
	18, // 18: queryservice.Query.MessageAck:input_type -> query.MessageAckRequest
	19, // 19: queryservice.Query.ReserveExecute:input_type -> query.ReserveExecuteRequest
	20, // 20: queryservice.Query.ReserveBeginExecute:input_type -> query.ReserveBeginExecuteRequest
	21, // 21: queryservice.Query.ReserveStreamExecute:input_type -> query.ReserveStreamExecuteRequest
	22, // 22: queryservice.Query.ReserveBeginStreamExecute:input_type -> query.ReserveBeginStreamExecuteRequest
	23, // 23: queryservice.Query.Release:input_type -> query.ReleaseRequest
	24, // 24: queryservice.Query.StreamHealth:input_type -> query.StreamHealthRequest
	25, // 25: queryservice.Query.VStream:input_type -> binlogdata.VStreamRequest
	26, // 26: queryservice.Query.VStreamRows:input_type -> binlogdata.VStreamRowsRequest
	27, // 27: queryservice.Query.VStreamResults:input_type -> binlogdata.VStreamResultsRequest
	28, // 28: queryservice.Query.Execute:output_type -> query.ExecuteResponse
	29, // 29: queryservice.Query.ExecuteBatch:output_type -> query.ExecuteBatchResponse
	30, // 30: queryservice.Query.StreamExecute:output_type -> query.StreamExecuteResponse
	31, // 31: queryservice.Query.Begin:output_type -> query.BeginResponse
	32, // 32: queryservice.Query.Commit:output_type -> query.CommitResponse
	33, // 33: queryservice.Query.Rollback:output_type -> query.RollbackResponse
	34, // 34: queryservice.Query.Prepare:output_type -> query.PrepareResponse
	35, // 35: queryservice.Query.CommitPrepared:output_type -> query.CommitPreparedResponse
	36, // 36: queryservice.Query.RollbackPrepared:output_type -> query.RollbackPreparedResponse
	37, // 37: queryservice.Query.CreateTransaction:output_type -> query.CreateTransactionResponse
	38, // 38: queryservice.Query.StartCommit:output_type -> query.StartCommitResponse
	39, // 39: queryservice.Query.SetRollback:output_type -> query.SetRollbackResponse
	40, // 40: queryservice.Query.ConcludeTransaction:output_type -> query.ConcludeTransactionResponse
	41, // 41: queryservice.Query.ReadTransaction:output_type -> query.ReadTransactionResponse
	42, // 42: queryservice.Query.BeginExecute:output_type -> query.BeginExecuteResponse
	43, // 43: queryservice.Query.BeginExecuteBatch:output_type -> query.BeginExecuteBatchResponse
	44, // 44: queryservice.Query.BeginStreamExecute:output_type -> query.BeginStreamExecuteResponse
	45, // 45: queryservice.Query.MessageStream:output_type -> query.MessageStreamResponse
	46, // 46: queryservice.Query.MessageAck:output_type -> query.MessageAckResponse
	47, // 47: queryservice.Query.ReserveExecute:output_type -> query.ReserveExecuteResponse
	48, // 48: queryservice.Query.ReserveBeginExecute:output_type -> query.ReserveBeginExecuteResponse
	49, // 49: queryservice.Query.ReserveStreamExecute:output_type -> query.ReserveStreamExecuteResponse
	50, // 50: queryservice.Query.ReserveBeginStreamExecute:output_type -> query.ReserveBeginStreamExecuteResponse
	51, // 51: queryservice.Query.Release:output_type -> query.ReleaseResponse
	52, // 52: queryservice.Query.StreamHealth:output_type -> query.StreamHealthResponse
	53, // 53: queryservice.Query.VStream:output_type -> binlogdata.VStreamResponse
	54, // 54: queryservice.Query.VStreamRows:output_type -> binlogdata.VStreamRowsResponse
	55, // 55: queryservice.Query.VStreamResults:output_type -> binlogdata.VStreamResultsResponse
	28, // [28:56] is the sub-list for method output_type
	0,  // [0:28] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_queryservice_proto_init() }
func file_queryservice_proto_init() {
	if File_queryservice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_queryservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_queryservice_proto_goTypes,
		DependencyIndexes: file_queryservice_proto_depIdxs,
	}.Build()
	File_queryservice_proto = out.File
	file_queryservice_proto_rawDesc = nil
	file_queryservice_proto_goTypes = nil
	file_queryservice_proto_depIdxs = nil
}
