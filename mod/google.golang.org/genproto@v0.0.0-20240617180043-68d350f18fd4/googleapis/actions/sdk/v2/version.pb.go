// Copyright 2024 Google LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.4
// source: google/actions/sdk/v2/version.proto

package sdk

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Enum indicating the states that a Version can take. This enum is not yet
// frozen and values maybe added later.
type Version_VersionState_State int32

const (
	// Default value of State.
	Version_VersionState_STATE_UNSPECIFIED Version_VersionState_State = 0
	// The version creation is in progress.
	Version_VersionState_CREATION_IN_PROGRESS Version_VersionState_State = 1
	// The version creation failed.
	Version_VersionState_CREATION_FAILED Version_VersionState_State = 2
	// The version has been successfully created.
	Version_VersionState_CREATED Version_VersionState_State = 3
	// The version is under policy review (aka Approval).
	Version_VersionState_REVIEW_IN_PROGRESS Version_VersionState_State = 4
	// The version has been approved for policy review and can be deployed.
	Version_VersionState_APPROVED Version_VersionState_State = 5
	// The version has been conditionally approved but is pending final
	// review. It may be rolled back if final review is denied.
	Version_VersionState_CONDITIONALLY_APPROVED Version_VersionState_State = 6
	// The version has been denied for policy review.
	Version_VersionState_DENIED Version_VersionState_State = 7
	// The version is taken down as entire agent and all versions are taken
	// down.
	Version_VersionState_UNDER_TAKEDOWN Version_VersionState_State = 8
	// The version has been deleted.
	Version_VersionState_DELETED Version_VersionState_State = 9
)

// Enum value maps for Version_VersionState_State.
var (
	Version_VersionState_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATION_IN_PROGRESS",
		2: "CREATION_FAILED",
		3: "CREATED",
		4: "REVIEW_IN_PROGRESS",
		5: "APPROVED",
		6: "CONDITIONALLY_APPROVED",
		7: "DENIED",
		8: "UNDER_TAKEDOWN",
		9: "DELETED",
	}
	Version_VersionState_State_value = map[string]int32{
		"STATE_UNSPECIFIED":      0,
		"CREATION_IN_PROGRESS":   1,
		"CREATION_FAILED":        2,
		"CREATED":                3,
		"REVIEW_IN_PROGRESS":     4,
		"APPROVED":               5,
		"CONDITIONALLY_APPROVED": 6,
		"DENIED":                 7,
		"UNDER_TAKEDOWN":         8,
		"DELETED":                9,
	}
)

func (x Version_VersionState_State) Enum() *Version_VersionState_State {
	p := new(Version_VersionState_State)
	*p = x
	return p
}

func (x Version_VersionState_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Version_VersionState_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_actions_sdk_v2_version_proto_enumTypes[0].Descriptor()
}

func (Version_VersionState_State) Type() protoreflect.EnumType {
	return &file_google_actions_sdk_v2_version_proto_enumTypes[0]
}

func (x Version_VersionState_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Version_VersionState_State.Descriptor instead.
func (Version_VersionState_State) EnumDescriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_version_proto_rawDescGZIP(), []int{0, 0, 0}
}

// Definition of version resource.
type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of the version in the following format.
	// `projects/{project}/versions/{version}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The current state of the version.
	VersionState *Version_VersionState `protobuf:"bytes,2,opt,name=version_state,json=versionState,proto3" json:"version_state,omitempty"`
	// Email of the user who created this version.
	Creator string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	// Timestamp of the last change to this version.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_version_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_version_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_version_proto_rawDescGZIP(), []int{0}
}

func (x *Version) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Version) GetVersionState() *Version_VersionState {
	if x != nil {
		return x.VersionState
	}
	return nil
}

func (x *Version) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *Version) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// Represents the current state of the version.
type Version_VersionState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The current state of the version.
	State Version_VersionState_State `protobuf:"varint,1,opt,name=state,proto3,enum=google.actions.sdk.v2.Version_VersionState_State" json:"state,omitempty"`
	// User-friendly message for the current state of the version.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Version_VersionState) Reset() {
	*x = Version_VersionState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_version_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version_VersionState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version_VersionState) ProtoMessage() {}

