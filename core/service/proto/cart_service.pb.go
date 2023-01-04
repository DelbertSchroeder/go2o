// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: cart_service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 购物车加入商品请求
type CartItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 买家编号
	Id *ShoppingCartId `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 商品编号
	ItemId int64 `protobuf:"zigzag64,2,opt,name=itemId,proto3" json:"itemId"`
	// SKU编号
	SkuId int64 `protobuf:"zigzag64,3,opt,name=skuId,proto3" json:"skuId"`
	// 数量
	Quantity int32 `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity"`
	// 是否只勾选当前商品(直接购买)
	CheckOnly bool `protobuf:"varint,5,opt,name=checkOnly,proto3" json:"checkOnly"`
}

func (x *CartItemRequest) Reset() {
	*x = CartItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemRequest) ProtoMessage() {}

func (x *CartItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemRequest.ProtoReflect.Descriptor instead.
func (*CartItemRequest) Descriptor() ([]byte, []int) {
	return file_cart_service_proto_rawDescGZIP(), []int{0}
}

func (x *CartItemRequest) GetId() *ShoppingCartId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *CartItemRequest) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *CartItemRequest) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *CartItemRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *CartItemRequest) GetCheckOnly() bool {
	if x != nil {
		return x.CheckOnly
	}
	return false
}

type CheckCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    *ShoppingCartId   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Items []*SCheckCartItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *CheckCartRequest) Reset() {
	*x = CheckCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckCartRequest) ProtoMessage() {}

func (x *CheckCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckCartRequest.ProtoReflect.Descriptor instead.
func (*CheckCartRequest) Descriptor() ([]byte, []int) {
	return file_cart_service_proto_rawDescGZIP(), []int{1}
}

func (x *CheckCartRequest) GetId() *ShoppingCartId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *CheckCartRequest) GetItems() []*SCheckCartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

// 购物车编号
type ShoppingCartId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 会员/用户编号
	UserId int64 `protobuf:"zigzag64,1,opt,name=userId,proto3" json:"userId"`
	// 购物车标识,当未指定用户时候使用
	CartCode string `protobuf:"bytes,2,opt,name=cartCode,proto3" json:"cartCode"`
	// 是否为批发销售的购物车
	IsWholesale bool `protobuf:"varint,3,opt,name=isWholesale,proto3" json:"isWholesale"`
}

func (x *ShoppingCartId) Reset() {
	*x = ShoppingCartId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShoppingCartId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShoppingCartId) ProtoMessage() {}

func (x *ShoppingCartId) ProtoReflect() protoreflect.Message {
	mi := &file_cart_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShoppingCartId.ProtoReflect.Descriptor instead.
func (*ShoppingCartId) Descriptor() ([]byte, []int) {
	return file_cart_service_proto_rawDescGZIP(), []int{2}
}

func (x *ShoppingCartId) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ShoppingCartId) GetCartCode() string {
	if x != nil {
		return x.CartCode
	}
	return ""
}

func (x *ShoppingCartId) GetIsWholesale() bool {
	if x != nil {
		return x.IsWholesale
	}
	return false
}

// 购物车商品更新请求
type CartUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 购物车编号
	CartId *ShoppingCartId `protobuf:"bytes,1,opt,name=cartId,proto3" json:"cartId"`
	// Sku列表
	Items []*SCartItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *CartUpdateRequest) Reset() {
	*x = CartUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartUpdateRequest) ProtoMessage() {}

