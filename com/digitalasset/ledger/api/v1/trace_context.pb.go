// Code generated by protoc-gen-go. DO NOT EDIT.
// source: com/digitalasset/ledger/api/v1/trace_context.proto

package com_digitalasset_ledger_api_v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protobuf "google/protobuf"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Data structure to propagate Zipkin trace information.
// See https://github.com/openzipkin/b3-propagation
// Trace identifiers are 64 or 128-bit, but all span identifiers within a trace are 64-bit. All identifiers are opaque.
type TraceContext struct {
	// If present, this is the high 64 bits of the 128-bit identifier. Otherwise the trace ID is 64 bits long.
	TraceIdHigh uint64 `protobuf:"varint,1,opt,name=trace_id_high,json=traceIdHigh,proto3" json:"trace_id_high,omitempty"`
	// The TraceId is 64 or 128-bit in length and indicates the overall ID of the trace. Every span in a trace shares this ID.
	TraceId uint64 `protobuf:"varint,2,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// The SpanId is 64-bit in length and indicates the position of the current operation in the trace tree.
	// The value should not be interpreted: it may or may not be derived from the value of the TraceId.
	SpanId uint64 `protobuf:"varint,3,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	// The ParentSpanId is 64-bit in length and indicates the position of the parent operation in the trace tree.
	// When the span is the root of the trace tree, the ParentSpanId is absent.
	ParentSpanId *protobuf.UInt64Value `protobuf:"bytes,4,opt,name=parent_span_id,json=parentSpanId,proto3" json:"parent_span_id,omitempty"`
	// When the sampled decision is accept, report this span to the tracing system. When it is reject, do not.
	// When B3 attributes are sent without a sampled decision, the receiver should make one.
	// Once the sampling decision is made, the same value should be consistently sent downstream.
	Sampled              bool     `protobuf:"varint,5,opt,name=sampled,proto3" json:"sampled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TraceContext) Reset()         { *m = TraceContext{} }
func (m *TraceContext) String() string { return proto.CompactTextString(m) }
func (*TraceContext) ProtoMessage()    {}
func (*TraceContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fe4bcba0d124a11, []int{0}
}

func (m *TraceContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceContext.Unmarshal(m, b)
}
func (m *TraceContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceContext.Marshal(b, m, deterministic)
}
func (m *TraceContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceContext.Merge(m, src)
}
func (m *TraceContext) XXX_Size() int {
	return xxx_messageInfo_TraceContext.Size(m)
}
func (m *TraceContext) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceContext.DiscardUnknown(m)
}

var xxx_messageInfo_TraceContext proto.InternalMessageInfo

func (m *TraceContext) GetTraceIdHigh() uint64 {
	if m != nil {
		return m.TraceIdHigh
	}
	return 0
}

func (m *TraceContext) GetTraceId() uint64 {
	if m != nil {
		return m.TraceId
	}
	return 0
}

func (m *TraceContext) GetSpanId() uint64 {
	if m != nil {
		return m.SpanId
	}
	return 0
}

func (m *TraceContext) GetParentSpanId() *protobuf.UInt64Value {
	if m != nil {
		return m.ParentSpanId
	}
	return nil
}

func (m *TraceContext) GetSampled() bool {
	if m != nil {
		return m.Sampled
	}
	return false
}

func init() {
	proto.RegisterType((*TraceContext)(nil), "com.digitalasset.ledger.api.v1.TraceContext")
}

func init() {
	proto.RegisterFile("com/digitalasset/ledger/api/v1/trace_context.proto", fileDescriptor_2fe4bcba0d124a11)
}

var fileDescriptor_2fe4bcba0d124a11 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x86, 0xe5, 0x7e, 0xfd, 0x9a, 0xca, 0x2d, 0x0c, 0x19, 0x50, 0xc4, 0x4f, 0x15, 0x31, 0xa0,
	0x4c, 0x36, 0x2d, 0x08, 0x31, 0xa7, 0x0b, 0x99, 0x90, 0xc2, 0xcf, 0x1a, 0xb9, 0xf1, 0xc1, 0xb1,
	0x94, 0xc4, 0x96, 0xed, 0x14, 0x2e, 0x93, 0x4b, 0x42, 0x89, 0x49, 0xd5, 0x8d, 0xd1, 0x7e, 0x9f,
	0xe7, 0x3d, 0xc7, 0xc6, 0x9b, 0x52, 0x35, 0x94, 0x4b, 0x21, 0x1d, 0xab, 0x99, 0xb5, 0xe0, 0x68,
	0x0d, 0x5c, 0x80, 0xa1, 0x4c, 0x4b, 0xba, 0x5f, 0x53, 0x67, 0x58, 0x09, 0x45, 0xa9, 0x5a, 0x07,
	0x5f, 0x8e, 0x68, 0xa3, 0x9c, 0x0a, 0x57, 0xa5, 0x6a, 0xc8, 0xb1, 0x43, 0xbc, 0x43, 0x98, 0x96,
	0x64, 0xbf, 0x3e, 0x5f, 0x09, 0xa5, 0x44, 0x0d, 0x74, 0xa0, 0x77, 0xdd, 0x07, 0xfd, 0x34, 0x4c,
	0x6b, 0x30, 0xd6, 0xfb, 0xd7, 0xdf, 0x08, 0x2f, 0x5f, 0xfb, 0xde, 0xad, 0xaf, 0x0d, 0x6f, 0xf0,
	0x89, 0x9f, 0x23, 0x79, 0x51, 0x49, 0x51, 0x45, 0x28, 0x46, 0xc9, 0x34, 0x9d, 0xdc, 0xa2, 0x7c,
	0x31, 0x04, 0x19, 0x7f, 0x92, 0xa2, 0x0a, 0xaf, 0xf0, 0x7c, 0xe4, 0xa2, 0xc9, 0x01, 0x09, 0x7e,
	0x91, 0xf0, 0x02, 0x07, 0x56, 0xb3, 0xb6, 0x4f, 0xff, 0x1d, 0xd2, 0x59, 0x7f, 0x95, 0xf1, 0x30,
	0xc5, 0xa7, 0x9a, 0x19, 0x68, 0x5d, 0x31, 0x32, 0xd3, 0x18, 0x25, 0x8b, 0xcd, 0x25, 0xf1, 0xdb,
	0x92, 0x71, 0x5b, 0xf2, 0x96, 0xb5, 0xee, 0xe1, 0xfe, 0x9d, 0xd5, 0x1d, 0xe4, 0x4b, 0xef, 0xbc,
	0xf8, 0x8e, 0x08, 0x07, 0x96, 0x35, 0xba, 0x06, 0x1e, 0xfd, 0x8f, 0x51, 0x32, 0xcf, 0xc7, 0x63,
	0xfa, 0x88, 0xff, 0xf8, 0x94, 0xf4, 0xec, 0xf8, 0xc5, 0xcf, 0x9d, 0x03, 0xb3, 0xed, 0xb9, 0xdd,
	0x6c, 0x98, 0x7b, 0xf7, 0x13, 0x00, 0x00, 0xff, 0xff, 0xab, 0x83, 0x7e, 0x87, 0x89, 0x01, 0x00,
	0x00,
}