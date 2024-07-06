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
// source: google/actions/sdk/v2/theme_customization.proto

package sdk

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Describes how the borders of images should be rendered.
type ThemeCustomization_ImageCornerStyle int32

const (
	// Undefined / Unspecified.
	ThemeCustomization_IMAGE_CORNER_STYLE_UNSPECIFIED ThemeCustomization_ImageCornerStyle = 0
	// Round corner for image.
	ThemeCustomization_CURVED ThemeCustomization_ImageCornerStyle = 1
	// Rectangular corner for image.
	ThemeCustomization_ANGLED ThemeCustomization_ImageCornerStyle = 2
)

// Enum value maps for ThemeCustomization_ImageCornerStyle.
var (
	ThemeCustomization_ImageCornerStyle_name = map[int32]string{
		0: "IMAGE_CORNER_STYLE_UNSPECIFIED",
		1: "CURVED",
		2: "ANGLED",
	}
	ThemeCustomization_ImageCornerStyle_value = map[string]int32{
		"IMAGE_CORNER_STYLE_UNSPECIFIED": 0,
		"CURVED":                         1,
		"ANGLED":                         2,
	}
)

func (x ThemeCustomization_ImageCornerStyle) Enum() *ThemeCustomization_ImageCornerStyle {
	p := new(ThemeCustomization_ImageCornerStyle)
	*p = x
	return p
}

func (x ThemeCustomization_ImageCornerStyle) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ThemeCustomization_ImageCornerStyle) Descriptor() protoreflect.EnumDescriptor {
	return file_google_actions_sdk_v2_theme_customization_proto_enumTypes[0].Descriptor()
}

func (ThemeCustomization_ImageCornerStyle) Type() protoreflect.EnumType {
	return &file_google_actions_sdk_v2_theme_customization_proto_enumTypes[0]
}

func (x ThemeCustomization_ImageCornerStyle) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ThemeCustomization_ImageCornerStyle.Descriptor instead.
func (ThemeCustomization_ImageCornerStyle) EnumDescriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_theme_customization_proto_rawDescGZIP(), []int{0, 0}
}

// Styles applied to cards that are presented to users
type ThemeCustomization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Background color of cards. Acts as a fallback if `background_image` is
	// not provided by developers or `background_image` doesn't fit for certain
	// surfaces.
	// Example usage: #FAFAFA
	BackgroundColor string `protobuf:"bytes,1,opt,name=background_color,json=backgroundColor,proto3" json:"background_color,omitempty"`
	// Primary theme color of the Action will be used to set text color of title,
	// action item background color for Actions on Google cards.
	// Example usage: #FAFAFA
	PrimaryColor string `protobuf:"bytes,2,opt,name=primary_color,json=primaryColor,proto3" json:"primary_color,omitempty"`
	// The font family that will be used for title of cards.
	// Supported fonts:
	// - Sans Serif
	// - Sans Serif Medium
	// - Sans Serif Bold
	// - Sans Serif Black
	// - Sans Serif Condensed
	// - Sans Serif Condensed Medium
	// - Serif
	// - Serif Bold
	// - Monospace
	// - Cursive
	// - Sans Serif Smallcaps
	FontFamily string `protobuf:"bytes,3,opt,name=font_family,json=fontFamily,proto3" json:"font_family,omitempty"`
	// Border style of foreground image of cards. For example, can be applied on
	// the foreground image of a basic card or carousel card.
	ImageCornerStyle ThemeCustomization_ImageCornerStyle `protobuf:"varint,4,opt,name=image_corner_style,json=imageCornerStyle,proto3,enum=google.actions.sdk.v2.ThemeCustomization_ImageCornerStyle" json:"image_corner_style,omitempty"`
	// Landscape mode (minimum 1920x1200 pixels).
	// This should be specified as a reference to the corresponding image in the
	// `resources/images/` directory. Eg: `$resources.images.foo` (without the
	// extension) for image in `resources/images/foo.jpg`
	// When working on a project pulled from Console the Google managed url pulled
	// could be used.
	LandscapeBackgroundImage string `protobuf:"bytes,5,opt,name=landscape_background_image,json=landscapeBackgroundImage,proto3" json:"landscape_background_image,omitempty"`
	// Portrait mode (minimum 1200x1920 pixels).
	// This should be specified as a reference to the corresponding image in the
	// `resources/images/` directory. Eg: `$resources.images.foo` (without the
	// extension) for image in `resources/images/foo.jpg`
	// When working on a project pulled from Console the Google managed url pulled
	// could be used.
	PortraitBackgroundImage string `protobuf:"bytes,6,opt,name=portrait_background_image,json=portraitBackgroundImage,proto3" json:"portrait_background_image,omitempty"`
}

func (x *ThemeCustomization) Reset() {
	*x = ThemeCustomization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_theme_customization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThemeCustomization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThemeCustomization) ProtoMessage() {}

