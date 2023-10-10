// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: multiplier.proto

package multiplier

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

type Multiplicands struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int64 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int64 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Multiplicands) Reset() {
	*x = Multiplicands{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiplier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Multiplicands) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Multiplicands) ProtoMessage() {}

func (x *Multiplicands) ProtoReflect() protoreflect.Message {
	mi := &file_multiplier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Multiplicands.ProtoReflect.Descriptor instead.
func (*Multiplicands) Descriptor() ([]byte, []int) {
	return file_multiplier_proto_rawDescGZIP(), []int{0}
}

func (x *Multiplicands) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Multiplicands) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product int64 `protobuf:"varint,1,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiplier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_multiplier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_multiplier_proto_rawDescGZIP(), []int{1}
}

func (x *Product) GetProduct() int64 {
	if x != nil {
		return x.Product
	}
	return 0
}

var File_multiplier_proto protoreflect.FileDescriptor

var file_multiplier_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0d, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x6e, 0x64, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01,
	0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x79, 0x22,
	0x23, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x32, 0x34, 0x0a, 0x0a, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69,
	0x65, 0x72, 0x12, 0x26, 0x0a, 0x08, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x12, 0x0e,
	0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x6e, 0x64, 0x73, 0x1a, 0x08,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x68, 0x69, 0x74, 0x74, 0x61,
	0x6b, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x32,
	0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_multiplier_proto_rawDescOnce sync.Once
	file_multiplier_proto_rawDescData = file_multiplier_proto_rawDesc
)

func file_multiplier_proto_rawDescGZIP() []byte {
	file_multiplier_proto_rawDescOnce.Do(func() {
		file_multiplier_proto_rawDescData = protoimpl.X.CompressGZIP(file_multiplier_proto_rawDescData)
	})
	return file_multiplier_proto_rawDescData
}

var file_multiplier_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_multiplier_proto_goTypes = []interface{}{
	(*Multiplicands)(nil), // 0: Multiplicands
	(*Product)(nil),       // 1: Product
}
var file_multiplier_proto_depIdxs = []int32{
	0, // 0: Multiplier.Multiply:input_type -> Multiplicands
	1, // 1: Multiplier.Multiply:output_type -> Product
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_multiplier_proto_init() }
func file_multiplier_proto_init() {
	if File_multiplier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_multiplier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Multiplicands); i {
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
		file_multiplier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
			RawDescriptor: file_multiplier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_multiplier_proto_goTypes,
		DependencyIndexes: file_multiplier_proto_depIdxs,
		MessageInfos:      file_multiplier_proto_msgTypes,
	}.Build()
	File_multiplier_proto = out.File
	file_multiplier_proto_rawDesc = nil
	file_multiplier_proto_goTypes = nil
	file_multiplier_proto_depIdxs = nil
}