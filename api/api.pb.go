// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	Line
	LineList
	BufferList
	BufferIdentity
	BufferQuery
*/
package api

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

type Line struct {
	Id       int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Contents string `protobuf:"bytes,2,opt,name=contents" json:"contents,omitempty"`
	Number   int32  `protobuf:"varint,3,opt,name=number" json:"number,omitempty"`
}

func (m *Line) Reset()                    { *m = Line{} }
func (m *Line) String() string            { return proto.CompactTextString(m) }
func (*Line) ProtoMessage()               {}
func (*Line) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Line) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Line) GetContents() string {
	if m != nil {
		return m.Contents
	}
	return ""
}

func (m *Line) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type LineList struct {
	Lines []*Line `protobuf:"bytes,1,rep,name=lines" json:"lines,omitempty"`
}

func (m *LineList) Reset()                    { *m = LineList{} }
func (m *LineList) String() string            { return proto.CompactTextString(m) }
func (*LineList) ProtoMessage()               {}
func (*LineList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LineList) GetLines() []*Line {
	if m != nil {
		return m.Lines
	}
	return nil
}

type BufferList struct {
	Buffers []*BufferIdentity `protobuf:"bytes,1,rep,name=buffers" json:"buffers,omitempty"`
}

func (m *BufferList) Reset()                    { *m = BufferList{} }
func (m *BufferList) String() string            { return proto.CompactTextString(m) }
func (*BufferList) ProtoMessage()               {}
func (*BufferList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BufferList) GetBuffers() []*BufferIdentity {
	if m != nil {
		return m.Buffers
	}
	return nil
}

type BufferIdentity struct {
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
}

func (m *BufferIdentity) Reset()                    { *m = BufferIdentity{} }
func (m *BufferIdentity) String() string            { return proto.CompactTextString(m) }
func (*BufferIdentity) ProtoMessage()               {}
func (*BufferIdentity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BufferIdentity) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type BufferQuery struct {
	Prefix string `protobuf:"bytes,1,opt,name=prefix" json:"prefix,omitempty"`
}