func (x *ThemeCustomization) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_theme_customization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThemeCustomization.ProtoReflect.Descriptor instead.
func (*ThemeCustomization) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_theme_customization_proto_rawDescGZIP(), []int{0}
}

func (x *ThemeCustomization) GetBackgroundColor() string {
	if x != nil {
		return x.BackgroundColor
	}
	return ""
}

func (x *ThemeCustomization) GetPrimaryColor() string {
	if x != nil {
		return x.PrimaryColor
	}
	return ""
}

func (x *ThemeCustomization) GetFontFamily() string {
	if x != nil {
		return x.FontFamily
	}
	return ""
}

func (x *ThemeCustomization) GetImageCornerStyle() ThemeCustomization_ImageCornerStyle {
	if x != nil {
		return x.ImageCornerStyle
	}
	return ThemeCustomization_IMAGE_CORNER_STYLE_UNSPECIFIED
}

func (x *ThemeCustomization) GetLandscapeBackgroundImage() string {
	if x != nil {
		return x.LandscapeBackgroundImage
	}
	return ""
}

func (x *ThemeCustomization) GetPortraitBackgroundImage() string {
	if x != nil {
		return x.PortraitBackgroundImage
	}
	return ""
}

var File_google_actions_sdk_v2_theme_customization_proto protoreflect.FileDescriptor

var file_google_actions_sdk_v2_theme_customization_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x5f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x22, 0xb9, 0x03, 0x0a, 0x12, 0x54, 0x68, 0x65,
	0x6d, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x29, 0x0a, 0x10, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12,
	0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x6e, 0x74, 0x5f, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x6e, 0x74, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79,
	0x12, 0x68, 0x0a, 0x12, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x72, 0x6e, 0x65, 0x72,
	0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64,
	0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72,
	0x6e, 0x65, 0x72, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x52, 0x10, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x72, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x1a, 0x6c, 0x61,
	0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18,
	0x6c, 0x61, 0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x19, 0x70, 0x6f, 0x72, 0x74,
	0x72, 0x61, 0x69, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x70, 0x6f, 0x72,
	0x74, 0x72, 0x61, 0x69, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x22, 0x4e, 0x0a, 0x10, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72,
	0x6e, 0x65, 0x72, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x1e, 0x49, 0x4d, 0x41, 0x47,
	0x45, 0x5f, 0x43, 0x4f, 0x52, 0x4e, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x59, 0x4c, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x43, 0x55, 0x52, 0x56, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x4e, 0x47, 0x4c,
	0x45, 0x44, 0x10, 0x02, 0x42, 0x70, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76,
	0x32, 0x42, 0x17, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f,
	0x76, 0x32, 0x3b, 0x73, 0x64, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_actions_sdk_v2_theme_customization_proto_rawDescOnce sync.Once
	file_google_actions_sdk_v2_theme_customization_proto_rawDescData = file_google_actions_sdk_v2_theme_customization_proto_rawDesc
)

func file_google_actions_sdk_v2_theme_customization_proto_rawDescGZIP() []byte {
	file_google_actions_sdk_v2_theme_customization_proto_rawDescOnce.Do(func() {
		file_google_actions_sdk_v2_theme_customization_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_actions_sdk_v2_theme_customization_proto_rawDescData)
	})
	return file_google_actions_sdk_v2_theme_customization_proto_rawDescData
}

var file_google_actions_sdk_v2_theme_customization_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_actions_sdk_v2_theme_customization_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_actions_sdk_v2_theme_customization_proto_goTypes = []interface{}{
	(ThemeCustomization_ImageCornerStyle)(0), // 0: google.actions.sdk.v2.ThemeCustomization.ImageCornerStyle
	(*ThemeCustomization)(nil),               // 1: google.actions.sdk.v2.ThemeCustomization
}
var file_google_actions_sdk_v2_theme_customization_proto_depIdxs = []int32{
	0, // 0: google.actions.sdk.v2.ThemeCustomization.image_corner_style:type_name -> google.actions.sdk.v2.ThemeCustomization.ImageCornerStyle
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_actions_sdk_v2_theme_customization_proto_init() }
func file_google_actions_sdk_v2_theme_customization_proto_init() {
	if File_google_actions_sdk_v2_theme_customization_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_actions_sdk_v2_theme_customization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThemeCustomization); i {
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
			RawDescriptor: file_google_actions_sdk_v2_theme_customization_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_actions_sdk_v2_theme_customization_proto_goTypes,
		DependencyIndexes: file_google_actions_sdk_v2_theme_customization_proto_depIdxs,
		EnumInfos:         file_google_actions_sdk_v2_theme_customization_proto_enumTypes,
		MessageInfos:      file_google_actions_sdk_v2_theme_customization_proto_msgTypes,
	}.Build()
	File_google_actions_sdk_v2_theme_customization_proto = out.File
	file_google_actions_sdk_v2_theme_customization_proto_rawDesc = nil
	file_google_actions_sdk_v2_theme_customization_proto_goTypes = nil
	file_google_actions_sdk_v2_theme_customization_proto_depIdxs = nil
}
