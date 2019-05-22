// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package stock

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

type SkuMessage struct {
	Sku                  string   `protobuf:"bytes,1,opt,name=sku,proto3" json:"sku,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SkuMessage) Reset()         { *m = SkuMessage{} }
func (m *SkuMessage) String() string { return proto.CompactTextString(m) }
func (*SkuMessage) ProtoMessage()    {}
func (*SkuMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *SkuMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SkuMessage.Unmarshal(m, b)
}
func (m *SkuMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SkuMessage.Marshal(b, m, deterministic)
}
func (m *SkuMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SkuMessage.Merge(m, src)
}
func (m *SkuMessage) XXX_Size() int {
	return xxx_messageInfo_SkuMessage.Size(m)
}
func (m *SkuMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SkuMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SkuMessage proto.InternalMessageInfo

func (m *SkuMessage) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

type StockMessage struct {
	Stock                int32    `protobuf:"varint,1,opt,name=stock,proto3" json:"stock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StockMessage) Reset()         { *m = StockMessage{} }
func (m *StockMessage) String() string { return proto.CompactTextString(m) }
func (*StockMessage) ProtoMessage()    {}
func (*StockMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *StockMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StockMessage.Unmarshal(m, b)
}
func (m *StockMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StockMessage.Marshal(b, m, deterministic)
}
func (m *StockMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StockMessage.Merge(m, src)
}
func (m *StockMessage) XXX_Size() int {
	return xxx_messageInfo_StockMessage.Size(m)
}
func (m *StockMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StockMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StockMessage proto.InternalMessageInfo

func (m *StockMessage) GetStock() int32 {
	if m != nil {
		return m.Stock
	}
	return 0
}

func init() {
	proto.RegisterType((*SkuMessage)(nil), "service.bidogo.product.stock.SkuMessage")
	proto.RegisterType((*StockMessage)(nil), "service.bidogo.product.stock.StockMessage")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x81, 0x71, 0x93, 0x32, 0x53,
	0xf2, 0xd3, 0xf3, 0x41, 0xa2, 0x29, 0xa5, 0xc9, 0x25, 0x7a, 0xc5, 0x25, 0xf9, 0xc9, 0xd9, 0x4a,
	0x72, 0x5c, 0x5c, 0xc1, 0xd9, 0xa5, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x42, 0x02, 0x5c,
	0xcc, 0xc5, 0xd9, 0xa5, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x92, 0x0a, 0x17,
	0x4f, 0x30, 0x48, 0x21, 0x4c, 0x85, 0x08, 0x17, 0x2b, 0x58, 0x23, 0x58, 0x0d, 0x6b, 0x10, 0x84,
	0x63, 0x94, 0xc7, 0xc5, 0x0a, 0x56, 0x25, 0x94, 0xca, 0xc5, 0x9d, 0x9e, 0x5a, 0x12, 0x9c, 0x5d,
	0x0a, 0xe1, 0x6a, 0xe8, 0xe1, 0xb3, 0x5c, 0x0f, 0x61, 0xb3, 0x94, 0x16, 0x01, 0x95, 0x48, 0x6e,
	0x50, 0x62, 0x48, 0x62, 0x03, 0x7b, 0xcd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xf1, 0x2c,
	0x0e, 0xeb, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StockClient is the client API for Stock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StockClient interface {
	GetSkuStock(ctx context.Context, in *SkuMessage, opts ...grpc.CallOption) (*StockMessage, error)
}

type stockClient struct {
	cc *grpc.ClientConn
}

func NewStockClient(cc *grpc.ClientConn) StockClient {
	return &stockClient{cc}
}

func (c *stockClient) GetSkuStock(ctx context.Context, in *SkuMessage, opts ...grpc.CallOption) (*StockMessage, error) {
	out := new(StockMessage)
	err := c.cc.Invoke(ctx, "/service.bidogo.product.stock.Stock/getSkuStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockServer is the server API for Stock service.
type StockServer interface {
	GetSkuStock(context.Context, *SkuMessage) (*StockMessage, error)
}

// UnimplementedStockServer can be embedded to have forward compatible implementations.
type UnimplementedStockServer struct {
}

func (*UnimplementedStockServer) GetSkuStock(ctx context.Context, req *SkuMessage) (*StockMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSkuStock not implemented")
}

func RegisterStockServer(s *grpc.Server, srv StockServer) {
	s.RegisterService(&_Stock_serviceDesc, srv)
}

func _Stock_GetSkuStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SkuMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).GetSkuStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.bidogo.product.stock.Stock/GetSkuStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).GetSkuStock(ctx, req.(*SkuMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Stock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.bidogo.product.stock.Stock",
	HandlerType: (*StockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getSkuStock",
			Handler:    _Stock_GetSkuStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
