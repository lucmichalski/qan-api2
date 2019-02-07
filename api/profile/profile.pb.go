// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/profile/profile.proto

package profile

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type ProfileRequest struct {
	From                 string           `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string           `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Keyword              string           `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`
	FirstSeen            bool             `protobuf:"varint,4,opt,name=first_seen,json=firstSeen,proto3" json:"first_seen,omitempty"`
	GroupBy              string           `protobuf:"bytes,5,opt,name=group_by,json=groupBy,proto3" json:"group_by,omitempty"`
	Labels               []*MapFieldEntry `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty"`
	Include              []string         `protobuf:"bytes,7,rep,name=include,proto3" json:"include,omitempty"`
	Order                string           `protobuf:"bytes,8,opt,name=order,proto3" json:"order,omitempty"`
	Offset               uint32           `protobuf:"varint,9,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                uint32           `protobuf:"varint,10,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ProfileRequest) Reset()         { *m = ProfileRequest{} }
func (m *ProfileRequest) String() string { return proto.CompactTextString(m) }
func (*ProfileRequest) ProtoMessage()    {}
func (*ProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_profile_00a01d03d7c8bd1d, []int{0}
}
func (m *ProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileRequest.Unmarshal(m, b)
}
func (m *ProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileRequest.Marshal(b, m, deterministic)
}
func (dst *ProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileRequest.Merge(dst, src)
}
func (m *ProfileRequest) XXX_Size() int {
	return xxx_messageInfo_ProfileRequest.Size(m)
}
func (m *ProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileRequest proto.InternalMessageInfo

func (m *ProfileRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *ProfileRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *ProfileRequest) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func (m *ProfileRequest) GetFirstSeen() bool {
	if m != nil {
		return m.FirstSeen
	}
	return false
}

func (m *ProfileRequest) GetGroupBy() string {
	if m != nil {
		return m.GroupBy
	}
	return ""
}

func (m *ProfileRequest) GetLabels() []*MapFieldEntry {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ProfileRequest) GetInclude() []string {
	if m != nil {
		return m.Include
	}
	return nil
}

func (m *ProfileRequest) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

func (m *ProfileRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ProfileRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// MapFieldEntry
type MapFieldEntry struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []string `protobuf:"bytes,2,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MapFieldEntry) Reset()         { *m = MapFieldEntry{} }
func (m *MapFieldEntry) String() string { return proto.CompactTextString(m) }
func (*MapFieldEntry) ProtoMessage()    {}
func (*MapFieldEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_profile_00a01d03d7c8bd1d, []int{1}
}
func (m *MapFieldEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapFieldEntry.Unmarshal(m, b)
}
func (m *MapFieldEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapFieldEntry.Marshal(b, m, deterministic)
}
func (dst *MapFieldEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapFieldEntry.Merge(dst, src)
}
func (m *MapFieldEntry) XXX_Size() int {
	return xxx_messageInfo_MapFieldEntry.Size(m)
}
func (m *MapFieldEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_MapFieldEntry.DiscardUnknown(m)
}

var xxx_messageInfo_MapFieldEntry proto.InternalMessageInfo

func (m *MapFieldEntry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MapFieldEntry) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type ProfileReply struct {
	Rows                 []*ProfileRow `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ProfileReply) Reset()         { *m = ProfileReply{} }
func (m *ProfileReply) String() string { return proto.CompactTextString(m) }
func (*ProfileReply) ProtoMessage()    {}
func (*ProfileReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_profile_00a01d03d7c8bd1d, []int{2}
}
func (m *ProfileReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileReply.Unmarshal(m, b)
}
func (m *ProfileReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileReply.Marshal(b, m, deterministic)
}
func (dst *ProfileReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileReply.Merge(dst, src)
}
func (m *ProfileReply) XXX_Size() int {
	return xxx_messageInfo_ProfileReply.Size(m)
}
func (m *ProfileReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileReply.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileReply proto.InternalMessageInfo

func (m *ProfileReply) GetRows() []*ProfileRow {
	if m != nil {
		return m.Rows
	}
	return nil
}

type ProfileRow struct {
	Rank                 uint32   `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	Percentage           float32  `protobuf:"fixed32,2,opt,name=percentage,proto3" json:"percentage,omitempty"`
	Digest               string   `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
	DigestText           string   `protobuf:"bytes,4,opt,name=digest_text,json=digestText,proto3" json:"digest_text,omitempty"`
	FirstSeen            string   `protobuf:"bytes,5,opt,name=first_seen,json=firstSeen,proto3" json:"first_seen,omitempty"`
	Qps                  float32  `protobuf:"fixed32,6,opt,name=qps,proto3" json:"qps,omitempty"`
	Load                 float32  `protobuf:"fixed32,7,opt,name=load,proto3" json:"load,omitempty"`
	Stats                *Stats   `protobuf:"bytes,8,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProfileRow) Reset()         { *m = ProfileRow{} }
func (m *ProfileRow) String() string { return proto.CompactTextString(m) }
func (*ProfileRow) ProtoMessage()    {}
func (*ProfileRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_profile_00a01d03d7c8bd1d, []int{3}
}
func (m *ProfileRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileRow.Unmarshal(m, b)
}
func (m *ProfileRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileRow.Marshal(b, m, deterministic)
}
func (dst *ProfileRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileRow.Merge(dst, src)
}
func (m *ProfileRow) XXX_Size() int {
	return xxx_messageInfo_ProfileRow.Size(m)
}
func (m *ProfileRow) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileRow.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileRow proto.InternalMessageInfo

func (m *ProfileRow) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *ProfileRow) GetPercentage() float32 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

func (m *ProfileRow) GetDigest() string {
	if m != nil {
		return m.Digest
	}
	return ""
}

func (m *ProfileRow) GetDigestText() string {
	if m != nil {
		return m.DigestText
	}
	return ""
}

func (m *ProfileRow) GetFirstSeen() string {
	if m != nil {
		return m.FirstSeen
	}
	return ""
}

func (m *ProfileRow) GetQps() float32 {
	if m != nil {
		return m.Qps
	}
	return 0
}

func (m *ProfileRow) GetLoad() float32 {
	if m != nil {
		return m.Load
	}
	return 0
}

func (m *ProfileRow) GetStats() *Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type Stats struct {
	NumQueries           uint64   `protobuf:"varint,1,opt,name=num_queries,json=numQueries,proto3" json:"num_queries,omitempty"`
	MQueryTimeSum        float32  `protobuf:"fixed32,2,opt,name=m_query_time_sum,json=mQueryTimeSum,proto3" json:"m_query_time_sum,omitempty"`
	MQueryTimeMin        float32  `protobuf:"fixed32,3,opt,name=m_query_time_min,json=mQueryTimeMin,proto3" json:"m_query_time_min,omitempty"`
	MQueryTimeMax        float32  `protobuf:"fixed32,4,opt,name=m_query_time_max,json=mQueryTimeMax,proto3" json:"m_query_time_max,omitempty"`
	MQueryTimeP99        float32  `protobuf:"fixed32,5,opt,name=m_query_time_p99,json=mQueryTimeP99,proto3" json:"m_query_time_p99,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stats) Reset()         { *m = Stats{} }
func (m *Stats) String() string { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()    {}
func (*Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_profile_00a01d03d7c8bd1d, []int{4}
}
func (m *Stats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stats.Unmarshal(m, b)
}
func (m *Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stats.Marshal(b, m, deterministic)
}
func (dst *Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stats.Merge(dst, src)
}
func (m *Stats) XXX_Size() int {
	return xxx_messageInfo_Stats.Size(m)
}
func (m *Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_Stats proto.InternalMessageInfo

func (m *Stats) GetNumQueries() uint64 {
	if m != nil {
		return m.NumQueries
	}
	return 0
}

func (m *Stats) GetMQueryTimeSum() float32 {
	if m != nil {
		return m.MQueryTimeSum
	}
	return 0
}

func (m *Stats) GetMQueryTimeMin() float32 {
	if m != nil {
		return m.MQueryTimeMin
	}
	return 0
}

func (m *Stats) GetMQueryTimeMax() float32 {
	if m != nil {
		return m.MQueryTimeMax
	}
	return 0
}

func (m *Stats) GetMQueryTimeP99() float32 {
	if m != nil {
		return m.MQueryTimeP99
	}
	return 0
}

func init() {
	proto.RegisterType((*ProfileRequest)(nil), "profile.ProfileRequest")
	proto.RegisterType((*MapFieldEntry)(nil), "profile.MapFieldEntry")
	proto.RegisterType((*ProfileReply)(nil), "profile.ProfileReply")
	proto.RegisterType((*ProfileRow)(nil), "profile.ProfileRow")
	proto.RegisterType((*Stats)(nil), "profile.Stats")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProfileClient is the client API for Profile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfileClient interface {
	// GetDigestGroup returns list of metrics group by digest.
	GetDigestGroup(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*ProfileReply, error)
}

type profileClient struct {
	cc *grpc.ClientConn
}

func NewProfileClient(cc *grpc.ClientConn) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) GetDigestGroup(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, "/profile.Profile/GetDigestGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServer is the server API for Profile service.
type ProfileServer interface {
	// GetDigestGroup returns list of metrics group by digest.
	GetDigestGroup(context.Context, *ProfileRequest) (*ProfileReply, error)
}

func RegisterProfileServer(s *grpc.Server, srv ProfileServer) {
	s.RegisterService(&_Profile_serviceDesc, srv)
}

func _Profile_GetDigestGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetDigestGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/GetDigestGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetDigestGroup(ctx, req.(*ProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Profile_serviceDesc = grpc.ServiceDesc{
	ServiceName: "profile.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDigestGroup",
			Handler:    _Profile_GetDigestGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/profile/profile.proto",
}

func init() { proto.RegisterFile("api/profile/profile.proto", fileDescriptor_profile_00a01d03d7c8bd1d) }

var fileDescriptor_profile_00a01d03d7c8bd1d = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x9d, 0xaf, 0x66, 0x4a, 0xa2, 0x68, 0x81, 0xb2, 0x45, 0x02, 0xa2, 0x08, 0xa9, 0x39,
	0x05, 0x29, 0x1c, 0xaa, 0xdc, 0x50, 0x05, 0xf4, 0x80, 0x2a, 0x95, 0x4d, 0x4f, 0x5c, 0x2c, 0xa7,
	0x9e, 0x44, 0xab, 0xd8, 0x5e, 0x67, 0x77, 0x4d, 0xe2, 0x3f, 0xc2, 0xaf, 0xe2, 0x67, 0xf0, 0x43,
	0xd0, 0xce, 0x3a, 0xa1, 0x21, 0xa7, 0xcc, 0x7b, 0xfb, 0x36, 0x33, 0xf3, 0x9e, 0x17, 0x2e, 0xe3,
	0x42, 0x7e, 0x28, 0xb4, 0x5a, 0xca, 0x14, 0xf7, 0xbf, 0x93, 0x42, 0x2b, 0xab, 0x58, 0xa7, 0x86,
	0xa3, 0x5f, 0x21, 0xf4, 0xef, 0x7d, 0x2d, 0x70, 0x53, 0xa2, 0xb1, 0x8c, 0x41, 0x73, 0xa9, 0x55,
	0xc6, 0x83, 0x61, 0x30, 0xee, 0x0a, 0xaa, 0x59, 0x1f, 0x42, 0xab, 0x78, 0x48, 0x4c, 0x68, 0x15,
	0xe3, 0xd0, 0x59, 0x63, 0xb5, 0x55, 0x3a, 0xe1, 0x0d, 0x22, 0xf7, 0x90, 0xbd, 0x01, 0x58, 0x4a,
	0x6d, 0x6c, 0x64, 0x10, 0x73, 0xde, 0x1c, 0x06, 0xe3, 0x33, 0xd1, 0x25, 0x66, 0x8e, 0x98, 0xb3,
	0x4b, 0x38, 0x5b, 0x69, 0x55, 0x16, 0xd1, 0xa2, 0xe2, 0x2d, 0x7f, 0x93, 0xf0, 0x4d, 0xc5, 0x26,
	0xd0, 0x4e, 0xe3, 0x05, 0xa6, 0x86, 0xb7, 0x87, 0x8d, 0xf1, 0xf9, 0xf4, 0x62, 0xb2, 0x9f, 0xf9,
	0x2e, 0x2e, 0xbe, 0x4a, 0x4c, 0x93, 0x2f, 0xb9, 0xd5, 0x95, 0xa8, 0x55, 0x6e, 0x06, 0x99, 0x3f,
	0xa6, 0x65, 0x82, 0xbc, 0x33, 0x6c, 0xb8, 0x7f, 0xaa, 0x21, 0x7b, 0x01, 0x2d, 0xa5, 0x13, 0xd4,
	0xfc, 0x8c, 0x3a, 0x78, 0xc0, 0x2e, 0xa0, 0xad, 0x96, 0x4b, 0x83, 0x96, 0x77, 0x87, 0xc1, 0xb8,
	0x27, 0x6a, 0xe4, 0xd4, 0xa9, 0xcc, 0xa4, 0xe5, 0x40, 0xb4, 0x07, 0xa3, 0x6b, 0xe8, 0x1d, 0xb5,
	0x65, 0x03, 0x68, 0xac, 0xb1, 0xaa, 0x5d, 0x71, 0xa5, 0xbb, 0xf8, 0x33, 0x4e, 0x4b, 0xe4, 0x21,
	0xb5, 0xf7, 0x60, 0x74, 0x0d, 0xcf, 0x0e, 0x86, 0x16, 0x69, 0xc5, 0xae, 0xa0, 0xa9, 0xd5, 0xd6,
	0xf0, 0x80, 0x96, 0x7a, 0x7e, 0x58, 0x6a, 0x2f, 0x52, 0x5b, 0x41, 0x82, 0xd1, 0x9f, 0x00, 0xe0,
	0x1f, 0xe9, 0x62, 0xd0, 0x71, 0xbe, 0xa6, 0x86, 0x3d, 0x41, 0x35, 0x7b, 0x0b, 0x50, 0xa0, 0x7e,
	0xc4, 0xdc, 0xc6, 0x2b, 0xa4, 0x38, 0x42, 0xf1, 0x84, 0x71, 0x2b, 0x26, 0x72, 0x85, 0xc6, 0xd6,
	0xa9, 0xd4, 0x88, 0xbd, 0x83, 0x73, 0x5f, 0x45, 0x16, 0x77, 0x96, 0x52, 0xe9, 0x0a, 0xf0, 0xd4,
	0x03, 0xee, 0xec, 0x7f, 0xa9, 0xf9, 0x60, 0x9e, 0xa4, 0x36, 0x80, 0xc6, 0xa6, 0x70, 0xb9, 0xb8,
	0x86, 0xae, 0x74, 0xd3, 0xa5, 0x2a, 0x4e, 0x78, 0x87, 0x28, 0xaa, 0xd9, 0x7b, 0x68, 0x19, 0x1b,
	0x5b, 0x43, 0xb6, 0x9f, 0x4f, 0xfb, 0x87, 0x55, 0xe7, 0x8e, 0x15, 0xfe, 0x70, 0xf4, 0x3b, 0x80,
	0x16, 0x11, 0x6e, 0xaa, 0xbc, 0xcc, 0xa2, 0x4d, 0x89, 0x5a, 0xa2, 0xa1, 0x45, 0x9b, 0x02, 0xf2,
	0x32, 0xfb, 0xee, 0x19, 0x76, 0x05, 0x03, 0x7f, 0x5c, 0x45, 0x56, 0x66, 0x18, 0x99, 0x32, 0xab,
	0x97, 0xee, 0x91, 0xa6, 0x7a, 0x90, 0x19, 0xce, 0xcb, 0xec, 0x44, 0x98, 0xc9, 0x9c, 0x1c, 0x38,
	0x12, 0xde, 0xc9, 0xfc, 0x54, 0x18, 0xef, 0xc8, 0x8d, 0x63, 0x61, 0xbc, 0x3b, 0x11, 0x16, 0xb3,
	0x19, 0xd9, 0x72, 0x24, 0xbc, 0x9f, 0xcd, 0xa6, 0xdf, 0xa0, 0x53, 0x87, 0xc6, 0x3e, 0x41, 0xff,
	0x16, 0xed, 0x67, 0x72, 0xf5, 0xd6, 0x7d, 0xd4, 0xec, 0xd5, 0x49, 0xda, 0xfe, 0x8d, 0xbd, 0x7e,
	0x79, 0x7a, 0x50, 0xa4, 0xd5, 0x4d, 0xf7, 0xc7, 0xfe, 0x61, 0x2e, 0xda, 0xf4, 0x50, 0x3f, 0xfe,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0x50, 0x26, 0x1f, 0x19, 0xc5, 0x03, 0x00, 0x00,
}
