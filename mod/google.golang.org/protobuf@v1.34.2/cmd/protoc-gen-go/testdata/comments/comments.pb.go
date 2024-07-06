// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd/protoc-gen-go/testdata/comments/comments.proto

// COMMENT: package goproto.protoc.comments;

package comments

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

// COMMENT: Enum1.Leading
type Enum1 int32

const (
	// COMMENT: FOO.Leading
	Enum1_FOO Enum1 = 0 // COMMENT: FOO.InlineTrailing
	// COMMENT: BAR.Leading
	Enum1_BAR Enum1 = 1
)

// Enum value maps for Enum1.
var (
	Enum1_name = map[int32]string{
		0: "FOO",
		1: "BAR",
	}
	Enum1_value = map[string]int32{
		"FOO": 0,
		"BAR": 1,
	}
)

func (x Enum1) Enum() *Enum1 {
	p := new(Enum1)
	*p = x
	return p
}

func (x Enum1) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum1) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_comments_comments_proto_enumTypes[0].Descriptor()
}

func (Enum1) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_comments_comments_proto_enumTypes[0]
}

func (x Enum1) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Enum1) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Enum1(num)
	return nil
}

// Deprecated: Use Enum1.Descriptor instead.
func (Enum1) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_comments_comments_proto_rawDescGZIP(), []int{0}
}

// COMMENT: Message1.Leading
type Message1 struct {
	state           protoimpl.MessageState
	sizeCache       protoimpl.SizeCache
	unknownFields   protoimpl.UnknownFields
	extensionFields protoimpl.ExtensionFields

	// COMMENT: Field1A.Leading
	Field1A *string `protobuf:"bytes,1,opt,name=Field1A" json:"Field1A,omitempty"` // COMMENT: Field1A.Trailing
	// COMMENT: Oneof1A.Leading
	//
	// Types that are assignable to Oneof1A:
	//
	//	*Message1_Oneof1AField1
	Oneof1A isMessage1_Oneof1A `protobuf_oneof:"Oneof1a"`
}

func (x *Message1) Reset() {
	*x = Message1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_testdata_comments_comments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message1) ProtoMessage() {}

func (x *Message1) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_comments_comments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message1.ProtoReflect.Descriptor instead.
func (*Message1) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_comments_comments_proto_rawDescGZIP(), []int{0}
}

func (x *Message1) GetField1A() string {
	if x != nil && x.Field1A != nil {
		return *x.Field1A
	}
	return ""
}

func (m *Message1) GetOneof1A() isMessage1_Oneof1A {
	if m != nil {
		return m.Oneof1A
	}
	return nil
}

func (x *Message1) GetOneof1AField1() string {
	if x, ok := x.GetOneof1A().(*Message1_Oneof1AField1); ok {
		return x.Oneof1AField1
	}
	return ""
}

type isMessage1_Oneof1A interface {
	isMessage1_Oneof1A()
}

type Message1_Oneof1AField1 struct {
	// COMMENT: Oneof1AField1.Leading
	Oneof1AField1 string `protobuf:"bytes,2,opt,name=Oneof1AField1,oneof"` // COMMENT: Oneof1AField1.Trailing
}

func (*Message1_Oneof1AField1) isMessage1_Oneof1A() {}

// COMMENT: Message2
type Message2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Message2) Reset() {
	*x = Message2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_testdata_comments_comments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message2) ProtoMessage() {}

func (x *Message2) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_comments_comments_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message2.ProtoReflect.Descriptor instead.
func (*Message2) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_comments_comments_proto_rawDescGZIP(), []int{1}
}

// COMMENT: Message1A.Leading
type Message1_Message1A struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Message1_Message1A) Reset() {
	*x = Message1_Message1A{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_testdata_comments_comments_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message1_Message1A) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message1_Message1A) ProtoMessage() {}

func (x *Message1_Message1A) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_comments_comments_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message1_Message1A.ProtoReflect.Descriptor instead.
func (*Message1_Message1A) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_comments_comments_proto_rawDescGZIP(), []int{0, 0}
}

// COMMENT: Message1B
type Message1_Message1B struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Message1_Message1B) Reset() {
	*x = Message1_Message1B{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_testdata_comments_comments_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message1_Message1B) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message1_Message1B) ProtoMessage() {}

func (x *Message1_Message1B) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_comments_comments_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message1_Message1B.ProtoReflect.Descriptor instead.
func (*Message1_Message1B) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_comments_comments_proto_rawDescGZIP(), []int{0, 1}
}

// COMMENT: Message2A
type Message2_Message2A struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Message2_Message2A) Reset() {
	*x = Message2_Message2A{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_testdata_comments_comments_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message2_Message2A) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message2_Message2A) ProtoMessage() {}

func (x *Message2_Message2A) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_comments_comments_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message2_Message2A.ProtoReflect.Descriptor instead.
func (*Message2_Message2A) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_comments_comments_proto_rawDescGZIP(), []int{1, 0}
}

// COMMENT: Message2B
type Message2_Message2B struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Message2_Message2B) Reset() {
	*x = Message2_Message2B{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_testdata_comments_comments_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message2_Message2B) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message2_Message2B) ProtoMessage() {}

func (x *Message2_Message2B) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_comments_comments_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message2_Message2B.ProtoReflect.Descriptor instead.
func (*Message2_Message2B) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_comments_comments_proto_rawDescGZIP(), []int{1, 1}
}

