// Code generated by protoc-gen-go. DO NOT EDIT.
// source: merchant_service.proto

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

// 商家
type SMerchant struct {
	// * 编号
	Id int32 `protobuf:"zigzag32,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// * 会员编号
	MemberId int64 `protobuf:"zigzag64,2,opt,name=MemberId,proto3" json:"MemberId,omitempty"`
	// * 登录用户
	LoginUser string `protobuf:"bytes,3,opt,name=LoginUser,proto3" json:"LoginUser,omitempty"`
	// * 登录密码
	LoginPwd string `protobuf:"bytes,4,opt,name=LoginPwd,proto3" json:"LoginPwd,omitempty"`
	// * 名称
	Name string `protobuf:"bytes,5,opt,name=Name,proto3" json:"Name,omitempty"`
	// * 公司名称
	CompanyName string `protobuf:"bytes,6,opt,name=CompanyName,proto3" json:"CompanyName,omitempty"`
	// * 是否字营
	SelfSales int32 `protobuf:"zigzag32,7,opt,name=SelfSales,proto3" json:"SelfSales,omitempty"`
	// * 商户等级
	Level int32 `protobuf:"zigzag32,8,opt,name=Level,proto3" json:"Level,omitempty"`
	// * 标志
	Logo string `protobuf:"bytes,9,opt,name=Logo,proto3" json:"Logo,omitempty"`
	// * 省
	Province int32 `protobuf:"zigzag32,10,opt,name=Province,proto3" json:"Province,omitempty"`
	// * 市
	City int32 `protobuf:"zigzag32,11,opt,name=City,proto3" json:"City,omitempty"`
	// * 区
	District int32 `protobuf:"zigzag32,12,opt,name=District,proto3" json:"District,omitempty"`
	// * 标志
	Flag int32 `protobuf:"zigzag32,13,opt,name=Flag,proto3" json:"Flag,omitempty"`
	// * 是否启用
	Enabled int32 `protobuf:"zigzag32,14,opt,name=Enabled,proto3" json:"Enabled,omitempty"`
	// * 最后登录时间
	LastLoginTime        int32    `protobuf:"zigzag32,15,opt,name=LastLoginTime,proto3" json:"LastLoginTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SMerchant) Reset()         { *m = SMerchant{} }
func (m *SMerchant) String() string { return proto.CompactTextString(m) }
func (*SMerchant) ProtoMessage()    {}
func (*SMerchant) Descriptor() ([]byte, []int) {
	return fileDescriptor_merchant_service_926c8c486a879d74, []int{0}
}
func (m *SMerchant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SMerchant.Unmarshal(m, b)
}
func (m *SMerchant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SMerchant.Marshal(b, m, deterministic)
}
func (dst *SMerchant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SMerchant.Merge(dst, src)
}
func (m *SMerchant) XXX_Size() int {
	return xxx_messageInfo_SMerchant.Size(m)
}
func (m *SMerchant) XXX_DiscardUnknown() {
	xxx_messageInfo_SMerchant.DiscardUnknown(m)
}

var xxx_messageInfo_SMerchant proto.InternalMessageInfo

func (m *SMerchant) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SMerchant) GetMemberId() int64 {
	if m != nil {
		return m.MemberId
	}
	return 0
}

func (m *SMerchant) GetLoginUser() string {
	if m != nil {
		return m.LoginUser
	}
	return ""
}

func (m *SMerchant) GetLoginPwd() string {
	if m != nil {
		return m.LoginPwd
	}
	return ""
}

func (m *SMerchant) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SMerchant) GetCompanyName() string {
	if m != nil {
		return m.CompanyName
	}
	return ""
}

func (m *SMerchant) GetSelfSales() int32 {
	if m != nil {
		return m.SelfSales
	}
	return 0
}

func (m *SMerchant) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *SMerchant) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *SMerchant) GetProvince() int32 {
	if m != nil {
		return m.Province
	}
	return 0
}

func (m *SMerchant) GetCity() int32 {
	if m != nil {
		return m.City
	}
	return 0
}

func (m *SMerchant) GetDistrict() int32 {
	if m != nil {
		return m.District
	}
	return 0
}

