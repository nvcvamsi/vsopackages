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
// source: google/actions/sdk/v2/webhook.proto

package sdk

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Metadata for different types of webhooks. If you're using
// `inlineCloudFunction`, your source code must be in a directory with the same
// name as the value for the `executeFunction` key.
// For example, a value of `my_webhook` for the`executeFunction` key would have
// a code structure like this:
//   - `/webhooks/my_webhook.yaml`
//   - `/webhooks/my_webhook/index.js`
//   - `/webhooks/my_webhook/package.json`
type Webhook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of handlers for this webhook.
	Handlers []*Webhook_Handler `protobuf:"bytes,1,rep,name=handlers,proto3" json:"handlers,omitempty"`
	// Only one webhook type is supported.
	//
	// Types that are assignable to WebhookType:
	//
	//	*Webhook_HttpsEndpoint_
	//	*Webhook_InlineCloudFunction_
	WebhookType isWebhook_WebhookType `protobuf_oneof:"webhook_type"`
}

func (x *Webhook) Reset() {
	*x = Webhook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_webhook_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Webhook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Webhook) ProtoMessage() {}

func (x *Webhook) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_webhook_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Webhook.ProtoReflect.Descriptor instead.
func (*Webhook) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_webhook_proto_rawDescGZIP(), []int{0}
}

func (x *Webhook) GetHandlers() []*Webhook_Handler {
	if x != nil {
		return x.Handlers
	}
	return nil
}

func (m *Webhook) GetWebhookType() isWebhook_WebhookType {
	if m != nil {
		return m.WebhookType
	}
	return nil
}

func (x *Webhook) GetHttpsEndpoint() *Webhook_HttpsEndpoint {
	if x, ok := x.GetWebhookType().(*Webhook_HttpsEndpoint_); ok {
		return x.HttpsEndpoint
	}
	return nil
}

func (x *Webhook) GetInlineCloudFunction() *Webhook_InlineCloudFunction {
	if x, ok := x.GetWebhookType().(*Webhook_InlineCloudFunction_); ok {
		return x.InlineCloudFunction
	}
	return nil
}

type isWebhook_WebhookType interface {
	isWebhook_WebhookType()
}

type Webhook_HttpsEndpoint_ struct {
	// Custom webhook HTTPS endpoint.
	HttpsEndpoint *Webhook_HttpsEndpoint `protobuf:"bytes,2,opt,name=https_endpoint,json=httpsEndpoint,proto3,oneof"`
}

type Webhook_InlineCloudFunction_ struct {
	// Metadata for cloud function deployed from code in the webhooks folder.
	InlineCloudFunction *Webhook_InlineCloudFunction `protobuf:"bytes,3,opt,name=inline_cloud_function,json=inlineCloudFunction,proto3,oneof"`
}

func (*Webhook_HttpsEndpoint_) isWebhook_WebhookType() {}

func (*Webhook_InlineCloudFunction_) isWebhook_WebhookType() {}

// Declares the name of the webhoook handler. A webhook can have
// multiple handlers registered. These handlers can be called from multiple
// places in your Actions project.
type Webhook_Handler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Name of the handler. Must be unique across all handlers the Actions
	// project. You can check the name of this handler to invoke the correct
	// function in your fulfillment source code.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Webhook_Handler) Reset() {
	*x = Webhook_Handler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_webhook_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Webhook_Handler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Webhook_Handler) ProtoMessage() {}