func (x *Version_VersionState) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_version_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version_VersionState.ProtoReflect.Descriptor instead.
func (*Version_VersionState) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_version_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Version_VersionState) GetState() Version_VersionState_State {
	if x != nil {
		return x.State
	}
	return Version_VersionState_STATE_UNSPECIFIED
}

func (x *Version_VersionState) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_google_actions_sdk_v2_version_proto protoreflect.FileDescriptor

var file_google_actions_sdk_v2_version_proto_rawDesc = []byte{
	0x0a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x04, 0x0a, 0x07, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x1a, 0xbd, 0x02, 0x0a, 0x0c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x47, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xc9, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x01,
	0x12, 0x13, 0x0a, 0x0f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44,
	0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x49, 0x4e, 0x5f,
	0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x50,
	0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x05, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4f, 0x4e, 0x44,
	0x49, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56,
	0x45, 0x44, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x07,
	0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x41, 0x4b, 0x45, 0x44, 0x4f,
	0x57, 0x4e, 0x10, 0x08, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10,
	0x09, 0x3a, 0x4a, 0xea, 0x41, 0x47, 0x0a, 0x1e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0x42, 0x65, 0x0a,
	0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x42, 0x0c, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32,
	0x3b, 0x73, 0x64, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_actions_sdk_v2_version_proto_rawDescOnce sync.Once
	file_google_actions_sdk_v2_version_proto_rawDescData = file_google_actions_sdk_v2_version_proto_rawDesc
)

func file_google_actions_sdk_v2_version_proto_rawDescGZIP() []byte {
	file_google_actions_sdk_v2_version_proto_rawDescOnce.Do(func() {
		file_google_actions_sdk_v2_version_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_actions_sdk_v2_version_proto_rawDescData)
	})
	return file_google_actions_sdk_v2_version_proto_rawDescData
}

var file_google_actions_sdk_v2_version_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_actions_sdk_v2_version_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_actions_sdk_v2_version_proto_goTypes = []interface{}{
	(Version_VersionState_State)(0), // 0: google.actions.sdk.v2.Version.VersionState.State
	(*Version)(nil),                 // 1: google.actions.sdk.v2.Version
	(*Version_VersionState)(nil),    // 2: google.actions.sdk.v2.Version.VersionState
	(*timestamppb.Timestamp)(nil),   // 3: google.protobuf.Timestamp
}
var file_google_actions_sdk_v2_version_proto_depIdxs = []int32{
	2, // 0: google.actions.sdk.v2.Version.version_state:type_name -> google.actions.sdk.v2.Version.VersionState
	3, // 1: google.actions.sdk.v2.Version.update_time:type_name -> google.protobuf.Timestamp
	0, // 2: google.actions.sdk.v2.Version.VersionState.state:type_name -> google.actions.sdk.v2.Version.VersionState.State
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_actions_sdk_v2_version_proto_init() }
func file_google_actions_sdk_v2_version_proto_init() {
	if File_google_actions_sdk_v2_version_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_actions_sdk_v2_version_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
		file_google_actions_sdk_v2_version_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version_VersionState); i {
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
			RawDescriptor: file_google_actions_sdk_v2_version_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_actions_sdk_v2_version_proto_goTypes,
		DependencyIndexes: file_google_actions_sdk_v2_version_proto_depIdxs,
		EnumInfos:         file_google_actions_sdk_v2_version_proto_enumTypes,
		MessageInfos:      file_google_actions_sdk_v2_version_proto_msgTypes,
	}.Build()
	File_google_actions_sdk_v2_version_proto = out.File
	file_google_actions_sdk_v2_version_proto_rawDesc = nil
	file_google_actions_sdk_v2_version_proto_goTypes = nil
	file_google_actions_sdk_v2_version_proto_depIdxs = nil
}