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
// source: google/bigtable/v2/request_stats.proto

package bigtable

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ReadIterationStats captures information about the iteration of rows or cells
// over the course of a read, e.g. how many results were scanned in a read
// operation versus the results returned.
type ReadIterationStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The rows seen (scanned) as part of the request. This includes the count of
	// rows returned, as captured below.
	RowsSeenCount int64 `protobuf:"varint,1,opt,name=rows_seen_count,json=rowsSeenCount,proto3" json:"rows_seen_count,omitempty"`
	// The rows returned as part of the request.
	RowsReturnedCount int64 `protobuf:"varint,2,opt,name=rows_returned_count,json=rowsReturnedCount,proto3" json:"rows_returned_count,omitempty"`
	// The cells seen (scanned) as part of the request. This includes the count of
	// cells returned, as captured below.
	CellsSeenCount int64 `protobuf:"varint,3,opt,name=cells_seen_count,json=cellsSeenCount,proto3" json:"cells_seen_count,omitempty"`
	// The cells returned as part of the request.
	CellsReturnedCount int64 `protobuf:"varint,4,opt,name=cells_returned_count,json=cellsReturnedCount,proto3" json:"cells_returned_count,omitempty"`
}

func (x *ReadIterationStats) Reset() {
	*x = ReadIterationStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_bigtable_v2_request_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadIterationStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadIterationStats) ProtoMessage() {}

func (x *ReadIterationStats) ProtoReflect() protoreflect.Message {
	mi := &file_google_bigtable_v2_request_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadIterationStats.ProtoReflect.Descriptor instead.
func (*ReadIterationStats) Descriptor() ([]byte, []int) {
	return file_google_bigtable_v2_request_stats_proto_rawDescGZIP(), []int{0}
}

func (x *ReadIterationStats) GetRowsSeenCount() int64 {
	if x != nil {
		return x.RowsSeenCount
	}
	return 0
}

func (x *ReadIterationStats) GetRowsReturnedCount() int64 {
	if x != nil {
		return x.RowsReturnedCount
	}
	return 0
}

func (x *ReadIterationStats) GetCellsSeenCount() int64 {
	if x != nil {
		return x.CellsSeenCount
	}
	return 0
}

func (x *ReadIterationStats) GetCellsReturnedCount() int64 {
	if x != nil {
		return x.CellsReturnedCount
	}
	return 0
}

// RequestLatencyStats provides a measurement of the latency of the request as
// it interacts with different systems over its lifetime, e.g. how long the
// request took to execute within a frontend server.
type RequestLatencyStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The latency measured by the frontend server handling this request, from
	// when the request was received, to when this value is sent back in the
	// response. For more context on the component that is measuring this latency,
	// see: https://cloud.google.com/bigtable/docs/overview
	//
	// Note: This value may be slightly shorter than the value reported into
	// aggregate latency metrics in Monitoring for this request
	// (https://cloud.google.com/bigtable/docs/monitoring-instance) as this value
	// needs to be sent in the response before the latency measurement including
	// that transmission is finalized.
	//
	// Note: This value includes the end-to-end latency of contacting nodes in
	// the targeted cluster, e.g. measuring from when the first byte arrives at
	// the frontend server, to when this value is sent back as the last value in
	// the response, including any latency incurred by contacting nodes, waiting
	// for results from nodes, and finally sending results from nodes back to the
	// caller.
	FrontendServerLatency *durationpb.Duration `protobuf:"bytes,1,opt,name=frontend_server_latency,json=frontendServerLatency,proto3" json:"frontend_server_latency,omitempty"`
}

func (x *RequestLatencyStats) Reset() {
	*x = RequestLatencyStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_bigtable_v2_request_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestLatencyStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestLatencyStats) ProtoMessage() {}

