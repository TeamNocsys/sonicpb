// openconfig.accton_loopback_interface is generated by proto_generator as a protobuf
// representation of a YANG schema.
//
// Input schema modules:
//  - ../api/yang/sonic/nocsys-types.yang
//  - ../api/yang/sonic/nocsys-interface.yang
//  - ../api/yang/sonic/nocsys-vlan.yang
//  - ../api/yang/sonic/nocsys-mdns.yang
//  - ../api/yang/sonic/nocsys-extension.yang
//  - ../api/yang/sonic/nocsys-acl.yang
//  - ../api/yang/sonic/nocsys-loopback-interface.yang
//  - ../api/yang/sonic/nocsys-vrf.yang
//  - ../api/yang/sonic/nocsys-portchannel.yang
//  - ../api/yang/sonic/nocsys-system.yang
//  - ../api/yang/sonic/nocsys-neighbor.yang
//  - ../api/yang/sonic/nocsys-vxlan.yang
//  - ../api/yang/sonic/nocsys-port.yang
//  - ../api/yang/sonic/nocsys-route.yang
//  - ../api/yang/sonic/nocsys-ntp.yang
//  - ../api/yang/sonic/nocsys-fdb.yang
//  - ../api/yang/sonic/nocsys-platform-types.yang
//  - ../api/yang/sonic/nocsys-platform.yang
//  - ../api/yang/sonic/nocsys-lldp.yang
//  - ../api/yang/sonic/nocsys-mirror-session.yang
//  - ../api/yang/sonic/third_party/ietf/ietf-inet-types.yang
//  - ../api/yang/sonic/third_party/ietf/ietf-yang-types.yang
//  - ../api/yang/sonic/third_party/ietf/ietf-interfaces.yang
//  - ../api/yang/sonic/third_party/ietf/iana-if-type.yang
//  - ../api/yang/sonic/third_party/openconfig/openconfig-types.yang
//  - ../api/yang/sonic/third_party/openconfig/openconfig-extensions.yang
//  - ../api/yang/sonic/nocsys-todo.yang

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: accton_loopback_interface.proto

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

type AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_Scope int32

const (
	AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_SCOPE_UNSET  AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_Scope = 0
	AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_SCOPE_global AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_Scope = 1
	AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_SCOPE_local  AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_Scope = 2
)

// Enum value maps for AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_Scope.
var (
	AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_Scope_name = map[int32]string{
		0: "SCOPE_UNSET",
		1: "SCOPE_global",
		2: "SCOPE_local",
	}
	AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_Scope_value = map[string]int32{
		"SCOPE_UNSET":  0,
		"SCOPE_global": 1,
		"SCOPE_local":  2,
	}
)

func (x AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_Scope) Enum() *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_Scope {
	p := new(AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_Scope)
	*p = x
	return p
}

func (x AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_Scope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_Scope) Descriptor() protoreflect.EnumDescriptor {
	return file_accton_loopback_interface_proto_enumTypes[0].Descriptor()
}

func (AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_Scope) Type() protoreflect.EnumType {
	return &file_accton_loopback_interface_proto_enumTypes[0]
}

