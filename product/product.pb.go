// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product.proto

package product

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type FetchQuery struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchQuery) Reset()         { *m = FetchQuery{} }
func (m *FetchQuery) String() string { return proto.CompactTextString(m) }
func (*FetchQuery) ProtoMessage()    {}
func (*FetchQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{0}
}

func (m *FetchQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchQuery.Unmarshal(m, b)
}
func (m *FetchQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchQuery.Marshal(b, m, deterministic)
}
func (m *FetchQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchQuery.Merge(m, src)
}
func (m *FetchQuery) XXX_Size() int {
	return xxx_messageInfo_FetchQuery.Size(m)
}
func (m *FetchQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchQuery.DiscardUnknown(m)
}

var xxx_messageInfo_FetchQuery proto.InternalMessageInfo

func (m *FetchQuery) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type FetchReply struct {
	Fetched              uint32   `protobuf:"varint,1,opt,name=fetched,proto3" json:"fetched,omitempty"`
	Created              uint32   `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	Updated              uint32   `protobuf:"varint,3,opt,name=updated,proto3" json:"updated,omitempty"`
	Node                 uint32   `protobuf:"varint,4,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchReply) Reset()         { *m = FetchReply{} }
func (m *FetchReply) String() string { return proto.CompactTextString(m) }
func (*FetchReply) ProtoMessage()    {}
func (*FetchReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{1}
}

func (m *FetchReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchReply.Unmarshal(m, b)
}
func (m *FetchReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchReply.Marshal(b, m, deterministic)
}
func (m *FetchReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchReply.Merge(m, src)
}
func (m *FetchReply) XXX_Size() int {
	return xxx_messageInfo_FetchReply.Size(m)
}
func (m *FetchReply) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchReply.DiscardUnknown(m)
}

var xxx_messageInfo_FetchReply proto.InternalMessageInfo

func (m *FetchReply) GetFetched() uint32 {
	if m != nil {
		return m.Fetched
	}
	return 0
}

func (m *FetchReply) GetCreated() uint32 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *FetchReply) GetUpdated() uint32 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *FetchReply) GetNode() uint32 {
	if m != nil {
		return m.Node
	}
	return 0
}

type Page struct {
	Limit                uint32   `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               uint32   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Page) Reset()         { *m = Page{} }
func (m *Page) String() string { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()    {}
func (*Page) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{2}
}

func (m *Page) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Page.Unmarshal(m, b)
}
func (m *Page) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Page.Marshal(b, m, deterministic)
}
func (m *Page) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Page.Merge(m, src)
}
func (m *Page) XXX_Size() int {
	return xxx_messageInfo_Page.Size(m)
}
func (m *Page) XXX_DiscardUnknown() {
	xxx_messageInfo_Page.DiscardUnknown(m)
}

var xxx_messageInfo_Page proto.InternalMessageInfo

func (m *Page) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Page) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type Sort struct {
	Order                bool     `protobuf:"varint,1,opt,name=order,proto3" json:"order,omitempty"`
	By                   string   `protobuf:"bytes,2,opt,name=by,proto3" json:"by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sort) Reset()         { *m = Sort{} }
func (m *Sort) String() string { return proto.CompactTextString(m) }
func (*Sort) ProtoMessage()    {}
func (*Sort) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{3}
}

func (m *Sort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sort.Unmarshal(m, b)
}
func (m *Sort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sort.Marshal(b, m, deterministic)
}
func (m *Sort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sort.Merge(m, src)
}
func (m *Sort) XXX_Size() int {
	return xxx_messageInfo_Sort.Size(m)
}
func (m *Sort) XXX_DiscardUnknown() {
	xxx_messageInfo_Sort.DiscardUnknown(m)
}

var xxx_messageInfo_Sort proto.InternalMessageInfo

func (m *Sort) GetOrder() bool {
	if m != nil {
		return m.Order
	}
	return false
}

func (m *Sort) GetBy() string {
	if m != nil {
		return m.By
	}
	return ""
}

