// Code generated by protoc-gen-go. DO NOT EDIT.
// source: egress.proto

/*
Package logcache_v1 is a generated protocol buffer package.

It is generated from these files:
	egress.proto
	group_reader.proto
	ingress.proto
	orchestration.proto

It has these top-level messages:
	ReadRequest
	ReadResponse
	MetaRequest
	MetaResponse
	MetaInfo
	AddToGroupRequest
	AddToGroupResponse
	RemoveFromGroupRequest
	RemoveFromGroupResponse
	GroupReadRequest
	GroupReadResponse
	GroupRequest
	GroupResponse
	SendRequest
	SendResponse
	Range
	Ranges
	AddRangeRequest
	AddRangeResponse
	ListRangesRequest
	ListRangesResponse
	SetRangesRequest
	SetRangesResponse
*/
package logcache_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import loggregator_v2 "code.cloudfoundry.org/go-loggregator/rpc/loggregator_v2"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EnvelopeType int32

const (
	EnvelopeType_ANY     EnvelopeType = 0
	EnvelopeType_LOG     EnvelopeType = 1
	EnvelopeType_COUNTER EnvelopeType = 2
	EnvelopeType_GAUGE   EnvelopeType = 3
	EnvelopeType_TIMER   EnvelopeType = 4
	EnvelopeType_EVENT   EnvelopeType = 5
)

var EnvelopeType_name = map[int32]string{
	0: "ANY",
	1: "LOG",
	2: "COUNTER",
	3: "GAUGE",
	4: "TIMER",
	5: "EVENT",
}
var EnvelopeType_value = map[string]int32{
	"ANY":     0,
	"LOG":     1,
	"COUNTER": 2,
	"GAUGE":   3,
	"TIMER":   4,
	"EVENT":   5,
}

