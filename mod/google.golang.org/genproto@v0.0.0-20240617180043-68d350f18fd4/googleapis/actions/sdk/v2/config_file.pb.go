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
// source: google/actions/sdk/v2/config_file.proto

package sdk

import (
	reflect "reflect"
	sync "sync"

	interactionmodel "google.golang.org/genproto/googleapis/actions/sdk/v2/interactionmodel"
	prompt "google.golang.org/genproto/googleapis/actions/sdk/v2/interactionmodel/prompt"
	_type "google.golang.org/genproto/googleapis/actions/sdk/v2/interactionmodel/type"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Wrapper for repeated config files. Repeated fields cannot exist in a oneof.
type ConfigFiles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Multiple config files.
	ConfigFiles []*ConfigFile `protobuf:"bytes,1,rep,name=config_files,json=configFiles,proto3" json:"config_files,omitempty"`
}

func (x *ConfigFiles) Reset() {
	*x = ConfigFiles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_config_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigFiles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigFiles) ProtoMessage() {}

func (x *ConfigFiles) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_config_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigFiles.ProtoReflect.Descriptor instead.
func (*ConfigFiles) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_config_file_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigFiles) GetConfigFiles() []*ConfigFile {
	if x != nil {
		return x.ConfigFiles
	}
	return nil
}

