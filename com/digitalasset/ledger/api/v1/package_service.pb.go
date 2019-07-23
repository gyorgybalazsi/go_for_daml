// Code generated by protoc-gen-go. DO NOT EDIT.
// source: com/digitalasset/ledger/api/v1/package_service.proto

package com_digitalasset_ledger_api_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PackageStatus int32

const (
	// The server is not aware of such a package.
	PackageStatus_UNKNOWN PackageStatus = 0
	// The server is able to execute DAML commands operating on this package.
	PackageStatus_REGISTERED PackageStatus = 1
)

var PackageStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "REGISTERED",
}

var PackageStatus_value = map[string]int32{
	"UNKNOWN":    0,
	"REGISTERED": 1,
}

func (x PackageStatus) String() string {
	return proto.EnumName(PackageStatus_name, int32(x))
}

func (PackageStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d92e37ad641cd809, []int{0}
}

type HashFunction int32

const (
	HashFunction_SHA256 HashFunction = 0
)

var HashFunction_name = map[int32]string{
	0: "SHA256",
}

var HashFunction_value = map[string]int32{
	"SHA256": 0,
}

func (x HashFunction) String() string {
	return proto.EnumName(HashFunction_name, int32(x))
}

func (HashFunction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d92e37ad641cd809, []int{1}
}

