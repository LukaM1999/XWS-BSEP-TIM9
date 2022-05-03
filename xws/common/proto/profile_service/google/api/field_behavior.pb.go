// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: common/proto/profile_service/google/api/field_behavior.proto

package annotations

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// An indicator of the behavior of a given field (for example, that a field
// is required in requests, or given as output but ignored as input).
// This **does not** change the behavior in protocol buffers itself; it only
// denotes the behavior and may affect how API tooling handles the field.
//
// Note: This enum **may** receive new values in the future.
type FieldBehavior int32

const (
	// Conventional default for enums. Do not use this.
	FieldBehavior_FIELD_BEHAVIOR_UNSPECIFIED FieldBehavior = 0
	// Specifically denotes a field as optional.
	// While all fields in protocol buffers are optional, this may be specified
	// for emphasis if appropriate.
	FieldBehavior_OPTIONAL FieldBehavior = 1
	// Denotes a field as required.
	// This indicates that the field **must** be provided as part of the request,
	// and failure to do so will cause an error (usually `INVALID_ARGUMENT`).
	FieldBehavior_REQUIRED FieldBehavior = 2
	// Denotes a field as output only.
	// This indicates that the field is provided in responses, but including the
	// field in a request does nothing (the server *must* ignore it and
	// *must not* throw an error as a result of the field's presence).
	FieldBehavior_OUTPUT_ONLY FieldBehavior = 3
	// Denotes a field as input only.
	// This indicates that the field is provided in requests, and the
	// corresponding field is not included in output.
	FieldBehavior_INPUT_ONLY FieldBehavior = 4
	// Denotes a field as immutable.
	// This indicates that the field may be set once in a request to create a
	// resource, but may not be changed thereafter.
	FieldBehavior_IMMUTABLE FieldBehavior = 5
	// Denotes that a (repeated) field is an unordered list.
	// This indicates that the service may provide the elements of the list
	// in any arbitrary  order, rather than the order the user originally
	// provided. Additionally, the list's order may or may not be stable.
	FieldBehavior_UNORDERED_LIST FieldBehavior = 6
	// Denotes that this field returns a non-empty default value if not set.
	// This indicates that if the user provides the empty value in a request,
	// a non-empty value will be returned. The user will not be aware of what
	// non-empty value to expect.
	FieldBehavior_NON_EMPTY_DEFAULT FieldBehavior = 7
)

// Enum value maps for FieldBehavior.
var (
	FieldBehavior_name = map[int32]string{
		0: "FIELD_BEHAVIOR_UNSPECIFIED",
		1: "OPTIONAL",
		2: "REQUIRED",
		3: "OUTPUT_ONLY",
		4: "INPUT_ONLY",
		5: "IMMUTABLE",
		6: "UNORDERED_LIST",
		7: "NON_EMPTY_DEFAULT",
	}
	FieldBehavior_value = map[string]int32{
		"FIELD_BEHAVIOR_UNSPECIFIED": 0,
		"OPTIONAL":                   1,
		"REQUIRED":                   2,
		"OUTPUT_ONLY":                3,
		"INPUT_ONLY":                 4,
		"IMMUTABLE":                  5,
		"UNORDERED_LIST":             6,
		"NON_EMPTY_DEFAULT":          7,
	}
)

func (x FieldBehavior) Enum() *FieldBehavior {
	p := new(FieldBehavior)
	*p = x
	return p
}

func (x FieldBehavior) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FieldBehavior) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_profile_service_google_api_field_behavior_proto_enumTypes[0].Descriptor()
}

func (FieldBehavior) Type() protoreflect.EnumType {
	return &file_common_proto_profile_service_google_api_field_behavior_proto_enumTypes[0]
}

func (x FieldBehavior) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FieldBehavior.Descriptor instead.
func (FieldBehavior) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_profile_service_google_api_field_behavior_proto_rawDescGZIP(), []int{0}
}

