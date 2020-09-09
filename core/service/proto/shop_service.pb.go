// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shop_service.proto

package proto // import "."

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

// 商店,需重构
type SShop2 struct {
	ID                   int32             `protobuf:"zigzag32,1,opt,name=ID,proto3" json:"ID,omitempty"`
	VendorId             int32             `protobuf:"zigzag32,2,opt,name=VendorId,proto3" json:"VendorId,omitempty"`
	ShopType             int32             `protobuf:"zigzag32,3,opt,name=ShopType,proto3" json:"ShopType,omitempty"`
	Name                 string            `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	State                int32             `protobuf:"zigzag32,5,opt,name=State,proto3" json:"State,omitempty"`
	OpeningState         int32             `protobuf:"zigzag32,6,opt,name=OpeningState,proto3" json:"OpeningState,omitempty"`
	Data                 map[string]string `protobuf:"bytes,7,rep,name=Data,proto3" json:"Data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SShop2) Reset()         { *m = SShop2{} }
func (m *SShop2) String() string { return proto.CompactTextString(m) }
func (*SShop2) ProtoMessage()    {}
func (*SShop2) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_service_90ace9cc188534b0, []int{0}
}
func (m *SShop2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SShop2.Unmarshal(m, b)
}
func (m *SShop2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SShop2.Marshal(b, m, deterministic)
}
func (dst *SShop2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SShop2.Merge(dst, src)
}
func (m *SShop2) XXX_Size() int {
	return xxx_messageInfo_SShop2.Size(m)
}
func (m *SShop2) XXX_DiscardUnknown() {
	xxx_messageInfo_SShop2.DiscardUnknown(m)
}

var xxx_messageInfo_SShop2 proto.InternalMessageInfo

func (m *SShop2) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *SShop2) GetVendorId() int32 {
	if m != nil {
		return m.VendorId
	}
	return 0
}

func (m *SShop2) GetShopType() int32 {
	if m != nil {
		return m.ShopType
	}
	return 0
}

func (m *SShop2) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SShop2) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *SShop2) GetOpeningState() int32 {
	if m != nil {
		return m.OpeningState
	}
	return 0
}

func (m *SShop2) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

// 店铺
type SShop struct {
	// * 店铺编号
	Id int32 `protobuf:"zigzag32,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// * 商户编号
	VendorId int32 `protobuf:"zigzag32,2,opt,name=VendorId,proto3" json:"VendorId,omitempty"`
	// * 店铺名称
	ShopName string `protobuf:"bytes,3,opt,name=ShopName,proto3" json:"ShopName,omitempty"`
	// * 店铺标志
	Logo string `protobuf:"bytes,4,opt,name=Logo,proto3" json:"Logo,omitempty"`
	// * 自定义 域名
	Host string `protobuf:"bytes,5,opt,name=Host,proto3" json:"Host,omitempty"`
	// * 个性化域名
	Alias string `protobuf:"bytes,6,opt,name=Alias,proto3" json:"Alias,omitempty"`
	// * 电话
	Tel string `protobuf:"bytes,7,opt,name=Tel,proto3" json:"Tel,omitempty"`
	// * 地址
	Addr string `protobuf:"bytes,8,opt,name=Addr,proto3" json:"Addr,omitempty"`
	// * 店铺标题
	ShopTitle string `protobuf:"bytes,9,opt,name=ShopTitle,proto3" json:"ShopTitle,omitempty"`
	// * 店铺公告
	ShopNotice string `protobuf:"bytes,10,opt,name=ShopNotice,proto3" json:"ShopNotice,omitempty"`
	// * 标志
	Flag int32 `protobuf:"zigzag32,11,opt,name=Flag,proto3" json:"Flag,omitempty"`
	// * 状态
	State                int32    `protobuf:"zigzag32,12,opt,name=State,proto3" json:"State,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SShop) Reset()         { *m = SShop{} }
func (m *SShop) String() string { return proto.CompactTextString(m) }
func (*SShop) ProtoMessage()    {}
func (*SShop) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_service_90ace9cc188534b0, []int{1}
}
func (m *SShop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SShop.Unmarshal(m, b)
}
func (m *SShop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SShop.Marshal(b, m, deterministic)
}
func (dst *SShop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SShop.Merge(dst, src)
}
func (m *SShop) XXX_Size() int {
	return xxx_messageInfo_SShop.Size(m)
}
func (m *SShop) XXX_DiscardUnknown() {
	xxx_messageInfo_SShop.DiscardUnknown(m)
}

