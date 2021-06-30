// openconfig.nocsys_fdb is generated by proto_generator as a protobuf
// representation of a YANG schema.
//
// Input schema modules:
//  - ../api/yang/sonic/nocsys-fdb.yang
//  - ../api/yang/sonic/nocsys-ntp.yang
//  - ../api/yang/sonic/nocsys-todo.yang
//  - ../api/yang/sonic/nocsys-types.yang
//  - ../api/yang/sonic/nocsys-route.yang
//  - ../api/yang/sonic/nocsys-system.yang
//  - ../api/yang/sonic/nocsys-extension.yang
//  - ../api/yang/sonic/nocsys-loopback-interface.yang
//  - ../api/yang/sonic/nocsys-portchannel.yang
//  - ../api/yang/sonic/nocsys-vlan.yang
//  - ../api/yang/sonic/nocsys-port.yang
//  - ../api/yang/sonic/nocsys-platform-types.yang
//  - ../api/yang/sonic/nocsys-mirror-session.yang
//  - ../api/yang/sonic/nocsys-vrf.yang
//  - ../api/yang/sonic/third_party/openconfig/openconfig-extensions.yang
//  - ../api/yang/sonic/third_party/openconfig/openconfig-types.yang
//  - ../api/yang/sonic/third_party/ietf/ietf-yang-types.yang
//  - ../api/yang/sonic/third_party/ietf/ietf-inet-types.yang
//  - ../api/yang/sonic/third_party/ietf/iana-if-type.yang
//  - ../api/yang/sonic/third_party/ietf/ietf-interfaces.yang
//  - ../api/yang/sonic/nocsys-neighbor.yang
//  - ../api/yang/sonic/nocsys-acl.yang
//  - ../api/yang/sonic/nocsys-vxlan.yang
//  - ../api/yang/sonic/nocsys-platform.yang
//  - ../api/yang/sonic/nocsys-interface.yang
//  - ../api/yang/sonic/nocsys-lldp.yang

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: nocsys_fdb.proto

package sonic

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/openconfig/ygot/proto/yext"
	ywrapper "github.com/openconfig/ygot/proto/ywrapper"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type NocsysFdb_Fdb_FdbList_Type int32

const (
	NocsysFdb_Fdb_FdbList_TYPE_UNSET   NocsysFdb_Fdb_FdbList_Type = 0
	NocsysFdb_Fdb_FdbList_TYPE_STATIC  NocsysFdb_Fdb_FdbList_Type = 1
	NocsysFdb_Fdb_FdbList_TYPE_DYNAMIC NocsysFdb_Fdb_FdbList_Type = 2
)

// Enum value maps for NocsysFdb_Fdb_FdbList_Type.
var (
	NocsysFdb_Fdb_FdbList_Type_name = map[int32]string{
		0: "TYPE_UNSET",
		1: "TYPE_STATIC",
		2: "TYPE_DYNAMIC",
	}
	NocsysFdb_Fdb_FdbList_Type_value = map[string]int32{
		"TYPE_UNSET":   0,
		"TYPE_STATIC":  1,
		"TYPE_DYNAMIC": 2,
	}
)

func (x NocsysFdb_Fdb_FdbList_Type) Enum() *NocsysFdb_Fdb_FdbList_Type {
	p := new(NocsysFdb_Fdb_FdbList_Type)
	*p = x
	return p
}

func (x NocsysFdb_Fdb_FdbList_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NocsysFdb_Fdb_FdbList_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_nocsys_fdb_proto_enumTypes[0].Descriptor()
}

func (NocsysFdb_Fdb_FdbList_Type) Type() protoreflect.EnumType {
	return &file_nocsys_fdb_proto_enumTypes[0]
}

