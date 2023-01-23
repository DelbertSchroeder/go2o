// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.13.0
// source: payment_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentServiceClient interface {
	// 创建支付单并提交
	SubmitPaymentOrder(ctx context.Context, in *SPaymentOrder, opts ...grpc.CallOption) (*Result, error)
	// 根据支付单号获取支付单,orderNo
	GetPaymentOrder(ctx context.Context, in *String, opts ...grpc.CallOption) (*SPaymentOrder, error)
	// 根据交易号获取支付单编号,tradeNo
	GetPaymentOrderId(ctx context.Context, in *String, opts ...grpc.CallOption) (*Int32, error)
	// 根据编号获取支付单
	GetPaymentOrderById(ctx context.Context, in *Int32, opts ...grpc.CallOption) (*SPaymentOrder, error)
	// 调整支付单金额
	AdjustOrder(ctx context.Context, in *AdjustOrderRequest, opts ...grpc.CallOption) (*Result, error)
	// 余额抵扣
	DiscountByBalance(ctx context.Context, in *DiscountBalanceRequest, opts ...grpc.CallOption) (*Result, error)
	// 积分抵扣支付单
	DiscountByIntegral(ctx context.Context, in *DiscountIntegralRequest, opts ...grpc.CallOption) (*Result, error)
	// 钱包账户支付
	PaymentByWallet(ctx context.Context, in *WalletPaymentRequest, opts ...grpc.CallOption) (*Result, error)
	// 余额钱包混合支付，优先扣除余额。
	HybridPayment(ctx context.Context, in *HyperPaymentRequest, opts ...grpc.CallOption) (*Result, error)
	// 完成支付单支付，并传入支付方式及外部订单号
	FinishPayment(ctx context.Context, in *FinishPaymentRequest, opts ...grpc.CallOption) (*Result, error)
	// 支付网关(仅交易单使用)
	GatewayV1(ctx context.Context, in *PayGatewayRequest, opts ...grpc.CallOption) (*Result, error)
	// 获取支付预交易数据
	GetPreparePaymentInfo(ctx context.Context, in *OrderInfoRequest, opts ...grpc.CallOption) (*SPrepareTradeData, error)
	// 支付网关V2
	GatewayV2(ctx context.Context, in *PayGatewayV2Request, opts ...grpc.CallOption) (*PayGatewayResponse, error)
	// *
	// 支付单混合支付
	//
	// @param storeCode 店铺编号
	// @param tradeNo   交易号
	// @param Data  支付数据
	// @return 支付结果,返回:order_state
	MixedPayment(ctx context.Context, in *MixedPaymentRequest, opts ...grpc.CallOption) (*Result, error)
	// * 保存集成支付应用
	SaveIntegrateApp(ctx context.Context, in *SIntegrateApp, opts ...grpc.CallOption) (*Result, error)
	// * 获取集成支付应用列表
	QueryIntegrateAppList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*QueryIntegrateAppResponse, error)
	// 准备集成支付的参数
	PrepareIntegrateParams(ctx context.Context, in *IntegrateParamsRequest, opts ...grpc.CallOption) (*IntegrateParamsResponse, error)
	// * 删除集成支付应用
	DeleteIntegrateApp(ctx context.Context, in *PayIntegrateAppId, opts ...grpc.CallOption) (*Result, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) SubmitPaymentOrder(ctx context.Context, in *SPaymentOrder, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/PaymentService/SubmitPaymentOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetPaymentOrder(ctx context.Context, in *String, opts ...grpc.CallOption) (*SPaymentOrder, error) {
	out := new(SPaymentOrder)
	err := c.cc.Invoke(ctx, "/PaymentService/GetPaymentOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetPaymentOrderId(ctx context.Context, in *String, opts ...grpc.CallOption) (*Int32, error) {
	out := new(Int32)
	err := c.cc.Invoke(ctx, "/PaymentService/GetPaymentOrderId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetPaymentOrderById(ctx context.Context, in *Int32, opts ...grpc.CallOption) (*SPaymentOrder, error) {
	out := new(SPaymentOrder)
	err := c.cc.Invoke(ctx, "/PaymentService/GetPaymentOrderById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) AdjustOrder(ctx context.Context, in *AdjustOrderRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/PaymentService/AdjustOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) DiscountByBalance(ctx context.Context, in *DiscountBalanceRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/PaymentService/DiscountByBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) DiscountByIntegral(ctx context.Context, in *DiscountIntegralRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/PaymentService/DiscountByIntegral", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) PaymentByWallet(ctx context.Context, in *WalletPaymentRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/PaymentService/PaymentByWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) HybridPayment(ctx context.Context, in *HyperPaymentRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/PaymentService/HybridPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) FinishPayment(ctx context.Context, in *FinishPaymentRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/PaymentService/FinishPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GatewayV1(ctx context.Context, in *PayGatewayRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/PaymentService/GatewayV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetPreparePaymentInfo(ctx context.Context, in *OrderInfoRequest, opts ...grpc.CallOption) (*SPrepareTradeData, error) {
	out := new(SPrepareTradeData)
	err := c.cc.Invoke(ctx, "/PaymentService/GetPreparePaymentInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GatewayV2(ctx context.Context, in *PayGatewayV2Request, opts ...grpc.CallOption) (*PayGatewayResponse, error) {
	out := new(PayGatewayResponse)
	err := c.cc.Invoke(ctx, "/PaymentService/GatewayV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) MixedPayment(ctx context.Context, in *MixedPaymentRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/PaymentService/MixedPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) SaveIntegrateApp(ctx context.Context, in *SIntegrateApp, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/PaymentService/SaveIntegrateApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) QueryIntegrateAppList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*QueryIntegrateAppResponse, error) {
	out := new(QueryIntegrateAppResponse)
	err := c.cc.Invoke(ctx, "/PaymentService/QueryIntegrateAppList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) PrepareIntegrateParams(ctx context.Context, in *IntegrateParamsRequest, opts ...grpc.CallOption) (*IntegrateParamsResponse, error) {
	out := new(IntegrateParamsResponse)
	err := c.cc.Invoke(ctx, "/PaymentService/PrepareIntegrateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) DeleteIntegrateApp(ctx context.Context, in *PayIntegrateAppId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/PaymentService/DeleteIntegrateApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
// All implementations must embed UnimplementedPaymentServiceServer
// for forward compatibility
type PaymentServiceServer interface {
	// 创建支付单并提交
	SubmitPaymentOrder(context.Context, *SPaymentOrder) (*Result, error)
	// 根据支付单号获取支付单,orderNo
	GetPaymentOrder(context.Context, *String) (*SPaymentOrder, error)
	// 根据交易号获取支付单编号,tradeNo
	GetPaymentOrderId(context.Context, *String) (*Int32, error)
	// 根据编号获取支付单
	GetPaymentOrderById(context.Context, *Int32) (*SPaymentOrder, error)
	// 调整支付单金额
	AdjustOrder(context.Context, *AdjustOrderRequest) (*Result, error)
	// 余额抵扣
	DiscountByBalance(context.Context, *DiscountBalanceRequest) (*Result, error)
	// 积分抵扣支付单
	DiscountByIntegral(context.Context, *DiscountIntegralRequest) (*Result, error)
	// 钱包账户支付
	PaymentByWallet(context.Context, *WalletPaymentRequest) (*Result, error)
	// 余额钱包混合支付，优先扣除余额。
	HybridPayment(context.Context, *HyperPaymentRequest) (*Result, error)
	// 完成支付单支付，并传入支付方式及外部订单号
	FinishPayment(context.Context, *FinishPaymentRequest) (*Result, error)
	// 支付网关(仅交易单使用)
	GatewayV1(context.Context, *PayGatewayRequest) (*Result, error)
	// 获取支付预交易数据
	GetPreparePaymentInfo(context.Context, *OrderInfoRequest) (*SPrepareTradeData, error)
	// 支付网关V2
	GatewayV2(context.Context, *PayGatewayV2Request) (*PayGatewayResponse, error)
	// *
	// 支付单混合支付
	//
	// @param storeCode 店铺编号
	// @param tradeNo   交易号
	// @param Data  支付数据
	// @return 支付结果,返回:order_state
	MixedPayment(context.Context, *MixedPaymentRequest) (*Result, error)
	// * 保存集成支付应用
	SaveIntegrateApp(context.Context, *SIntegrateApp) (*Result, error)
	// * 获取集成支付应用列表
	QueryIntegrateAppList(context.Context, *Empty) (*QueryIntegrateAppResponse, error)
	// 准备集成支付的参数
	PrepareIntegrateParams(context.Context, *IntegrateParamsRequest) (*IntegrateParamsResponse, error)
	// * 删除集成支付应用
	DeleteIntegrateApp(context.Context, *PayIntegrateAppId) (*Result, error)
	mustEmbedUnimplementedPaymentServiceServer()
}

// UnimplementedPaymentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentServiceServer struct {
}

func (UnimplementedPaymentServiceServer) SubmitPaymentOrder(context.Context, *SPaymentOrder) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitPaymentOrder not implemented")
}
func (UnimplementedPaymentServiceServer) GetPaymentOrder(context.Context, *String) (*SPaymentOrder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentOrder not implemented")
}
func (UnimplementedPaymentServiceServer) GetPaymentOrderId(context.Context, *String) (*Int32, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentOrderId not implemented")
}
func (UnimplementedPaymentServiceServer) GetPaymentOrderById(context.Context, *Int32) (*SPaymentOrder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentOrderById not implemented")
}
func (UnimplementedPaymentServiceServer) AdjustOrder(context.Context, *AdjustOrderRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdjustOrder not implemented")
}
func (UnimplementedPaymentServiceServer) DiscountByBalance(context.Context, *DiscountBalanceRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscountByBalance not implemented")
}
func (UnimplementedPaymentServiceServer) DiscountByIntegral(context.Context, *DiscountIntegralRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscountByIntegral not implemented")
}
func (UnimplementedPaymentServiceServer) PaymentByWallet(context.Context, *WalletPaymentRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PaymentByWallet not implemented")
}
func (UnimplementedPaymentServiceServer) HybridPayment(context.Context, *HyperPaymentRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HybridPayment not implemented")
}
func (UnimplementedPaymentServiceServer) FinishPayment(context.Context, *FinishPaymentRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishPayment not implemented")
}
func (UnimplementedPaymentServiceServer) GatewayV1(context.Context, *PayGatewayRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GatewayV1 not implemented")
}
func (UnimplementedPaymentServiceServer) GetPreparePaymentInfo(context.Context, *OrderInfoRequest) (*SPrepareTradeData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPreparePaymentInfo not implemented")
}
func (UnimplementedPaymentServiceServer) GatewayV2(context.Context, *PayGatewayV2Request) (*PayGatewayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GatewayV2 not implemented")
}
func (UnimplementedPaymentServiceServer) MixedPayment(context.Context, *MixedPaymentRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MixedPayment not implemented")
}
func (UnimplementedPaymentServiceServer) SaveIntegrateApp(context.Context, *SIntegrateApp) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveIntegrateApp not implemented")
}
func (UnimplementedPaymentServiceServer) QueryIntegrateAppList(context.Context, *Empty) (*QueryIntegrateAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryIntegrateAppList not implemented")
}
func (UnimplementedPaymentServiceServer) PrepareIntegrateParams(context.Context, *IntegrateParamsRequest) (*IntegrateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareIntegrateParams not implemented")
}
func (UnimplementedPaymentServiceServer) DeleteIntegrateApp(context.Context, *PayIntegrateAppId) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIntegrateApp not implemented")
}
func (UnimplementedPaymentServiceServer) mustEmbedUnimplementedPaymentServiceServer() {}

// UnsafePaymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServiceServer will
// result in compilation errors.
type UnsafePaymentServiceServer interface {
	mustEmbedUnimplementedPaymentServiceServer()
}

func RegisterPaymentServiceServer(s grpc.ServiceRegistrar, srv PaymentServiceServer) {
	s.RegisterService(&PaymentService_ServiceDesc, srv)
}

func _PaymentService_SubmitPaymentOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SPaymentOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).SubmitPaymentOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentService/SubmitPaymentOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).SubmitPaymentOrder(ctx, req.(*SPaymentOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetPaymentOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetPaymentOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentService/GetPaymentOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetPaymentOrder(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetPaymentOrderId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetPaymentOrderId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentService/GetPaymentOrderId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetPaymentOrderId(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetPaymentOrderById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int32)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetPaymentOrderById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentService/GetPaymentOrderById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetPaymentOrderById(ctx, req.(*Int32))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_AdjustOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdjustOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).AdjustOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentService/AdjustOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).AdjustOrder(ctx, req.(*AdjustOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_DiscountByBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscountBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).DiscountByBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentService/DiscountByBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).DiscountByBalance(ctx, req.(*DiscountBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_DiscountByIntegral_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscountIntegralRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).DiscountByIntegral(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentService/DiscountByIntegral",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).DiscountByIntegral(ctx, req.(*DiscountIntegralRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_PaymentByWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WalletPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).PaymentByWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentService/PaymentByWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).PaymentByWallet(ctx, req.(*WalletPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_HybridPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HyperPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).HybridPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentService/HybridPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).HybridPayment(ctx, req.(*HyperPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_FinishPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).FinishPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentService/FinishPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).FinishPayment(ctx, req.(*FinishPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GatewayV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GatewayV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentService/GatewayV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GatewayV1(ctx, req.(*PayGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetPreparePaymentInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetPreparePaymentInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentService/GetPreparePaymentInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetPreparePaymentInfo(ctx, req.(*OrderInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GatewayV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayGatewayV2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GatewayV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentService/GatewayV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GatewayV2(ctx, req.(*PayGatewayV2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_MixedPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MixedPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).MixedPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentService/MixedPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).MixedPayment(ctx, req.(*MixedPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_SaveIntegrateApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SIntegrateApp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).SaveIntegrateApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentService/SaveIntegrateApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).SaveIntegrateApp(ctx, req.(*SIntegrateApp))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_QueryIntegrateAppList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).QueryIntegrateAppList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentService/QueryIntegrateAppList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).QueryIntegrateAppList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_PrepareIntegrateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntegrateParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).PrepareIntegrateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentService/PrepareIntegrateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).PrepareIntegrateParams(ctx, req.(*IntegrateParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_DeleteIntegrateApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayIntegrateAppId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).DeleteIntegrateApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentService/DeleteIntegrateApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).DeleteIntegrateApp(ctx, req.(*PayIntegrateAppId))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentService_ServiceDesc is the grpc.ServiceDesc for PaymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitPaymentOrder",
			Handler:    _PaymentService_SubmitPaymentOrder_Handler,
		},
		{
			MethodName: "GetPaymentOrder",
			Handler:    _PaymentService_GetPaymentOrder_Handler,
		},
		{
			MethodName: "GetPaymentOrderId",
			Handler:    _PaymentService_GetPaymentOrderId_Handler,
		},
		{
			MethodName: "GetPaymentOrderById",
			Handler:    _PaymentService_GetPaymentOrderById_Handler,
		},
		{
			MethodName: "AdjustOrder",
			Handler:    _PaymentService_AdjustOrder_Handler,
		},
		{
			MethodName: "DiscountByBalance",
			Handler:    _PaymentService_DiscountByBalance_Handler,
		},
		{
			MethodName: "DiscountByIntegral",
			Handler:    _PaymentService_DiscountByIntegral_Handler,
		},
		{
			MethodName: "PaymentByWallet",
			Handler:    _PaymentService_PaymentByWallet_Handler,
		},
		{
			MethodName: "HybridPayment",
			Handler:    _PaymentService_HybridPayment_Handler,
		},
		{
			MethodName: "FinishPayment",
			Handler:    _PaymentService_FinishPayment_Handler,
		},
		{
			MethodName: "GatewayV1",
			Handler:    _PaymentService_GatewayV1_Handler,
		},
		{
			MethodName: "GetPreparePaymentInfo",
			Handler:    _PaymentService_GetPreparePaymentInfo_Handler,
		},
		{
			MethodName: "GatewayV2",
			Handler:    _PaymentService_GatewayV2_Handler,
		},
		{
			MethodName: "MixedPayment",
			Handler:    _PaymentService_MixedPayment_Handler,
		},
		{
			MethodName: "SaveIntegrateApp",
			Handler:    _PaymentService_SaveIntegrateApp_Handler,
		},
		{
			MethodName: "QueryIntegrateAppList",
			Handler:    _PaymentService_QueryIntegrateAppList_Handler,
		},
		{
			MethodName: "PrepareIntegrateParams",
			Handler:    _PaymentService_PrepareIntegrateParams_Handler,
		},
		{
			MethodName: "DeleteIntegrateApp",
			Handler:    _PaymentService_DeleteIntegrateApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment_service.proto",
}