var xxx_messageInfo_SShop proto.InternalMessageInfo

func (m *SShop) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SShop) GetVendorId() int32 {
	if m != nil {
		return m.VendorId
	}
	return 0
}

func (m *SShop) GetShopName() string {
	if m != nil {
		return m.ShopName
	}
	return ""
}

func (m *SShop) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *SShop) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *SShop) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *SShop) GetTel() string {
	if m != nil {
		return m.Tel
	}
	return ""
}

func (m *SShop) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *SShop) GetShopTitle() string {
	if m != nil {
		return m.ShopTitle
	}
	return ""
}

func (m *SShop) GetShopNotice() string {
	if m != nil {
		return m.ShopNotice
	}
	return ""
}

func (m *SShop) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *SShop) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

// 店铺
type SStore struct {
	ID                   int32    `protobuf:"zigzag32,1,opt,name=ID,proto3" json:"ID,omitempty"`
	VendorId             int32    `protobuf:"zigzag32,2,opt,name=VendorId,proto3" json:"VendorId,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Alias                string   `protobuf:"bytes,4,opt,name=Alias,proto3" json:"Alias,omitempty"`
	Host                 string   `protobuf:"bytes,5,opt,name=Host,proto3" json:"Host,omitempty"`
	Logo                 string   `protobuf:"bytes,6,opt,name=Logo,proto3" json:"Logo,omitempty"`
	State                int32    `protobuf:"zigzag32,7,opt,name=State,proto3" json:"State,omitempty"`
	OpeningState         int32    `protobuf:"zigzag32,8,opt,name=OpeningState,proto3" json:"OpeningState,omitempty"`
	StorePhone           string   `protobuf:"bytes,9,opt,name=StorePhone,proto3" json:"StorePhone,omitempty"`
	StoreTitle           string   `protobuf:"bytes,10,opt,name=StoreTitle,proto3" json:"StoreTitle,omitempty"`
	StoreNotice          string   `protobuf:"bytes,11,opt,name=StoreNotice,proto3" json:"StoreNotice,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SStore) Reset()         { *m = SStore{} }
func (m *SStore) String() string { return proto.CompactTextString(m) }
func (*SStore) ProtoMessage()    {}
func (*SStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_service_90ace9cc188534b0, []int{2}
}
func (m *SStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SStore.Unmarshal(m, b)
}
func (m *SStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SStore.Marshal(b, m, deterministic)
}
func (dst *SStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SStore.Merge(dst, src)
}
func (m *SStore) XXX_Size() int {
	return xxx_messageInfo_SStore.Size(m)
}
func (m *SStore) XXX_DiscardUnknown() {
	xxx_messageInfo_SStore.DiscardUnknown(m)
}

var xxx_messageInfo_SStore proto.InternalMessageInfo

func (m *SStore) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *SStore) GetVendorId() int32 {
	if m != nil {
		return m.VendorId
	}
	return 0
}

func (m *SStore) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SStore) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *SStore) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *SStore) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *SStore) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *SStore) GetOpeningState() int32 {
	if m != nil {
		return m.OpeningState
	}
	return 0
}

func (m *SStore) GetStorePhone() string {
	if m != nil {
		return m.StorePhone
	}
	return ""
}

func (m *SStore) GetStoreTitle() string {
	if m != nil {
		return m.StoreTitle
	}
	return ""
}

func (m *SStore) GetStoreNotice() string {
	if m != nil {
		return m.StoreNotice
	}
	return ""
}