func (x AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_Scope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_Scope.Descriptor instead.
func (AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_Scope) EnumDescriptor() ([]byte, []int) {
	return file_accton_loopback_interface_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

type AcctonLoopbackInterface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoopbackInterface *AcctonLoopbackInterface_LoopbackInterface `protobuf:"bytes,1,opt,name=loopback_interface,json=loopbackInterface,proto3" json:"loopback_interface,omitempty"`
}

func (x *AcctonLoopbackInterface) Reset() {
	*x = AcctonLoopbackInterface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accton_loopback_interface_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcctonLoopbackInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcctonLoopbackInterface) ProtoMessage() {}

func (x *AcctonLoopbackInterface) ProtoReflect() protoreflect.Message {
	mi := &file_accton_loopback_interface_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcctonLoopbackInterface.ProtoReflect.Descriptor instead.
func (*AcctonLoopbackInterface) Descriptor() ([]byte, []int) {
	return file_accton_loopback_interface_proto_rawDescGZIP(), []int{0}
}

func (x *AcctonLoopbackInterface) GetLoopbackInterface() *AcctonLoopbackInterface_LoopbackInterface {
	if x != nil {
		return x.LoopbackInterface
	}
	return nil
}

type AcctonLoopbackInterface_LoopbackInterface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoopbackInterfaceIpprefixList []*AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixListKey `protobuf:"bytes,1,rep,name=loopback_interface_ipprefix_list,json=loopbackInterfaceIpprefixList,proto3" json:"loopback_interface_ipprefix_list,omitempty"`
	LoopbackInterfaceList         []*AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceListKey         `protobuf:"bytes,2,rep,name=loopback_interface_list,json=loopbackInterfaceList,proto3" json:"loopback_interface_list,omitempty"`
}

func (x *AcctonLoopbackInterface_LoopbackInterface) Reset() {
	*x = AcctonLoopbackInterface_LoopbackInterface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accton_loopback_interface_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcctonLoopbackInterface_LoopbackInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcctonLoopbackInterface_LoopbackInterface) ProtoMessage() {}

func (x *AcctonLoopbackInterface_LoopbackInterface) ProtoReflect() protoreflect.Message {
	mi := &file_accton_loopback_interface_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcctonLoopbackInterface_LoopbackInterface.ProtoReflect.Descriptor instead.
func (*AcctonLoopbackInterface_LoopbackInterface) Descriptor() ([]byte, []int) {
	return file_accton_loopback_interface_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AcctonLoopbackInterface_LoopbackInterface) GetLoopbackInterfaceIpprefixList() []*AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixListKey {
	if x != nil {
		return x.LoopbackInterfaceIpprefixList
	}
	return nil
}

func (x *AcctonLoopbackInterface_LoopbackInterface) GetLoopbackInterfaceList() []*AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceListKey {
	if x != nil {
		return x.LoopbackInterfaceList
	}
	return nil
}

type AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Family AcctonTypesIpFamily                                                           `protobuf:"varint,1,opt,name=family,proto3,enum=sonic.AcctonTypesIpFamily" json:"family,omitempty"`
	Scope  AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_Scope `protobuf:"varint,2,opt,name=scope,proto3,enum=sonic.AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_Scope" json:"scope,omitempty"`
}

func (x *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList) Reset() {
	*x = AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accton_loopback_interface_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList) ProtoMessage() {}

func (x *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList) ProtoReflect() protoreflect.Message {
	mi := &file_accton_loopback_interface_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList.ProtoReflect.Descriptor instead.
func (*AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList) Descriptor() ([]byte, []int) {
	return file_accton_loopback_interface_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList) GetFamily() AcctonTypesIpFamily {
	if x != nil {
		return x.Family
	}
	return AcctonTypesIpFamily_ACCTONTYPESIPFAMILY_UNSET
}

func (x *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList) GetScope() AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_Scope {
	if x != nil {
		return x.Scope
	}
	return AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_SCOPE_UNSET
}

type AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixListKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoopbackInterfaceName         string                                                                   `protobuf:"bytes,1,opt,name=loopback_interface_name,json=loopbackInterfaceName,proto3" json:"loopback_interface_name,omitempty"`
	IpPrefix                      string                                                                   `protobuf:"bytes,2,opt,name=ip_prefix,json=ipPrefix,proto3" json:"ip_prefix,omitempty"`
	LoopbackInterfaceIpprefixList *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList `protobuf:"bytes,3,opt,name=loopback_interface_ipprefix_list,json=loopbackInterfaceIpprefixList,proto3" json:"loopback_interface_ipprefix_list,omitempty"`
}

func (x *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixListKey) Reset() {
	*x = AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixListKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accton_loopback_interface_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixListKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixListKey) ProtoMessage() {}

func (x *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixListKey) ProtoReflect() protoreflect.Message {
	mi := &file_accton_loopback_interface_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixListKey.ProtoReflect.Descriptor instead.
func (*AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixListKey) Descriptor() ([]byte, []int) {
	return file_accton_loopback_interface_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixListKey) GetLoopbackInterfaceName() string {
	if x != nil {
		return x.LoopbackInterfaceName
	}
	return ""
}

func (x *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixListKey) GetIpPrefix() string {
	if x != nil {
		return x.IpPrefix
	}
	return ""
}

func (x *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixListKey) GetLoopbackInterfaceIpprefixList() *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList {
	if x != nil {
		return x.LoopbackInterfaceIpprefixList
	}
	return nil
}

type AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VrfName *ywrapper.StringValue `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
}

func (x *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceList) Reset() {
	*x = AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accton_loopback_interface_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceList) ProtoMessage() {}

func (x *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceList) ProtoReflect() protoreflect.Message {
	mi := &file_accton_loopback_interface_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceList.ProtoReflect.Descriptor instead.
func (*AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceList) Descriptor() ([]byte, []int) {
	return file_accton_loopback_interface_proto_rawDescGZIP(), []int{0, 0, 2}
}

func (x *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceList) GetVrfName() *ywrapper.StringValue {
	if x != nil {
		return x.VrfName
	}
	return nil
}

type AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceListKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoopbackInterfaceName string                                                           `protobuf:"bytes,1,opt,name=loopback_interface_name,json=loopbackInterfaceName,proto3" json:"loopback_interface_name,omitempty"`
	LoopbackInterfaceList *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceList `protobuf:"bytes,2,opt,name=loopback_interface_list,json=loopbackInterfaceList,proto3" json:"loopback_interface_list,omitempty"`
}

func (x *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceListKey) Reset() {
	*x = AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceListKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accton_loopback_interface_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceListKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceListKey) ProtoMessage() {}

func (x *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceListKey) ProtoReflect() protoreflect.Message {
	mi := &file_accton_loopback_interface_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceListKey.ProtoReflect.Descriptor instead.
func (*AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceListKey) Descriptor() ([]byte, []int) {
	return file_accton_loopback_interface_proto_rawDescGZIP(), []int{0, 0, 3}
}

func (x *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceListKey) GetLoopbackInterfaceName() string {
	if x != nil {
		return x.LoopbackInterfaceName
	}
	return ""
}

func (x *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceListKey) GetLoopbackInterfaceList() *AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceList {
	if x != nil {
		return x.LoopbackInterfaceList
	}
	return nil
}

var File_accton_loopback_interface_proto protoreflect.FileDescriptor

var file_accton_loopback_interface_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x63, 0x63, 0x74, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63,
	0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x1a, 0x0e, 0x79, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x79, 0x65, 0x78, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xff, 0x0f, 0x0a, 0x17, 0x41, 0x63, 0x63, 0x74, 0x6f, 0x6e, 0x4c, 0x6f, 0x6f, 0x70,
	0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x91, 0x01,
	0x0a, 0x12, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x73, 0x6f, 0x6e,
	0x69, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x74, 0x6f, 0x6e, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63,
	0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x70, 0x62,
	0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x42, 0x30, 0x82, 0x41,
	0x2d, 0x2f, 0x61, 0x63, 0x63, 0x74, 0x6f, 0x6e, 0x2d, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63,
	0x6b, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x6f, 0x70,
	0x62, 0x61, 0x63, 0x6b, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x11,
	0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x1a, 0xcf, 0x0e, 0x0a, 0x11, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0xed, 0x01, 0x0a, 0x20, 0x6c, 0x6f, 0x6f, 0x70,
	0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x69,
	0x70, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x51, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x74, 0x6f,
	0x6e, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x70, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x4c, 0x69,
	0x73, 0x74, 0x4b, 0x65, 0x79, 0x42, 0x51, 0x82, 0x41, 0x4e, 0x2f, 0x61, 0x63, 0x63, 0x74, 0x6f,
	0x6e, 0x2d, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b,
	0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2d, 0x69, 0x70, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x1d, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61,
	0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x70, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x12, 0xcb, 0x01, 0x0a, 0x17, 0x6c, 0x6f, 0x6f, 0x70,
	0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x73, 0x6f, 0x6e, 0x69,
	0x63, 0x2e, 0x41, 0x63, 0x63, 0x74, 0x6f, 0x6e, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61,
	0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x70,
	0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x4b, 0x65, 0x79, 0x42, 0x48, 0x82, 0x41, 0x45, 0x2f, 0x61, 0x63, 0x63, 0x74, 0x6f, 0x6e,
	0x2d, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x2d,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x15,
	0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0xc6, 0x03, 0x0a, 0x1d, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61,
	0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x70, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x8c, 0x01, 0x0a, 0x06, 0x66, 0x61, 0x6d, 0x69,
	0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63,
	0x2e, 0x41, 0x63, 0x63, 0x74, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x49, 0x70, 0x46, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x42, 0x58, 0x82, 0x41, 0x55, 0x2f, 0x61, 0x63, 0x63, 0x74, 0x6f, 0x6e,
	0x2d, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x2d,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2d, 0x69, 0x70, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x06,
	0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0xc3, 0x01, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x54, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x41,
	0x63, 0x63, 0x74, 0x6f, 0x6e, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63,
	0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x70, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x42, 0x57, 0x82, 0x41,
	0x54, 0x2f, 0x61, 0x63, 0x63, 0x74, 0x6f, 0x6e, 0x2d, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63,
	0x6b, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x6f, 0x70,
	0x62, 0x61, 0x63, 0x6b, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x6c,
	0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x2d, 0x69, 0x70, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x22, 0x50, 0x0a, 0x05,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x0c, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x10, 0x01, 0x1a, 0x09, 0x82, 0x41, 0x06, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x0b, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x10, 0x02, 0x1a, 0x08, 0x82, 0x41, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x1a, 0xda,
	0x03, 0x0a, 0x20, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x49, 0x70, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x4c, 0x69, 0x73, 0x74,
	0x4b, 0x65, 0x79, 0x12, 0xa1, 0x01, 0x0a, 0x17, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x69, 0x82, 0x41, 0x66, 0x2f, 0x61, 0x63, 0x63, 0x74, 0x6f,
	0x6e, 0x2d, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b,
	0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2d, 0x69, 0x70, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63,
	0x6b, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2d, 0x6e, 0x61, 0x6d, 0x65,
	0x52, 0x15, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x78, 0x0a, 0x09, 0x69, 0x70, 0x5f, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x5b, 0x82, 0x41, 0x58, 0x2f,
	0x61, 0x63, 0x63, 0x74, 0x6f, 0x6e, 0x2d, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x2d,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61,
	0x63, 0x6b, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x6f,
	0x70, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2d,
	0x69, 0x70, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x69, 0x70,
	0x2d, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x08, 0x69, 0x70, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x12, 0x97, 0x01, 0x0a, 0x20, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x70, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x73,
	0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x74, 0x6f, 0x6e, 0x4c, 0x6f, 0x6f, 0x70, 0x62,
	0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x6f,
	0x70, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x4c,
	0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x49, 0x70, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x1d, 0x6c, 0x6f,
	0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49,
	0x70, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x9d, 0x01, 0x0a, 0x15,
	0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x83, 0x01, 0x0a, 0x08, 0x76, 0x72, 0x66, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x79, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x51, 0x82, 0x41, 0x4e, 0x2f, 0x61, 0x63, 0x63, 0x74, 0x6f, 0x6e, 0x2d, 0x6c, 0x6f, 0x6f, 0x70,
	0x62, 0x61, 0x63, 0x6b, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x6c,
	0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x76, 0x72, 0x66, 0x2d, 0x6e, 0x61,
	0x6d, 0x65, 0x52, 0x07, 0x76, 0x72, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0xb5, 0x02, 0x0a, 0x18,
	0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x98, 0x01, 0x0a, 0x17, 0x6c, 0x6f, 0x6f,
	0x70, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x60, 0x82, 0x41, 0x5d, 0x2f,
	0x61, 0x63, 0x63, 0x74, 0x6f, 0x6e, 0x2d, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x2d,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61,
	0x63, 0x6b, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x6f,
	0x70, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2d,
	0x6c, 0x69, 0x73, 0x74, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x15, 0x6c, 0x6f,
	0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x7e, 0x0a, 0x17, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x41, 0x63, 0x63,
	0x74, 0x6f, 0x6e, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x15, 0x6c, 0x6f,
	0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x1b, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x63, 0x63, 0x74, 0x6f,
	0x6e, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x5a, 0x07, 0x2e, 0x3b, 0x73, 0x6f, 0x6e, 0x69, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accton_loopback_interface_proto_rawDescOnce sync.Once
	file_accton_loopback_interface_proto_rawDescData = file_accton_loopback_interface_proto_rawDesc
)

func file_accton_loopback_interface_proto_rawDescGZIP() []byte {
	file_accton_loopback_interface_proto_rawDescOnce.Do(func() {
		file_accton_loopback_interface_proto_rawDescData = protoimpl.X.CompressGZIP(file_accton_loopback_interface_proto_rawDescData)
	})
	return file_accton_loopback_interface_proto_rawDescData
}

var file_accton_loopback_interface_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_accton_loopback_interface_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_accton_loopback_interface_proto_goTypes = []interface{}{
	(AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList_Scope)(0), // 0: sonic.AcctonLoopbackInterface.LoopbackInterface.LoopbackInterfaceIpprefixList.Scope
	(*AcctonLoopbackInterface)(nil),                                                    // 1: sonic.AcctonLoopbackInterface
	(*AcctonLoopbackInterface_LoopbackInterface)(nil),                                  // 2: sonic.AcctonLoopbackInterface.LoopbackInterface
	(*AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList)(nil),    // 3: sonic.AcctonLoopbackInterface.LoopbackInterface.LoopbackInterfaceIpprefixList
	(*AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixListKey)(nil), // 4: sonic.AcctonLoopbackInterface.LoopbackInterface.LoopbackInterfaceIpprefixListKey
	(*AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceList)(nil),            // 5: sonic.AcctonLoopbackInterface.LoopbackInterface.LoopbackInterfaceList
	(*AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceListKey)(nil),         // 6: sonic.AcctonLoopbackInterface.LoopbackInterface.LoopbackInterfaceListKey
	(AcctonTypesIpFamily)(0),     // 7: sonic.AcctonTypesIpFamily
	(*ywrapper.StringValue)(nil), // 8: ywrapper.StringValue
}
var file_accton_loopback_interface_proto_depIdxs = []int32{
	2, // 0: sonic.AcctonLoopbackInterface.loopback_interface:type_name -> sonic.AcctonLoopbackInterface.LoopbackInterface
	4, // 1: sonic.AcctonLoopbackInterface.LoopbackInterface.loopback_interface_ipprefix_list:type_name -> sonic.AcctonLoopbackInterface.LoopbackInterface.LoopbackInterfaceIpprefixListKey
	6, // 2: sonic.AcctonLoopbackInterface.LoopbackInterface.loopback_interface_list:type_name -> sonic.AcctonLoopbackInterface.LoopbackInterface.LoopbackInterfaceListKey
	7, // 3: sonic.AcctonLoopbackInterface.LoopbackInterface.LoopbackInterfaceIpprefixList.family:type_name -> sonic.AcctonTypesIpFamily
	0, // 4: sonic.AcctonLoopbackInterface.LoopbackInterface.LoopbackInterfaceIpprefixList.scope:type_name -> sonic.AcctonLoopbackInterface.LoopbackInterface.LoopbackInterfaceIpprefixList.Scope
	3, // 5: sonic.AcctonLoopbackInterface.LoopbackInterface.LoopbackInterfaceIpprefixListKey.loopback_interface_ipprefix_list:type_name -> sonic.AcctonLoopbackInterface.LoopbackInterface.LoopbackInterfaceIpprefixList
	8, // 6: sonic.AcctonLoopbackInterface.LoopbackInterface.LoopbackInterfaceList.vrf_name:type_name -> ywrapper.StringValue
	5, // 7: sonic.AcctonLoopbackInterface.LoopbackInterface.LoopbackInterfaceListKey.loopback_interface_list:type_name -> sonic.AcctonLoopbackInterface.LoopbackInterface.LoopbackInterfaceList
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_accton_loopback_interface_proto_init() }
func file_accton_loopback_interface_proto_init() {
	if File_accton_loopback_interface_proto != nil {
		return
	}
	file_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_accton_loopback_interface_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcctonLoopbackInterface); i {
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
		file_accton_loopback_interface_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcctonLoopbackInterface_LoopbackInterface); i {
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
		file_accton_loopback_interface_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixList); i {
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
		file_accton_loopback_interface_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceIpprefixListKey); i {
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
		file_accton_loopback_interface_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceList); i {
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
		file_accton_loopback_interface_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcctonLoopbackInterface_LoopbackInterface_LoopbackInterfaceListKey); i {
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
			RawDescriptor: file_accton_loopback_interface_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_accton_loopback_interface_proto_goTypes,
		DependencyIndexes: file_accton_loopback_interface_proto_depIdxs,
		EnumInfos:         file_accton_loopback_interface_proto_enumTypes,
		MessageInfos:      file_accton_loopback_interface_proto_msgTypes,
	}.Build()
	File_accton_loopback_interface_proto = out.File
	file_accton_loopback_interface_proto_rawDesc = nil
	file_accton_loopback_interface_proto_goTypes = nil
	file_accton_loopback_interface_proto_depIdxs = nil
}