func (x *CartUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartUpdateRequest.ProtoReflect.Descriptor instead.
func (*CartUpdateRequest) Descriptor() ([]byte, []int) {
	return file_cart_service_proto_rawDescGZIP(), []int{3}
}

func (x *CartUpdateRequest) GetCartId() *ShoppingCartId {
	if x != nil {
		return x.CartId
	}
	return nil
}

func (x *CartUpdateRequest) GetItems() []*SCartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_cart_service_proto protoreflect.FileDescriptor

var file_cart_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x61, 0x72, 0x74,
	0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x0f, 0x43,
	0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x53, 0x68, 0x6f,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x74, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x12, 0x52,
	0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x12, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x4f, 0x6e, 0x6c, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0x5a, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x43, 0x61, 0x72, 0x74, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x53, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x66, 0x0a, 0x0e, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43,
	0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x57,
	0x68, 0x6f, 0x6c, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x69, 0x73, 0x57, 0x68, 0x6f, 0x6c, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x22, 0x5e, 0x0a, 0x11, 0x43,
	0x61, 0x72, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x27, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x64, 0x52, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xb7, 0x02, 0x0a, 0x0b,
	0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x0f, 0x57,
	0x68, 0x6f, 0x6c, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x43, 0x61, 0x72, 0x74, 0x56, 0x31, 0x12, 0x0e,
	0x2e, 0x57, 0x73, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x74, 0x12, 0x0f, 0x2e, 0x53,
	0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x74, 0x49, 0x64, 0x1a, 0x0e, 0x2e,
	0x53, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x74, 0x22, 0x00, 0x12,
	0x32, 0x0a, 0x09, 0x50, 0x75, 0x74, 0x49, 0x6e, 0x43, 0x61, 0x72, 0x74, 0x12, 0x10, 0x2e, 0x43,
	0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x12, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x12, 0x37, 0x0a, 0x0e, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x10, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x09, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x74, 0x12, 0x11, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x1f, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x67, 0x6f, 0x32, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cart_service_proto_rawDescOnce sync.Once
	file_cart_service_proto_rawDescData = file_cart_service_proto_rawDesc
)

func file_cart_service_proto_rawDescGZIP() []byte {
	file_cart_service_proto_rawDescOnce.Do(func() {
		file_cart_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cart_service_proto_rawDescData)
	})
	return file_cart_service_proto_rawDescData
}

var file_cart_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cart_service_proto_goTypes = []interface{}{
	(*CartItemRequest)(nil),   // 0: CartItemRequest
	(*CheckCartRequest)(nil),  // 1: CheckCartRequest
	(*ShoppingCartId)(nil),    // 2: ShoppingCartId
	(*CartUpdateRequest)(nil), // 3: CartUpdateRequest
	(*SCheckCartItem)(nil),    // 4: SCheckCartItem
	(*SCartItem)(nil),         // 5: SCartItem
	(*WsCartRequest)(nil),     // 6: WsCartRequest
	(*Result)(nil),            // 7: Result
	(*SShoppingCart)(nil),     // 8: SShoppingCart
	(*CartItemResponse)(nil),  // 9: CartItemResponse
}
var file_cart_service_proto_depIdxs = []int32{
	2,  // 0: CartItemRequest.id:type_name -> ShoppingCartId
	2,  // 1: CheckCartRequest.id:type_name -> ShoppingCartId
	4,  // 2: CheckCartRequest.items:type_name -> SCheckCartItem
	2,  // 3: CartUpdateRequest.cartId:type_name -> ShoppingCartId
	5,  // 4: CartUpdateRequest.items:type_name -> SCartItem
	6,  // 5: CartService.WholesaleCartV1:input_type -> WsCartRequest
	2,  // 6: CartService.GetShoppingCart:input_type -> ShoppingCartId
	0,  // 7: CartService.PutInCart:input_type -> CartItemRequest
	3,  // 8: CartService.UpdateItems:input_type -> CartUpdateRequest
	0,  // 9: CartService.ReduceCartItem:input_type -> CartItemRequest
	1,  // 10: CartService.CheckCart:input_type -> CheckCartRequest
	7,  // 11: CartService.WholesaleCartV1:output_type -> Result
	8,  // 12: CartService.GetShoppingCart:output_type -> SShoppingCart
	9,  // 13: CartService.PutInCart:output_type -> CartItemResponse
	7,  // 14: CartService.UpdateItems:output_type -> Result
	9,  // 15: CartService.ReduceCartItem:output_type -> CartItemResponse
	7,  // 16: CartService.CheckCart:output_type -> Result
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_cart_service_proto_init() }
func file_cart_service_proto_init() {
	if File_cart_service_proto != nil {
		return
	}
	file_global_proto_init()
	file_message_cart_dto_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cart_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cart_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckCartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cart_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShoppingCartId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cart_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartUpdateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cart_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cart_service_proto_goTypes,
		DependencyIndexes: file_cart_service_proto_depIdxs,
		MessageInfos:      file_cart_service_proto_msgTypes,
	}.Build()
	File_cart_service_proto = out.File
	file_cart_service_proto_rawDesc = nil
	file_cart_service_proto_goTypes = nil
	file_cart_service_proto_depIdxs = nil
}
