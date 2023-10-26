// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: account.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ModelUsages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ModelUsages_Usage `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ModelUsages) Reset() {
	*x = ModelUsages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelUsages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelUsages) ProtoMessage() {}

func (x *ModelUsages) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelUsages.ProtoReflect.Descriptor instead.
func (*ModelUsages) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{0}
}

func (x *ModelUsages) GetItems() []*ModelUsages_Usage {
	if x != nil {
		return x.Items
	}
	return nil
}

type Payments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Payments_Payment `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Payments) Reset() {
	*x = Payments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payments) ProtoMessage() {}

func (x *Payments) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payments.ProtoReflect.Descriptor instead.
func (*Payments) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{1}
}

func (x *Payments) GetItems() []*Payments_Payment {
	if x != nil {
		return x.Items
	}
	return nil
}

type ModelUsages_Usage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Model  string `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	Input  uint32 `protobuf:"varint,2,opt,name=input,proto3" json:"input,omitempty"`
	Output uint32 `protobuf:"varint,3,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *ModelUsages_Usage) Reset() {
	*x = ModelUsages_Usage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelUsages_Usage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelUsages_Usage) ProtoMessage() {}

func (x *ModelUsages_Usage) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelUsages_Usage.ProtoReflect.Descriptor instead.
func (*ModelUsages_Usage) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ModelUsages_Usage) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *ModelUsages_Usage) GetInput() uint32 {
	if x != nil {
		return x.Input
	}
	return 0
}

func (x *ModelUsages_Usage) GetOutput() uint32 {
	if x != nil {
		return x.Output
	}
	return 0
}

type Payments_Payment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Date   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Amount uint32                 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Payments_Payment) Reset() {
	*x = Payments_Payment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payments_Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payments_Payment) ProtoMessage() {}

func (x *Payments_Payment) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payments_Payment.ProtoReflect.Descriptor instead.
func (*Payments_Payment) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Payments_Payment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Payments_Payment) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Payments_Payment) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_account_proto protoreflect.FileDescriptor

var file_account_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x62, 0x72, 0x61, 0x69, 0x6e, 0x62,
	0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x01,
	0x0a, 0x0b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x55, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x47, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x62, 0x72, 0x61, 0x69, 0x6e, 0x62, 0x6f, 0x6f,
	0x73, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x55, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x4b, 0x0a, 0x05, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0xb5, 0x01, 0x0a, 0x08, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x46, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x30, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x62, 0x72, 0x61, 0x69, 0x6e,
	0x62, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x61, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xb8, 0x01, 0x0a, 0x0e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x55, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2b, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x62, 0x72, 0x61, 0x69, 0x6e, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x4f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x28, 0x2e, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x62, 0x72, 0x61, 0x69, 0x6e, 0x62, 0x6f, 0x6f,
	0x73, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_proto_rawDescOnce sync.Once
	file_account_proto_rawDescData = file_account_proto_rawDesc
)

func file_account_proto_rawDescGZIP() []byte {
	file_account_proto_rawDescOnce.Do(func() {
		file_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_proto_rawDescData)
	})
	return file_account_proto_rawDescData
}

var file_account_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_account_proto_goTypes = []interface{}{
	(*ModelUsages)(nil),           // 0: endpoint.brainboost.account.v1.ModelUsages
	(*Payments)(nil),              // 1: endpoint.brainboost.account.v1.Payments
	(*ModelUsages_Usage)(nil),     // 2: endpoint.brainboost.account.v1.ModelUsages.Usage
	(*Payments_Payment)(nil),      // 3: endpoint.brainboost.account.v1.Payments.Payment
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 5: google.protobuf.Empty
}
var file_account_proto_depIdxs = []int32{
	2, // 0: endpoint.brainboost.account.v1.ModelUsages.items:type_name -> endpoint.brainboost.account.v1.ModelUsages.Usage
	3, // 1: endpoint.brainboost.account.v1.Payments.items:type_name -> endpoint.brainboost.account.v1.Payments.Payment
	4, // 2: endpoint.brainboost.account.v1.Payments.Payment.date:type_name -> google.protobuf.Timestamp
	5, // 3: endpoint.brainboost.account.v1.AccountService.GetModelUsages:input_type -> google.protobuf.Empty
	5, // 4: endpoint.brainboost.account.v1.AccountService.GetPayments:input_type -> google.protobuf.Empty
	0, // 5: endpoint.brainboost.account.v1.AccountService.GetModelUsages:output_type -> endpoint.brainboost.account.v1.ModelUsages
	1, // 6: endpoint.brainboost.account.v1.AccountService.GetPayments:output_type -> endpoint.brainboost.account.v1.Payments
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_account_proto_init() }
func file_account_proto_init() {
	if File_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelUsages); i {
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
		file_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payments); i {
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
		file_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelUsages_Usage); i {
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
		file_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payments_Payment); i {
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
			RawDescriptor: file_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_proto_goTypes,
		DependencyIndexes: file_account_proto_depIdxs,
		MessageInfos:      file_account_proto_msgTypes,
	}.Build()
	File_account_proto = out.File
	file_account_proto_rawDesc = nil
	file_account_proto_goTypes = nil
	file_account_proto_depIdxs = nil
}