func (x *RequestLatencyStats) ProtoReflect() protoreflect.Message {
	mi := &file_google_bigtable_v2_request_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestLatencyStats.ProtoReflect.Descriptor instead.
func (*RequestLatencyStats) Descriptor() ([]byte, []int) {
	return file_google_bigtable_v2_request_stats_proto_rawDescGZIP(), []int{1}
}

func (x *RequestLatencyStats) GetFrontendServerLatency() *durationpb.Duration {
	if x != nil {
		return x.FrontendServerLatency
	}
	return nil
}

// FullReadStatsView captures all known information about a read.
type FullReadStatsView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Iteration stats describe how efficient the read is, e.g. comparing
	// rows seen vs. rows returned or cells seen vs cells returned can provide an
	// indication of read efficiency (the higher the ratio of seen to retuned the
	// better).
	ReadIterationStats *ReadIterationStats `protobuf:"bytes,1,opt,name=read_iteration_stats,json=readIterationStats,proto3" json:"read_iteration_stats,omitempty"`
	// Request latency stats describe the time taken to complete a request, from
	// the server side.
	RequestLatencyStats *RequestLatencyStats `protobuf:"bytes,2,opt,name=request_latency_stats,json=requestLatencyStats,proto3" json:"request_latency_stats,omitempty"`
}

func (x *FullReadStatsView) Reset() {
	*x = FullReadStatsView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_bigtable_v2_request_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FullReadStatsView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FullReadStatsView) ProtoMessage() {}

func (x *FullReadStatsView) ProtoReflect() protoreflect.Message {
	mi := &file_google_bigtable_v2_request_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FullReadStatsView.ProtoReflect.Descriptor instead.
func (*FullReadStatsView) Descriptor() ([]byte, []int) {
	return file_google_bigtable_v2_request_stats_proto_rawDescGZIP(), []int{2}
}

func (x *FullReadStatsView) GetReadIterationStats() *ReadIterationStats {
	if x != nil {
		return x.ReadIterationStats
	}
	return nil
}

func (x *FullReadStatsView) GetRequestLatencyStats() *RequestLatencyStats {
	if x != nil {
		return x.RequestLatencyStats
	}
	return nil
}

// RequestStats is the container for additional information pertaining to a
// single request, helpful for evaluating the performance of the sent request.
// Currently, there are the following supported methods:
//   - google.bigtable.v2.ReadRows
type RequestStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Information pertaining to each request type received. The type is chosen
	// based on the requested view.
	//
	// See the messages above for additional context.
	//
	// Types that are assignable to StatsView:
	//
	//	*RequestStats_FullReadStatsView
	StatsView isRequestStats_StatsView `protobuf_oneof:"stats_view"`
}

func (x *RequestStats) Reset() {
	*x = RequestStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_bigtable_v2_request_stats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestStats) ProtoMessage() {}

func (x *RequestStats) ProtoReflect() protoreflect.Message {
	mi := &file_google_bigtable_v2_request_stats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestStats.ProtoReflect.Descriptor instead.
func (*RequestStats) Descriptor() ([]byte, []int) {
	return file_google_bigtable_v2_request_stats_proto_rawDescGZIP(), []int{3}
}

func (m *RequestStats) GetStatsView() isRequestStats_StatsView {
	if m != nil {
		return m.StatsView
	}
	return nil
}

func (x *RequestStats) GetFullReadStatsView() *FullReadStatsView {
	if x, ok := x.GetStatsView().(*RequestStats_FullReadStatsView); ok {
		return x.FullReadStatsView
	}
	return nil
}

type isRequestStats_StatsView interface {
	isRequestStats_StatsView()
}

type RequestStats_FullReadStatsView struct {
	// Available with the ReadRowsRequest.RequestStatsView.REQUEST_STATS_FULL
	// view, see package google.bigtable.v2.
	FullReadStatsView *FullReadStatsView `protobuf:"bytes,1,opt,name=full_read_stats_view,json=fullReadStatsView,proto3,oneof"`
}

func (*RequestStats_FullReadStatsView) isRequestStats_StatsView() {}

var File_google_bigtable_v2_request_stats_proto protoreflect.FileDescriptor

var file_google_bigtable_v2_request_stats_proto_rawDesc = []byte{
	0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x32, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x01, 0x0a,
	0x12, 0x52, 0x65, 0x61, 0x64, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x6f, 0x77, 0x73, 0x5f, 0x73, 0x65, 0x65, 0x6e,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x72, 0x6f,
	0x77, 0x73, 0x53, 0x65, 0x65, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x72,
	0x6f, 0x77, 0x73, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x72, 0x6f, 0x77, 0x73, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x63,
	0x65, 0x6c, 0x6c, 0x73, 0x5f, 0x73, 0x65, 0x65, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x53, 0x65, 0x65, 0x6e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x5f, 0x72,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x12, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x68, 0x0a, 0x13, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x51,
	0x0a, 0x17, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x15, 0x66, 0x72, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63,
	0x79, 0x22, 0xca, 0x01, 0x0a, 0x11, 0x46, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x56, 0x69, 0x65, 0x77, 0x12, 0x58, 0x0a, 0x14, 0x72, 0x65, 0x61, 0x64, 0x5f,
	0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62,
	0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x49,
	0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x12, 0x72,
	0x65, 0x61, 0x64, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x5b, 0x0a, 0x15, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x61, 0x74,
	0x65, 0x6e, 0x63, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x61, 0x74,
	0x65, 0x6e, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x22, 0x76,
	0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x58,
	0x0a, 0x14, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76,
	0x32, 0x2e, 0x46, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x73, 0x56,
	0x69, 0x65, 0x77, 0x48, 0x00, 0x52, 0x11, 0x66, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x56, 0x69, 0x65, 0x77, 0x42, 0x0c, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x42, 0xbd, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76,
	0x32, 0x42, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x62, 0x69,
	0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x32, 0x3b, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0xaa, 0x02, 0x18, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x42, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x18,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x42, 0x69, 0x67,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x5c, 0x56, 0x32, 0xea, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x42, 0x69, 0x67, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_bigtable_v2_request_stats_proto_rawDescOnce sync.Once
	file_google_bigtable_v2_request_stats_proto_rawDescData = file_google_bigtable_v2_request_stats_proto_rawDesc
)

func file_google_bigtable_v2_request_stats_proto_rawDescGZIP() []byte {
	file_google_bigtable_v2_request_stats_proto_rawDescOnce.Do(func() {
		file_google_bigtable_v2_request_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_bigtable_v2_request_stats_proto_rawDescData)
	})
	return file_google_bigtable_v2_request_stats_proto_rawDescData
}