type TurnShopRequest struct {
	ShopId               int32    `protobuf:"zigzag32,1,opt,name=shopId,proto3" json:"shopId,omitempty"`
	On                   bool     `protobuf:"varint,2,opt,name=on,proto3" json:"on,omitempty"`
	Reason               string   `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TurnShopRequest) Reset()         { *m = TurnShopRequest{} }
func (m *TurnShopRequest) String() string { return proto.CompactTextString(m) }
func (*TurnShopRequest) ProtoMessage()    {}
func (*TurnShopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_service_90ace9cc188534b0, []int{3}
}
func (m *TurnShopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TurnShopRequest.Unmarshal(m, b)
}
func (m *TurnShopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TurnShopRequest.Marshal(b, m, deterministic)
}
func (dst *TurnShopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TurnShopRequest.Merge(dst, src)
}
func (m *TurnShopRequest) XXX_Size() int {
	return xxx_messageInfo_TurnShopRequest.Size(m)
}
func (m *TurnShopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TurnShopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TurnShopRequest proto.InternalMessageInfo

func (m *TurnShopRequest) GetShopId() int32 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *TurnShopRequest) GetOn() bool {
	if m != nil {
		return m.On
	}
	return false
}

func (m *TurnShopRequest) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterType((*SShop2)(nil), "SShop2")
	proto.RegisterMapType((map[string]string)(nil), "SShop2.DataEntry")
	proto.RegisterType((*SShop)(nil), "SShop")
	proto.RegisterType((*SStore)(nil), "SStore")
	proto.RegisterType((*TurnShopRequest)(nil), "TurnShopRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShopServiceClient is the client API for ShopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShopServiceClient interface {
	// * 获取店铺,shopId
	GetShop(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SShop, error)
	// * 获取商户的店铺,vendorId
	GetVendorShop(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SShop, error)
	// * 获取门店,storeId
	GetStore(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SStore, error)
	// * 根据主机头获取店铺编号,host
	QueryShopByHost(ctx context.Context, in *String, opts ...grpc.CallOption) (*Int64, error)
	// 获取门店
	// rpc GetOfflineShop(1:sint32 shopId)returns(Shop)
	// 打开或关闭商店
	TurnShop(ctx context.Context, in *TurnShopRequest, opts ...grpc.CallOption) (*Result, error)
}

type shopServiceClient struct {
	cc *grpc.ClientConn
}

func NewShopServiceClient(cc *grpc.ClientConn) ShopServiceClient {
	return &shopServiceClient{cc}
}

func (c *shopServiceClient) GetShop(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SShop, error) {
	out := new(SShop)
	err := c.cc.Invoke(ctx, "/ShopService/GetShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) GetVendorShop(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SShop, error) {
	out := new(SShop)
	err := c.cc.Invoke(ctx, "/ShopService/GetVendorShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) GetStore(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SStore, error) {
	out := new(SStore)
	err := c.cc.Invoke(ctx, "/ShopService/GetStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) QueryShopByHost(ctx context.Context, in *String, opts ...grpc.CallOption) (*Int64, error) {
	out := new(Int64)
	err := c.cc.Invoke(ctx, "/ShopService/QueryShopByHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) TurnShop(ctx context.Context, in *TurnShopRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ShopService/TurnShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServiceServer is the server API for ShopService service.
type ShopServiceServer interface {
	// * 获取店铺,shopId
	GetShop(context.Context, *Int64) (*SShop, error)
	// * 获取商户的店铺,vendorId
	GetVendorShop(context.Context, *Int64) (*SShop, error)
	// * 获取门店,storeId
	GetStore(context.Context, *Int64) (*SStore, error)
	// * 根据主机头获取店铺编号,host
	QueryShopByHost(context.Context, *String) (*Int64, error)
	// 获取门店
	// rpc GetOfflineShop(1:sint32 shopId)returns(Shop)
	// 打开或关闭商店
	TurnShop(context.Context, *TurnShopRequest) (*Result, error)
}

func RegisterShopServiceServer(s *grpc.Server, srv ShopServiceServer) {
	s.RegisterService(&_ShopService_serviceDesc, srv)
}

func _ShopService_GetShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/GetShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetShop(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_GetVendorShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetVendorShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/GetVendorShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetVendorShop(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_GetStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/GetStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetStore(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_QueryShopByHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).QueryShopByHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/QueryShopByHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).QueryShopByHost(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_TurnShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TurnShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).TurnShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/TurnShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).TurnShop(ctx, req.(*TurnShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShopService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ShopService",
	HandlerType: (*ShopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShop",
			Handler:    _ShopService_GetShop_Handler,
		},
		{
			MethodName: "GetVendorShop",
			Handler:    _ShopService_GetVendorShop_Handler,
		},
		{
			MethodName: "GetStore",
			Handler:    _ShopService_GetStore_Handler,
		},
		{
			MethodName: "QueryShopByHost",
			Handler:    _ShopService_QueryShopByHost_Handler,
		},
		{
			MethodName: "TurnShop",
			Handler:    _ShopService_TurnShop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop_service.proto",
}

func init() { proto.RegisterFile("shop_service.proto", fileDescriptor_shop_service_90ace9cc188534b0) }

var fileDescriptor_shop_service_90ace9cc188534b0 = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x8d, 0x9d, 0xd4, 0x7f, 0xc6, 0xfd, 0xfd, 0xda, 0xae, 0x10, 0x5a, 0x99, 0x3f, 0x0a, 0x16,
	0x88, 0x9e, 0x7c, 0x08, 0x08, 0x10, 0x9c, 0x5a, 0x05, 0x4a, 0x24, 0x04, 0xd4, 0x89, 0x38, 0x70,
	0x41, 0xa6, 0x5e, 0x25, 0x16, 0x66, 0xd7, 0xd8, 0x9b, 0x4a, 0xfe, 0x2c, 0xdc, 0xf8, 0x1a, 0x7c,
	0x37, 0x84, 0x66, 0xd6, 0x76, 0x5d, 0x54, 0x10, 0x9c, 0x32, 0x6f, 0xde, 0xec, 0x66, 0xde, 0x9b,
	0x59, 0x03, 0xab, 0x37, 0xaa, 0xfc, 0x50, 0x8b, 0xea, 0x3c, 0x3f, 0x13, 0x71, 0x59, 0x29, 0xad,
	0xc2, 0x40, 0xeb, 0xa6, 0x6c, 0x41, 0xf4, 0xc3, 0x02, 0x67, 0xb9, 0xdc, 0xa8, 0x72, 0xc6, 0xfe,
	0x07, 0x7b, 0x31, 0xe7, 0xd6, 0xd4, 0x3a, 0x3c, 0x48, 0xec, 0xc5, 0x9c, 0x85, 0xe0, 0xbd, 0x13,
	0x32, 0x53, 0xd5, 0x22, 0xe3, 0x36, 0x65, 0x7b, 0x8c, 0x1c, 0x1e, 0x5a, 0x35, 0xa5, 0xe0, 0x63,
	0xc3, 0x75, 0x98, 0x31, 0x98, 0xbc, 0x4e, 0x3f, 0x0b, 0x3e, 0x99, 0x5a, 0x87, 0x7e, 0x42, 0x31,
	0xbb, 0x06, 0x3b, 0x4b, 0x9d, 0x6a, 0xc1, 0x77, 0xa8, 0xd8, 0x00, 0x16, 0xc1, 0xee, 0x9b, 0x52,
	0xc8, 0x5c, 0xae, 0x0d, 0xe9, 0x10, 0x79, 0x29, 0xc7, 0xee, 0xc1, 0x64, 0x9e, 0xea, 0x94, 0xbb,
	0xd3, 0xf1, 0x61, 0x30, 0x3b, 0x88, 0x4d, 0xb3, 0x31, 0xe6, 0x9e, 0x4b, 0x5d, 0x35, 0x09, 0xd1,
	0xe1, 0x63, 0xf0, 0xfb, 0x14, 0xdb, 0x87, 0xf1, 0x27, 0xd1, 0x90, 0x14, 0x3f, 0xc1, 0x10, 0xff,
	0xff, 0x3c, 0x2d, 0xb6, 0x82, 0x84, 0xf8, 0x89, 0x01, 0x4f, 0xed, 0x27, 0x56, 0xf4, 0xd5, 0x86,
	0x1d, 0xba, 0x93, 0xf4, 0x67, 0xbd, 0xfe, 0xec, 0x6f, 0xf4, 0x93, 0xce, 0x31, 0x5d, 0xd9, 0x63,
	0xd4, 0xff, 0x4a, 0xad, 0x55, 0xa7, 0x1f, 0x63, 0xcc, 0xbd, 0x54, 0xb5, 0x26, 0xf9, 0x7e, 0x42,
	0x31, 0xf6, 0x74, 0x54, 0xe4, 0x69, 0x4d, 0xb2, 0xfd, 0xc4, 0x00, 0xec, 0x7d, 0x25, 0x0a, 0xee,
	0x9a, 0xde, 0x57, 0xa2, 0xc0, 0xb3, 0x47, 0x59, 0x56, 0x71, 0xcf, 0x9c, 0xc5, 0x98, 0xdd, 0x04,
	0x9f, 0xfc, 0xce, 0x75, 0x21, 0xb8, 0x4f, 0xc4, 0x45, 0x82, 0xdd, 0x06, 0xa0, 0x6e, 0x94, 0xce,
	0xcf, 0x04, 0x07, 0xa2, 0x07, 0x19, 0xbc, 0xf1, 0x45, 0x91, 0xae, 0x79, 0x40, 0xaa, 0x28, 0xbe,
	0x98, 0xd0, 0xee, 0x60, 0x42, 0xd1, 0x37, 0x1b, 0xd7, 0x43, 0xab, 0x4a, 0xfc, 0xd3, 0x7a, 0x74,
	0x2b, 0x30, 0xbe, 0xbc, 0x02, 0x46, 0xee, 0x64, 0x28, 0xf7, 0x2a, 0x63, 0x3a, 0x03, 0x9d, 0x81,
	0x81, 0x7d, 0x7b, 0xee, 0x9f, 0x16, 0xc8, 0xbb, 0x62, 0x81, 0xd0, 0x0c, 0x14, 0xf0, 0x76, 0xa3,
	0x64, 0xe7, 0xd5, 0x20, 0xd3, 0xf3, 0xc6, 0x4b, 0x18, 0xf0, 0xc6, 0xcc, 0x29, 0x04, 0x84, 0x5a,
	0x37, 0x03, 0x2a, 0x18, 0xa6, 0xa2, 0x53, 0xd8, 0x5b, 0x6d, 0x2b, 0x89, 0x06, 0x27, 0xe2, 0xcb,
	0x56, 0xd4, 0x9a, 0x5d, 0x07, 0x07, 0x5f, 0x5e, 0xbf, 0x4f, 0x2d, 0x42, 0x13, 0x95, 0x24, 0xbb,
	0xbc, 0xc4, 0x56, 0x12, 0xeb, 0x2a, 0x91, 0xd6, 0x4a, 0xb6, 0x56, 0xb5, 0x68, 0xf6, 0xdd, 0x82,
	0x00, 0xef, 0x5b, 0x9a, 0x97, 0xcb, 0x6e, 0x80, 0x7b, 0x22, 0x34, 0xad, 0xa9, 0x13, 0x2f, 0xa4,
	0x7e, 0xf4, 0x30, 0x74, 0xcc, 0x53, 0x88, 0x46, 0xec, 0x0e, 0xfc, 0x77, 0x22, 0xb4, 0x31, 0xff,
	0x37, 0x25, 0xb7, 0xc0, 0xc3, 0xf3, 0x34, 0xc8, 0x8e, 0x75, 0x63, 0x33, 0xd9, 0x68, 0xc4, 0xee,
	0xc2, 0xde, 0xe9, 0x56, 0x54, 0x0d, 0x56, 0x1f, 0x37, 0x34, 0x04, 0x37, 0x5e, 0xea, 0x2a, 0x97,
	0xeb, 0xb0, 0x2d, 0x8f, 0x46, 0xec, 0x3e, 0x78, 0x9d, 0x4e, 0xb6, 0x1f, 0xff, 0x22, 0x39, 0x74,
	0xe3, 0x44, 0xd4, 0xdb, 0x42, 0x47, 0xa3, 0x63, 0xff, 0xbd, 0x1b, 0x3f, 0xa3, 0xef, 0xcb, 0x47,
	0x87, 0x7e, 0x1e, 0xfc, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xcd, 0x6d, 0x41, 0x89, 0x04, 0x00,
	0x00,
}