// Represents a single file which contains structured data. Developers can
// define most of their project using structured config including Actions,
// Settings, Fulfillment.
type ConfigFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Relative path of the config file from the project root in the SDK file
	// structure. Each file types below have an allowed file path.
	// Eg: settings/settings.yaml
	FilePath string `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	// Each type of config file should have a corresponding field in the oneof.
	//
	// Types that are assignable to File:
	//
	//	*ConfigFile_Manifest
	//	*ConfigFile_Actions
	//	*ConfigFile_Settings
	//	*ConfigFile_Webhook
	//	*ConfigFile_Intent
	//	*ConfigFile_Type
	//	*ConfigFile_EntitySet
	//	*ConfigFile_GlobalIntentEvent
	//	*ConfigFile_Scene
	//	*ConfigFile_StaticPrompt
	//	*ConfigFile_AccountLinkingSecret
	//	*ConfigFile_ResourceBundle
	File isConfigFile_File `protobuf_oneof:"file"`
}

func (x *ConfigFile) Reset() {
	*x = ConfigFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_config_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigFile) ProtoMessage() {}

func (x *ConfigFile) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_config_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigFile.ProtoReflect.Descriptor instead.
func (*ConfigFile) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_config_file_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigFile) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (m *ConfigFile) GetFile() isConfigFile_File {
	if m != nil {
		return m.File
	}
	return nil
}

func (x *ConfigFile) GetManifest() *Manifest {
	if x, ok := x.GetFile().(*ConfigFile_Manifest); ok {
		return x.Manifest
	}
	return nil
}

func (x *ConfigFile) GetActions() *Actions {
	if x, ok := x.GetFile().(*ConfigFile_Actions); ok {
		return x.Actions
	}
	return nil
}

func (x *ConfigFile) GetSettings() *Settings {
	if x, ok := x.GetFile().(*ConfigFile_Settings); ok {
		return x.Settings
	}
	return nil
}

func (x *ConfigFile) GetWebhook() *Webhook {
	if x, ok := x.GetFile().(*ConfigFile_Webhook); ok {
		return x.Webhook
	}
	return nil
}

func (x *ConfigFile) GetIntent() *interactionmodel.Intent {
	if x, ok := x.GetFile().(*ConfigFile_Intent); ok {
		return x.Intent
	}
	return nil
}

func (x *ConfigFile) GetType() *_type.Type {
	if x, ok := x.GetFile().(*ConfigFile_Type); ok {
		return x.Type
	}
	return nil
}

func (x *ConfigFile) GetEntitySet() *interactionmodel.EntitySet {
	if x, ok := x.GetFile().(*ConfigFile_EntitySet); ok {
		return x.EntitySet
	}
	return nil
}

func (x *ConfigFile) GetGlobalIntentEvent() *interactionmodel.GlobalIntentEvent {
	if x, ok := x.GetFile().(*ConfigFile_GlobalIntentEvent); ok {
		return x.GlobalIntentEvent
	}
	return nil
}

func (x *ConfigFile) GetScene() *interactionmodel.Scene {
	if x, ok := x.GetFile().(*ConfigFile_Scene); ok {
		return x.Scene
	}
	return nil
}

func (x *ConfigFile) GetStaticPrompt() *prompt.StaticPrompt {
	if x, ok := x.GetFile().(*ConfigFile_StaticPrompt); ok {
		return x.StaticPrompt
	}
	return nil
}

func (x *ConfigFile) GetAccountLinkingSecret() *AccountLinkingSecret {
	if x, ok := x.GetFile().(*ConfigFile_AccountLinkingSecret); ok {
		return x.AccountLinkingSecret
	}
	return nil
}

func (x *ConfigFile) GetResourceBundle() *structpb.Struct {
	if x, ok := x.GetFile().(*ConfigFile_ResourceBundle); ok {
		return x.ResourceBundle
	}
	return nil
}

type isConfigFile_File interface {
	isConfigFile_File()
}

type ConfigFile_Manifest struct {
	// Single manifest file.
	// Allowed file path: `manifest.yaml`
	Manifest *Manifest `protobuf:"bytes,2,opt,name=manifest,proto3,oneof"`
}

type ConfigFile_Actions struct {
	// Single actions file with all the actions defined.
	// Allowed file paths: `actions/{language}?/actions.yaml`
	Actions *Actions `protobuf:"bytes,3,opt,name=actions,proto3,oneof"`
}

type ConfigFile_Settings struct {
	// Single settings config which includes non-localizable settings and
	// settings for the project's default locale (if specified).
	// For a locale override file, only localized_settings field will be
	// populated.
	// Allowed file paths: `settings/{language}?/settings.yaml`
	// Note that the non-localized settings file `settings/settings.yaml` must
	// be present in the write flow requests.
	Settings *Settings `protobuf:"bytes,4,opt,name=settings,proto3,oneof"`
}

type ConfigFile_Webhook struct {
	// Single webhook definition.
	// Allowed file path: `webhooks/{WebhookName}.yaml`
	Webhook *Webhook `protobuf:"bytes,6,opt,name=webhook,proto3,oneof"`
}

type ConfigFile_Intent struct {
	// Single intent definition.
	// Allowed file paths: `custom/intents/{language}?/{IntentName}.yaml`
	Intent *interactionmodel.Intent `protobuf:"bytes,7,opt,name=intent,proto3,oneof"`
}

type ConfigFile_Type struct {
	// Single type definition.
	// Allowed file paths: `custom/types/{language}?/{TypeName}.yaml`
	Type *_type.Type `protobuf:"bytes,8,opt,name=type,proto3,oneof"`
}

type ConfigFile_EntitySet struct {
	// Single entity set definition.
	// Allowed file paths: `custom/entitySets/{language}?/{EntitySetName}.yaml`
	EntitySet *interactionmodel.EntitySet `protobuf:"bytes,15,opt,name=entity_set,json=entitySet,proto3,oneof"`
}

type ConfigFile_GlobalIntentEvent struct {
	// Single global intent event definition.
	// Allowed file paths: `custom/global/{GlobalIntentEventName}.yaml`
	// The file name (GlobalIntentEventName) should be the name of the intent
	// that this global intent event corresponds to.
	GlobalIntentEvent *interactionmodel.GlobalIntentEvent `protobuf:"bytes,9,opt,name=global_intent_event,json=globalIntentEvent,proto3,oneof"`
}

type ConfigFile_Scene struct {
	// Single scene definition.
	// Allowed file paths: `custom/scenes/{SceneName}.yaml`
	Scene *interactionmodel.Scene `protobuf:"bytes,10,opt,name=scene,proto3,oneof"`
}

type ConfigFile_StaticPrompt struct {
	// Single static prompt definition.
	// Allowed file paths: `custom/prompts/{language}?/{StaticPromptName}.yaml`
	StaticPrompt *prompt.StaticPrompt `protobuf:"bytes,11,opt,name=static_prompt,json=staticPrompt,proto3,oneof"`
}

type ConfigFile_AccountLinkingSecret struct {
	// Metadata corresponding to the client secret used in account linking.
	// Allowed file path: `settings/accountLinkingSecret.yaml`
	AccountLinkingSecret *AccountLinkingSecret `protobuf:"bytes,13,opt,name=account_linking_secret,json=accountLinkingSecret,proto3,oneof"`
}

type ConfigFile_ResourceBundle struct {
	// Single resource bundle, which is a map from a string to a string or list
	// of strings. Resource bundles could be used for localizing strings in
	// static prompts.
	// Allowed file paths: `resources/strings/{language}?/{multiple
	// directories}?/{BundleName}.yaml`
	ResourceBundle *structpb.Struct `protobuf:"bytes,12,opt,name=resource_bundle,json=resourceBundle,proto3,oneof"`
}