func (m *BufferQuery) Reset()                    { *m = BufferQuery{} }
func (m *BufferQuery) String() string            { return proto.CompactTextString(m) }
func (*BufferQuery) ProtoMessage()               {}
func (*BufferQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BufferQuery) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func init() {
	proto.RegisterType((*Line)(nil), "api.Line")
	proto.RegisterType((*LineList)(nil), "api.LineList")
	proto.RegisterType((*BufferList)(nil), "api.BufferList")
	proto.RegisterType((*BufferIdentity)(nil), "api.BufferIdentity")
	proto.RegisterType((*BufferQuery)(nil), "api.BufferQuery")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Buffers service

type BuffersClient interface {
	GetContent(ctx context.Context, in *BufferIdentity, opts ...grpc.CallOption) (*LineList, error)
	WatchLines(ctx context.Context, in *BufferIdentity, opts ...grpc.CallOption) (Buffers_WatchLinesClient, error)
}

type buffersClient struct {
	cc *grpc.ClientConn
}

func NewBuffersClient(cc *grpc.ClientConn) BuffersClient {
	return &buffersClient{cc}
}

func (c *buffersClient) GetContent(ctx context.Context, in *BufferIdentity, opts ...grpc.CallOption) (*LineList, error) {
	out := new(LineList)
	err := grpc.Invoke(ctx, "/api.Buffers/GetContent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buffersClient) WatchLines(ctx context.Context, in *BufferIdentity, opts ...grpc.CallOption) (Buffers_WatchLinesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Buffers_serviceDesc.Streams[0], c.cc, "/api.Buffers/WatchLines", opts...)
	if err != nil {
		return nil, err
	}
	x := &buffersWatchLinesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Buffers_WatchLinesClient interface {
	Recv() (*Line, error)
	grpc.ClientStream
}

type buffersWatchLinesClient struct {
	grpc.ClientStream
}

func (x *buffersWatchLinesClient) Recv() (*Line, error) {
	m := new(Line)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Buffers service

type BuffersServer interface {
	GetContent(context.Context, *BufferIdentity) (*LineList, error)
	WatchLines(*BufferIdentity, Buffers_WatchLinesServer) error
}

func RegisterBuffersServer(s *grpc.Server, srv BuffersServer) {
	s.RegisterService(&_Buffers_serviceDesc, srv)
}

func _Buffers_GetContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BufferIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuffersServer).GetContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Buffers/GetContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuffersServer).GetContent(ctx, req.(*BufferIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Buffers_WatchLines_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BufferIdentity)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BuffersServer).WatchLines(m, &buffersWatchLinesServer{stream})
}

type Buffers_WatchLinesServer interface {
	Send(*Line) error
	grpc.ServerStream
}

type buffersWatchLinesServer struct {
	grpc.ServerStream
}

func (x *buffersWatchLinesServer) Send(m *Line) error {
	return x.ServerStream.SendMsg(m)
}

var _Buffers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Buffers",
	HandlerType: (*BuffersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetContent",
			Handler:    _Buffers_GetContent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchLines",
			Handler:       _Buffers_WatchLines_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}

// Client API for Editor service

type EditorClient interface {
	GetBuffers(ctx context.Context, in *BufferQuery, opts ...grpc.CallOption) (*BufferList, error)
}

type editorClient struct {
	cc *grpc.ClientConn
}

func NewEditorClient(cc *grpc.ClientConn) EditorClient {
	return &editorClient{cc}
}

func (c *editorClient) GetBuffers(ctx context.Context, in *BufferQuery, opts ...grpc.CallOption) (*BufferList, error) {
	out := new(BufferList)
	err := grpc.Invoke(ctx, "/api.Editor/GetBuffers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Editor service

type EditorServer interface {
	GetBuffers(context.Context, *BufferQuery) (*BufferList, error)
}

func RegisterEditorServer(s *grpc.Server, srv EditorServer) {
	s.RegisterService(&_Editor_serviceDesc, srv)
}

func _Editor_GetBuffers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BufferQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EditorServer).GetBuffers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Editor/GetBuffers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EditorServer).GetBuffers(ctx, req.(*BufferQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _Editor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Editor",
	HandlerType: (*EditorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBuffers",
			Handler:    _Editor_GetBuffers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x41, 0x4b, 0xf3, 0x40,
	0x14, 0x64, 0x93, 0x26, 0x6d, 0x5e, 0xf9, 0xfa, 0xc9, 0x0a, 0x12, 0x72, 0x31, 0x04, 0x85, 0x80,
	0x18, 0x4b, 0x3c, 0x89, 0xb7, 0x8a, 0x88, 0xd2, 0x8b, 0x7b, 0x11, 0xbc, 0x25, 0xcd, 0x86, 0x2e,
	0xea, 0x26, 0x6c, 0x5e, 0xc0, 0xfe, 0x7b, 0xd9, 0xdd, 0xa4, 0x56, 0xe8, 0x6d, 0xe7, 0xbd, 0x99,
	0x9d, 0x19, 0x1e, 0x04, 0x45, 0x2b, 0xb2, 0x56, 0x35, 0xd8, 0x50, 0xb7, 0x68, 0x45, 0xf2, 0x02,
	0x93, 0xb5, 0x90, 0x9c, 0x2e, 0xc0, 0x11, 0x55, 0x48, 0x62, 0x92, 0xba, 0xcc, 0x11, 0x15, 0x8d,
	0x60, 0xb6, 0x69, 0x24, 0x72, 0x89, 0x5d, 0xe8, 0xc4, 0x24, 0x0d, 0xd8, 0x1e, 0xd3, 0x33, 0xf0,
	0x65, 0xff, 0x55, 0x72, 0x15, 0xba, 0x31, 0x49, 0x3d, 0x36, 0xa0, 0xe4, 0x0a, 0x66, 0xfa, 0xaf,
	0xb5, 0xe8, 0x90, 0x9e, 0x83, 0xf7, 0x29, 0x24, 0xef, 0x42, 0x12, 0xbb, 0xe9, 0x3c, 0x0f, 0x32,
	0xed, 0xab, 0xb7, 0xcc, 0xce, 0x93, 0x7b, 0x80, 0x55, 0x5f, 0xd7, 0x5c, 0x19, 0xfa, 0x35, 0x4c,
	0x4b, 0x83, 0x46, 0xc1, 0xa9, 0x11, 0x58, 0xc6, 0x73, 0xc5, 0x25, 0x0a, 0xdc, 0xb1, 0x91, 0x93,
	0x5c, 0xc0, 0xe2, 0xef, 0x8a, 0x52, 0x98, 0xb4, 0x05, 0x6e, 0x4d, 0x83, 0x80, 0x99, 0x77, 0x72,
	0x09, 0x73, 0xcb, 0x7a, 0xed, 0xb9, 0xda, 0xe9, 0xd8, 0xad, 0xe2, 0xb5, 0xf8, 0x1e, 0x48, 0x03,
	0xca, 0x3f, 0x60, 0x6a, 0x69, 0x1d, 0x5d, 0x02, 0x3c, 0x71, 0x7c, 0xb0, 0x45, 0xe9, 0xb1, 0x0c,
	0xd1, 0xbf, 0x7d, 0x13, 0x13, 0x3c, 0x03, 0x78, 0x2b, 0x70, 0xb3, 0xd5, 0x83, 0xee, 0xb8, 0xe2,
	0xb7, 0xfb, 0x92, 0xe4, 0x77, 0xe0, 0x3f, 0x56, 0x02, 0x1b, 0x45, 0x6f, 0x8c, 0xd7, 0xe8, 0x7c,
	0x72, 0xa0, 0x34, 0x71, 0xa3, 0xff, 0x07, 0x13, 0x6d, 0xb5, 0xf2, 0xde, 0xf5, 0xc5, 0x4a, 0xdf,
	0x5c, 0xef, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x90, 0xd1, 0x1e, 0x0e, 0xca, 0x01, 0x00, 0x00,
}