var file_cmd_protoc_gen_go_testdata_comments_comments_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*Message1)(nil),
		ExtensionType: (*Message1)(nil),
		Field:         100,
		Name:          "goproto.protoc.comments.extension",
		Tag:           "bytes,100,opt,name=extension",
		Filename:      "cmd/protoc-gen-go/testdata/comments/comments.proto",
	},
}

// Extension fields to Message1.
var (
	// COMMENT: Extension.Leading
	//
	// optional goproto.protoc.comments.Message1 extension = 100;
	E_Extension = &file_cmd_protoc_gen_go_testdata_comments_comments_proto_extTypes[0] // COMMENT: Extension.Trailing
)

var File_cmd_protoc_gen_go_testdata_comments_comments_proto protoreflect.FileDescriptor

var file_cmd_protoc_gen_go_testdata_comments_comments_proto_rawDesc = []byte{
	0x0a, 0x32, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x7b, 0x0a,
	0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x31, 0x41, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x31, 0x41, 0x12, 0x26, 0x0a, 0x0d, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x31, 0x41, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x4f, 0x6e,
	0x65, 0x6f, 0x66, 0x31, 0x41, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x1a, 0x0b, 0x0a, 0x09, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x41, 0x1a, 0x0b, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x31, 0x42, 0x2a, 0x08, 0x08, 0x64, 0x10, 0x80, 0x80, 0x80, 0x80, 0x02, 0x42,
	0x09, 0x0a, 0x07, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x31, 0x61, 0x22, 0x24, 0x0a, 0x08, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x1a, 0x0b, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x41, 0x1a, 0x0b, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x42,
	0x2a, 0x19, 0x0a, 0x05, 0x45, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x07, 0x0a, 0x03, 0x46, 0x4f, 0x4f,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x41, 0x52, 0x10, 0x01, 0x3a, 0x62, 0x0a, 0x09, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x31, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x42,
	0x40, 0x5a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6d,
	0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73,
}

var (
	file_cmd_protoc_gen_go_testdata_comments_comments_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_testdata_comments_comments_proto_rawDescData = file_cmd_protoc_gen_go_testdata_comments_comments_proto_rawDesc
)

func file_cmd_protoc_gen_go_testdata_comments_comments_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_testdata_comments_comments_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_testdata_comments_comments_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_protoc_gen_go_testdata_comments_comments_proto_rawDescData)
	})
	return file_cmd_protoc_gen_go_testdata_comments_comments_proto_rawDescData
}

var file_cmd_protoc_gen_go_testdata_comments_comments_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cmd_protoc_gen_go_testdata_comments_comments_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cmd_protoc_gen_go_testdata_comments_comments_proto_goTypes = []any{
	(Enum1)(0),                 // 0: goproto.protoc.comments.Enum1
	(*Message1)(nil),           // 1: goproto.protoc.comments.Message1
	(*Message2)(nil),           // 2: goproto.protoc.comments.Message2
	(*Message1_Message1A)(nil), // 3: goproto.protoc.comments.Message1.Message1A
	(*Message1_Message1B)(nil), // 4: goproto.protoc.comments.Message1.Message1B
	(*Message2_Message2A)(nil), // 5: goproto.protoc.comments.Message2.Message2A
	(*Message2_Message2B)(nil), // 6: goproto.protoc.comments.Message2.Message2B
}
var file_cmd_protoc_gen_go_testdata_comments_comments_proto_depIdxs = []int32{
	1, // 0: goproto.protoc.comments.extension:extendee -> goproto.protoc.comments.Message1
	1, // 1: goproto.protoc.comments.extension:type_name -> goproto.protoc.comments.Message1
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_testdata_comments_comments_proto_init() }
func file_cmd_protoc_gen_go_testdata_comments_comments_proto_init() {
	if File_cmd_protoc_gen_go_testdata_comments_comments_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_protoc_gen_go_testdata_comments_comments_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Message1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			case 3:
				return &v.extensionFields
			default:
				return nil
			}
		}
		file_cmd_protoc_gen_go_testdata_comments_comments_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Message2); i {
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
		file_cmd_protoc_gen_go_testdata_comments_comments_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Message1_Message1A); i {
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
		file_cmd_protoc_gen_go_testdata_comments_comments_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Message1_Message1B); i {
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
		file_cmd_protoc_gen_go_testdata_comments_comments_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Message2_Message2A); i {
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
		file_cmd_protoc_gen_go_testdata_comments_comments_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Message2_Message2B); i {
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
	file_cmd_protoc_gen_go_testdata_comments_comments_proto_msgTypes[0].OneofWrappers = []any{
		(*Message1_Oneof1AField1)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_protoc_gen_go_testdata_comments_comments_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_testdata_comments_comments_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_testdata_comments_comments_proto_depIdxs,
		EnumInfos:         file_cmd_protoc_gen_go_testdata_comments_comments_proto_enumTypes,
		MessageInfos:      file_cmd_protoc_gen_go_testdata_comments_comments_proto_msgTypes,
		ExtensionInfos:    file_cmd_protoc_gen_go_testdata_comments_comments_proto_extTypes,
	}.Build()
	File_cmd_protoc_gen_go_testdata_comments_comments_proto = out.File
	file_cmd_protoc_gen_go_testdata_comments_comments_proto_rawDesc = nil
	file_cmd_protoc_gen_go_testdata_comments_comments_proto_goTypes = nil
	file_cmd_protoc_gen_go_testdata_comments_comments_proto_depIdxs = nil
}
