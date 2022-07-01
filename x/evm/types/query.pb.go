// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evm/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_11c0d37eed5339f7, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_11c0d37eed5339f7, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetValsetByIDRequest struct {
	ValsetID         uint64 `protobuf:"varint,1,opt,name=valsetID,proto3" json:"valsetID,omitempty"`
	ChainReferenceID string `protobuf:"bytes,2,opt,name=chainReferenceID,proto3" json:"chainReferenceID,omitempty"`
}

func (m *QueryGetValsetByIDRequest) Reset()         { *m = QueryGetValsetByIDRequest{} }
func (m *QueryGetValsetByIDRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetValsetByIDRequest) ProtoMessage()    {}
func (*QueryGetValsetByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_11c0d37eed5339f7, []int{2}
}
func (m *QueryGetValsetByIDRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetValsetByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetValsetByIDRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetValsetByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetValsetByIDRequest.Merge(m, src)
}
func (m *QueryGetValsetByIDRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetValsetByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetValsetByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetValsetByIDRequest proto.InternalMessageInfo

func (m *QueryGetValsetByIDRequest) GetValsetID() uint64 {
	if m != nil {
		return m.ValsetID
	}
	return 0
}

func (m *QueryGetValsetByIDRequest) GetChainReferenceID() string {
	if m != nil {
		return m.ChainReferenceID
	}
	return ""
}

type QueryGetValsetByIDResponse struct {
	Valset *Valset `protobuf:"bytes,1,opt,name=valset,proto3" json:"valset,omitempty"`
}

func (m *QueryGetValsetByIDResponse) Reset()         { *m = QueryGetValsetByIDResponse{} }
func (m *QueryGetValsetByIDResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetValsetByIDResponse) ProtoMessage()    {}
func (*QueryGetValsetByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_11c0d37eed5339f7, []int{3}
}
func (m *QueryGetValsetByIDResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetValsetByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetValsetByIDResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetValsetByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetValsetByIDResponse.Merge(m, src)
}
func (m *QueryGetValsetByIDResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetValsetByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetValsetByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetValsetByIDResponse proto.InternalMessageInfo

func (m *QueryGetValsetByIDResponse) GetValset() *Valset {
	if m != nil {
		return m.Valset
	}
	return nil
}

type QueryChainsInfosRequest struct {
}

func (m *QueryChainsInfosRequest) Reset()         { *m = QueryChainsInfosRequest{} }
func (m *QueryChainsInfosRequest) String() string { return proto.CompactTextString(m) }
func (*QueryChainsInfosRequest) ProtoMessage()    {}
func (*QueryChainsInfosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_11c0d37eed5339f7, []int{4}
}
func (m *QueryChainsInfosRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryChainsInfosRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryChainsInfosRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryChainsInfosRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryChainsInfosRequest.Merge(m, src)
}
func (m *QueryChainsInfosRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryChainsInfosRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryChainsInfosRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryChainsInfosRequest proto.InternalMessageInfo

type QueryChainsInfosResponse struct {
}

func (m *QueryChainsInfosResponse) Reset()         { *m = QueryChainsInfosResponse{} }
func (m *QueryChainsInfosResponse) String() string { return proto.CompactTextString(m) }
func (*QueryChainsInfosResponse) ProtoMessage()    {}
func (*QueryChainsInfosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_11c0d37eed5339f7, []int{5}
}
func (m *QueryChainsInfosResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryChainsInfosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryChainsInfosResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryChainsInfosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryChainsInfosResponse.Merge(m, src)
}
func (m *QueryChainsInfosResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryChainsInfosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryChainsInfosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryChainsInfosResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "palomachain.paloma.evm.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "palomachain.paloma.evm.QueryParamsResponse")
	proto.RegisterType((*QueryGetValsetByIDRequest)(nil), "palomachain.paloma.evm.QueryGetValsetByIDRequest")
	proto.RegisterType((*QueryGetValsetByIDResponse)(nil), "palomachain.paloma.evm.QueryGetValsetByIDResponse")
	proto.RegisterType((*QueryChainsInfosRequest)(nil), "palomachain.paloma.evm.QueryChainsInfosRequest")
	proto.RegisterType((*QueryChainsInfosResponse)(nil), "palomachain.paloma.evm.QueryChainsInfosResponse")
}

func init() { proto.RegisterFile("evm/query.proto", fileDescriptor_11c0d37eed5339f7) }

var fileDescriptor_11c0d37eed5339f7 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcf, 0x6e, 0xd3, 0x30,
	0x18, 0xaf, 0xc7, 0xa8, 0xc0, 0x13, 0x62, 0xf2, 0x26, 0xe8, 0x22, 0x64, 0xaa, 0x08, 0x4d, 0xa3,
	0xa0, 0x98, 0x76, 0x82, 0x0b, 0x9c, 0xba, 0x4a, 0xa8, 0x37, 0x08, 0x88, 0x03, 0x97, 0xca, 0x09,
	0xdf, 0xb2, 0x48, 0x8d, 0x9d, 0xc5, 0x6e, 0x45, 0x85, 0xb8, 0x70, 0xe3, 0x86, 0xc4, 0x9d, 0x97,
	0xe0, 0x25, 0x76, 0x9c, 0xc4, 0x85, 0x13, 0x42, 0x2d, 0x27, 0x9e, 0x02, 0xc5, 0xf6, 0xc6, 0xa6,
	0xa6, 0xa3, 0xdc, 0x9c, 0x9f, 0x7f, 0xff, 0xfc, 0xd9, 0xc1, 0xd7, 0x61, 0x9c, 0xb1, 0xc3, 0x11,
	0x14, 0x93, 0x20, 0x2f, 0xa4, 0x96, 0xe4, 0x46, 0xce, 0x87, 0x32, 0xe3, 0xf1, 0x01, 0x4f, 0x45,
	0x60, 0xd7, 0x01, 0x8c, 0x33, 0x6f, 0x33, 0x91, 0x89, 0x34, 0x14, 0x56, 0xae, 0x2c, 0xdb, 0xbb,
	0x95, 0x48, 0x99, 0x0c, 0x81, 0xf1, 0x3c, 0x65, 0x5c, 0x08, 0xa9, 0xb9, 0x4e, 0xa5, 0x50, 0x6e,
	0xb7, 0x15, 0x4b, 0x95, 0x49, 0xc5, 0x22, 0xae, 0xc0, 0x86, 0xb0, 0x71, 0x3b, 0x02, 0xcd, 0xdb,
	0x2c, 0xe7, 0x49, 0x2a, 0x0c, 0xd9, 0x71, 0xd7, 0xcb, 0x22, 0x39, 0x2f, 0x78, 0x76, 0xa2, 0xde,
	0x28, 0x11, 0x3d, 0x2a, 0x84, 0xd2, 0x52, 0x80, 0x05, 0xfd, 0x4d, 0x4c, 0x9e, 0x97, 0x46, 0xcf,
	0x0c, 0x33, 0x84, 0xc3, 0x11, 0x28, 0xed, 0xbf, 0xc0, 0x1b, 0xe7, 0x50, 0x95, 0x4b, 0xa1, 0x80,
	0x3c, 0xc1, 0x75, 0xeb, 0xd8, 0x40, 0x4d, 0xb4, 0xb3, 0xd6, 0xa1, 0x41, 0xf5, 0xe1, 0x02, 0xab,
	0xeb, 0xae, 0x1e, 0xfd, 0xb8, 0x5d, 0x0b, 0x9d, 0xc6, 0x8f, 0xf1, 0x96, 0x31, 0x7d, 0x0a, 0xfa,
	0x15, 0x1f, 0x2a, 0xd0, 0xdd, 0x49, 0xbf, 0xe7, 0x12, 0x89, 0x87, 0xaf, 0x8c, 0x0d, 0xd8, 0xef,
	0x19, 0xf3, 0xd5, 0xf0, 0xf4, 0x9b, 0xb4, 0xf0, 0xba, 0x49, 0x08, 0x61, 0x1f, 0x0a, 0x10, 0x31,
	0xf4, 0x7b, 0x8d, 0x95, 0x26, 0xda, 0xb9, 0x1a, 0xce, 0xe1, 0xfe, 0x4b, 0xec, 0x55, 0x85, 0xb8,
	0x03, 0x3c, 0xc2, 0x75, 0xeb, 0xfa, 0xaf, 0x03, 0x58, 0x6d, 0xe8, 0xd8, 0xfe, 0x16, 0xbe, 0x69,
	0x5c, 0xf7, 0x4a, 0x9e, 0xea, 0x8b, 0x7d, 0x79, 0x3a, 0x2a, 0x0f, 0x37, 0xe6, 0xb7, 0x6c, 0x5c,
	0xe7, 0xf7, 0x25, 0x7c, 0xd9, 0x6c, 0x92, 0x8f, 0x08, 0xd7, 0xed, 0x50, 0x48, 0x6b, 0x51, 0xe6,
	0xfc, 0x3d, 0x78, 0xf7, 0x96, 0xe2, 0xda, 0x34, 0x7f, 0xfb, 0xc3, 0xb7, 0x5f, 0x9f, 0x57, 0x9a,
	0x84, 0xb2, 0x33, 0x22, 0xb7, 0x66, 0x7f, 0x5f, 0x03, 0xf9, 0x8a, 0xf0, 0xb5, 0x73, 0xe3, 0x21,
	0xed, 0x0b, 0x63, 0xaa, 0xee, 0xcb, 0xeb, 0xfc, 0x8f, 0xc4, 0x15, 0x7c, 0x6c, 0x0a, 0x3e, 0x24,
	0xbb, 0x8b, 0x0a, 0x26, 0xa0, 0x07, 0x76, 0xe2, 0x83, 0x68, 0x32, 0x48, 0xdf, 0xb0, 0x77, 0x27,
	0x6f, 0xe0, 0x3d, 0xf9, 0x82, 0xf0, 0xda, 0x99, 0x19, 0x13, 0x76, 0x61, 0x81, 0xf9, 0x8b, 0xf2,
	0x1e, 0x2c, 0x2f, 0x70, 0x7d, 0xef, 0x9b, 0xbe, 0xdb, 0xe4, 0xce, 0xa2, 0xbe, 0x06, 0x50, 0x83,
	0xb4, 0x54, 0x75, 0xf7, 0x8e, 0xa6, 0x14, 0x1d, 0x4f, 0x29, 0xfa, 0x39, 0xa5, 0xe8, 0xd3, 0x8c,
	0xd6, 0x8e, 0x67, 0xb4, 0xf6, 0x7d, 0x46, 0x6b, 0xaf, 0xef, 0x26, 0xa9, 0x3e, 0x18, 0x45, 0x41,
	0x2c, 0xb3, 0x2a, 0xa7, 0xb7, 0xc6, 0x4b, 0x4f, 0x72, 0x50, 0x51, 0xdd, 0xfc, 0x95, 0xbb, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x48, 0x5e, 0x4a, 0xd1, 0x47, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of GetValsetByID items.
	GetValsetByID(ctx context.Context, in *QueryGetValsetByIDRequest, opts ...grpc.CallOption) (*QueryGetValsetByIDResponse, error)
	// Queries a list of ChainsInfos items.
	ChainsInfos(ctx context.Context, in *QueryChainsInfosRequest, opts ...grpc.CallOption) (*QueryChainsInfosResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/palomachain.paloma.evm.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetValsetByID(ctx context.Context, in *QueryGetValsetByIDRequest, opts ...grpc.CallOption) (*QueryGetValsetByIDResponse, error) {
	out := new(QueryGetValsetByIDResponse)
	err := c.cc.Invoke(ctx, "/palomachain.paloma.evm.Query/GetValsetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ChainsInfos(ctx context.Context, in *QueryChainsInfosRequest, opts ...grpc.CallOption) (*QueryChainsInfosResponse, error) {
	out := new(QueryChainsInfosResponse)
	err := c.cc.Invoke(ctx, "/palomachain.paloma.evm.Query/ChainsInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of GetValsetByID items.
	GetValsetByID(context.Context, *QueryGetValsetByIDRequest) (*QueryGetValsetByIDResponse, error)
	// Queries a list of ChainsInfos items.
	ChainsInfos(context.Context, *QueryChainsInfosRequest) (*QueryChainsInfosResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) GetValsetByID(ctx context.Context, req *QueryGetValsetByIDRequest) (*QueryGetValsetByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValsetByID not implemented")
}
func (*UnimplementedQueryServer) ChainsInfos(ctx context.Context, req *QueryChainsInfosRequest) (*QueryChainsInfosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChainsInfos not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/palomachain.paloma.evm.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetValsetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetValsetByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetValsetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/palomachain.paloma.evm.Query/GetValsetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetValsetByID(ctx, req.(*QueryGetValsetByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ChainsInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryChainsInfosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ChainsInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/palomachain.paloma.evm.Query/ChainsInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ChainsInfos(ctx, req.(*QueryChainsInfosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "palomachain.paloma.evm.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "GetValsetByID",
			Handler:    _Query_GetValsetByID_Handler,
		},
		{
			MethodName: "ChainsInfos",
			Handler:    _Query_ChainsInfos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "evm/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetValsetByIDRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetValsetByIDRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetValsetByIDRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainReferenceID) > 0 {
		i -= len(m.ChainReferenceID)
		copy(dAtA[i:], m.ChainReferenceID)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ChainReferenceID)))
		i--
		dAtA[i] = 0x12
	}
	if m.ValsetID != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ValsetID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetValsetByIDResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetValsetByIDResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetValsetByIDResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Valset != nil {
		{
			size, err := m.Valset.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryChainsInfosRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryChainsInfosRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryChainsInfosRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryChainsInfosResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryChainsInfosResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryChainsInfosResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetValsetByIDRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ValsetID != 0 {
		n += 1 + sovQuery(uint64(m.ValsetID))
	}
	l = len(m.ChainReferenceID)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetValsetByIDResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Valset != nil {
		l = m.Valset.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryChainsInfosRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryChainsInfosResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetValsetByIDRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetValsetByIDRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetValsetByIDRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValsetID", wireType)
			}
			m.ValsetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValsetID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainReferenceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainReferenceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetValsetByIDResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetValsetByIDResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetValsetByIDResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Valset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Valset == nil {
				m.Valset = &Valset{}
			}
			if err := m.Valset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryChainsInfosRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryChainsInfosRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryChainsInfosRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryChainsInfosResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryChainsInfosResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryChainsInfosResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