func (x EnvelopeType) String() string {
	return proto.EnumName(EnvelopeType_name, int32(x))
}
func (EnvelopeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ReadRequest struct {
	SourceId      string         `protobuf:"bytes,1,opt,name=source_id,json=sourceId" json:"source_id,omitempty"`
	StartTime     int64          `protobuf:"varint,2,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	EndTime       int64          `protobuf:"varint,3,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	Limit         int64          `protobuf:"varint,4,opt,name=limit" json:"limit,omitempty"`
	EnvelopeTypes []EnvelopeType `protobuf:"varint,5,rep,packed,name=envelope_types,json=envelopeTypes,enum=logcache.v1.EnvelopeType" json:"envelope_types,omitempty"`
	Descending    bool           `protobuf:"varint,6,opt,name=descending" json:"descending,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ReadRequest) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

func (m *ReadRequest) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *ReadRequest) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *ReadRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ReadRequest) GetEnvelopeTypes() []EnvelopeType {
	if m != nil {
		return m.EnvelopeTypes
	}
	return nil
}

func (m *ReadRequest) GetDescending() bool {
	if m != nil {
		return m.Descending
	}
	return false
}

type ReadResponse struct {
	Envelopes *loggregator_v2.EnvelopeBatch `protobuf:"bytes,1,opt,name=envelopes" json:"envelopes,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReadResponse) GetEnvelopes() *loggregator_v2.EnvelopeBatch {
	if m != nil {
		return m.Envelopes
	}
	return nil
}

type MetaRequest struct {
	LocalOnly bool `protobuf:"varint,1,opt,name=local_only,json=localOnly" json:"local_only,omitempty"`
}

func (m *MetaRequest) Reset()                    { *m = MetaRequest{} }
func (m *MetaRequest) String() string            { return proto.CompactTextString(m) }
func (*MetaRequest) ProtoMessage()               {}
func (*MetaRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MetaRequest) GetLocalOnly() bool {
	if m != nil {
		return m.LocalOnly
	}
	return false
}

type MetaResponse struct {
	Meta map[string]*MetaInfo `protobuf:"bytes,1,rep,name=meta" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *MetaResponse) Reset()                    { *m = MetaResponse{} }
func (m *MetaResponse) String() string            { return proto.CompactTextString(m) }
func (*MetaResponse) ProtoMessage()               {}
func (*MetaResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MetaResponse) GetMeta() map[string]*MetaInfo {
	if m != nil {
		return m.Meta
	}
	return nil
}

type MetaInfo struct {
	Count           int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	Expired         int64 `protobuf:"varint,2,opt,name=expired" json:"expired,omitempty"`
	OldestTimestamp int64 `protobuf:"varint,3,opt,name=oldest_timestamp,json=oldestTimestamp" json:"oldest_timestamp,omitempty"`
	NewestTimestamp int64 `protobuf:"varint,4,opt,name=newest_timestamp,json=newestTimestamp" json:"newest_timestamp,omitempty"`
}

func (m *MetaInfo) Reset()                    { *m = MetaInfo{} }
func (m *MetaInfo) String() string            { return proto.CompactTextString(m) }
func (*MetaInfo) ProtoMessage()               {}
func (*MetaInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MetaInfo) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *MetaInfo) GetExpired() int64 {
	if m != nil {
		return m.Expired
	}
	return 0
}

func (m *MetaInfo) GetOldestTimestamp() int64 {
	if m != nil {
		return m.OldestTimestamp
	}
	return 0
}

func (m *MetaInfo) GetNewestTimestamp() int64 {
	if m != nil {
		return m.NewestTimestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*ReadRequest)(nil), "logcache.v1.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "logcache.v1.ReadResponse")
	proto.RegisterType((*MetaRequest)(nil), "logcache.v1.MetaRequest")
	proto.RegisterType((*MetaResponse)(nil), "logcache.v1.MetaResponse")
	proto.RegisterType((*MetaInfo)(nil), "logcache.v1.MetaInfo")
	proto.RegisterEnum("logcache.v1.EnvelopeType", EnvelopeType_name, EnvelopeType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Egress service

type EgressClient interface {
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Meta(ctx context.Context, in *MetaRequest, opts ...grpc.CallOption) (*MetaResponse, error)
}

type egressClient struct {
	cc *grpc.ClientConn
}

func NewEgressClient(cc *grpc.ClientConn) EgressClient {
	return &egressClient{cc}
}

func (c *egressClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := grpc.Invoke(ctx, "/logcache.v1.Egress/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *egressClient) Meta(ctx context.Context, in *MetaRequest, opts ...grpc.CallOption) (*MetaResponse, error) {
	out := new(MetaResponse)
	err := grpc.Invoke(ctx, "/logcache.v1.Egress/Meta", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Egress service

type EgressServer interface {
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Meta(context.Context, *MetaRequest) (*MetaResponse, error)
}

func RegisterEgressServer(s *grpc.Server, srv EgressServer) {
	s.RegisterService(&_Egress_serviceDesc, srv)
}

func _Egress_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EgressServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logcache.v1.Egress/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EgressServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Egress_Meta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EgressServer).Meta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logcache.v1.Egress/Meta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EgressServer).Meta(ctx, req.(*MetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Egress_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logcache.v1.Egress",
	HandlerType: (*EgressServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _Egress_Read_Handler,
		},
		{
			MethodName: "Meta",
			Handler:    _Egress_Meta_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "egress.proto",
}

func init() { proto.RegisterFile("egress.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x51, 0x8b, 0xd3, 0x40,
	0x10, 0x36, 0x4d, 0x7b, 0x6d, 0x26, 0xf5, 0x2e, 0x2e, 0x8a, 0x69, 0xf5, 0xb4, 0xc4, 0x97, 0x7a,
	0x4a, 0xcb, 0xc5, 0x07, 0x45, 0x11, 0x3c, 0x25, 0x1c, 0x87, 0x5e, 0x0f, 0x42, 0x4f, 0x10, 0x84,
	0xb2, 0x26, 0x63, 0x2e, 0x98, 0xee, 0xc6, 0xec, 0xb6, 0x5a, 0xc4, 0x17, 0x7f, 0x81, 0xe0, 0x83,
	0xbf, 0xc4, 0x5f, 0xe2, 0xab, 0x8f, 0xfe, 0x10, 0xc9, 0x6e, 0x5b, 0xd3, 0xd3, 0xb7, 0x99, 0x6f,
	0xbe, 0x7c, 0x9b, 0x99, 0x6f, 0x06, 0xda, 0x98, 0x14, 0x28, 0xc4, 0x20, 0x2f, 0xb8, 0xe4, 0xc4,
	0xce, 0x78, 0x12, 0xd1, 0xe8, 0x0c, 0x07, 0xf3, 0xfd, 0xee, 0xa5, 0xb9, 0x3f, 0x44, 0x36, 0xc7,
	0x8c, 0xe7, 0xa8, 0xeb, 0xdd, 0xeb, 0x09, 0xe7, 0x49, 0x86, 0x43, 0x9a, 0xa7, 0x43, 0xca, 0x18,
	0x97, 0x54, 0xa6, 0x9c, 0x2d, 0xbf, 0xf6, 0x7e, 0x19, 0x60, 0x87, 0x48, 0xe3, 0x10, 0xdf, 0xcf,
	0x50, 0x48, 0x72, 0x0d, 0x2c, 0xc1, 0x67, 0x45, 0x84, 0x93, 0x34, 0x76, 0x8d, 0x9e, 0xd1, 0xb7,
	0xc2, 0x96, 0x06, 0x8e, 0x62, 0xb2, 0x0b, 0x20, 0x24, 0x2d, 0xe4, 0x44, 0xa6, 0x53, 0x74, 0x6b,
	0x3d, 0xa3, 0x6f, 0x86, 0x96, 0x42, 0xc6, 0xe9, 0x14, 0x49, 0x07, 0x5a, 0xc8, 0x62, 0x5d, 0x34,
	0x55, 0xb1, 0x89, 0x2c, 0x56, 0xa5, 0xcb, 0xd0, 0xc8, 0xd2, 0x69, 0x2a, 0xdd, 0xba, 0xc2, 0x75,
	0x42, 0x9e, 0xc0, 0xf6, 0xea, 0x67, 0x27, 0x72, 0x91, 0xa3, 0x70, 0x1b, 0x3d, 0xb3, 0xbf, 0xed,
	0x77, 0x06, 0x95, 0x9e, 0x06, 0xc1, 0x92, 0x32, 0x5e, 0xe4, 0x18, 0x5e, 0xc4, 0x4a, 0x26, 0xc8,
	0x0d, 0x80, 0x18, 0x45, 0x84, 0x2c, 0x4e, 0x59, 0xe2, 0x6e, 0xf5, 0x8c, 0x7e, 0x2b, 0xac, 0x20,
	0xde, 0x73, 0x68, 0xeb, 0xee, 0x44, 0xce, 0x99, 0x40, 0xf2, 0x08, 0xac, 0x95, 0x80, 0x50, 0xed,
	0xd9, 0xfe, 0x6e, 0xf9, 0x58, 0x52, 0x60, 0x42, 0x25, 0x2f, 0x06, 0x73, 0x7f, 0xfd, 0xde, 0x53,
	0x2a, 0xa3, 0xb3, 0xf0, 0x2f, 0xdf, 0xbb, 0x0b, 0xf6, 0x31, 0x4a, 0xba, 0x1a, 0xd5, 0x2e, 0x40,
	0xc6, 0x23, 0x9a, 0x4d, 0x38, 0xcb, 0x16, 0x4a, 0xac, 0x15, 0x5a, 0x0a, 0x39, 0x61, 0xd9, 0xc2,
	0xfb, 0x6e, 0x40, 0x5b, 0xd3, 0x97, 0x6f, 0xdf, 0x87, 0xfa, 0x14, 0x25, 0x75, 0x8d, 0x9e, 0xd9,
	0xb7, 0xfd, 0x5b, 0x1b, 0x3d, 0x56, 0x89, 0x2a, 0x09, 0x98, 0x2c, 0x16, 0xa1, 0xfa, 0xa0, 0x3b,
	0x02, 0x6b, 0x0d, 0x11, 0x07, 0xcc, 0x77, 0xb8, 0x58, 0x5a, 0x53, 0x86, 0xe4, 0x0e, 0x34, 0xe6,
	0x34, 0x9b, 0x69, 0x43, 0x6c, 0xff, 0xca, 0x3f, 0xc2, 0x47, 0xec, 0x2d, 0x0f, 0x35, 0xe7, 0x61,
	0xed, 0x81, 0xe1, 0x7d, 0x35, 0xa0, 0xb5, 0xc2, 0x4b, 0x67, 0x22, 0x3e, 0x63, 0x52, 0x29, 0x9a,
	0xa1, 0x4e, 0x88, 0x0b, 0x4d, 0xfc, 0x98, 0xa7, 0x05, 0xc6, 0x4b, 0x9b, 0x57, 0x29, 0xb9, 0x0d,
	0x0e, 0xcf, 0x62, 0x14, 0x7a, 0x09, 0x84, 0xa4, 0xd3, 0x7c, 0x69, 0xf6, 0x8e, 0xc6, 0xc7, 0x2b,
	0xb8, 0xa4, 0x32, 0xfc, 0xb0, 0x49, 0xd5, 0xfe, 0xef, 0x68, 0x7c, 0x4d, 0xdd, 0x1b, 0x41, 0xbb,
	0x6a, 0x33, 0x69, 0x82, 0x79, 0x30, 0x7a, 0xe5, 0x5c, 0x28, 0x83, 0x17, 0x27, 0x87, 0x8e, 0x41,
	0x6c, 0x68, 0x3e, 0x3b, 0x39, 0x1d, 0x8d, 0x83, 0xd0, 0xa9, 0x11, 0x0b, 0x1a, 0x87, 0x07, 0xa7,
	0x87, 0x81, 0x63, 0x96, 0xe1, 0xf8, 0xe8, 0x38, 0x08, 0x9d, 0x7a, 0x19, 0x06, 0x2f, 0x83, 0xd1,
	0xd8, 0x69, 0xf8, 0x3f, 0x0c, 0xd8, 0x0a, 0xd4, 0x95, 0x90, 0xd7, 0x50, 0x2f, 0x57, 0x80, 0xb8,
	0x1b, 0x73, 0xa9, 0xec, 0x7c, 0xb7, 0xf3, 0x9f, 0x8a, 0xb6, 0xc2, 0xbb, 0xf9, 0xe5, 0xe7, 0xef,
	0x6f, 0xb5, 0x0e, 0xb9, 0x3a, 0x9c, 0xef, 0x0f, 0x0b, 0xa4, 0xf1, 0xf0, 0xd3, 0xfa, 0x3c, 0x1e,
	0xef, 0xed, 0x7d, 0x26, 0xc7, 0x50, 0x2f, 0x47, 0x79, 0x4e, 0xbd, 0xb2, 0x26, 0xe7, 0xd4, 0xab,
	0x46, 0x7b, 0x8e, 0x52, 0x07, 0xd2, 0x2a, 0xd5, 0x4b, 0xab, 0xdf, 0x6c, 0xa9, 0xab, 0xbc, 0xf7,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x79, 0x9e, 0x0a, 0xe3, 0x03, 0x00, 0x00,
}
