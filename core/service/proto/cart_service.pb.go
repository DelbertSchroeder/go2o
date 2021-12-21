// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cart_service.proto

package proto // import "./"

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

type SettlePersistRequest struct {
	// 买家编号
	BuyerId int64 `protobuf:"varint,1,opt,name=BuyerId,proto3" json:"BuyerId"`
	// 店铺编号
	ShopId               int64    `protobuf:"varint,2,opt,name=ShopId,proto3" json:"ShopId"`
	PaymentOpt           int64    `protobuf:"varint,3,opt,name=PaymentOpt,proto3" json:"PaymentOpt"`
	DeliverOpt           int64    `protobuf:"varint,4,opt,name=DeliverOpt,proto3" json:"DeliverOpt"`
	AddressId            int64    `protobuf:"varint,5,opt,name=AddressId,proto3" json:"AddressId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SettlePersistRequest) Reset()         { *m = SettlePersistRequest{} }
func (m *SettlePersistRequest) String() string { return proto.CompactTextString(m) }
func (*SettlePersistRequest) ProtoMessage()    {}
func (*SettlePersistRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cart_service_e4a312d7fef09cb2, []int{0}
}
func (m *SettlePersistRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettlePersistRequest.Unmarshal(m, b)
}
func (m *SettlePersistRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettlePersistRequest.Marshal(b, m, deterministic)
}
func (dst *SettlePersistRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettlePersistRequest.Merge(dst, src)
}
func (m *SettlePersistRequest) XXX_Size() int {
	return xxx_messageInfo_SettlePersistRequest.Size(m)
}
func (m *SettlePersistRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SettlePersistRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SettlePersistRequest proto.InternalMessageInfo

func (m *SettlePersistRequest) GetBuyerId() int64 {
	if m != nil {
		return m.BuyerId
	}
	return 0
}

func (m *SettlePersistRequest) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *SettlePersistRequest) GetPaymentOpt() int64 {
	if m != nil {
		return m.PaymentOpt
	}
	return 0
}

func (m *SettlePersistRequest) GetDeliverOpt() int64 {
	if m != nil {
		return m.DeliverOpt
	}
	return 0
}

func (m *SettlePersistRequest) GetAddressId() int64 {
	if m != nil {
		return m.AddressId
	}
	return 0
}

type CartItemRequest struct {
	// 买家编号
	Id *ShoppingCartId `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id"`
	// 商品编号
	ItemId int64 `protobuf:"zigzag64,2,opt,name=ItemId,proto3" json:"ItemId"`
	// SKU编号
	SkuId int64 `protobuf:"zigzag64,3,opt,name=SkuId,proto3" json:"SkuId"`
	// 数量
	Quantity             int32    `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CartItemRequest) Reset()         { *m = CartItemRequest{} }
func (m *CartItemRequest) String() string { return proto.CompactTextString(m) }
func (*CartItemRequest) ProtoMessage()    {}
func (*CartItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cart_service_e4a312d7fef09cb2, []int{1}
}
func (m *CartItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartItemRequest.Unmarshal(m, b)
}
func (m *CartItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartItemRequest.Marshal(b, m, deterministic)
}
func (dst *CartItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartItemRequest.Merge(dst, src)
}
func (m *CartItemRequest) XXX_Size() int {
	return xxx_messageInfo_CartItemRequest.Size(m)
}
func (m *CartItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CartItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CartItemRequest proto.InternalMessageInfo

func (m *CartItemRequest) GetId() *ShoppingCartId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *CartItemRequest) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *CartItemRequest) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *CartItemRequest) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type CheckSignRequest struct {
	BuyerId              int64             `protobuf:"varint,1,opt,name=BuyerId,proto3" json:"BuyerId"`
	CartCode             string            `protobuf:"bytes,2,opt,name=CartCode,proto3" json:"CartCode"`
	Items                []*SCheckCartItem `protobuf:"bytes,3,rep,name=Items,proto3" json:"Items"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CheckSignRequest) Reset()         { *m = CheckSignRequest{} }