type ListPackagesRequest struct {
	// Must correspond to the ledger ID reported by the Ledger Identification Service.
	// Required
	LedgerId string `protobuf:"bytes,1,opt,name=ledger_id,json=ledgerId,proto3" json:"ledger_id,omitempty"`
	// Server side tracing will be registered as a child of the submitted context.
	// This field is a future extension point and is currently not supported.
	// Optional
	TraceContext         *TraceContext `protobuf:"bytes,1000,opt,name=trace_context,json=traceContext,proto3" json:"trace_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListPackagesRequest) Reset()         { *m = ListPackagesRequest{} }
func (m *ListPackagesRequest) String() string { return proto.CompactTextString(m) }
func (*ListPackagesRequest) ProtoMessage()    {}
func (*ListPackagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d92e37ad641cd809, []int{0}
}

func (m *ListPackagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPackagesRequest.Unmarshal(m, b)
}
func (m *ListPackagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPackagesRequest.Marshal(b, m, deterministic)
}
func (m *ListPackagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPackagesRequest.Merge(m, src)
}
func (m *ListPackagesRequest) XXX_Size() int {
	return xxx_messageInfo_ListPackagesRequest.Size(m)
}
func (m *ListPackagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPackagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPackagesRequest proto.InternalMessageInfo

func (m *ListPackagesRequest) GetLedgerId() string {
	if m != nil {
		return m.LedgerId
	}
	return ""
}

func (m *ListPackagesRequest) GetTraceContext() *TraceContext {
	if m != nil {
		return m.TraceContext
	}
	return nil
}

type ListPackagesResponse struct {
	// The IDs of all DAML-LF packages supported by the server.
	// Required
	PackageIds           []string `protobuf:"bytes,1,rep,name=package_ids,json=packageIds,proto3" json:"package_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPackagesResponse) Reset()         { *m = ListPackagesResponse{} }
func (m *ListPackagesResponse) String() string { return proto.CompactTextString(m) }
func (*ListPackagesResponse) ProtoMessage()    {}
func (*ListPackagesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d92e37ad641cd809, []int{1}
}

func (m *ListPackagesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPackagesResponse.Unmarshal(m, b)
}
func (m *ListPackagesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPackagesResponse.Marshal(b, m, deterministic)
}
func (m *ListPackagesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPackagesResponse.Merge(m, src)
}
func (m *ListPackagesResponse) XXX_Size() int {
	return xxx_messageInfo_ListPackagesResponse.Size(m)
}
func (m *ListPackagesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPackagesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPackagesResponse proto.InternalMessageInfo

func (m *ListPackagesResponse) GetPackageIds() []string {
	if m != nil {
		return m.PackageIds
	}
	return nil
}

type GetPackageRequest struct {
	// Must correspond to the ledger ID reported by the Ledger Identification Service.
	// Required
	LedgerId string `protobuf:"bytes,1,opt,name=ledger_id,json=ledgerId,proto3" json:"ledger_id,omitempty"`
	// The ID of the requested package.
	// Required
	PackageId string `protobuf:"bytes,2,opt,name=package_id,json=packageId,proto3" json:"package_id,omitempty"`
	// Server side tracing will be registered as a child of the submitted context.
	// This field is a future extension point and is currently not supported.
	// Optional
	TraceContext         *TraceContext `protobuf:"bytes,1000,opt,name=trace_context,json=traceContext,proto3" json:"trace_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetPackageRequest) Reset()         { *m = GetPackageRequest{} }
func (m *GetPackageRequest) String() string { return proto.CompactTextString(m) }
func (*GetPackageRequest) ProtoMessage()    {}
func (*GetPackageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d92e37ad641cd809, []int{2}
}

func (m *GetPackageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPackageRequest.Unmarshal(m, b)
}
func (m *GetPackageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPackageRequest.Marshal(b, m, deterministic)
}
func (m *GetPackageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPackageRequest.Merge(m, src)
}
func (m *GetPackageRequest) XXX_Size() int {
	return xxx_messageInfo_GetPackageRequest.Size(m)
}
func (m *GetPackageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPackageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPackageRequest proto.InternalMessageInfo

func (m *GetPackageRequest) GetLedgerId() string {
	if m != nil {
		return m.LedgerId
	}
	return ""
}

func (m *GetPackageRequest) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *GetPackageRequest) GetTraceContext() *TraceContext {
	if m != nil {
		return m.TraceContext
	}
	return nil
}

type GetPackageResponse struct {
	// The hash function we use to calculate the hash.
	// Required
	HashFunction HashFunction `protobuf:"varint,1,opt,name=hash_function,json=hashFunction,proto3,enum=com.digitalasset.ledger.api.v1.HashFunction" json:"hash_function,omitempty"`
	// Contains a ``daml_lf`` ArchivePayload. See further details in ``daml_lf.proto``.
	// Required
	ArchivePayload []byte `protobuf:"bytes,2,opt,name=archive_payload,json=archivePayload,proto3" json:"archive_payload,omitempty"`
	// The hash of the archive payload, can also used as a ``package_id``.
	// Required
	Hash                 string   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPackageResponse) Reset()         { *m = GetPackageResponse{} }
func (m *GetPackageResponse) String() string { return proto.CompactTextString(m) }
func (*GetPackageResponse) ProtoMessage()    {}
func (*GetPackageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d92e37ad641cd809, []int{3}
}

func (m *GetPackageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPackageResponse.Unmarshal(m, b)
}
func (m *GetPackageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPackageResponse.Marshal(b, m, deterministic)
}
func (m *GetPackageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPackageResponse.Merge(m, src)
}
func (m *GetPackageResponse) XXX_Size() int {
	return xxx_messageInfo_GetPackageResponse.Size(m)
}
func (m *GetPackageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPackageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPackageResponse proto.InternalMessageInfo

func (m *GetPackageResponse) GetHashFunction() HashFunction {
	if m != nil {
		return m.HashFunction
	}
	return HashFunction_SHA256
}

func (m *GetPackageResponse) GetArchivePayload() []byte {
	if m != nil {
		return m.ArchivePayload
	}
	return nil
}

func (m *GetPackageResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type GetPackageStatusRequest struct {
	// Must correspond to the ledger ID reported by the Ledger Identification Service.
	// Required
	LedgerId string `protobuf:"bytes,1,opt,name=ledger_id,json=ledgerId,proto3" json:"ledger_id,omitempty"`
	// The ID of the requested package.
	// Required
	PackageId string `protobuf:"bytes,2,opt,name=package_id,json=packageId,proto3" json:"package_id,omitempty"`
	// Server side tracing will be registered as a child of the submitted context.
	// This field is a future extension point and is currently not supported.
	// Optional
	TraceContext         *TraceContext `protobuf:"bytes,1000,opt,name=trace_context,json=traceContext,proto3" json:"trace_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetPackageStatusRequest) Reset()         { *m = GetPackageStatusRequest{} }
func (m *GetPackageStatusRequest) String() string { return proto.CompactTextString(m) }
func (*GetPackageStatusRequest) ProtoMessage()    {}
func (*GetPackageStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d92e37ad641cd809, []int{4}
}

func (m *GetPackageStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPackageStatusRequest.Unmarshal(m, b)
}
func (m *GetPackageStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPackageStatusRequest.Marshal(b, m, deterministic)
}
func (m *GetPackageStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPackageStatusRequest.Merge(m, src)
}
func (m *GetPackageStatusRequest) XXX_Size() int {
	return xxx_messageInfo_GetPackageStatusRequest.Size(m)
}
func (m *GetPackageStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPackageStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPackageStatusRequest proto.InternalMessageInfo

func (m *GetPackageStatusRequest) GetLedgerId() string {
	if m != nil {
		return m.LedgerId
	}
	return ""
}

func (m *GetPackageStatusRequest) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

func (m *GetPackageStatusRequest) GetTraceContext() *TraceContext {
	if m != nil {
		return m.TraceContext
	}
	return nil
}

type GetPackageStatusResponse struct {
	// The status of the package.
	PackageStatus        PackageStatus `protobuf:"varint,1,opt,name=package_status,json=packageStatus,proto3,enum=com.digitalasset.ledger.api.v1.PackageStatus" json:"package_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetPackageStatusResponse) Reset()         { *m = GetPackageStatusResponse{} }
func (m *GetPackageStatusResponse) String() string { return proto.CompactTextString(m) }
func (*GetPackageStatusResponse) ProtoMessage()    {}
func (*GetPackageStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d92e37ad641cd809, []int{5}
}

func (m *GetPackageStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPackageStatusResponse.Unmarshal(m, b)
}
func (m *GetPackageStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPackageStatusResponse.Marshal(b, m, deterministic)
}
func (m *GetPackageStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPackageStatusResponse.Merge(m, src)
}
func (m *GetPackageStatusResponse) XXX_Size() int {
	return xxx_messageInfo_GetPackageStatusResponse.Size(m)
}
func (m *GetPackageStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPackageStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPackageStatusResponse proto.InternalMessageInfo

func (m *GetPackageStatusResponse) GetPackageStatus() PackageStatus {
	if m != nil {
		return m.PackageStatus
	}
	return PackageStatus_UNKNOWN
}

func init() {
	proto.RegisterEnum("com.digitalasset.ledger.api.v1.PackageStatus", PackageStatus_name, PackageStatus_value)
	proto.RegisterEnum("com.digitalasset.ledger.api.v1.HashFunction", HashFunction_name, HashFunction_value)
	proto.RegisterType((*ListPackagesRequest)(nil), "com.digitalasset.ledger.api.v1.ListPackagesRequest")
	proto.RegisterType((*ListPackagesResponse)(nil), "com.digitalasset.ledger.api.v1.ListPackagesResponse")
	proto.RegisterType((*GetPackageRequest)(nil), "com.digitalasset.ledger.api.v1.GetPackageRequest")
	proto.RegisterType((*GetPackageResponse)(nil), "com.digitalasset.ledger.api.v1.GetPackageResponse")
	proto.RegisterType((*GetPackageStatusRequest)(nil), "com.digitalasset.ledger.api.v1.GetPackageStatusRequest")
	proto.RegisterType((*GetPackageStatusResponse)(nil), "com.digitalasset.ledger.api.v1.GetPackageStatusResponse")
}

func init() {
	proto.RegisterFile("com/digitalasset/ledger/api/v1/package_service.proto", fileDescriptor_d92e37ad641cd809)
}

var fileDescriptor_d92e37ad641cd809 = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xed, 0x12, 0xd4, 0x92, 0x89, 0x63, 0xc2, 0x82, 0x84, 0x65, 0x04, 0x44, 0xb9, 0x10, 0x55,
	0xc5, 0x56, 0xdc, 0x42, 0x11, 0x37, 0x5a, 0x42, 0x1b, 0x81, 0xd2, 0xb2, 0x09, 0xe2, 0x68, 0x2d,
	0xf6, 0x12, 0xaf, 0x48, 0x63, 0xe3, 0xdd, 0x44, 0xf4, 0x03, 0x40, 0xfc, 0x07, 0x5c, 0xf8, 0x33,
	0x3e, 0x03, 0xc5, 0xbb, 0x69, 0x6c, 0x8a, 0x70, 0x72, 0x40, 0xe2, 0x96, 0x99, 0xc9, 0x7b, 0xf3,
	0xe6, 0x8d, 0x67, 0x61, 0x2f, 0x88, 0xcf, 0xdc, 0x90, 0x8f, 0xb8, 0xa4, 0x63, 0x2a, 0x04, 0x93,
	0xee, 0x98, 0x85, 0x23, 0x96, 0xba, 0x34, 0xe1, 0xee, 0xac, 0xe3, 0x26, 0x34, 0xf8, 0x40, 0x47,
	0xcc, 0x17, 0x2c, 0x9d, 0xf1, 0x80, 0x39, 0x49, 0x1a, 0xcb, 0x18, 0xdf, 0x0b, 0xe2, 0x33, 0x27,
	0x8f, 0x72, 0x14, 0xca, 0xa1, 0x09, 0x77, 0x66, 0x1d, 0xdb, 0x2b, 0x61, 0x95, 0x29, 0x0d, 0x98,
	0x1f, 0xc4, 0x13, 0xc9, 0x3e, 0x49, 0xc5, 0xd9, 0xfa, 0x82, 0xe0, 0xe6, 0x2b, 0x2e, 0xe4, 0xa9,
	0xea, 0x28, 0x08, 0xfb, 0x38, 0x65, 0x42, 0xe2, 0x3b, 0x50, 0x55, 0x60, 0x9f, 0x87, 0x16, 0x6a,
	0xa2, 0x76, 0x95, 0x5c, 0x53, 0x89, 0x5e, 0x88, 0x09, 0xd4, 0x0b, 0x5c, 0xd6, 0xcf, 0xad, 0x26,
	0x6a, 0xd7, 0xbc, 0x1d, 0xe7, 0xef, 0x0a, 0x9d, 0xe1, 0x1c, 0x75, 0xa8, 0x40, 0xc4, 0x90, 0xb9,
	0xa8, 0xb5, 0x0f, 0xb7, 0x8a, 0x3a, 0x44, 0x12, 0x4f, 0x04, 0xc3, 0xf7, 0xa1, 0xb6, 0x70, 0x83,
	0x87, 0xc2, 0x42, 0xcd, 0x4a, 0xbb, 0x4a, 0x40, 0xa7, 0x7a, 0xa1, 0x68, 0x7d, 0x43, 0x70, 0xe3,
	0x88, 0x2d, 0x80, 0x2b, 0xe9, 0xbf, 0x0b, 0xb0, 0xe4, 0xb4, 0xae, 0x64, 0xd5, 0xea, 0x05, 0xe5,
	0x3f, 0x19, 0xef, 0x3b, 0x02, 0x9c, 0x57, 0xa9, 0xa7, 0x7b, 0x0d, 0xf5, 0x88, 0x8a, 0xc8, 0x7f,
	0x3f, 0x9d, 0x04, 0x92, 0xc7, 0x93, 0x4c, 0xaa, 0x59, 0xde, 0xe9, 0x98, 0x8a, 0xe8, 0x85, 0xc6,
	0x10, 0x23, 0xca, 0x45, 0xf8, 0x01, 0x5c, 0xa7, 0x69, 0x10, 0xf1, 0x19, 0xf3, 0x13, 0x7a, 0x3e,
	0x8e, 0xa9, 0x9a, 0xd0, 0x20, 0xa6, 0x4e, 0x9f, 0xaa, 0x2c, 0xc6, 0x70, 0x75, 0x0e, 0xb4, 0x2a,
	0xd9, 0xfc, 0xd9, 0xef, 0xd6, 0x0f, 0x04, 0xb7, 0x97, 0x32, 0x07, 0x92, 0xca, 0xa9, 0xf8, 0x5f,
	0x2d, 0x4d, 0xc0, 0xba, 0x2c, 0x55, 0xfb, 0x3a, 0x04, 0xf3, 0xe2, 0x86, 0xb2, 0x8a, 0x36, 0xf6,
	0x61, 0x59, 0xbf, 0x22, 0x5d, 0x3d, 0xc9, 0x87, 0xdb, 0x3b, 0x50, 0x2f, 0xd4, 0x71, 0x0d, 0xb6,
	0xde, 0xf4, 0x5f, 0xf6, 0x4f, 0xde, 0xf6, 0x1b, 0x1b, 0xd8, 0x04, 0x20, 0xdd, 0xa3, 0xde, 0x60,
	0xd8, 0x25, 0xdd, 0xe7, 0x0d, 0xb4, 0x6d, 0x83, 0x91, 0x5f, 0x13, 0x06, 0xd8, 0x1c, 0x1c, 0x3f,
	0xf3, 0x1e, 0x3d, 0x6e, 0x6c, 0x78, 0x5f, 0x2b, 0x60, 0x2e, 0xa8, 0xd4, 0x8d, 0xe3, 0x73, 0x30,
	0xf2, 0x07, 0x80, 0x77, 0xcb, 0xa4, 0xfe, 0xe1, 0x6c, 0xed, 0xbd, 0xf5, 0x40, 0xda, 0x2d, 0x01,
	0xb0, 0x74, 0x12, 0x77, 0xca, 0x38, 0x2e, 0x5d, 0x9b, 0xed, 0xad, 0x03, 0xd1, 0x4d, 0x3f, 0x23,
	0x68, 0xfc, 0xbe, 0x3f, 0xbc, 0xbf, 0x3a, 0x51, 0xe1, 0xe3, 0xb4, 0x9f, 0xac, 0x0f, 0x54, 0x3a,
	0x0e, 0x9e, 0x42, 0xc9, 0xbb, 0x7a, 0x60, 0x15, 0x37, 0x75, 0x32, 0x95, 0x2c, 0x3d, 0x9c, 0xff,
	0xf3, 0xdd, 0x66, 0xf6, 0x88, 0xee, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x28, 0xae, 0xc2, 0x1e,
	0xd0, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PackageServiceClient is the client API for PackageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PackageServiceClient interface {
	// Returns the identifiers of all supported packages.
	ListPackages(ctx context.Context, in *ListPackagesRequest, opts ...grpc.CallOption) (*ListPackagesResponse, error)
	// Returns the contents of a single package, or a ``NOT_FOUND`` error if the requested package is unknown.
	GetPackage(ctx context.Context, in *GetPackageRequest, opts ...grpc.CallOption) (*GetPackageResponse, error)
	// Returns the status of a single package.
	GetPackageStatus(ctx context.Context, in *GetPackageStatusRequest, opts ...grpc.CallOption) (*GetPackageStatusResponse, error)
}

type packageServiceClient struct {
	cc *grpc.ClientConn
}

func NewPackageServiceClient(cc *grpc.ClientConn) PackageServiceClient {
	return &packageServiceClient{cc}
}

func (c *packageServiceClient) ListPackages(ctx context.Context, in *ListPackagesRequest, opts ...grpc.CallOption) (*ListPackagesResponse, error) {
	out := new(ListPackagesResponse)
	err := c.cc.Invoke(ctx, "/com.digitalasset.ledger.api.v1.PackageService/ListPackages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageServiceClient) GetPackage(ctx context.Context, in *GetPackageRequest, opts ...grpc.CallOption) (*GetPackageResponse, error) {
	out := new(GetPackageResponse)
	err := c.cc.Invoke(ctx, "/com.digitalasset.ledger.api.v1.PackageService/GetPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageServiceClient) GetPackageStatus(ctx context.Context, in *GetPackageStatusRequest, opts ...grpc.CallOption) (*GetPackageStatusResponse, error) {
	out := new(GetPackageStatusResponse)
	err := c.cc.Invoke(ctx, "/com.digitalasset.ledger.api.v1.PackageService/GetPackageStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PackageServiceServer is the server API for PackageService service.
type PackageServiceServer interface {
	// Returns the identifiers of all supported packages.
	ListPackages(context.Context, *ListPackagesRequest) (*ListPackagesResponse, error)
	// Returns the contents of a single package, or a ``NOT_FOUND`` error if the requested package is unknown.
	GetPackage(context.Context, *GetPackageRequest) (*GetPackageResponse, error)
	// Returns the status of a single package.
	GetPackageStatus(context.Context, *GetPackageStatusRequest) (*GetPackageStatusResponse, error)
}

// UnimplementedPackageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPackageServiceServer struct {
}

func (*UnimplementedPackageServiceServer) ListPackages(ctx context.Context, req *ListPackagesRequest) (*ListPackagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPackages not implemented")
}
func (*UnimplementedPackageServiceServer) GetPackage(ctx context.Context, req *GetPackageRequest) (*GetPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPackage not implemented")
}
func (*UnimplementedPackageServiceServer) GetPackageStatus(ctx context.Context, req *GetPackageStatusRequest) (*GetPackageStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPackageStatus not implemented")
}

func RegisterPackageServiceServer(s *grpc.Server, srv PackageServiceServer) {
	s.RegisterService(&_PackageService_serviceDesc, srv)
}

func _PackageService_ListPackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPackagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).ListPackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.ledger.api.v1.PackageService/ListPackages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).ListPackages(ctx, req.(*ListPackagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageService_GetPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).GetPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.ledger.api.v1.PackageService/GetPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).GetPackage(ctx, req.(*GetPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageService_GetPackageStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPackageStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).GetPackageStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.digitalasset.ledger.api.v1.PackageService/GetPackageStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).GetPackageStatus(ctx, req.(*GetPackageStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PackageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.digitalasset.ledger.api.v1.PackageService",
	HandlerType: (*PackageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPackages",
			Handler:    _PackageService_ListPackages_Handler,
		},
		{
			MethodName: "GetPackage",
			Handler:    _PackageService_GetPackage_Handler,
		},
		{
			MethodName: "GetPackageStatus",
			Handler:    _PackageService_GetPackageStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/digitalasset/ledger/api/v1/package_service.proto",
}