func (*ConfigFile_Manifest) isConfigFile_File() {}

func (*ConfigFile_Actions) isConfigFile_File() {}

func (*ConfigFile_Settings) isConfigFile_File() {}

func (*ConfigFile_Webhook) isConfigFile_File() {}

func (*ConfigFile_Intent) isConfigFile_File() {}

func (*ConfigFile_Type) isConfigFile_File() {}

func (*ConfigFile_EntitySet) isConfigFile_File() {}

func (*ConfigFile_GlobalIntentEvent) isConfigFile_File() {}

func (*ConfigFile_Scene) isConfigFile_File() {}

func (*ConfigFile_StaticPrompt) isConfigFile_File() {}

func (*ConfigFile_AccountLinkingSecret) isConfigFile_File() {}

func (*ConfigFile_ResourceBundle) isConfigFile_File() {}

var File_google_actions_sdk_v2_config_file_proto protoreflect.FileDescriptor

var file_google_actions_sdk_v2_config_file_proto_rawDesc = []byte{
	0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32,
	0x1a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x6c, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x70,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f,
	0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x6d,
	0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64,
	0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x77, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b,
	0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x0b,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x22, 0xcf, 0x07, 0x0a, 0x0a,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x3d, 0x0a, 0x08, 0x6d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76,
	0x32, 0x2e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x61,
	0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x3d, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x48, 0x00, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x3a, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x48, 0x00, 0x52, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x48, 0x0a,
	0x06, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52,
	0x06, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x52, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x53, 0x65, 0x74, 0x48, 0x00, 0x52, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x53, 0x65, 0x74, 0x12, 0x6b, 0x0a, 0x13, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x11,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x45, 0x0a, 0x05, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x48,
	0x00, 0x52, 0x05, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x12, 0x62, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x48, 0x00, 0x52, 0x0c,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x63, 0x0a, 0x16,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x5f,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64,
	0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x69, 0x6e, 0x67, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x48, 0x00, 0x52, 0x14, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x12, 0x42, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x62, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x68, 0x0a,
	0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x42, 0x0f, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b,
	0x2f, 0x76, 0x32, 0x3b, 0x73, 0x64, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_actions_sdk_v2_config_file_proto_rawDescOnce sync.Once
	file_google_actions_sdk_v2_config_file_proto_rawDescData = file_google_actions_sdk_v2_config_file_proto_rawDesc
)

func file_google_actions_sdk_v2_config_file_proto_rawDescGZIP() []byte {
	file_google_actions_sdk_v2_config_file_proto_rawDescOnce.Do(func() {
		file_google_actions_sdk_v2_config_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_actions_sdk_v2_config_file_proto_rawDescData)
	})
	return file_google_actions_sdk_v2_config_file_proto_rawDescData
}

