// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: registry_service.proto

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

// RegistryServiceClient is the client API for RegistryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegistryServiceClient interface {
	// 获取分组
	GetGroups(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RegistryGroupResponse, error)
	// * 获取注册项
	GetRegistry(ctx context.Context, in *String, opts ...grpc.CallOption) (*SRegistry, error)
	// * 获取注册表键值,key
	GetValue(ctx context.Context, in *String, opts ...grpc.CallOption) (*RegistryValueResponse, error)
	// * 更新注册表值
	UpdateValue(ctx context.Context, in *RegistryPair, opts ...grpc.CallOption) (*Result, error)
	// * 获取键值存储数据字典,keys
	GetValues(ctx context.Context, in *StringArray, opts ...grpc.CallOption) (*StringMap, error)
	// * 更新注册表键值
	UpdateValues(ctx context.Context, in *StringMap, opts ...grpc.CallOption) (*Result, error)
	// * 搜索键值
	Search(ctx context.Context, in *RegistrySearchRequest, opts ...grpc.CallOption) (*StringMap, error)
	// * 按键前缀获取键数据,prefix
	FindRegistries(ctx context.Context, in *String, opts ...grpc.CallOption) (*StringMap, error)
	// * 搜索注册表,keyword
	SearchRegistry(ctx context.Context, in *String, opts ...grpc.CallOption) (*RegistriesResponse, error)
	// * 创建自定义注册表项,@defaultValue 默认值,如需更改,使用UpdateRegistry方法
	CreateRegistry(ctx context.Context, in *RegistryCreateRequest, opts ...grpc.CallOption) (*Result, error)
}

type registryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistryServiceClient(cc grpc.ClientConnInterface) RegistryServiceClient {
	return &registryServiceClient{cc}
}