func (x NocsysFdb_Fdb_FdbList_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NocsysFdb_Fdb_FdbList_Type.Descriptor instead.
func (NocsysFdb_Fdb_FdbList_Type) EnumDescriptor() ([]byte, []int) {
	return file_nocsys_fdb_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

type NocsysFdb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fdb *NocsysFdb_Fdb `protobuf:"bytes,1,opt,name=fdb,proto3" json:"fdb,omitempty"`
}

func (x *NocsysFdb) Reset() {
	*x = NocsysFdb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nocsys_fdb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NocsysFdb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NocsysFdb) ProtoMessage() {}

func (x *NocsysFdb) ProtoReflect() protoreflect.Message {
	mi := &file_nocsys_fdb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NocsysFdb.ProtoReflect.Descriptor instead.
func (*NocsysFdb) Descriptor() ([]byte, []int) {
	return file_nocsys_fdb_proto_rawDescGZIP(), []int{0}
}

func (x *NocsysFdb) GetFdb() *NocsysFdb_Fdb {
	if x != nil {
		return x.Fdb
	}
	return nil
}

type NocsysFdb_Fdb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FdbList []*NocsysFdb_Fdb_FdbListKey `protobuf:"bytes,1,rep,name=fdb_list,json=fdbList,proto3" json:"fdb_list,omitempty"`
}

func (x *NocsysFdb_Fdb) Reset() {
	*x = NocsysFdb_Fdb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nocsys_fdb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NocsysFdb_Fdb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NocsysFdb_Fdb) ProtoMessage() {}

func (x *NocsysFdb_Fdb) ProtoReflect() protoreflect.Message {
	mi := &file_nocsys_fdb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NocsysFdb_Fdb.ProtoReflect.Descriptor instead.
func (*NocsysFdb_Fdb) Descriptor() ([]byte, []int) {
	return file_nocsys_fdb_proto_rawDescGZIP(), []int{0, 0}
}

func (x *NocsysFdb_Fdb) GetFdbList() []*NocsysFdb_Fdb_FdbListKey {
	if x != nil {
		return x.FdbList
	}
	return nil
}

type NocsysFdb_Fdb_FdbList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port *ywrapper.StringValue      `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
	Type NocsysFdb_Fdb_FdbList_Type `protobuf:"varint,2,opt,name=type,proto3,enum=sonic.NocsysFdb_Fdb_FdbList_Type" json:"type,omitempty"`
}

func (x *NocsysFdb_Fdb_FdbList) Reset() {
	*x = NocsysFdb_Fdb_FdbList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nocsys_fdb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NocsysFdb_Fdb_FdbList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NocsysFdb_Fdb_FdbList) ProtoMessage() {}

func (x *NocsysFdb_Fdb_FdbList) ProtoReflect() protoreflect.Message {
	mi := &file_nocsys_fdb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NocsysFdb_Fdb_FdbList.ProtoReflect.Descriptor instead.
func (*NocsysFdb_Fdb_FdbList) Descriptor() ([]byte, []int) {
	return file_nocsys_fdb_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *NocsysFdb_Fdb_FdbList) GetPort() *ywrapper.StringValue {
	if x != nil {
		return x.Port
	}
	return nil
}

func (x *NocsysFdb_Fdb_FdbList) GetType() NocsysFdb_Fdb_FdbList_Type {
	if x != nil {
		return x.Type
	}
	return NocsysFdb_Fdb_FdbList_TYPE_UNSET
}

type NocsysFdb_Fdb_FdbListKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FdbName string                 `protobuf:"bytes,1,opt,name=fdb_name,json=fdbName,proto3" json:"fdb_name,omitempty"`
	FdbList *NocsysFdb_Fdb_FdbList `protobuf:"bytes,2,opt,name=fdb_list,json=fdbList,proto3" json:"fdb_list,omitempty"`
}

func (x *NocsysFdb_Fdb_FdbListKey) Reset() {
	*x = NocsysFdb_Fdb_FdbListKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nocsys_fdb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NocsysFdb_Fdb_FdbListKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NocsysFdb_Fdb_FdbListKey) ProtoMessage() {}

func (x *NocsysFdb_Fdb_FdbListKey) ProtoReflect() protoreflect.Message {
	mi := &file_nocsys_fdb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NocsysFdb_Fdb_FdbListKey.ProtoReflect.Descriptor instead.
func (*NocsysFdb_Fdb_FdbListKey) Descriptor() ([]byte, []int) {
	return file_nocsys_fdb_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *NocsysFdb_Fdb_FdbListKey) GetFdbName() string {
	if x != nil {
		return x.FdbName
	}
	return ""
}

func (x *NocsysFdb_Fdb_FdbListKey) GetFdbList() *NocsysFdb_Fdb_FdbList {
	if x != nil {
		return x.FdbList
	}
	return nil
}

var File_nocsys_fdb_proto protoreflect.FileDescriptor

var file_nocsys_fdb_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6e, 0x6f, 0x63, 0x73, 0x79, 0x73, 0x5f, 0x66, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x1a, 0x0e, 0x79, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x79, 0x65, 0x78, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x04, 0x0a, 0x09, 0x4e, 0x6f, 0x63, 0x73, 0x79, 0x73,
	0x46, 0x64, 0x62, 0x12, 0x3a, 0x0a, 0x03, 0x66, 0x64, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x4e, 0x6f, 0x63, 0x73, 0x79, 0x73, 0x46,
	0x64, 0x62, 0x2e, 0x46, 0x64, 0x62, 0x42, 0x12, 0x82, 0x41, 0x0f, 0x2f, 0x6e, 0x6f, 0x63, 0x73,
	0x79, 0x73, 0x2d, 0x66, 0x64, 0x62, 0x2f, 0x66, 0x64, 0x62, 0x52, 0x03, 0x66, 0x64, 0x62, 0x1a,
	0xeb, 0x03, 0x0a, 0x03, 0x46, 0x64, 0x62, 0x12, 0x57, 0x0a, 0x08, 0x66, 0x64, 0x62, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x6f, 0x6e, 0x69,
	0x63, 0x2e, 0x4e, 0x6f, 0x63, 0x73, 0x79, 0x73, 0x46, 0x64, 0x62, 0x2e, 0x46, 0x64, 0x62, 0x2e,
	0x46, 0x64, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x42, 0x1b, 0x82, 0x41, 0x18, 0x2f,
	0x6e, 0x6f, 0x63, 0x73, 0x79, 0x73, 0x2d, 0x66, 0x64, 0x62, 0x2f, 0x66, 0x64, 0x62, 0x2f, 0x66,
	0x64, 0x62, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x07, 0x66, 0x64, 0x62, 0x4c, 0x69, 0x73, 0x74,
	0x1a, 0x81, 0x02, 0x0a, 0x07, 0x46, 0x64, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x79, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x20, 0x82, 0x41, 0x1d, 0x2f, 0x6e, 0x6f, 0x63, 0x73, 0x79, 0x73, 0x2d, 0x66, 0x64,
	0x62, 0x2f, 0x66, 0x64, 0x62, 0x2f, 0x66, 0x64, 0x62, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x57, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e,
	0x4e, 0x6f, 0x63, 0x73, 0x79, 0x73, 0x46, 0x64, 0x62, 0x2e, 0x46, 0x64, 0x62, 0x2e, 0x46, 0x64,
	0x62, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x20, 0x82, 0x41, 0x1d, 0x2f,
	0x6e, 0x6f, 0x63, 0x73, 0x79, 0x73, 0x2d, 0x66, 0x64, 0x62, 0x2f, 0x66, 0x64, 0x62, 0x2f, 0x66,
	0x64, 0x62, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x50, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x0b, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x49, 0x43, 0x10, 0x01, 0x1a, 0x09, 0x82, 0x41, 0x06,
	0x53, 0x54, 0x41, 0x54, 0x49, 0x43, 0x12, 0x1c, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44,
	0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x10, 0x02, 0x1a, 0x0a, 0x82, 0x41, 0x07, 0x44, 0x59, 0x4e,
	0x41, 0x4d, 0x49, 0x43, 0x1a, 0x86, 0x01, 0x0a, 0x0a, 0x46, 0x64, 0x62, 0x4c, 0x69, 0x73, 0x74,
	0x4b, 0x65, 0x79, 0x12, 0x3f, 0x0a, 0x08, 0x66, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x24, 0x82, 0x41, 0x21, 0x2f, 0x6e, 0x6f, 0x63, 0x73, 0x79,
	0x73, 0x2d, 0x66, 0x64, 0x62, 0x2f, 0x66, 0x64, 0x62, 0x2f, 0x66, 0x64, 0x62, 0x2d, 0x6c, 0x69,
	0x73, 0x74, 0x2f, 0x66, 0x64, 0x62, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x07, 0x66, 0x64, 0x62,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x66, 0x64, 0x62, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x4e,
	0x6f, 0x63, 0x73, 0x79, 0x73, 0x46, 0x64, 0x62, 0x2e, 0x46, 0x64, 0x62, 0x2e, 0x46, 0x64, 0x62,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x07, 0x66, 0x64, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x1b, 0x0a,
	0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x6f, 0x63, 0x73, 0x79, 0x73, 0x2e, 0x73, 0x6f, 0x6e, 0x69,
	0x63, 0x5a, 0x07, 0x2e, 0x3b, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_nocsys_fdb_proto_rawDescOnce sync.Once
	file_nocsys_fdb_proto_rawDescData = file_nocsys_fdb_proto_rawDesc
)

func file_nocsys_fdb_proto_rawDescGZIP() []byte {
	file_nocsys_fdb_proto_rawDescOnce.Do(func() {
		file_nocsys_fdb_proto_rawDescData = protoimpl.X.CompressGZIP(file_nocsys_fdb_proto_rawDescData)
	})
	return file_nocsys_fdb_proto_rawDescData
}

var file_nocsys_fdb_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nocsys_fdb_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nocsys_fdb_proto_goTypes = []interface{}{
	(NocsysFdb_Fdb_FdbList_Type)(0),  // 0: sonic.NocsysFdb.Fdb.FdbList.Type
	(*NocsysFdb)(nil),                // 1: sonic.NocsysFdb
	(*NocsysFdb_Fdb)(nil),            // 2: sonic.NocsysFdb.Fdb
	(*NocsysFdb_Fdb_FdbList)(nil),    // 3: sonic.NocsysFdb.Fdb.FdbList
	(*NocsysFdb_Fdb_FdbListKey)(nil), // 4: sonic.NocsysFdb.Fdb.FdbListKey
	(*ywrapper.StringValue)(nil),     // 5: ywrapper.StringValue
}
var file_nocsys_fdb_proto_depIdxs = []int32{
	2, // 0: sonic.NocsysFdb.fdb:type_name -> sonic.NocsysFdb.Fdb
	4, // 1: sonic.NocsysFdb.Fdb.fdb_list:type_name -> sonic.NocsysFdb.Fdb.FdbListKey
	5, // 2: sonic.NocsysFdb.Fdb.FdbList.port:type_name -> ywrapper.StringValue
	0, // 3: sonic.NocsysFdb.Fdb.FdbList.type:type_name -> sonic.NocsysFdb.Fdb.FdbList.Type
	3, // 4: sonic.NocsysFdb.Fdb.FdbListKey.fdb_list:type_name -> sonic.NocsysFdb.Fdb.FdbList
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_nocsys_fdb_proto_init() }
func file_nocsys_fdb_proto_init() {
	if File_nocsys_fdb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nocsys_fdb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NocsysFdb); i {
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
		file_nocsys_fdb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NocsysFdb_Fdb); i {
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
		file_nocsys_fdb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NocsysFdb_Fdb_FdbList); i {
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
		file_nocsys_fdb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NocsysFdb_Fdb_FdbListKey); i {
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
			RawDescriptor: file_nocsys_fdb_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nocsys_fdb_proto_goTypes,
		DependencyIndexes: file_nocsys_fdb_proto_depIdxs,
		EnumInfos:         file_nocsys_fdb_proto_enumTypes,
		MessageInfos:      file_nocsys_fdb_proto_msgTypes,
	}.Build()
	File_nocsys_fdb_proto = out.File
	file_nocsys_fdb_proto_rawDesc = nil
	file_nocsys_fdb_proto_goTypes = nil
	file_nocsys_fdb_proto_depIdxs = nil
}