type ListQuery struct {
	Page                 *Page    `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Sort                 []*Sort  `protobuf:"bytes,2,rep,name=sort,proto3" json:"sort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListQuery) Reset()         { *m = ListQuery{} }
func (m *ListQuery) String() string { return proto.CompactTextString(m) }
func (*ListQuery) ProtoMessage()    {}
func (*ListQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{4}
}

func (m *ListQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListQuery.Unmarshal(m, b)
}
func (m *ListQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListQuery.Marshal(b, m, deterministic)
}
func (m *ListQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListQuery.Merge(m, src)
}
func (m *ListQuery) XXX_Size() int {
	return xxx_messageInfo_ListQuery.Size(m)
}
func (m *ListQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ListQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ListQuery proto.InternalMessageInfo

func (m *ListQuery) GetPage() *Page {
	if m != nil {
		return m.Page
	}
	return nil
}

func (m *ListQuery) GetSort() []*Sort {
	if m != nil {
		return m.Sort
	}
	return nil
}

type Product struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price                float64              `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	Changes              uint32               `protobuf:"varint,3,opt,name=changes,proto3" json:"changes,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{5}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Product) GetChanges() uint32 {
	if m != nil {
		return m.Changes
	}
	return 0
}

func (m *Product) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

type ListReply struct {
	Products             []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	Node                 uint32     `protobuf:"varint,2,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListReply) Reset()         { *m = ListReply{} }
func (m *ListReply) String() string { return proto.CompactTextString(m) }
func (*ListReply) ProtoMessage()    {}
func (*ListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{6}
}

func (m *ListReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListReply.Unmarshal(m, b)
}
func (m *ListReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListReply.Marshal(b, m, deterministic)
}
func (m *ListReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListReply.Merge(m, src)
}
func (m *ListReply) XXX_Size() int {
	return xxx_messageInfo_ListReply.Size(m)
}
func (m *ListReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListReply proto.InternalMessageInfo

func (m *ListReply) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *ListReply) GetNode() uint32 {
	if m != nil {
		return m.Node
	}
	return 0
}

func init() {
	proto.RegisterType((*FetchQuery)(nil), "product.FetchQuery")
	proto.RegisterType((*FetchReply)(nil), "product.FetchReply")
	proto.RegisterType((*Page)(nil), "product.Page")
	proto.RegisterType((*Sort)(nil), "product.Sort")
	proto.RegisterType((*ListQuery)(nil), "product.ListQuery")
	proto.RegisterType((*Product)(nil), "product.Product")
	proto.RegisterType((*ListReply)(nil), "product.ListReply")
}

func init() { proto.RegisterFile("product.proto", fileDescriptor_f0fd8b59378f44a5) }

var fileDescriptor_f0fd8b59378f44a5 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x52, 0x3d, 0x8f, 0x9b, 0x40,
	0x10, 0x0d, 0x98, 0x3b, 0x1f, 0x63, 0xf9, 0x64, 0x6d, 0xa2, 0x08, 0x51, 0xe4, 0x63, 0xab, 0x14,
	0x16, 0x8e, 0xec, 0x74, 0xe9, 0x52, 0xa4, 0x4a, 0x24, 0x7b, 0x9d, 0x2a, 0x1d, 0x1f, 0x03, 0x46,
	0x02, 0x2f, 0x5a, 0x96, 0x44, 0x14, 0xf9, 0xef, 0xd1, 0xce, 0x2e, 0x58, 0xbe, 0xca, 0xfb, 0xe6,
	0xbd, 0xf1, 0xbc, 0x79, 0x03, 0xac, 0x3b, 0x25, 0x8b, 0x21, 0xd7, 0x49, 0xa7, 0xa4, 0x96, 0x6c,
	0xe9, 0x60, 0xfc, 0xbe, 0x92, 0xb2, 0x6a, 0x70, 0x47, 0xe5, 0x6c, 0x28, 0x77, 0xba, 0x6e, 0xb1,
	0xd7, 0x69, 0xdb, 0x59, 0x25, 0x7f, 0x07, 0xf0, 0x1d, 0x75, 0x7e, 0x39, 0x0d, 0xa8, 0x46, 0xb6,
	0x81, 0xc5, 0xa0, 0x9a, 0xc8, 0xfb, 0xe0, 0x7d, 0x0a, 0x85, 0x79, 0xf2, 0xab, 0xe3, 0x05, 0x76,
	0xcd, 0xc8, 0x22, 0x58, 0x96, 0x06, 0x61, 0x41, 0x9a, 0xb5, 0x98, 0xa0, 0x61, 0x72, 0x85, 0xa9,
	0xc6, 0x22, 0xf2, 0x2d, 0xe3, 0xa0, 0x61, 0x86, 0xae, 0x20, 0x66, 0x61, 0x19, 0x07, 0x19, 0x83,
	0xe0, 0x2a, 0x0b, 0x8c, 0x02, 0x2a, 0xd3, 0x9b, 0x7f, 0x81, 0xe0, 0x98, 0x56, 0xc8, 0xde, 0xc0,
	0x43, 0x53, 0xb7, 0xb5, 0x76, 0x73, 0x2c, 0x60, 0x6f, 0xe1, 0x51, 0x96, 0x65, 0x8f, 0xda, 0x0d,
	0x71, 0x88, 0x6f, 0x21, 0x38, 0x4b, 0xa5, 0x4d, 0x97, 0x54, 0x05, 0x2a, 0xea, 0x7a, 0x12, 0x16,
	0xb0, 0x67, 0xf0, 0xb3, 0x91, 0x3a, 0x42, 0xe1, 0x67, 0x23, 0x3f, 0x41, 0xf8, 0xa3, 0xee, 0xb5,
	0x5d, 0xf9, 0x23, 0x04, 0x5d, 0x5a, 0x21, 0x75, 0xac, 0xf6, 0xeb, 0x64, 0x0a, 0xd2, 0xb8, 0x10,
	0x44, 0x19, 0x49, 0x2f, 0x95, 0x99, 0xb9, 0xb8, 0x93, 0x98, 0x91, 0x82, 0x28, 0xfe, 0x0f, 0x96,
	0x47, 0x5b, 0xa5, 0xad, 0xd2, 0x16, 0x5d, 0x88, 0xf4, 0x36, 0xbe, 0x3a, 0x55, 0xe7, 0x48, 0x26,
	0x3c, 0x61, 0x01, 0x65, 0x76, 0x49, 0xaf, 0x15, 0xf6, 0x53, 0x32, 0x0e, 0xb2, 0x04, 0x02, 0x13,
	0x11, 0x25, 0xb3, 0xda, 0xc7, 0x89, 0xbd, 0x62, 0x32, 0x5d, 0x31, 0xf9, 0x35, 0x5d, 0x51, 0x90,
	0x8e, 0xff, 0xb4, 0x1b, 0xd9, 0x23, 0x6d, 0xe1, 0xc9, 0x39, 0xec, 0x23, 0x8f, 0x2c, 0x6f, 0x6e,
	0x5b, 0xd9, 0x5f, 0x31, 0x2b, 0xe6, 0x23, 0xf8, 0xb7, 0x23, 0xec, 0xff, 0xc2, 0xb3, 0x13, 0x9e,
	0x51, 0xfd, 0x31, 0x56, 0x0f, 0xf0, 0x40, 0x9f, 0x01, 0x7b, 0x3d, 0xff, 0xd5, 0xed, 0xb3, 0x89,
	0x5f, 0x14, 0xc9, 0x06, 0x7f, 0xc5, 0x3e, 0x43, 0x60, 0x5c, 0x31, 0x36, 0xd3, 0x73, 0xec, 0xf1,
	0x7d, 0xcd, 0x75, 0x7c, 0x5b, 0xfd, 0x0e, 0x93, 0xaf, 0x8e, 0xc8, 0x1e, 0x69, 0xdd, 0xc3, 0xff,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x73, 0xf3, 0x57, 0xdc, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	Fetch(ctx context.Context, in *FetchQuery, opts ...grpc.CallOption) (*FetchReply, error)
	List(ctx context.Context, in *ListQuery, opts ...grpc.CallOption) (*ListReply, error)
}

type productServiceClient struct {
	cc *grpc.ClientConn
}

func NewProductServiceClient(cc *grpc.ClientConn) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) Fetch(ctx context.Context, in *FetchQuery, opts ...grpc.CallOption) (*FetchReply, error) {
	out := new(FetchReply)
	err := c.cc.Invoke(ctx, "/product.ProductService/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) List(ctx context.Context, in *ListQuery, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := c.cc.Invoke(ctx, "/product.ProductService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	Fetch(context.Context, *FetchQuery) (*FetchReply, error)
	List(context.Context, *ListQuery) (*ListReply, error)
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) Fetch(ctx context.Context, req *FetchQuery) (*FetchReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (*UnimplementedProductServiceServer) List(ctx context.Context, req *ListQuery) (*ListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).Fetch(ctx, req.(*FetchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).List(ctx, req.(*ListQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _ProductService_Fetch_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ProductService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
