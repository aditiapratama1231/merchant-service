// Code generated by protoc-gen-go. DO NOT EDIT.
// source: products.proto

package pb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type ProductRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductRequest) Reset()         { *m = ProductRequest{} }
func (m *ProductRequest) String() string { return proto.CompactTextString(m) }
func (*ProductRequest) ProtoMessage()    {}
func (*ProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c6e54f42122eb82, []int{0}
}

func (m *ProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductRequest.Unmarshal(m, b)
}
func (m *ProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductRequest.Marshal(b, m, deterministic)
}
func (m *ProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductRequest.Merge(m, src)
}
func (m *ProductRequest) XXX_Size() int {
	return xxx_messageInfo_ProductRequest.Size(m)
}
func (m *ProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProductRequest proto.InternalMessageInfo

func (m *ProductRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ProductResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductResponse) Reset()         { *m = ProductResponse{} }
func (m *ProductResponse) String() string { return proto.CompactTextString(m) }
func (*ProductResponse) ProtoMessage()    {}
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c6e54f42122eb82, []int{1}
}

func (m *ProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductResponse.Unmarshal(m, b)
}
func (m *ProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductResponse.Marshal(b, m, deterministic)
}
func (m *ProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductResponse.Merge(m, src)
}
func (m *ProductResponse) XXX_Size() int {
	return xxx_messageInfo_ProductResponse.Size(m)
}
func (m *ProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProductResponse proto.InternalMessageInfo

func (m *ProductResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ProductResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*ProductRequest)(nil), "pb.ProductRequest")
	proto.RegisterType((*ProductResponse)(nil), "pb.ProductResponse")
}

func init() { proto.RegisterFile("products.proto", fileDescriptor_8c6e54f42122eb82) }

var fileDescriptor_8c6e54f42122eb82 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x28, 0xca, 0x4f,
	0x29, 0x4d, 0x2e, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x52,
	0xe0, 0xe2, 0x0b, 0x80, 0x88, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xf1, 0x71, 0x31,
	0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x31, 0x65, 0xa6, 0x28, 0xd9, 0x72, 0xf1,
	0xc3, 0x55, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x49, 0x70, 0xb1, 0xe7, 0xa6, 0x16, 0x17,
	0x27, 0xa6, 0xa7, 0x82, 0xd5, 0x71, 0x06, 0xc1, 0xb8, 0x42, 0x02, 0x5c, 0xcc, 0xa9, 0x45, 0x45,
	0x12, 0x4c, 0x60, 0x51, 0x10, 0xd3, 0xc8, 0x85, 0x8b, 0x03, 0xaa, 0xbd, 0x58, 0xc8, 0x82, 0x8b,
	0xdb, 0x3d, 0xb5, 0x04, 0xce, 0x15, 0xd2, 0x2b, 0x48, 0xd2, 0x43, 0xb5, 0x5d, 0x4a, 0x18, 0x45,
	0x0c, 0x62, 0x9f, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0xc5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x5b, 0xbe, 0xc0, 0x08, 0xc3, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductsClient is the client API for Products service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductsClient interface {
	GetProducts(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
}

type productsClient struct {
	cc *grpc.ClientConn
}

func NewProductsClient(cc *grpc.ClientConn) ProductsClient {
	return &productsClient{cc}
}

func (c *productsClient) GetProducts(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/pb.Products/GetProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductsServer is the server API for Products service.
type ProductsServer interface {
	GetProducts(context.Context, *ProductRequest) (*ProductResponse, error)
}

func RegisterProductsServer(s *grpc.Server, srv ProductsServer) {
	s.RegisterService(&_Products_serviceDesc, srv)
}

func _Products_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Products/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).GetProducts(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Products_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Products",
	HandlerType: (*ProductsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProducts",
			Handler:    _Products_GetProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "products.proto",
}