var file_google_actions_sdk_v2_config_file_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_actions_sdk_v2_config_file_proto_goTypes = []interface{}{
	(*ConfigFiles)(nil),                        // 0: google.actions.sdk.v2.ConfigFiles
	(*ConfigFile)(nil),                         // 1: google.actions.sdk.v2.ConfigFile
	(*Manifest)(nil),                           // 2: google.actions.sdk.v2.Manifest
	(*Actions)(nil),                            // 3: google.actions.sdk.v2.Actions
	(*Settings)(nil),                           // 4: google.actions.sdk.v2.Settings
	(*Webhook)(nil),                            // 5: google.actions.sdk.v2.Webhook
	(*interactionmodel.Intent)(nil),            // 6: google.actions.sdk.v2.interactionmodel.Intent
	(*_type.Type)(nil),                         // 7: google.actions.sdk.v2.interactionmodel.type.Type
	(*interactionmodel.EntitySet)(nil),         // 8: google.actions.sdk.v2.interactionmodel.EntitySet
	(*interactionmodel.GlobalIntentEvent)(nil), // 9: google.actions.sdk.v2.interactionmodel.GlobalIntentEvent
	(*interactionmodel.Scene)(nil),             // 10: google.actions.sdk.v2.interactionmodel.Scene
	(*prompt.StaticPrompt)(nil),                // 11: google.actions.sdk.v2.interactionmodel.prompt.StaticPrompt
	(*AccountLinkingSecret)(nil),               // 12: google.actions.sdk.v2.AccountLinkingSecret
	(*structpb.Struct)(nil),                    // 13: google.protobuf.Struct
}
var file_google_actions_sdk_v2_config_file_proto_depIdxs = []int32{
	1,  // 0: google.actions.sdk.v2.ConfigFiles.config_files:type_name -> google.actions.sdk.v2.ConfigFile
	2,  // 1: google.actions.sdk.v2.ConfigFile.manifest:type_name -> google.actions.sdk.v2.Manifest
	3,  // 2: google.actions.sdk.v2.ConfigFile.actions:type_name -> google.actions.sdk.v2.Actions
	4,  // 3: google.actions.sdk.v2.ConfigFile.settings:type_name -> google.actions.sdk.v2.Settings
	5,  // 4: google.actions.sdk.v2.ConfigFile.webhook:type_name -> google.actions.sdk.v2.Webhook
	6,  // 5: google.actions.sdk.v2.ConfigFile.intent:type_name -> google.actions.sdk.v2.interactionmodel.Intent
	7,  // 6: google.actions.sdk.v2.ConfigFile.type:type_name -> google.actions.sdk.v2.interactionmodel.type.Type
	8,  // 7: google.actions.sdk.v2.ConfigFile.entity_set:type_name -> google.actions.sdk.v2.interactionmodel.EntitySet
	9,  // 8: google.actions.sdk.v2.ConfigFile.global_intent_event:type_name -> google.actions.sdk.v2.interactionmodel.GlobalIntentEvent
	10, // 9: google.actions.sdk.v2.ConfigFile.scene:type_name -> google.actions.sdk.v2.interactionmodel.Scene
	11, // 10: google.actions.sdk.v2.ConfigFile.static_prompt:type_name -> google.actions.sdk.v2.interactionmodel.prompt.StaticPrompt
	12, // 11: google.actions.sdk.v2.ConfigFile.account_linking_secret:type_name -> google.actions.sdk.v2.AccountLinkingSecret
	13, // 12: google.actions.sdk.v2.ConfigFile.resource_bundle:type_name -> google.protobuf.Struct
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_google_actions_sdk_v2_config_file_proto_init() }
func file_google_actions_sdk_v2_config_file_proto_init() {
	if File_google_actions_sdk_v2_config_file_proto != nil {
		return
	}
	file_google_actions_sdk_v2_account_linking_secret_proto_init()
	file_google_actions_sdk_v2_action_proto_init()
	file_google_actions_sdk_v2_manifest_proto_init()
	file_google_actions_sdk_v2_settings_proto_init()
	file_google_actions_sdk_v2_webhook_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_actions_sdk_v2_config_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigFiles); i {
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
		file_google_actions_sdk_v2_config_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigFile); i {
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
	file_google_actions_sdk_v2_config_file_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ConfigFile_Manifest)(nil),
		(*ConfigFile_Actions)(nil),
		(*ConfigFile_Settings)(nil),
		(*ConfigFile_Webhook)(nil),
		(*ConfigFile_Intent)(nil),
		(*ConfigFile_Type)(nil),
		(*ConfigFile_EntitySet)(nil),
		(*ConfigFile_GlobalIntentEvent)(nil),
		(*ConfigFile_Scene)(nil),
		(*ConfigFile_StaticPrompt)(nil),
		(*ConfigFile_AccountLinkingSecret)(nil),
		(*ConfigFile_ResourceBundle)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_actions_sdk_v2_config_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_actions_sdk_v2_config_file_proto_goTypes,
		DependencyIndexes: file_google_actions_sdk_v2_config_file_proto_depIdxs,
		MessageInfos:      file_google_actions_sdk_v2_config_file_proto_msgTypes,
	}.Build()
	File_google_actions_sdk_v2_config_file_proto = out.File
	file_google_actions_sdk_v2_config_file_proto_rawDesc = nil
	file_google_actions_sdk_v2_config_file_proto_goTypes = nil
	file_google_actions_sdk_v2_config_file_proto_depIdxs = nil
}