func (x *Webhook_Handler) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_webhook_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Webhook_Handler.ProtoReflect.Descriptor instead.
func (*Webhook_Handler) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_webhook_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Webhook_Handler) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// REST endpoint to notify if you're not using the inline editor.
type Webhook_HttpsEndpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The HTTPS base URL for your fulfillment endpoint (HTTP is not supported).
	// Handler names are appended to the base URL path after a colon
	// (following the style guide in
	// https://cloud.google.com/apis/design/custom_methods).
	// For example a base URL of 'https://gactions.service.com/api' would
	// receive requests with URL 'https://gactions.service.com/api:{method}'.
	BaseUrl string `protobuf:"bytes,1,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty"`
	// Map of HTTP parameters to be included in the POST request.
	HttpHeaders map[string]string `protobuf:"bytes,2,rep,name=http_headers,json=httpHeaders,proto3" json:"http_headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Version of the protocol used by the endpoint. This is the protocol shared
	// by all fulfillment types and not specific to Google fulfillment type.
	EndpointApiVersion int32 `protobuf:"varint,3,opt,name=endpoint_api_version,json=endpointApiVersion,proto3" json:"endpoint_api_version,omitempty"`
}

func (x *Webhook_HttpsEndpoint) Reset() {
	*x = Webhook_HttpsEndpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_webhook_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Webhook_HttpsEndpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Webhook_HttpsEndpoint) ProtoMessage() {}

func (x *Webhook_HttpsEndpoint) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_webhook_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Webhook_HttpsEndpoint.ProtoReflect.Descriptor instead.
func (*Webhook_HttpsEndpoint) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_webhook_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Webhook_HttpsEndpoint) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

func (x *Webhook_HttpsEndpoint) GetHttpHeaders() map[string]string {
	if x != nil {
		return x.HttpHeaders
	}
	return nil
}

func (x *Webhook_HttpsEndpoint) GetEndpointApiVersion() int32 {
	if x != nil {
		return x.EndpointApiVersion
	}
	return 0
}

// Holds the metadata of an inline Cloud Function deployed from the
// webhooks folder.
type Webhook_InlineCloudFunction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the Cloud Function entry point. The value of this field
	// should match the name of the method exported from the source code.
	ExecuteFunction string `protobuf:"bytes,1,opt,name=execute_function,json=executeFunction,proto3" json:"execute_function,omitempty"`
}

func (x *Webhook_InlineCloudFunction) Reset() {
	*x = Webhook_InlineCloudFunction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_webhook_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Webhook_InlineCloudFunction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Webhook_InlineCloudFunction) ProtoMessage() {}

func (x *Webhook_InlineCloudFunction) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_webhook_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Webhook_InlineCloudFunction.ProtoReflect.Descriptor instead.
func (*Webhook_InlineCloudFunction) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_webhook_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Webhook_InlineCloudFunction) GetExecuteFunction() string {
	if x != nil {
		return x.ExecuteFunction
	}
	return ""
}

var File_google_actions_sdk_v2_webhook_proto protoreflect.FileDescriptor

var file_google_actions_sdk_v2_webhook_proto_rawDesc = []byte{
	0x0a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x05,
	0x0a, 0x07, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x42, 0x0a, 0x08, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b,
	0x2e, 0x76, 0x32, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x52, 0x08, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x55, 0x0a,
	0x0e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x57, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x73, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x68, 0x74, 0x74, 0x70, 0x73, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x68, 0x0a, 0x15, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x57, 0x65, 0x62, 0x68,
	0x6f, 0x6f, 0x6b, 0x2e, 0x49, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x13, 0x69, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x22,
	0x0a, 0x07, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x1a, 0xfe, 0x01, 0x0a, 0x0d, 0x48, 0x74, 0x74, 0x70, 0x73, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x12,
	0x60, 0x0a, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x57, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x73, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x30, 0x0a, 0x14, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x61, 0x70,
	0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x12, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x1a, 0x3e, 0x0a, 0x10, 0x48, 0x74, 0x74, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x40, 0x0a, 0x13, 0x49, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0e, 0x0a, 0x0c, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x65, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x76, 0x32, 0x42, 0x0c, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x3b, 0x73, 0x64, 0x6b, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_actions_sdk_v2_webhook_proto_rawDescOnce sync.Once
	file_google_actions_sdk_v2_webhook_proto_rawDescData = file_google_actions_sdk_v2_webhook_proto_rawDesc
)

func file_google_actions_sdk_v2_webhook_proto_rawDescGZIP() []byte {
	file_google_actions_sdk_v2_webhook_proto_rawDescOnce.Do(func() {
		file_google_actions_sdk_v2_webhook_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_actions_sdk_v2_webhook_proto_rawDescData)
	})
	return file_google_actions_sdk_v2_webhook_proto_rawDescData
}

var file_google_actions_sdk_v2_webhook_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_actions_sdk_v2_webhook_proto_goTypes = []interface{}{
	(*Webhook)(nil),                     // 0: google.actions.sdk.v2.Webhook
	(*Webhook_Handler)(nil),             // 1: google.actions.sdk.v2.Webhook.Handler
	(*Webhook_HttpsEndpoint)(nil),       // 2: google.actions.sdk.v2.Webhook.HttpsEndpoint
	(*Webhook_InlineCloudFunction)(nil), // 3: google.actions.sdk.v2.Webhook.InlineCloudFunction
	nil,                                 // 4: google.actions.sdk.v2.Webhook.HttpsEndpoint.HttpHeadersEntry
}
var file_google_actions_sdk_v2_webhook_proto_depIdxs = []int32{
	1, // 0: google.actions.sdk.v2.Webhook.handlers:type_name -> google.actions.sdk.v2.Webhook.Handler
	2, // 1: google.actions.sdk.v2.Webhook.https_endpoint:type_name -> google.actions.sdk.v2.Webhook.HttpsEndpoint
	3, // 2: google.actions.sdk.v2.Webhook.inline_cloud_function:type_name -> google.actions.sdk.v2.Webhook.InlineCloudFunction
	4, // 3: google.actions.sdk.v2.Webhook.HttpsEndpoint.http_headers:type_name -> google.actions.sdk.v2.Webhook.HttpsEndpoint.HttpHeadersEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_actions_sdk_v2_webhook_proto_init() }
func file_google_actions_sdk_v2_webhook_proto_init() {
	if File_google_actions_sdk_v2_webhook_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_actions_sdk_v2_webhook_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Webhook); i {
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
		file_google_actions_sdk_v2_webhook_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Webhook_Handler); i {
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
		file_google_actions_sdk_v2_webhook_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Webhook_HttpsEndpoint); i {
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
		file_google_actions_sdk_v2_webhook_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Webhook_InlineCloudFunction); i {
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
	file_google_actions_sdk_v2_webhook_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Webhook_HttpsEndpoint_)(nil),
		(*Webhook_InlineCloudFunction_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_actions_sdk_v2_webhook_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_actions_sdk_v2_webhook_proto_goTypes,
		DependencyIndexes: file_google_actions_sdk_v2_webhook_proto_depIdxs,
		MessageInfos:      file_google_actions_sdk_v2_webhook_proto_msgTypes,
	}.Build()
	File_google_actions_sdk_v2_webhook_proto = out.File
	file_google_actions_sdk_v2_webhook_proto_rawDesc = nil
	file_google_actions_sdk_v2_webhook_proto_goTypes = nil
	file_google_actions_sdk_v2_webhook_proto_depIdxs = nil
}