var file_common_proto_profile_service_google_api_field_behavior_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: ([]FieldBehavior)(nil),
		Field:         1052,
		Name:          "google.api.field_behavior",
		Tag:           "varint,1052,rep,name=field_behavior,enum=google.api.FieldBehavior",
		Filename:      "common/proto/profile_service/google/api/field_behavior.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// A designation of a specific field behavior (required, output only, etc.)
	// in protobuf messages.
	//
	// Examples:
	//
	//   string name = 1 [(google.api.field_behavior) = REQUIRED];
	//   State state = 1 [(google.api.field_behavior) = OUTPUT_ONLY];
	//   google.protobuf.Duration ttl = 1
	//     [(google.api.field_behavior) = INPUT_ONLY];
	//   google.protobuf.Timestamp expire_time = 1
	//     [(google.api.field_behavior) = OUTPUT_ONLY,
	//      (google.api.field_behavior) = IMMUTABLE];
	//
	// repeated google.api.FieldBehavior field_behavior = 1052;
	E_FieldBehavior = &file_common_proto_profile_service_google_api_field_behavior_proto_extTypes[0]
)

var File_common_proto_profile_service_google_api_field_behavior_proto protoreflect.FileDescriptor

var file_common_proto_profile_service_google_api_field_behavior_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xa6, 0x01, 0x0a,
	0x0d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x12, 0x1e,
	0x0a, 0x1a, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x42, 0x45, 0x48, 0x41, 0x56, 0x49, 0x4f, 0x52,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x55,
	0x54, 0x50, 0x55, 0x54, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x49,
	0x4e, 0x50, 0x55, 0x54, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x49,
	0x4d, 0x4d, 0x55, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e,
	0x4f, 0x52, 0x44, 0x45, 0x52, 0x45, 0x44, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x06, 0x12, 0x15,
	0x0a, 0x11, 0x4e, 0x4f, 0x4e, 0x5f, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x5f, 0x44, 0x45, 0x46, 0x41,
	0x55, 0x4c, 0x54, 0x10, 0x07, 0x3a, 0x60, 0x0a, 0x0e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9c, 0x08, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x52, 0x0d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x42, 0x70, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x12, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0xa2, 0x02, 0x04, 0x47, 0x41, 0x50, 0x49, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_common_proto_profile_service_google_api_field_behavior_proto_rawDescOnce sync.Once
	file_common_proto_profile_service_google_api_field_behavior_proto_rawDescData = file_common_proto_profile_service_google_api_field_behavior_proto_rawDesc
)

func file_common_proto_profile_service_google_api_field_behavior_proto_rawDescGZIP() []byte {
	file_common_proto_profile_service_google_api_field_behavior_proto_rawDescOnce.Do(func() {
		file_common_proto_profile_service_google_api_field_behavior_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_profile_service_google_api_field_behavior_proto_rawDescData)
	})
	return file_common_proto_profile_service_google_api_field_behavior_proto_rawDescData
}

var file_common_proto_profile_service_google_api_field_behavior_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_proto_profile_service_google_api_field_behavior_proto_goTypes = []interface{}{
	(FieldBehavior)(0),                // 0: google.api.FieldBehavior
	(*descriptorpb.FieldOptions)(nil), // 1: google.protobuf.FieldOptions
}
var file_common_proto_profile_service_google_api_field_behavior_proto_depIdxs = []int32{
	1, // 0: google.api.field_behavior:extendee -> google.protobuf.FieldOptions
	0, // 1: google.api.field_behavior:type_name -> google.api.FieldBehavior
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_proto_profile_service_google_api_field_behavior_proto_init() }
func file_common_proto_profile_service_google_api_field_behavior_proto_init() {
	if File_common_proto_profile_service_google_api_field_behavior_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proto_profile_service_google_api_field_behavior_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_profile_service_google_api_field_behavior_proto_goTypes,
		DependencyIndexes: file_common_proto_profile_service_google_api_field_behavior_proto_depIdxs,
		EnumInfos:         file_common_proto_profile_service_google_api_field_behavior_proto_enumTypes,
		ExtensionInfos:    file_common_proto_profile_service_google_api_field_behavior_proto_extTypes,
	}.Build()
	File_common_proto_profile_service_google_api_field_behavior_proto = out.File
	file_common_proto_profile_service_google_api_field_behavior_proto_rawDesc = nil
	file_common_proto_profile_service_google_api_field_behavior_proto_goTypes = nil
	file_common_proto_profile_service_google_api_field_behavior_proto_depIdxs = nil
}