var file_google_bigtable_v2_request_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_bigtable_v2_request_stats_proto_goTypes = []interface{}{
	(*ReadIterationStats)(nil),  // 0: google.bigtable.v2.ReadIterationStats
	(*RequestLatencyStats)(nil), // 1: google.bigtable.v2.RequestLatencyStats
	(*FullReadStatsView)(nil),   // 2: google.bigtable.v2.FullReadStatsView
	(*RequestStats)(nil),        // 3: google.bigtable.v2.RequestStats
	(*durationpb.Duration)(nil), // 4: google.protobuf.Duration
}
var file_google_bigtable_v2_request_stats_proto_depIdxs = []int32{
	4, // 0: google.bigtable.v2.RequestLatencyStats.frontend_server_latency:type_name -> google.protobuf.Duration
	0, // 1: google.bigtable.v2.FullReadStatsView.read_iteration_stats:type_name -> google.bigtable.v2.ReadIterationStats
	1, // 2: google.bigtable.v2.FullReadStatsView.request_latency_stats:type_name -> google.bigtable.v2.RequestLatencyStats
	2, // 3: google.bigtable.v2.RequestStats.full_read_stats_view:type_name -> google.bigtable.v2.FullReadStatsView
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_bigtable_v2_request_stats_proto_init() }
func file_google_bigtable_v2_request_stats_proto_init() {
	if File_google_bigtable_v2_request_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_bigtable_v2_request_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadIterationStats); i {
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
		file_google_bigtable_v2_request_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestLatencyStats); i {
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
		file_google_bigtable_v2_request_stats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FullReadStatsView); i {
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
		file_google_bigtable_v2_request_stats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestStats); i {
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
	file_google_bigtable_v2_request_stats_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*RequestStats_FullReadStatsView)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_bigtable_v2_request_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_bigtable_v2_request_stats_proto_goTypes,
		DependencyIndexes: file_google_bigtable_v2_request_stats_proto_depIdxs,
		MessageInfos:      file_google_bigtable_v2_request_stats_proto_msgTypes,
	}.Build()
	File_google_bigtable_v2_request_stats_proto = out.File
	file_google_bigtable_v2_request_stats_proto_rawDesc = nil
	file_google_bigtable_v2_request_stats_proto_goTypes = nil
	file_google_bigtable_v2_request_stats_proto_depIdxs = nil
}