func (m *SMerchant) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *SMerchant) GetEnabled() int32 {
	if m != nil {
		return m.Enabled
	}
	return 0
}

func (m *SMerchant) GetLastLoginTime() int32 {
	if m != nil {
		return m.LastLoginTime
	}
	return 0
}

// 商家
type SMerchantPack struct {
	// * 登录用户
	LoginUser string `protobuf:"bytes,1,opt,name=LoginUser,proto3" json:"LoginUser,omitempty"`
	// * 登录密码
	LoginPwd string `protobuf:"bytes,2,opt,name=LoginPwd,proto3" json:"LoginPwd,omitempty"`
	// * 名称
	Name string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	// * 是否字营
	SelfSales int32 `protobuf:"zigzag32,4,opt,name=SelfSales,proto3" json:"SelfSales,omitempty"`
	// * 店铺名称
	ShopName string `protobuf:"bytes,5,opt,name=ShopName,proto3" json:"ShopName,omitempty"`
	// * 标志
	ShopLogo string `protobuf:"bytes,6,opt,name=ShopLogo,proto3" json:"ShopLogo,omitempty"`
	// * 电话
	Tel string `protobuf:"bytes,7,opt,name=Tel,proto3" json:"Tel,omitempty"`
	// * 地址
	Addr                 string   `protobuf:"bytes,8,opt,name=Addr,proto3" json:"Addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SMerchantPack) Reset()         { *m = SMerchantPack{} }
func (m *SMerchantPack) String() string { return proto.CompactTextString(m) }
func (*SMerchantPack) ProtoMessage()    {}
func (*SMerchantPack) Descriptor() ([]byte, []int) {
	return fileDescriptor_merchant_service_926c8c486a879d74, []int{1}
}
func (m *SMerchantPack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SMerchantPack.Unmarshal(m, b)
}
func (m *SMerchantPack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SMerchantPack.Marshal(b, m, deterministic)
}
func (dst *SMerchantPack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SMerchantPack.Merge(dst, src)
}
func (m *SMerchantPack) XXX_Size() int {
	return xxx_messageInfo_SMerchantPack.Size(m)
}
func (m *SMerchantPack) XXX_DiscardUnknown() {
	xxx_messageInfo_SMerchantPack.DiscardUnknown(m)
}

var xxx_messageInfo_SMerchantPack proto.InternalMessageInfo

func (m *SMerchantPack) GetLoginUser() string {
	if m != nil {
		return m.LoginUser
	}
	return ""
}

func (m *SMerchantPack) GetLoginPwd() string {
	if m != nil {
		return m.LoginPwd
	}
	return ""
}

func (m *SMerchantPack) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SMerchantPack) GetSelfSales() int32 {
	if m != nil {
		return m.SelfSales
	}
	return 0
}

func (m *SMerchantPack) GetShopName() string {
	if m != nil {
		return m.ShopName
	}
	return ""
}

func (m *SMerchantPack) GetShopLogo() string {
	if m != nil {
		return m.ShopLogo
	}
	return ""
}

func (m *SMerchantPack) GetTel() string {
	if m != nil {
		return m.Tel
	}
	return ""
}

func (m *SMerchantPack) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type MerchantCreateRequest struct {
	Mch                  *SMerchantPack `protobuf:"bytes,1,opt,name=mch,proto3" json:"mch,omitempty"`
	RelMemberId          int64          `protobuf:"zigzag64,2,opt,name=relMemberId,proto3" json:"relMemberId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MerchantCreateRequest) Reset()         { *m = MerchantCreateRequest{} }
func (m *MerchantCreateRequest) String() string { return proto.CompactTextString(m) }
func (*MerchantCreateRequest) ProtoMessage()    {}
func (*MerchantCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_merchant_service_926c8c486a879d74, []int{2}
}
func (m *MerchantCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantCreateRequest.Unmarshal(m, b)
}
func (m *MerchantCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantCreateRequest.Marshal(b, m, deterministic)
}
func (dst *MerchantCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantCreateRequest.Merge(dst, src)
}
func (m *MerchantCreateRequest) XXX_Size() int {
	return xxx_messageInfo_MerchantCreateRequest.Size(m)
}
func (m *MerchantCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantCreateRequest proto.InternalMessageInfo

func (m *MerchantCreateRequest) GetMch() *SMerchantPack {
	if m != nil {
		return m.Mch
	}
	return nil
}

func (m *MerchantCreateRequest) GetRelMemberId() int64 {
	if m != nil {
		return m.RelMemberId
	}
	return 0
}

type MchUserPwd struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Pwd                  string   `protobuf:"bytes,2,opt,name=pwd,proto3" json:"pwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MchUserPwd) Reset()         { *m = MchUserPwd{} }
func (m *MchUserPwd) String() string { return proto.CompactTextString(m) }
func (*MchUserPwd) ProtoMessage()    {}
func (*MchUserPwd) Descriptor() ([]byte, []int) {
	return fileDescriptor_merchant_service_926c8c486a879d74, []int{3}
}
func (m *MchUserPwd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MchUserPwd.Unmarshal(m, b)
}
func (m *MchUserPwd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MchUserPwd.Marshal(b, m, deterministic)
}
func (dst *MchUserPwd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MchUserPwd.Merge(dst, src)
}
func (m *MchUserPwd) XXX_Size() int {
	return xxx_messageInfo_MchUserPwd.Size(m)
}
func (m *MchUserPwd) XXX_DiscardUnknown() {
	xxx_messageInfo_MchUserPwd.DiscardUnknown(m)
}

var xxx_messageInfo_MchUserPwd proto.InternalMessageInfo

func (m *MchUserPwd) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *MchUserPwd) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type SyncWSItemsResponse struct {
	Value                map[string]int32 `protobuf:"bytes,1,rep,name=Value,proto3" json:"Value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"zigzag32,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SyncWSItemsResponse) Reset()         { *m = SyncWSItemsResponse{} }
func (m *SyncWSItemsResponse) String() string { return proto.CompactTextString(m) }
func (*SyncWSItemsResponse) ProtoMessage()    {}
func (*SyncWSItemsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_merchant_service_926c8c486a879d74, []int{4}
}
func (m *SyncWSItemsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncWSItemsResponse.Unmarshal(m, b)
}
func (m *SyncWSItemsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncWSItemsResponse.Marshal(b, m, deterministic)
}
func (dst *SyncWSItemsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncWSItemsResponse.Merge(dst, src)
}
func (m *SyncWSItemsResponse) XXX_Size() int {
	return xxx_messageInfo_SyncWSItemsResponse.Size(m)
}
func (m *SyncWSItemsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncWSItemsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SyncWSItemsResponse proto.InternalMessageInfo

func (m *SyncWSItemsResponse) GetValue() map[string]int32 {
	if m != nil {
		return m.Value
	}
	return nil
}

type STradeConfListResponse struct {
	Value                []*STradeConf `protobuf:"bytes,1,rep,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *STradeConfListResponse) Reset()         { *m = STradeConfListResponse{} }
func (m *STradeConfListResponse) String() string { return proto.CompactTextString(m) }
func (*STradeConfListResponse) ProtoMessage()    {}
func (*STradeConfListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_merchant_service_926c8c486a879d74, []int{5}
}
func (m *STradeConfListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_STradeConfListResponse.Unmarshal(m, b)
}
func (m *STradeConfListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_STradeConfListResponse.Marshal(b, m, deterministic)
}
func (dst *STradeConfListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_STradeConfListResponse.Merge(dst, src)
}
func (m *STradeConfListResponse) XXX_Size() int {
	return xxx_messageInfo_STradeConfListResponse.Size(m)
}
func (m *STradeConfListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_STradeConfListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_STradeConfListResponse proto.InternalMessageInfo

func (m *STradeConfListResponse) GetValue() []*STradeConf {
	if m != nil {
		return m.Value
	}
	return nil
}

type TradeConfRequest struct {
	MchId                int32    `protobuf:"zigzag32,1,opt,name=mchId,proto3" json:"mchId,omitempty"`
	TradeType            int32    `protobuf:"zigzag32,2,opt,name=tradeType,proto3" json:"tradeType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeConfRequest) Reset()         { *m = TradeConfRequest{} }
func (m *TradeConfRequest) String() string { return proto.CompactTextString(m) }
func (*TradeConfRequest) ProtoMessage()    {}
func (*TradeConfRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_merchant_service_926c8c486a879d74, []int{6}
}
func (m *TradeConfRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeConfRequest.Unmarshal(m, b)
}
func (m *TradeConfRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeConfRequest.Marshal(b, m, deterministic)
}
func (dst *TradeConfRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeConfRequest.Merge(dst, src)
}
func (m *TradeConfRequest) XXX_Size() int {
	return xxx_messageInfo_TradeConfRequest.Size(m)
}
func (m *TradeConfRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeConfRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TradeConfRequest proto.InternalMessageInfo

func (m *TradeConfRequest) GetMchId() int32 {
	if m != nil {
		return m.MchId
	}
	return 0
}

func (m *TradeConfRequest) GetTradeType() int32 {
	if m != nil {
		return m.TradeType
	}
	return 0
}

type TradeConfSaveRequest struct {
	MchId                int32         `protobuf:"zigzag32,1,opt,name=mchId,proto3" json:"mchId,omitempty"`
	Arr                  []*STradeConf `protobuf:"bytes,2,rep,name=arr,proto3" json:"arr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TradeConfSaveRequest) Reset()         { *m = TradeConfSaveRequest{} }
func (m *TradeConfSaveRequest) String() string { return proto.CompactTextString(m) }
func (*TradeConfSaveRequest) ProtoMessage()    {}
func (*TradeConfSaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_merchant_service_926c8c486a879d74, []int{7}
}
func (m *TradeConfSaveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeConfSaveRequest.Unmarshal(m, b)
}
func (m *TradeConfSaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeConfSaveRequest.Marshal(b, m, deterministic)
}
func (dst *TradeConfSaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeConfSaveRequest.Merge(dst, src)
}
func (m *TradeConfSaveRequest) XXX_Size() int {
	return xxx_messageInfo_TradeConfSaveRequest.Size(m)
}
func (m *TradeConfSaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeConfSaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TradeConfSaveRequest proto.InternalMessageInfo

func (m *TradeConfSaveRequest) GetMchId() int32 {
	if m != nil {
		return m.MchId
	}
	return 0
}

func (m *TradeConfSaveRequest) GetArr() []*STradeConf {
	if m != nil {
		return m.Arr
	}
	return nil
}

// 商户交易设置
type STradeConf struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *STradeConf) Reset()         { *m = STradeConf{} }
func (m *STradeConf) String() string { return proto.CompactTextString(m) }
func (*STradeConf) ProtoMessage()    {}
func (*STradeConf) Descriptor() ([]byte, []int) {
	return fileDescriptor_merchant_service_926c8c486a879d74, []int{8}
}
func (m *STradeConf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_STradeConf.Unmarshal(m, b)
}
func (m *STradeConf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_STradeConf.Marshal(b, m, deterministic)
}
func (dst *STradeConf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_STradeConf.Merge(dst, src)
}
func (m *STradeConf) XXX_Size() int {
	return xxx_messageInfo_STradeConf.Size(m)
}
func (m *STradeConf) XXX_DiscardUnknown() {
	xxx_messageInfo_STradeConf.DiscardUnknown(m)
}

var xxx_messageInfo_STradeConf proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SMerchant)(nil), "SMerchant")
	proto.RegisterType((*SMerchantPack)(nil), "SMerchantPack")
	proto.RegisterType((*MerchantCreateRequest)(nil), "MerchantCreateRequest")
	proto.RegisterType((*MchUserPwd)(nil), "MchUserPwd")
	proto.RegisterType((*SyncWSItemsResponse)(nil), "SyncWSItemsResponse")
	proto.RegisterMapType((map[string]int32)(nil), "SyncWSItemsResponse.ValueEntry")
	proto.RegisterType((*STradeConfListResponse)(nil), "STradeConfListResponse")
	proto.RegisterType((*TradeConfRequest)(nil), "TradeConfRequest")
	proto.RegisterType((*TradeConfSaveRequest)(nil), "TradeConfSaveRequest")
	proto.RegisterType((*STradeConf)(nil), "STradeConf")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MerchantServiceClient is the client API for MerchantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MerchantServiceClient interface {
	// 获取商家的信息,mchId
	GetMerchant(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SMerchant, error)
	// 注册商户并开店
	CreateMerchant(ctx context.Context, in *MerchantCreateRequest, opts ...grpc.CallOption) (*Result, error)
	// 验证用户密码,并返回编号。可传入商户或会员的账号密码
	CheckLogin(ctx context.Context, in *MchUserPwd, opts ...grpc.CallOption) (*Result, error)
	// 验证商户状态,mchId
	Stat(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error)
	// 同步批发商品,mchId
	SyncWholesaleItem(ctx context.Context, in *Int32, opts ...grpc.CallOption) (*SyncWSItemsResponse, error)
	// 获取所有的交易设置,mchId
	GetAllTradeConf(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*STradeConfListResponse, error)
	// 根据交易类型获取交易设置
	GetTradeConf(ctx context.Context, in *TradeConfRequest, opts ...grpc.CallOption) (*STradeConf, error)
	// 保存交易设置
	SaveTradeConf(ctx context.Context, in *TradeConfSaveRequest, opts ...grpc.CallOption) (*Result, error)
}

type merchantServiceClient struct {
	cc *grpc.ClientConn
}

func NewMerchantServiceClient(cc *grpc.ClientConn) MerchantServiceClient {
	return &merchantServiceClient{cc}
}

func (c *merchantServiceClient) GetMerchant(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SMerchant, error) {
	out := new(SMerchant)
	err := c.cc.Invoke(ctx, "/MerchantService/GetMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) CreateMerchant(ctx context.Context, in *MerchantCreateRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/MerchantService/CreateMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) CheckLogin(ctx context.Context, in *MchUserPwd, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/MerchantService/CheckLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) Stat(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/MerchantService/Stat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) SyncWholesaleItem(ctx context.Context, in *Int32, opts ...grpc.CallOption) (*SyncWSItemsResponse, error) {
	out := new(SyncWSItemsResponse)
	err := c.cc.Invoke(ctx, "/MerchantService/SyncWholesaleItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) GetAllTradeConf(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*STradeConfListResponse, error) {
	out := new(STradeConfListResponse)
	err := c.cc.Invoke(ctx, "/MerchantService/GetAllTradeConf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) GetTradeConf(ctx context.Context, in *TradeConfRequest, opts ...grpc.CallOption) (*STradeConf, error) {
	out := new(STradeConf)
	err := c.cc.Invoke(ctx, "/MerchantService/GetTradeConf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) SaveTradeConf(ctx context.Context, in *TradeConfSaveRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/MerchantService/SaveTradeConf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MerchantServiceServer is the server API for MerchantService service.
type MerchantServiceServer interface {
	// 获取商家的信息,mchId
	GetMerchant(context.Context, *Int64) (*SMerchant, error)
	// 注册商户并开店
	CreateMerchant(context.Context, *MerchantCreateRequest) (*Result, error)
	// 验证用户密码,并返回编号。可传入商户或会员的账号密码
	CheckLogin(context.Context, *MchUserPwd) (*Result, error)
	// 验证商户状态,mchId
	Stat(context.Context, *Int64) (*Result, error)
	// 同步批发商品,mchId
	SyncWholesaleItem(context.Context, *Int32) (*SyncWSItemsResponse, error)
	// 获取所有的交易设置,mchId
	GetAllTradeConf(context.Context, *Int64) (*STradeConfListResponse, error)
	// 根据交易类型获取交易设置
	GetTradeConf(context.Context, *TradeConfRequest) (*STradeConf, error)
	// 保存交易设置
	SaveTradeConf(context.Context, *TradeConfSaveRequest) (*Result, error)
}

func RegisterMerchantServiceServer(s *grpc.Server, srv MerchantServiceServer) {
	s.RegisterService(&_MerchantService_serviceDesc, srv)
}

func _MerchantService_GetMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).GetMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MerchantService/GetMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).GetMerchant(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_CreateMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerchantCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).CreateMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MerchantService/CreateMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).CreateMerchant(ctx, req.(*MerchantCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_CheckLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MchUserPwd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).CheckLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MerchantService/CheckLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).CheckLogin(ctx, req.(*MchUserPwd))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MerchantService/Stat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).Stat(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_SyncWholesaleItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int32)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).SyncWholesaleItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MerchantService/SyncWholesaleItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).SyncWholesaleItem(ctx, req.(*Int32))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_GetAllTradeConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).GetAllTradeConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MerchantService/GetAllTradeConf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).GetAllTradeConf(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_GetTradeConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeConfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).GetTradeConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MerchantService/GetTradeConf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).GetTradeConf(ctx, req.(*TradeConfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_SaveTradeConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeConfSaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).SaveTradeConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MerchantService/SaveTradeConf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).SaveTradeConf(ctx, req.(*TradeConfSaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MerchantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MerchantService",
	HandlerType: (*MerchantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMerchant",
			Handler:    _MerchantService_GetMerchant_Handler,
		},
		{
			MethodName: "CreateMerchant",
			Handler:    _MerchantService_CreateMerchant_Handler,
		},
		{
			MethodName: "CheckLogin",
			Handler:    _MerchantService_CheckLogin_Handler,
		},
		{
			MethodName: "Stat",
			Handler:    _MerchantService_Stat_Handler,
		},
		{
			MethodName: "SyncWholesaleItem",
			Handler:    _MerchantService_SyncWholesaleItem_Handler,
		},
		{
			MethodName: "GetAllTradeConf",
			Handler:    _MerchantService_GetAllTradeConf_Handler,
		},
		{
			MethodName: "GetTradeConf",
			Handler:    _MerchantService_GetTradeConf_Handler,
		},
		{
			MethodName: "SaveTradeConf",
			Handler:    _MerchantService_SaveTradeConf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "merchant_service.proto",
}

func init() {
	proto.RegisterFile("merchant_service.proto", fileDescriptor_merchant_service_926c8c486a879d74)
}

var fileDescriptor_merchant_service_926c8c486a879d74 = []byte{
	// 727 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x5d, 0x6f, 0xd3, 0x4a,
	0x10, 0x8d, 0x9d, 0xaf, 0x66, 0xdc, 0xa4, 0xed, 0xde, 0xb4, 0x77, 0xaf, 0x75, 0x11, 0xc1, 0xf4,
	0x21, 0x4f, 0x16, 0x24, 0x05, 0x55, 0xf4, 0xa9, 0x84, 0xb6, 0x8a, 0x48, 0x51, 0x65, 0x07, 0x90,
	0xe0, 0x01, 0x6d, 0xed, 0x69, 0x63, 0xd5, 0x1f, 0xc1, 0xde, 0x04, 0xe5, 0x0f, 0xf0, 0xca, 0xef,
	0xe3, 0xbf, 0xf0, 0x80, 0x76, 0x9d, 0xd8, 0x4e, 0x09, 0x7d, 0xca, 0xcc, 0x99, 0xb3, 0x3b, 0xb3,
	0xe7, 0x4c, 0x0c, 0x07, 0x01, 0xc6, 0xce, 0x84, 0x85, 0xfc, 0x4b, 0x82, 0xf1, 0xdc, 0x73, 0xd0,
	0x9c, 0xc6, 0x11, 0x8f, 0x74, 0x8d, 0xf3, 0xc5, 0x74, 0x99, 0x18, 0x3f, 0xca, 0xd0, 0xb0, 0x2f,
	0x97, 0x44, 0xd2, 0x02, 0x75, 0xe8, 0x52, 0xa5, 0xa3, 0x74, 0xf7, 0x2c, 0x75, 0xe8, 0x12, 0x1d,
	0xb6, 0x2e, 0x31, 0xb8, 0xc6, 0x78, 0xe8, 0x52, 0xb5, 0xa3, 0x74, 0x89, 0x95, 0xe5, 0xe4, 0x7f,
	0x68, 0x8c, 0xa2, 0x5b, 0x2f, 0x7c, 0x9f, 0x60, 0x4c, 0xcb, 0x1d, 0xa5, 0xdb, 0xb0, 0x72, 0x40,
	0x9c, 0x94, 0xc9, 0xd5, 0x37, 0x97, 0x56, 0x64, 0x31, 0xcb, 0x09, 0x81, 0xca, 0x3b, 0x16, 0x20,
	0xad, 0x4a, 0x5c, 0xc6, 0xa4, 0x03, 0xda, 0x20, 0x0a, 0xa6, 0x2c, 0x5c, 0xc8, 0x52, 0x4d, 0x96,
	0x8a, 0x90, 0xe8, 0x67, 0xa3, 0x7f, 0x63, 0x33, 0x1f, 0x13, 0x5a, 0x97, 0x23, 0xe6, 0x00, 0x69,
	0x43, 0x75, 0x84, 0x73, 0xf4, 0xe9, 0x96, 0xac, 0xa4, 0x89, 0xe8, 0x34, 0x8a, 0x6e, 0x23, 0xda,
	0x48, 0x3b, 0x89, 0x58, 0x4c, 0x76, 0x15, 0x47, 0x73, 0x2f, 0x74, 0x90, 0x82, 0x24, 0x67, 0xb9,
	0xe0, 0x0f, 0x3c, 0xbe, 0xa0, 0x9a, 0xc4, 0x65, 0x2c, 0xf8, 0x6f, 0xbc, 0x84, 0xc7, 0x9e, 0xc3,
	0xe9, 0x76, 0xca, 0x5f, 0xe5, 0x82, 0x7f, 0xee, 0xb3, 0x5b, 0xda, 0x4c, 0xf9, 0x22, 0x26, 0x14,
	0xea, 0x67, 0x21, 0xbb, 0xf6, 0xd1, 0xa5, 0x2d, 0x09, 0xaf, 0x52, 0x72, 0x08, 0xcd, 0x11, 0x4b,
	0xb8, 0xd4, 0x61, 0xec, 0x05, 0x48, 0x77, 0x64, 0x7d, 0x1d, 0x34, 0x7e, 0x2a, 0xd0, 0xcc, 0x1c,
	0xb9, 0x62, 0xce, 0xdd, 0xba, 0xd2, 0xca, 0x43, 0x4a, 0xab, 0x7f, 0x51, 0xba, 0x5c, 0x50, 0x7a,
	0x4d, 0xc7, 0xca, 0x7d, 0x1d, 0x75, 0xd8, 0xb2, 0x27, 0xd1, 0xb4, 0xe0, 0x4f, 0x96, 0xaf, 0x6a,
	0x52, 0xd1, 0x5a, 0x5e, 0x93, 0xaa, 0xee, 0x42, 0x79, 0x8c, 0xbe, 0xf4, 0xa5, 0x61, 0x89, 0x50,
	0xf4, 0x3e, 0x75, 0xdd, 0x58, 0x1a, 0xd2, 0xb0, 0x64, 0x6c, 0x7c, 0x86, 0xfd, 0xd5, 0xcb, 0x06,
	0x31, 0x32, 0x8e, 0x16, 0x7e, 0x9d, 0x61, 0xc2, 0x49, 0x07, 0xca, 0x81, 0x33, 0x91, 0x8f, 0xd3,
	0x7a, 0x2d, 0x73, 0xed, 0xfd, 0x96, 0x28, 0x89, 0x05, 0x89, 0xd1, 0xbf, 0xb7, 0x8d, 0x45, 0xc8,
	0xe8, 0x01, 0x5c, 0x3a, 0x13, 0xa1, 0xc9, 0xf2, 0xe9, 0xb3, 0x5c, 0x2f, 0x19, 0x8b, 0x21, 0xa7,
	0x99, 0x4a, 0x22, 0x34, 0xbe, 0x2b, 0xf0, 0x8f, 0xbd, 0x08, 0x9d, 0x8f, 0xf6, 0x90, 0x63, 0x90,
	0x58, 0x98, 0x4c, 0xa3, 0x30, 0x41, 0xf2, 0x02, 0xaa, 0x1f, 0x98, 0x3f, 0x43, 0xaa, 0x74, 0xca,
	0x5d, 0xad, 0xf7, 0xd8, 0xdc, 0x40, 0x32, 0x25, 0xe3, 0x2c, 0xe4, 0xf1, 0xc2, 0x4a, 0xd9, 0xfa,
	0x31, 0x40, 0x0e, 0x8a, 0x76, 0x77, 0xb8, 0x58, 0x4e, 0x20, 0x42, 0xb1, 0xa5, 0x73, 0x79, 0xad,
	0x9a, 0x6e, 0xa9, 0x4c, 0x5e, 0xa9, 0xc7, 0x8a, 0x71, 0x02, 0x07, 0xf6, 0x38, 0x66, 0x2e, 0x0e,
	0xa2, 0xf0, 0x66, 0xe4, 0x25, 0x3c, 0x1b, 0xe5, 0xc9, 0xfa, 0x28, 0x9a, 0x99, 0xf3, 0x96, 0x6d,
	0x8d, 0x73, 0xd8, 0xcd, 0xb1, 0xa5, 0xa2, 0x6d, 0xa8, 0x06, 0xce, 0x24, 0xfb, 0x37, 0xa7, 0x89,
	0x30, 0x9f, 0x0b, 0xe6, 0x78, 0x31, 0x5d, 0x0d, 0x91, 0x03, 0xc6, 0x5b, 0x68, 0x67, 0xf7, 0xd8,
	0x6c, 0x8e, 0x0f, 0xdf, 0xf5, 0x08, 0xca, 0x2c, 0x8e, 0xa9, 0xfa, 0xe7, 0x58, 0x02, 0x37, 0xb6,
	0x01, 0x72, 0xa8, 0xf7, 0x4b, 0x85, 0x9d, 0x95, 0xa9, 0x76, 0xfa, 0x39, 0x22, 0x4f, 0x41, 0xbb,
	0x40, 0x9e, 0x7d, 0x7c, 0x6a, 0xe6, 0x30, 0xe4, 0x2f, 0x8f, 0x74, 0xc8, 0xed, 0x37, 0x4a, 0xa4,
	0x0f, 0xad, 0x74, 0x55, 0x32, 0xde, 0x81, 0xb9, 0x71, 0x87, 0xf4, 0xba, 0x69, 0x61, 0x32, 0xf3,
	0xc5, 0xa1, 0x43, 0x80, 0xc1, 0x04, 0x9d, 0x3b, 0xf9, 0x47, 0x20, 0x9a, 0x99, 0xef, 0x45, 0x91,
	0xf5, 0x1f, 0x54, 0x6c, 0xce, 0xf2, 0xc6, 0x85, 0x52, 0x1f, 0xf6, 0xa4, 0xe3, 0x93, 0xc8, 0xc7,
	0x84, 0xf9, 0x28, 0x8c, 0x4f, 0x79, 0xfd, 0x9e, 0xde, 0xde, 0xb4, 0x0d, 0x46, 0x89, 0x1c, 0xc1,
	0xce, 0x05, 0xf2, 0x53, 0xdf, 0xcf, 0x9e, 0x9d, 0x5d, 0xfd, 0xaf, 0xb9, 0xd9, 0x5d, 0xa3, 0x44,
	0x9e, 0xc1, 0xf6, 0x05, 0xf2, 0xfc, 0xc8, 0x9e, 0x79, 0xdf, 0x4b, 0xbd, 0x28, 0xae, 0x51, 0x22,
	0xcf, 0xa1, 0x29, 0xdc, 0xc9, 0x8f, 0xec, 0x9b, 0x9b, 0x6c, 0x2b, 0xbc, 0xe7, 0x75, 0xe3, 0x53,
	0xdd, 0x3c, 0x91, 0x5f, 0xfc, 0xeb, 0x9a, 0xfc, 0xe9, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x64,
	0x95, 0x5b, 0xb8, 0x1f, 0x06, 0x00, 0x00,
}