func (m *CheckSignRequest) String() string { return proto.CompactTextString(m) }
func (*CheckSignRequest) ProtoMessage()    {}
func (*CheckSignRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cart_service_e4a312d7fef09cb2, []int{2}
}
func (m *CheckSignRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckSignRequest.Unmarshal(m, b)
}
func (m *CheckSignRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckSignRequest.Marshal(b, m, deterministic)
}
func (dst *CheckSignRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckSignRequest.Merge(dst, src)
}
func (m *CheckSignRequest) XXX_Size() int {
	return xxx_messageInfo_CheckSignRequest.Size(m)
}
func (m *CheckSignRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckSignRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckSignRequest proto.InternalMessageInfo

func (m *CheckSignRequest) GetBuyerId() int64 {
	if m != nil {
		return m.BuyerId
	}
	return 0
}

func (m *CheckSignRequest) GetCartCode() string {
	if m != nil {
		return m.CartCode
	}
	return ""
}

func (m *CheckSignRequest) GetItems() []*SCheckCartItem {
	if m != nil {
		return m.Items
	}
	return nil
}

// 购物车编号
type ShoppingCartId struct {
	// 会员/用户编号
	UserId int64 `protobuf:"zigzag64,1,opt,name=UserId,proto3" json:"UserId"`
	// 购物车标识,当未指定用户时候使用
	CartCode string `protobuf:"bytes,2,opt,name=CartCode,proto3" json:"CartCode"`
	// 是否为批发销售的购物车
	IsWholesale          bool     `protobuf:"varint,3,opt,name=IsWholesale,proto3" json:"IsWholesale"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShoppingCartId) Reset()         { *m = ShoppingCartId{} }
func (m *ShoppingCartId) String() string { return proto.CompactTextString(m) }
func (*ShoppingCartId) ProtoMessage()    {}
func (*ShoppingCartId) Descriptor() ([]byte, []int) {
	return fileDescriptor_cart_service_e4a312d7fef09cb2, []int{3}
}
func (m *ShoppingCartId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShoppingCartId.Unmarshal(m, b)
}
func (m *ShoppingCartId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShoppingCartId.Marshal(b, m, deterministic)
}
func (dst *ShoppingCartId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShoppingCartId.Merge(dst, src)
}
func (m *ShoppingCartId) XXX_Size() int {
	return xxx_messageInfo_ShoppingCartId.Size(m)
}
func (m *ShoppingCartId) XXX_DiscardUnknown() {
	xxx_messageInfo_ShoppingCartId.DiscardUnknown(m)
}

var xxx_messageInfo_ShoppingCartId proto.InternalMessageInfo

func (m *ShoppingCartId) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ShoppingCartId) GetCartCode() string {
	if m != nil {
		return m.CartCode
	}
	return ""
}

func (m *ShoppingCartId) GetIsWholesale() bool {
	if m != nil {
		return m.IsWholesale
	}
	return false
}

type SettleMeta_ struct {
	PaymentOpt           int64               `protobuf:"varint,1,opt,name=PaymentOpt,proto3" json:"PaymentOpt"`
	DeliverOpt           int64               `protobuf:"varint,2,opt,name=DeliverOpt,proto3" json:"DeliverOpt"`
	Shop                 *SettleShopMeta_    `protobuf:"bytes,3,opt,name=Shop,proto3" json:"Shop"`
	Deliver              *SettleDeliverMeta_ `protobuf:"bytes,4,opt,name=Deliver,proto3" json:"Deliver"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SettleMeta_) Reset()         { *m = SettleMeta_{} }
func (m *SettleMeta_) String() string { return proto.CompactTextString(m) }
func (*SettleMeta_) ProtoMessage()    {}
func (*SettleMeta_) Descriptor() ([]byte, []int) {
	return fileDescriptor_cart_service_e4a312d7fef09cb2, []int{4}
}
func (m *SettleMeta_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettleMeta_.Unmarshal(m, b)
}
func (m *SettleMeta_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettleMeta_.Marshal(b, m, deterministic)
}
func (dst *SettleMeta_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettleMeta_.Merge(dst, src)
}
func (m *SettleMeta_) XXX_Size() int {
	return xxx_messageInfo_SettleMeta_.Size(m)
}
func (m *SettleMeta_) XXX_DiscardUnknown() {
	xxx_messageInfo_SettleMeta_.DiscardUnknown(m)
}

var xxx_messageInfo_SettleMeta_ proto.InternalMessageInfo

func (m *SettleMeta_) GetPaymentOpt() int64 {
	if m != nil {
		return m.PaymentOpt
	}
	return 0
}

func (m *SettleMeta_) GetDeliverOpt() int64 {
	if m != nil {
		return m.DeliverOpt
	}
	return 0
}

func (m *SettleMeta_) GetShop() *SettleShopMeta_ {
	if m != nil {
		return m.Shop
	}
	return nil
}

func (m *SettleMeta_) GetDeliver() *SettleDeliverMeta_ {
	if m != nil {
		return m.Deliver
	}
	return nil
}

type SettleShopMeta_ struct {
	ShopId               int64    `protobuf:"varint,1,opt,name=ShopId,proto3" json:"ShopId"`
	ShopName             string   `protobuf:"bytes,2,opt,name=ShopName,proto3" json:"ShopName"`
	Telephone            string   `protobuf:"bytes,3,opt,name=Telephone,proto3" json:"Telephone"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SettleShopMeta_) Reset()         { *m = SettleShopMeta_{} }
func (m *SettleShopMeta_) String() string { return proto.CompactTextString(m) }
func (*SettleShopMeta_) ProtoMessage()    {}
func (*SettleShopMeta_) Descriptor() ([]byte, []int) {
	return fileDescriptor_cart_service_e4a312d7fef09cb2, []int{5}
}
func (m *SettleShopMeta_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettleShopMeta_.Unmarshal(m, b)
}
func (m *SettleShopMeta_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettleShopMeta_.Marshal(b, m, deterministic)
}
func (dst *SettleShopMeta_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettleShopMeta_.Merge(dst, src)
}
func (m *SettleShopMeta_) XXX_Size() int {
	return xxx_messageInfo_SettleShopMeta_.Size(m)
}
func (m *SettleShopMeta_) XXX_DiscardUnknown() {
	xxx_messageInfo_SettleShopMeta_.DiscardUnknown(m)
}

var xxx_messageInfo_SettleShopMeta_ proto.InternalMessageInfo

func (m *SettleShopMeta_) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *SettleShopMeta_) GetShopName() string {
	if m != nil {
		return m.ShopName
	}
	return ""
}

func (m *SettleShopMeta_) GetTelephone() string {
	if m != nil {
		return m.Telephone
	}
	return ""
}

type SettleDeliverMeta_ struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id"`
	ConsigneeName        string   `protobuf:"bytes,2,opt,name=ConsigneeName,proto3" json:"ConsigneeName"`
	ConsigneePhone       string   `protobuf:"bytes,3,opt,name=ConsigneePhone,proto3" json:"ConsigneePhone"`
	Address              string   `protobuf:"bytes,4,opt,name=Address,proto3" json:"Address"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SettleDeliverMeta_) Reset()         { *m = SettleDeliverMeta_{} }
func (m *SettleDeliverMeta_) String() string { return proto.CompactTextString(m) }
func (*SettleDeliverMeta_) ProtoMessage()    {}
func (*SettleDeliverMeta_) Descriptor() ([]byte, []int) {
	return fileDescriptor_cart_service_e4a312d7fef09cb2, []int{6}
}
func (m *SettleDeliverMeta_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettleDeliverMeta_.Unmarshal(m, b)
}
func (m *SettleDeliverMeta_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettleDeliverMeta_.Marshal(b, m, deterministic)
}
func (dst *SettleDeliverMeta_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettleDeliverMeta_.Merge(dst, src)
}
func (m *SettleDeliverMeta_) XXX_Size() int {
	return xxx_messageInfo_SettleDeliverMeta_.Size(m)
}
func (m *SettleDeliverMeta_) XXX_DiscardUnknown() {
	xxx_messageInfo_SettleDeliverMeta_.DiscardUnknown(m)
}

var xxx_messageInfo_SettleDeliverMeta_ proto.InternalMessageInfo

func (m *SettleDeliverMeta_) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SettleDeliverMeta_) GetConsigneeName() string {
	if m != nil {
		return m.ConsigneeName
	}
	return ""
}

func (m *SettleDeliverMeta_) GetConsigneePhone() string {
	if m != nil {
		return m.ConsigneePhone
	}
	return ""
}

func (m *SettleDeliverMeta_) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*SettlePersistRequest)(nil), "SettlePersistRequest")
	proto.RegisterType((*CartItemRequest)(nil), "CartItemRequest")
	proto.RegisterType((*CheckSignRequest)(nil), "CheckSignRequest")
	proto.RegisterType((*ShoppingCartId)(nil), "ShoppingCartId")
	proto.RegisterType((*SettleMeta_)(nil), "SettleMeta_")
	proto.RegisterType((*SettleShopMeta_)(nil), "SettleShopMeta_")
	proto.RegisterType((*SettleDeliverMeta_)(nil), "SettleDeliverMeta_")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CartServiceClient interface {
	// 批发购物车接口
	WholesaleCartV1(ctx context.Context, in *WsCartRequest, opts ...grpc.CallOption) (*Result, error)
	// 获取购物车,当购物车编号不存在时,将返回一个新的购物车
	GetShoppingCart(ctx context.Context, in *ShoppingCartId, opts ...grpc.CallOption) (*SShoppingCart, error)
	// 放入购物车
	PutInCart(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*CartItemResponse, error)
	// 从购物车里删除指定数量商品
	SubCartItem(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*CartItemResponse, error)
	// 获取购物车结算信息
	GetCartSettle_(ctx context.Context, in *ShoppingCartId, opts ...grpc.CallOption) (*SettleMeta_, error)
	// 勾选商品结算
	CheckSign_(ctx context.Context, in *CheckSignRequest, opts ...grpc.CallOption) (*Result, error)
	// 更新购物车结算
	PrepareSettlePersist_(ctx context.Context, in *SettlePersistRequest, opts ...grpc.CallOption) (*Result, error)
}

type cartServiceClient struct {
	cc *grpc.ClientConn
}

func NewCartServiceClient(cc *grpc.ClientConn) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) WholesaleCartV1(ctx context.Context, in *WsCartRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/CartService/WholesaleCartV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetShoppingCart(ctx context.Context, in *ShoppingCartId, opts ...grpc.CallOption) (*SShoppingCart, error) {
	out := new(SShoppingCart)
	err := c.cc.Invoke(ctx, "/CartService/GetShoppingCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) PutInCart(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*CartItemResponse, error) {
	out := new(CartItemResponse)
	err := c.cc.Invoke(ctx, "/CartService/PutInCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) SubCartItem(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*CartItemResponse, error) {
	out := new(CartItemResponse)
	err := c.cc.Invoke(ctx, "/CartService/SubCartItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetCartSettle_(ctx context.Context, in *ShoppingCartId, opts ...grpc.CallOption) (*SettleMeta_, error) {
	out := new(SettleMeta_)
	err := c.cc.Invoke(ctx, "/CartService/GetCartSettle_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) CheckSign_(ctx context.Context, in *CheckSignRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/CartService/CheckSign_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) PrepareSettlePersist_(ctx context.Context, in *SettlePersistRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/CartService/PrepareSettlePersist_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
type CartServiceServer interface {
	// 批发购物车接口
	WholesaleCartV1(context.Context, *WsCartRequest) (*Result, error)
	// 获取购物车,当购物车编号不存在时,将返回一个新的购物车
	GetShoppingCart(context.Context, *ShoppingCartId) (*SShoppingCart, error)
	// 放入购物车
	PutInCart(context.Context, *CartItemRequest) (*CartItemResponse, error)
	// 从购物车里删除指定数量商品
	SubCartItem(context.Context, *CartItemRequest) (*CartItemResponse, error)
	// 获取购物车结算信息
	GetCartSettle_(context.Context, *ShoppingCartId) (*SettleMeta_, error)
	// 勾选商品结算
	CheckSign_(context.Context, *CheckSignRequest) (*Result, error)
	// 更新购物车结算
	PrepareSettlePersist_(context.Context, *SettlePersistRequest) (*Result, error)
}

func RegisterCartServiceServer(s *grpc.Server, srv CartServiceServer) {
	s.RegisterService(&_CartService_serviceDesc, srv)
}

func _CartService_WholesaleCartV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WsCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).WholesaleCartV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/WholesaleCartV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).WholesaleCartV1(ctx, req.(*WsCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetShoppingCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShoppingCartId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetShoppingCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/GetShoppingCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetShoppingCart(ctx, req.(*ShoppingCartId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_PutInCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).PutInCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/PutInCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).PutInCart(ctx, req.(*CartItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_SubCartItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).SubCartItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/SubCartItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).SubCartItem(ctx, req.(*CartItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetCartSettle__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShoppingCartId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetCartSettle_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/GetCartSettle_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetCartSettle_(ctx, req.(*ShoppingCartId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_CheckSign__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).CheckSign_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/CheckSign_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).CheckSign_(ctx, req.(*CheckSignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_PrepareSettlePersist__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SettlePersistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).PrepareSettlePersist_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/PrepareSettlePersist_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).PrepareSettlePersist_(ctx, req.(*SettlePersistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CartService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WholesaleCartV1",
			Handler:    _CartService_WholesaleCartV1_Handler,
		},
		{
			MethodName: "GetShoppingCart",
			Handler:    _CartService_GetShoppingCart_Handler,
		},
		{
			MethodName: "PutInCart",
			Handler:    _CartService_PutInCart_Handler,
		},
		{
			MethodName: "SubCartItem",
			Handler:    _CartService_SubCartItem_Handler,
		},
		{
			MethodName: "GetCartSettle_",
			Handler:    _CartService_GetCartSettle__Handler,
		},
		{
			MethodName: "CheckSign_",
			Handler:    _CartService_CheckSign__Handler,
		},
		{
			MethodName: "PrepareSettlePersist_",
			Handler:    _CartService_PrepareSettlePersist__Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cart_service.proto",
}

func init() { proto.RegisterFile("cart_service.proto", fileDescriptor_cart_service_e4a312d7fef09cb2) }

var fileDescriptor_cart_service_e4a312d7fef09cb2 = []byte{
	// 637 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0x93, 0xa6, 0x4d, 0xc6, 0x25, 0x69, 0xb7, 0x3f, 0x8a, 0x2c, 0x44, 0x2b, 0xab, 0x20,
	0x84, 0x60, 0xab, 0x06, 0x2e, 0x88, 0x13, 0x0d, 0x52, 0xe5, 0x03, 0x10, 0xad, 0x81, 0x4a, 0x5c,
	0x22, 0x27, 0x1e, 0x1c, 0xab, 0x8e, 0xd7, 0xf5, 0xae, 0x2b, 0x55, 0xe2, 0x11, 0x78, 0x0a, 0x0e,
	0xbc, 0x21, 0x77, 0xb4, 0xbb, 0x8e, 0xeb, 0xa4, 0x88, 0x72, 0xb2, 0xe7, 0x9b, 0xd9, 0x9d, 0x6f,
	0x66, 0xbe, 0x59, 0x20, 0xd3, 0x20, 0x97, 0x63, 0x81, 0xf9, 0x75, 0x3c, 0x45, 0x9a, 0xe5, 0x5c,
	0x72, 0x67, 0x2b, 0x4a, 0xf8, 0x24, 0x48, 0x4a, 0xeb, 0x60, 0x8e, 0x42, 0x04, 0x11, 0x9e, 0xe8,
	0xc8, 0x50, 0x72, 0x83, 0xbb, 0xbf, 0x2c, 0xd8, 0xf3, 0x51, 0xca, 0x04, 0x47, 0x98, 0x8b, 0x58,
	0x48, 0x86, 0x57, 0x05, 0x0a, 0x49, 0xfa, 0xb0, 0x79, 0x56, 0xdc, 0x60, 0xee, 0x85, 0x7d, 0xeb,
	0xc8, 0x7a, 0xda, 0x64, 0x0b, 0x93, 0x1c, 0xc0, 0x86, 0x3f, 0xe3, 0x99, 0x17, 0xf6, 0x1b, 0xda,
	0x51, 0x5a, 0xe4, 0x11, 0xc0, 0x28, 0xb8, 0x99, 0x63, 0x2a, 0x3f, 0x66, 0xb2, 0xdf, 0xd4, 0xbe,
	0x1a, 0xa2, 0xfc, 0xef, 0x30, 0x89, 0xaf, 0x31, 0x57, 0xfe, 0x75, 0xe3, 0xbf, 0x45, 0xc8, 0x43,
	0xe8, 0xbc, 0x0d, 0xc3, 0x1c, 0x85, 0xf0, 0xc2, 0x7e, 0x4b, 0xbb, 0x6f, 0x01, 0xf7, 0x3b, 0xf4,
	0x86, 0x41, 0x2e, 0x3d, 0x89, 0xf3, 0x05, 0xc5, 0x43, 0x68, 0x94, 0xec, 0xec, 0x41, 0x8f, 0x2a,
	0x16, 0x59, 0x9c, 0x46, 0x3a, 0x2a, 0x64, 0x0d, 0xc3, 0x54, 0xc5, 0x97, 0x4c, 0x09, 0x2b, 0x2d,
	0xb2, 0x07, 0x2d, 0xff, 0xb2, 0xf0, 0x42, 0x4d, 0x92, 0x30, 0x63, 0x10, 0x07, 0xda, 0x57, 0x45,
	0x90, 0xca, 0x58, 0xde, 0x68, 0x76, 0x2d, 0x56, 0xd9, 0x2e, 0x87, 0xed, 0xe1, 0x0c, 0xa7, 0x97,
	0x7e, 0x1c, 0xa5, 0xf7, 0x77, 0xc8, 0x81, 0xb6, 0x62, 0x31, 0xe4, 0x21, 0xea, 0xcc, 0x1d, 0x56,
	0xd9, 0xe4, 0x31, 0xb4, 0x14, 0x0b, 0xd1, 0x6f, 0x1e, 0x35, 0x0d, 0x6f, 0x7d, 0x71, 0x55, 0x9b,
	0xf1, 0xba, 0xdf, 0xa0, 0xbb, 0x5c, 0x90, 0x2a, 0xe6, 0xb3, 0xa8, 0xb2, 0x11, 0x56, 0x5a, 0xff,
	0x4c, 0x76, 0x04, 0xb6, 0x27, 0x2e, 0x66, 0x3c, 0x41, 0x11, 0x24, 0xa8, 0xcb, 0x6d, 0xb3, 0x3a,
	0xe4, 0xfe, 0xb4, 0xc0, 0x36, 0xf3, 0x7f, 0x8f, 0x32, 0x18, 0xaf, 0x0c, 0xd1, 0xba, 0x67, 0x88,
	0x8d, 0x3b, 0x43, 0x3c, 0x86, 0x75, 0xc5, 0x5b, 0xa7, 0xb2, 0x07, 0xdb, 0xd4, 0xdc, 0xad, 0x20,
	0x7d, 0x3f, 0xd3, 0x5e, 0xf2, 0x02, 0x36, 0xcb, 0x33, 0xba, 0xd3, 0xf6, 0x60, 0xb7, 0x0c, 0x2c,
	0x51, 0x13, 0xbb, 0x88, 0x71, 0xa7, 0xd0, 0x5b, 0xb9, 0xa7, 0x26, 0x42, 0x6b, 0x49, 0x84, 0x0e,
	0xb4, 0xd5, 0xdf, 0x87, 0x60, 0x5e, 0x75, 0x63, 0x61, 0x2b, 0x81, 0x7d, 0xc2, 0x04, 0xb3, 0x19,
	0x4f, 0x4d, 0x2f, 0x3a, 0xec, 0x16, 0x70, 0x7f, 0x58, 0x40, 0xee, 0x92, 0x20, 0xdd, 0x4a, 0x64,
	0x4d, 0xad, 0xa9, 0x63, 0x78, 0x30, 0xe4, 0xa9, 0x88, 0xa3, 0x14, 0xb1, 0x96, 0x65, 0x19, 0x24,
	0x4f, 0xa0, 0x5b, 0x01, 0xa3, 0x5a, 0xbe, 0x15, 0x54, 0x69, 0xa8, 0x94, 0xb8, 0x6e, 0x44, 0x87,
	0x2d, 0xcc, 0xc1, 0xef, 0x06, 0xd8, 0x6a, 0x8e, 0xbe, 0x59, 0x6a, 0xf2, 0x1c, 0x7a, 0xd5, 0xd4,
	0x14, 0xfe, 0xe5, 0x94, 0x74, 0xe9, 0x85, 0x50, 0xbf, 0xa5, 0x20, 0x9d, 0x4d, 0xca, 0x50, 0x14,
	0x89, 0x74, 0xd7, 0xc8, 0x2b, 0xe8, 0x9d, 0xa3, 0xac, 0x2b, 0x88, 0xac, 0x6e, 0x88, 0xd3, 0xa5,
	0x7e, 0x1d, 0x71, 0xd7, 0xc8, 0x00, 0x3a, 0xa3, 0x42, 0x7a, 0xa9, 0x8e, 0xdf, 0xa6, 0x2b, 0xfb,
	0xe6, 0xec, 0xd4, 0x10, 0x91, 0xf1, 0x54, 0xa0, 0xce, 0x64, 0xfb, 0xc5, 0x64, 0xe1, 0xf8, 0xdf,
	0x53, 0xa7, 0xd0, 0x3d, 0x47, 0x69, 0xea, 0x53, 0x2d, 0x1f, 0xdf, 0xa5, 0xb7, 0x45, 0x6b, 0xba,
	0x74, 0xd7, 0xc8, 0x33, 0x80, 0x6a, 0x05, 0xc7, 0x64, 0x87, 0xae, 0xee, 0x63, 0xbd, 0xfc, 0xd7,
	0xb0, 0x3f, 0xca, 0x31, 0x0b, 0x72, 0x5c, 0x7a, 0xdb, 0xc6, 0x64, 0x9f, 0xfe, 0xed, 0xb1, 0xab,
	0x1d, 0x3d, 0x3b, 0x84, 0xdd, 0x29, 0x9f, 0xd3, 0x28, 0x96, 0xb3, 0x62, 0x42, 0x23, 0x3e, 0xe0,
	0x34, 0xcf, 0xa6, 0x5f, 0xdb, 0xf4, 0xe4, 0x8d, 0x7e, 0x31, 0x27, 0x1b, 0xfa, 0xf3, 0xf2, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x29, 0x03, 0x5e, 0x74, 0x05, 0x00, 0x00,
}
