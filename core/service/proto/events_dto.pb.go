// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: message/events_dto.proto

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

// 商品分销设置
type EVItemAffiliteConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分销比例
	Rate float64 `protobuf:"fixed64,2,opt,name=rate,proto3" json:"rate"`
	// 分销金额
	Amount int64 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount"`
	// 是否按比例分销
	IsByRate bool `protobuf:"varint,4,opt,name=isByRate,proto3" json:"isByRate"`
}

func (x *EVItemAffiliteConfig) Reset() {
	*x = EVItemAffiliteConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_events_dto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EVItemAffiliteConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EVItemAffiliteConfig) ProtoMessage() {}

func (x *EVItemAffiliteConfig) ProtoReflect() protoreflect.Message {
	mi := &file_message_events_dto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EVItemAffiliteConfig.ProtoReflect.Descriptor instead.
func (*EVItemAffiliteConfig) Descriptor() ([]byte, []int) {
	return file_message_events_dto_proto_rawDescGZIP(), []int{0}
}

func (x *EVItemAffiliteConfig) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *EVItemAffiliteConfig) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *EVItemAffiliteConfig) GetIsByRate() bool {
	if x != nil {
		return x.IsByRate
	}
	return false
}

// 订单分销商品
type EVOrderAffiliteItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 商品编号
	ItemId int64 `protobuf:"varint,1,opt,name=itemId,proto3" json:"itemId"`
	// 商品SKU编号
	SkuId int64 `protobuf:"varint,2,opt,name=skuId,proto3" json:"skuId"`
	// 数量
	Quantity int32 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity"`
	// 金额
	Amount int64 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount"`
	// 最终金额, 可能会有优惠均摊抵扣的金额
	FinalAmount int64 `protobuf:"varint,5,opt,name=finalAmount,proto3" json:"finalAmount"`
	// 分销参数
	Params []*EVItemAffiliteConfig `protobuf:"bytes,6,rep,name=Params,proto3" json:"Params"`
}

func (x *EVOrderAffiliteItem) Reset() {
	*x = EVOrderAffiliteItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_events_dto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EVOrderAffiliteItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EVOrderAffiliteItem) ProtoMessage() {}

func (x *EVOrderAffiliteItem) ProtoReflect() protoreflect.Message {
	mi := &file_message_events_dto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EVOrderAffiliteItem.ProtoReflect.Descriptor instead.
func (*EVOrderAffiliteItem) Descriptor() ([]byte, []int) {
	return file_message_events_dto_proto_rawDescGZIP(), []int{1}
}

func (x *EVOrderAffiliteItem) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *EVOrderAffiliteItem) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *EVOrderAffiliteItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *EVOrderAffiliteItem) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *EVOrderAffiliteItem) GetFinalAmount() int64 {
	if x != nil {
		return x.FinalAmount
	}
	return 0
}

func (x *EVOrderAffiliteItem) GetParams() []*EVItemAffiliteConfig {
	if x != nil {
		return x.Params
	}
	return nil
}

// 分销订单
type EVOrderAffiliteRebateOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 订单号
	OrderNo string `protobuf:"bytes,1,opt,name=orderNo,proto3" json:"orderNo"`
	// 订单金额
	OrderAmount int64 `protobuf:"varint,2,opt,name=orderAmount,proto3" json:"orderAmount"`
	// 分销商品
	AffiliteItems []*EVOrderAffiliteItem `protobuf:"bytes,3,rep,name=affiliteItems,proto3" json:"affiliteItems"`
}

func (x *EVOrderAffiliteRebateOrder) Reset() {
	*x = EVOrderAffiliteRebateOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_events_dto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EVOrderAffiliteRebateOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EVOrderAffiliteRebateOrder) ProtoMessage() {}

func (x *EVOrderAffiliteRebateOrder) ProtoReflect() protoreflect.Message {
	mi := &file_message_events_dto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EVOrderAffiliteRebateOrder.ProtoReflect.Descriptor instead.
func (*EVOrderAffiliteRebateOrder) Descriptor() ([]byte, []int) {
	return file_message_events_dto_proto_rawDescGZIP(), []int{2}
}

func (x *EVOrderAffiliteRebateOrder) GetOrderNo() string {
	if x != nil {
		return x.OrderNo
	}
	return ""
}

func (x *EVOrderAffiliteRebateOrder) GetOrderAmount() int64 {
	if x != nil {
		return x.OrderAmount
	}
	return 0
}

func (x *EVOrderAffiliteRebateOrder) GetAffiliteItems() []*EVOrderAffiliteItem {
	if x != nil {
		return x.AffiliteItems
	}
	return nil
}

var File_message_events_dto_proto protoreflect.FileDescriptor

var file_message_events_dto_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x14, 0x45, 0x56,
	0x49, 0x74, 0x65, 0x6d, 0x41, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x73, 0x42, 0x79, 0x52, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x42, 0x79, 0x52, 0x61, 0x74, 0x65, 0x22, 0xc8, 0x01, 0x0a, 0x13, 0x45,
	0x56, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x74, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6b,
	0x75, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x61, 0x6c,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x45, 0x56, 0x49, 0x74, 0x65, 0x6d, 0x41,
	0x66, 0x66, 0x69, 0x6c, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x1a, 0x45, 0x56, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x41, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x74, 0x65, 0x52, 0x65, 0x62, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x12, 0x20,
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x3a, 0x0a, 0x0d, 0x61, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x45, 0x56, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x41, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0d, 0x61,
	0x66, 0x66, 0x69, 0x6c, 0x69, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x1f, 0x0a, 0x13,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x67, 0x6f, 0x32, 0x6f, 0x2e,
	0x72, 0x70, 0x63, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_events_dto_proto_rawDescOnce sync.Once
	file_message_events_dto_proto_rawDescData = file_message_events_dto_proto_rawDesc
)

func file_message_events_dto_proto_rawDescGZIP() []byte {
	file_message_events_dto_proto_rawDescOnce.Do(func() {
		file_message_events_dto_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_events_dto_proto_rawDescData)
	})
	return file_message_events_dto_proto_rawDescData
}

var file_message_events_dto_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_message_events_dto_proto_goTypes = []interface{}{
	(*EVItemAffiliteConfig)(nil),       // 0: EVItemAffiliteConfig
	(*EVOrderAffiliteItem)(nil),        // 1: EVOrderAffiliteItem
	(*EVOrderAffiliteRebateOrder)(nil), // 2: EVOrderAffiliteRebateOrder
}
var file_message_events_dto_proto_depIdxs = []int32{
	0, // 0: EVOrderAffiliteItem.Params:type_name -> EVItemAffiliteConfig
	1, // 1: EVOrderAffiliteRebateOrder.affiliteItems:type_name -> EVOrderAffiliteItem
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_message_events_dto_proto_init() }
func file_message_events_dto_proto_init() {
	if File_message_events_dto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_events_dto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EVItemAffiliteConfig); i {
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
		file_message_events_dto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EVOrderAffiliteItem); i {
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
		file_message_events_dto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EVOrderAffiliteRebateOrder); i {
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
			RawDescriptor: file_message_events_dto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_events_dto_proto_goTypes,
		DependencyIndexes: file_message_events_dto_proto_depIdxs,
		MessageInfos:      file_message_events_dto_proto_msgTypes,
	}.Build()
	File_message_events_dto_proto = out.File
	file_message_events_dto_proto_rawDesc = nil
	file_message_events_dto_proto_goTypes = nil
	file_message_events_dto_proto_depIdxs = nil
}