func (c *registryServiceClient) GetGroups(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RegistryGroupResponse, error) {
	out := new(RegistryGroupResponse)
	err := c.cc.Invoke(ctx, "/RegistryService/GetGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) GetRegistry(ctx context.Context, in *String, opts ...grpc.CallOption) (*SRegistry, error) {
	out := new(SRegistry)
	err := c.cc.Invoke(ctx, "/RegistryService/GetRegistry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) GetValue(ctx context.Context, in *String, opts ...grpc.CallOption) (*RegistryValueResponse, error) {
	out := new(RegistryValueResponse)
	err := c.cc.Invoke(ctx, "/RegistryService/GetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) UpdateValue(ctx context.Context, in *RegistryPair, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/RegistryService/UpdateValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) GetValues(ctx context.Context, in *StringArray, opts ...grpc.CallOption) (*StringMap, error) {
	out := new(StringMap)
	err := c.cc.Invoke(ctx, "/RegistryService/GetValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) UpdateValues(ctx context.Context, in *StringMap, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/RegistryService/UpdateValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) Search(ctx context.Context, in *RegistrySearchRequest, opts ...grpc.CallOption) (*StringMap, error) {
	out := new(StringMap)
	err := c.cc.Invoke(ctx, "/RegistryService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) FindRegistries(ctx context.Context, in *String, opts ...grpc.CallOption) (*StringMap, error) {
	out := new(StringMap)
	err := c.cc.Invoke(ctx, "/RegistryService/FindRegistries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) SearchRegistry(ctx context.Context, in *String, opts ...grpc.CallOption) (*RegistriesResponse, error) {
	out := new(RegistriesResponse)
	err := c.cc.Invoke(ctx, "/RegistryService/SearchRegistry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) CreateRegistry(ctx context.Context, in *RegistryCreateRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/RegistryService/CreateRegistry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryServiceServer is the server API for RegistryService service.
// All implementations must embed UnimplementedRegistryServiceServer
// for forward compatibility
type RegistryServiceServer interface {
	// 获取分组
	GetGroups(context.Context, *Empty) (*RegistryGroupResponse, error)
	// * 获取注册项
	GetRegistry(context.Context, *String) (*SRegistry, error)
	// * 获取注册表键值,key
	GetValue(context.Context, *String) (*RegistryValueResponse, error)
	// * 更新注册表值
	UpdateValue(context.Context, *RegistryPair) (*Result, error)
	// * 获取键值存储数据字典,keys
	GetValues(context.Context, *StringArray) (*StringMap, error)
	// * 更新注册表键值
	UpdateValues(context.Context, *StringMap) (*Result, error)
	// * 搜索键值
	Search(context.Context, *RegistrySearchRequest) (*StringMap, error)
	// * 按键前缀获取键数据,prefix
	FindRegistries(context.Context, *String) (*StringMap, error)
	// * 搜索注册表,keyword
	SearchRegistry(context.Context, *String) (*RegistriesResponse, error)
	// * 创建自定义注册表项,@defaultValue 默认值,如需更改,使用UpdateRegistry方法
	CreateRegistry(context.Context, *RegistryCreateRequest) (*Result, error)
	mustEmbedUnimplementedRegistryServiceServer()
}

// UnimplementedRegistryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRegistryServiceServer struct {
}

func (UnimplementedRegistryServiceServer) GetGroups(context.Context, *Empty) (*RegistryGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroups not implemented")
}
func (UnimplementedRegistryServiceServer) GetRegistry(context.Context, *String) (*SRegistry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegistry not implemented")
}
func (UnimplementedRegistryServiceServer) GetValue(context.Context, *String) (*RegistryValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValue not implemented")
}
func (UnimplementedRegistryServiceServer) UpdateValue(context.Context, *RegistryPair) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateValue not implemented")
}
func (UnimplementedRegistryServiceServer) GetValues(context.Context, *StringArray) (*StringMap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValues not implemented")
}
func (UnimplementedRegistryServiceServer) UpdateValues(context.Context, *StringMap) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateValues not implemented")
}
func (UnimplementedRegistryServiceServer) Search(context.Context, *RegistrySearchRequest) (*StringMap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedRegistryServiceServer) FindRegistries(context.Context, *String) (*StringMap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindRegistries not implemented")
}
func (UnimplementedRegistryServiceServer) SearchRegistry(context.Context, *String) (*RegistriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchRegistry not implemented")
}
func (UnimplementedRegistryServiceServer) CreateRegistry(context.Context, *RegistryCreateRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRegistry not implemented")
}
func (UnimplementedRegistryServiceServer) mustEmbedUnimplementedRegistryServiceServer() {}

// UnsafeRegistryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegistryServiceServer will
// result in compilation errors.
type UnsafeRegistryServiceServer interface {
	mustEmbedUnimplementedRegistryServiceServer()
}

func RegisterRegistryServiceServer(s grpc.ServiceRegistrar, srv RegistryServiceServer) {
	s.RegisterService(&RegistryService_ServiceDesc, srv)
}

func _RegistryService_GetGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).GetGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegistryService/GetGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).GetGroups(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_GetRegistry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).GetRegistry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegistryService/GetRegistry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).GetRegistry(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegistryService/GetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).GetValue(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_UpdateValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistryPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).UpdateValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegistryService/UpdateValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).UpdateValue(ctx, req.(*RegistryPair))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_GetValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringArray)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).GetValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegistryService/GetValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).GetValues(ctx, req.(*StringArray))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_UpdateValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMap)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).UpdateValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegistryService/UpdateValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).UpdateValues(ctx, req.(*StringMap))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrySearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegistryService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).Search(ctx, req.(*RegistrySearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_FindRegistries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).FindRegistries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegistryService/FindRegistries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).FindRegistries(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_SearchRegistry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).SearchRegistry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegistryService/SearchRegistry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).SearchRegistry(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_CreateRegistry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistryCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).CreateRegistry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegistryService/CreateRegistry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).CreateRegistry(ctx, req.(*RegistryCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegistryService_ServiceDesc is the grpc.ServiceDesc for RegistryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegistryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RegistryService",
	HandlerType: (*RegistryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroups",
			Handler:    _RegistryService_GetGroups_Handler,
		},
		{
			MethodName: "GetRegistry",
			Handler:    _RegistryService_GetRegistry_Handler,
		},
		{
			MethodName: "GetValue",
			Handler:    _RegistryService_GetValue_Handler,
		},
		{
			MethodName: "UpdateValue",
			Handler:    _RegistryService_UpdateValue_Handler,
		},
		{
			MethodName: "GetValues",
			Handler:    _RegistryService_GetValues_Handler,
		},
		{
			MethodName: "UpdateValues",
			Handler:    _RegistryService_UpdateValues_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _RegistryService_Search_Handler,
		},
		{
			MethodName: "FindRegistries",
			Handler:    _RegistryService_FindRegistries_Handler,
		},
		{
			MethodName: "SearchRegistry",
			Handler:    _RegistryService_SearchRegistry_Handler,
		},
		{
			MethodName: "CreateRegistry",
			Handler:    _RegistryService_CreateRegistry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry_service.proto",
}
