// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.7
// source: pkg/desc_test_pkg.proto

package pkg

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

type Foo int32

const (
	Foo_ABC Foo = 0
	Foo_DEF Foo = 1
	Foo_GHI Foo = 2
	Foo_JKL Foo = 3
	Foo_MNO Foo = 4
	Foo_PQR Foo = 5
	Foo_STU Foo = 6
	Foo_VWX Foo = 7
	Foo_Y_Z Foo = 8
)

// Enum value maps for Foo.
var (
	Foo_name = map[int32]string{
		0: "ABC",
		1: "DEF",
		2: "GHI",
		3: "JKL",
		4: "MNO",
		5: "PQR",
		6: "STU",
		7: "VWX",
		8: "Y_Z",
	}
	Foo_value = map[string]int32{
		"ABC": 0,
		"DEF": 1,
		"GHI": 2,
		"JKL": 3,
		"MNO": 4,
		"PQR": 5,
		"STU": 6,
		"VWX": 7,
		"Y_Z": 8,
	}
)

func (x Foo) Enum() *Foo {
	p := new(Foo)
	*p = x
	return p
}

func (x Foo) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Foo) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_desc_test_pkg_proto_enumTypes[0].Descriptor()
}

func (Foo) Type() protoreflect.EnumType {
	return &file_pkg_desc_test_pkg_proto_enumTypes[0]
}

func (x Foo) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Foo) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Foo(num)
	return nil
}

// Deprecated: Use Foo.Descriptor instead.
func (Foo) EnumDescriptor() ([]byte, []int) {
	return file_pkg_desc_test_pkg_proto_rawDescGZIP(), []int{0}
}

type Bar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Baz []Foo `protobuf:"varint,1,rep,name=baz,enum=jhump.protoreflect.desc.Foo" json:"baz,omitempty"`
}

func (x *Bar) Reset() {
	*x = Bar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_desc_test_pkg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bar) ProtoMessage() {}

func (x *Bar) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_desc_test_pkg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bar.ProtoReflect.Descriptor instead.
func (*Bar) Descriptor() ([]byte, []int) {
	return file_pkg_desc_test_pkg_proto_rawDescGZIP(), []int{0}
}

func (x *Bar) GetBaz() []Foo {
	if x != nil {
		return x.Baz
	}
	return nil
}

var File_pkg_desc_test_pkg_proto protoreflect.FileDescriptor

var file_pkg_desc_test_pkg_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6a, 0x68, 0x75, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x64, 0x65,
	0x73, 0x63, 0x22, 0x35, 0x0a, 0x03, 0x42, 0x61, 0x72, 0x12, 0x2e, 0x0a, 0x03, 0x62, 0x61, 0x7a,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6a, 0x68, 0x75, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x64, 0x65, 0x73, 0x63,
	0x2e, 0x46, 0x6f, 0x6f, 0x52, 0x03, 0x62, 0x61, 0x7a, 0x2a, 0x56, 0x0a, 0x03, 0x46, 0x6f, 0x6f,
	0x12, 0x07, 0x0a, 0x03, 0x41, 0x42, 0x43, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x45, 0x46,
	0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x48, 0x49, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x4a,
	0x4b, 0x4c, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x4e, 0x4f, 0x10, 0x04, 0x12, 0x07, 0x0a,
	0x03, 0x50, 0x51, 0x52, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x54, 0x55, 0x10, 0x06, 0x12,
	0x07, 0x0a, 0x03, 0x56, 0x57, 0x58, 0x10, 0x07, 0x12, 0x07, 0x0a, 0x03, 0x59, 0x5f, 0x5a, 0x10,
	0x08, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6a, 0x68, 0x75, 0x6d, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x65, 0x66, 0x6c, 0x65,
	0x63, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x3b, 0x70, 0x6b, 0x67,
}

var (
	file_pkg_desc_test_pkg_proto_rawDescOnce sync.Once
	file_pkg_desc_test_pkg_proto_rawDescData = file_pkg_desc_test_pkg_proto_rawDesc
)

func file_pkg_desc_test_pkg_proto_rawDescGZIP() []byte {
	file_pkg_desc_test_pkg_proto_rawDescOnce.Do(func() {
		file_pkg_desc_test_pkg_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_desc_test_pkg_proto_rawDescData)
	})
	return file_pkg_desc_test_pkg_proto_rawDescData
}

var file_pkg_desc_test_pkg_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_desc_test_pkg_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_desc_test_pkg_proto_goTypes = []interface{}{
	(Foo)(0),    // 0: jhump.protoreflect.desc.Foo
	(*Bar)(nil), // 1: jhump.protoreflect.desc.Bar
}
var file_pkg_desc_test_pkg_proto_depIdxs = []int32{
	0, // 0: jhump.protoreflect.desc.Bar.baz:type_name -> jhump.protoreflect.desc.Foo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_desc_test_pkg_proto_init() }
func file_pkg_desc_test_pkg_proto_init() {
	if File_pkg_desc_test_pkg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_desc_test_pkg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bar); i {
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
			RawDescriptor: file_pkg_desc_test_pkg_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_desc_test_pkg_proto_goTypes,
		DependencyIndexes: file_pkg_desc_test_pkg_proto_depIdxs,
		EnumInfos:         file_pkg_desc_test_pkg_proto_enumTypes,
		MessageInfos:      file_pkg_desc_test_pkg_proto_msgTypes,
	}.Build()
	File_pkg_desc_test_pkg_proto = out.File
	file_pkg_desc_test_pkg_proto_rawDesc = nil
	file_pkg_desc_test_pkg_proto_goTypes = nil
	file_pkg_desc_test_pkg_proto_depIdxs = nil
}
