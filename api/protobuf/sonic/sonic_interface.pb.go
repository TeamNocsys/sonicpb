// openconfig.sonic_interface is generated by proto_generator as a protobuf
// representation of a YANG schema.
//
// Input schema modules:
//  - ../api/yang/sonic/sonic-interface.yang
//  - ../api/yang/sonic/sonic-loopback-interface.yang
//  - ../api/yang/sonic/sonic-port.yang
//  - ../api/yang/sonic/sonic-vlan.yang
//  - ../api/yang/sonic/sonic-extension.yang
//  - ../api/yang/sonic/sonic-acl.yang
//  - ../api/yang/sonic/sonic-types.yang
//  - ../api/yang/sonic/third_party/ietf/ietf-inet-types.yang
//  - ../api/yang/sonic/third_party/ietf/ietf-yang-types.yang
//  - ../api/yang/sonic/third_party/ietf/iana-if-type.yang
//  - ../api/yang/sonic/third_party/ietf/ietf-interfaces.yang
//  - ../api/yang/sonic/sonic-portchannel.yang
//  - ../api/yang/sonic/sonic-lldp-types.yang
//  - ../api/yang/sonic/sonic-lldp.yang
//  - ../api/yang/sonic/sonic-platform.yang

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.0
// source: sonic_interface.proto

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

type SonicInterface_Interface_InterfaceIpprefixList_Scope int32

const (
	SonicInterface_Interface_InterfaceIpprefixList_SCOPE_UNSET  SonicInterface_Interface_InterfaceIpprefixList_Scope = 0
	SonicInterface_Interface_InterfaceIpprefixList_SCOPE_global SonicInterface_Interface_InterfaceIpprefixList_Scope = 1
	SonicInterface_Interface_InterfaceIpprefixList_SCOPE_local  SonicInterface_Interface_InterfaceIpprefixList_Scope = 2
)

// Enum value maps for SonicInterface_Interface_InterfaceIpprefixList_Scope.
var (
	SonicInterface_Interface_InterfaceIpprefixList_Scope_name = map[int32]string{
		0: "SCOPE_UNSET",
		1: "SCOPE_global",
		2: "SCOPE_local",
	}
	SonicInterface_Interface_InterfaceIpprefixList_Scope_value = map[string]int32{
		"SCOPE_UNSET":  0,
		"SCOPE_global": 1,
		"SCOPE_local":  2,
	}
)

func (x SonicInterface_Interface_InterfaceIpprefixList_Scope) Enum() *SonicInterface_Interface_InterfaceIpprefixList_Scope {
	p := new(SonicInterface_Interface_InterfaceIpprefixList_Scope)
	*p = x
	return p
}

func (x SonicInterface_Interface_InterfaceIpprefixList_Scope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SonicInterface_Interface_InterfaceIpprefixList_Scope) Descriptor() protoreflect.EnumDescriptor {
	return file_sonic_interface_proto_enumTypes[0].Descriptor()
}

func (SonicInterface_Interface_InterfaceIpprefixList_Scope) Type() protoreflect.EnumType {
	return &file_sonic_interface_proto_enumTypes[0]
}

func (x SonicInterface_Interface_InterfaceIpprefixList_Scope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SonicInterface_Interface_InterfaceIpprefixList_Scope.Descriptor instead.
func (SonicInterface_Interface_InterfaceIpprefixList_Scope) EnumDescriptor() ([]byte, []int) {
	return file_sonic_interface_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

type SonicInterface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interface *SonicInterface_Interface `protobuf:"bytes,87274325,opt,name=interface,proto3" json:"interface,omitempty"`
}

func (x *SonicInterface) Reset() {
	*x = SonicInterface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_interface_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicInterface) ProtoMessage() {}

func (x *SonicInterface) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_interface_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicInterface.ProtoReflect.Descriptor instead.
func (*SonicInterface) Descriptor() ([]byte, []int) {
	return file_sonic_interface_proto_rawDescGZIP(), []int{0}
}

func (x *SonicInterface) GetInterface() *SonicInterface_Interface {
	if x != nil {
		return x.Interface
	}
	return nil
}

type SonicInterface_Interface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InterfaceIpprefixList []*SonicInterface_Interface_InterfaceIpprefixListKey `protobuf:"bytes,163932438,rep,name=interface_ipprefix_list,json=interfaceIpprefixList,proto3" json:"interface_ipprefix_list,omitempty"`
	InterfaceList         []*SonicInterface_Interface_InterfaceListKey         `protobuf:"bytes,516396642,rep,name=interface_list,json=interfaceList,proto3" json:"interface_list,omitempty"`
}

func (x *SonicInterface_Interface) Reset() {
	*x = SonicInterface_Interface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_interface_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicInterface_Interface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicInterface_Interface) ProtoMessage() {}

func (x *SonicInterface_Interface) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_interface_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicInterface_Interface.ProtoReflect.Descriptor instead.
func (*SonicInterface_Interface) Descriptor() ([]byte, []int) {
	return file_sonic_interface_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SonicInterface_Interface) GetInterfaceIpprefixList() []*SonicInterface_Interface_InterfaceIpprefixListKey {
	if x != nil {
		return x.InterfaceIpprefixList
	}
	return nil
}

func (x *SonicInterface_Interface) GetInterfaceList() []*SonicInterface_Interface_InterfaceListKey {
	if x != nil {
		return x.InterfaceList
	}
	return nil
}

type SonicInterface_Interface_InterfaceIpprefixList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Family SonicInterfaceIpFamily                               `protobuf:"varint,157275605,opt,name=family,proto3,enum=sonic.SonicInterfaceIpFamily" json:"family,omitempty"`
	Scope  SonicInterface_Interface_InterfaceIpprefixList_Scope `protobuf:"varint,49312695,opt,name=scope,proto3,enum=sonic.SonicInterface_Interface_InterfaceIpprefixList_Scope" json:"scope,omitempty"`
}

func (x *SonicInterface_Interface_InterfaceIpprefixList) Reset() {
	*x = SonicInterface_Interface_InterfaceIpprefixList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_interface_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicInterface_Interface_InterfaceIpprefixList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicInterface_Interface_InterfaceIpprefixList) ProtoMessage() {}

func (x *SonicInterface_Interface_InterfaceIpprefixList) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_interface_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicInterface_Interface_InterfaceIpprefixList.ProtoReflect.Descriptor instead.
func (*SonicInterface_Interface_InterfaceIpprefixList) Descriptor() ([]byte, []int) {
	return file_sonic_interface_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *SonicInterface_Interface_InterfaceIpprefixList) GetFamily() SonicInterfaceIpFamily {
	if x != nil {
		return x.Family
	}
	return SonicInterfaceIpFamily_SONICINTERFACEIPFAMILY_UNSET
}

func (x *SonicInterface_Interface_InterfaceIpprefixList) GetScope() SonicInterface_Interface_InterfaceIpprefixList_Scope {
	if x != nil {
		return x.Scope
	}
	return SonicInterface_Interface_InterfaceIpprefixList_SCOPE_UNSET
}

type SonicInterface_Interface_InterfaceIpprefixListKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PortName              string                                          `protobuf:"bytes,1,opt,name=port_name,json=portName,proto3" json:"port_name,omitempty"`
	IpPrefix              string                                          `protobuf:"bytes,2,opt,name=ip_prefix,json=ipPrefix,proto3" json:"ip_prefix,omitempty"`
	InterfaceIpprefixList *SonicInterface_Interface_InterfaceIpprefixList `protobuf:"bytes,3,opt,name=interface_ipprefix_list,json=interfaceIpprefixList,proto3" json:"interface_ipprefix_list,omitempty"`
}

func (x *SonicInterface_Interface_InterfaceIpprefixListKey) Reset() {
	*x = SonicInterface_Interface_InterfaceIpprefixListKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_interface_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicInterface_Interface_InterfaceIpprefixListKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicInterface_Interface_InterfaceIpprefixListKey) ProtoMessage() {}

func (x *SonicInterface_Interface_InterfaceIpprefixListKey) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_interface_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicInterface_Interface_InterfaceIpprefixListKey.ProtoReflect.Descriptor instead.
func (*SonicInterface_Interface_InterfaceIpprefixListKey) Descriptor() ([]byte, []int) {
	return file_sonic_interface_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *SonicInterface_Interface_InterfaceIpprefixListKey) GetPortName() string {
	if x != nil {
		return x.PortName
	}
	return ""
}

func (x *SonicInterface_Interface_InterfaceIpprefixListKey) GetIpPrefix() string {
	if x != nil {
		return x.IpPrefix
	}
	return ""
}

func (x *SonicInterface_Interface_InterfaceIpprefixListKey) GetInterfaceIpprefixList() *SonicInterface_Interface_InterfaceIpprefixList {
	if x != nil {
		return x.InterfaceIpprefixList
	}
	return nil
}

type SonicInterface_Interface_InterfaceList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VrfName *ywrapper.StringValue `protobuf:"bytes,193384515,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
}

func (x *SonicInterface_Interface_InterfaceList) Reset() {
	*x = SonicInterface_Interface_InterfaceList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_interface_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicInterface_Interface_InterfaceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicInterface_Interface_InterfaceList) ProtoMessage() {}

func (x *SonicInterface_Interface_InterfaceList) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_interface_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicInterface_Interface_InterfaceList.ProtoReflect.Descriptor instead.
func (*SonicInterface_Interface_InterfaceList) Descriptor() ([]byte, []int) {
	return file_sonic_interface_proto_rawDescGZIP(), []int{0, 0, 2}
}

func (x *SonicInterface_Interface_InterfaceList) GetVrfName() *ywrapper.StringValue {
	if x != nil {
		return x.VrfName
	}
	return nil
}

type SonicInterface_Interface_InterfaceListKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PortName      string                                  `protobuf:"bytes,1,opt,name=port_name,json=portName,proto3" json:"port_name,omitempty"`
	InterfaceList *SonicInterface_Interface_InterfaceList `protobuf:"bytes,2,opt,name=interface_list,json=interfaceList,proto3" json:"interface_list,omitempty"`
}

func (x *SonicInterface_Interface_InterfaceListKey) Reset() {
	*x = SonicInterface_Interface_InterfaceListKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_interface_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicInterface_Interface_InterfaceListKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicInterface_Interface_InterfaceListKey) ProtoMessage() {}

func (x *SonicInterface_Interface_InterfaceListKey) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_interface_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicInterface_Interface_InterfaceListKey.ProtoReflect.Descriptor instead.
func (*SonicInterface_Interface_InterfaceListKey) Descriptor() ([]byte, []int) {
	return file_sonic_interface_proto_rawDescGZIP(), []int{0, 0, 3}
}

func (x *SonicInterface_Interface_InterfaceListKey) GetPortName() string {
	if x != nil {
		return x.PortName
	}
	return ""
}

func (x *SonicInterface_Interface_InterfaceListKey) GetInterfaceList() *SonicInterface_Interface_InterfaceList {
	if x != nil {
		return x.InterfaceList
	}
	return nil
}

var File_sonic_interface_proto protoreflect.FileDescriptor

var file_sonic_interface_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x1a, 0x0e,
	0x79, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a,
	0x79, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x0b, 0x0a, 0x0e, 0x53, 0x6f, 0x6e, 0x69,
	0x63, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x09, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0xd5, 0xe6, 0xce, 0x29, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x53, 0x6f, 0x6e, 0x69, 0x63, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x42, 0x1d, 0x82, 0x41, 0x1a, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x52, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x1a, 0xc1, 0x0a, 0x0a, 0x09,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0xaa, 0x01, 0x0a, 0x17, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x70, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x96, 0xd2, 0x95, 0x4e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38,
	0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x53, 0x6f, 0x6e, 0x69, 0x63, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x70, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x42, 0x35, 0x82, 0x41, 0x32, 0x2f, 0x73, 0x6f,
	0x6e, 0x69, 0x63, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x2d, 0x69, 0x70, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x52,
	0x15, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x70, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x89, 0x01, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0xe2, 0xac, 0x9e, 0xf6, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x53, 0x6f, 0x6e, 0x69,
	0x63, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x4b, 0x65, 0x79, 0x42, 0x2c, 0x82, 0x41, 0x29, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63,
	0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2d, 0x6c,
	0x69, 0x73, 0x74, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x1a, 0xf5, 0x02, 0x0a, 0x15, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x49, 0x70, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x76, 0x0a, 0x06,
	0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0xd5, 0xab, 0xff, 0x4a, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1d, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x53, 0x6f, 0x6e, 0x69, 0x63, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x70, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x42, 0x3c,
	0x82, 0x41, 0x39, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2d, 0x69, 0x70, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x06, 0x66, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x12, 0x91, 0x01, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0xb7,
	0xe7, 0xc1, 0x17, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e,
	0x53, 0x6f, 0x6e, 0x69, 0x63, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x49, 0x70, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x42, 0x3b, 0x82, 0x41, 0x38, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2d, 0x69, 0x70,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x22, 0x50, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x70,
	0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54,
	0x10, 0x00, 0x12, 0x1b, 0x0a, 0x0c, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x10, 0x01, 0x1a, 0x09, 0x82, 0x41, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x12,
	0x19, 0x0a, 0x0b, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x10, 0x02,
	0x1a, 0x08, 0x82, 0x41, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x1a, 0xc5, 0x02, 0x0a, 0x18, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x70, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x5c, 0x0a, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3f, 0x82, 0x41, 0x3c, 0x2f,
	0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x2d, 0x69, 0x70, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x2d, 0x6c, 0x69, 0x73,
	0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x08, 0x70, 0x6f, 0x72,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x5c, 0x0a, 0x09, 0x69, 0x70, 0x5f, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3f, 0x82, 0x41, 0x3c, 0x2f, 0x73, 0x6f,
	0x6e, 0x69, 0x63, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x2d, 0x69, 0x70, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f,
	0x69, 0x70, 0x2d, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x08, 0x69, 0x70, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x12, 0x6d, 0x0a, 0x17, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x5f, 0x69, 0x70, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x53, 0x6f, 0x6e,
	0x69, 0x63, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49,
	0x70, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x15, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x70, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x4c, 0x69,
	0x73, 0x74, 0x1a, 0x7b, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x6a, 0x0a, 0x08, 0x76, 0x72, 0x66, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0xc3, 0xa0, 0x9b, 0x5c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x79, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x35, 0x82, 0x41, 0x32, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x76, 0x72,
	0x66, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x07, 0x76, 0x72, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x1a,
	0xbd, 0x01, 0x0a, 0x10, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x4b, 0x65, 0x79, 0x12, 0x53, 0x0a, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x36, 0x82, 0x41, 0x33, 0x2f, 0x73, 0x6f, 0x6e,
	0x69, 0x63, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x52,
	0x08, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x54, 0x0a, 0x0e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x53, 0x6f, 0x6e, 0x69, 0x63, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x1b, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x6f, 0x63, 0x73, 0x79, 0x73, 0x2e, 0x73, 0x6f,
	0x6e, 0x69, 0x63, 0x5a, 0x07, 0x2e, 0x3b, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sonic_interface_proto_rawDescOnce sync.Once
	file_sonic_interface_proto_rawDescData = file_sonic_interface_proto_rawDesc
)

func file_sonic_interface_proto_rawDescGZIP() []byte {
	file_sonic_interface_proto_rawDescOnce.Do(func() {
		file_sonic_interface_proto_rawDescData = protoimpl.X.CompressGZIP(file_sonic_interface_proto_rawDescData)
	})
	return file_sonic_interface_proto_rawDescData
}

var file_sonic_interface_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sonic_interface_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sonic_interface_proto_goTypes = []interface{}{
	(SonicInterface_Interface_InterfaceIpprefixList_Scope)(0), // 0: sonic.SonicInterface.Interface.InterfaceIpprefixList.Scope
	(*SonicInterface)(nil),                                    // 1: sonic.SonicInterface
	(*SonicInterface_Interface)(nil),                          // 2: sonic.SonicInterface.Interface
	(*SonicInterface_Interface_InterfaceIpprefixList)(nil),    // 3: sonic.SonicInterface.Interface.InterfaceIpprefixList
	(*SonicInterface_Interface_InterfaceIpprefixListKey)(nil), // 4: sonic.SonicInterface.Interface.InterfaceIpprefixListKey
	(*SonicInterface_Interface_InterfaceList)(nil),            // 5: sonic.SonicInterface.Interface.InterfaceList
	(*SonicInterface_Interface_InterfaceListKey)(nil),         // 6: sonic.SonicInterface.Interface.InterfaceListKey
	(SonicInterfaceIpFamily)(0),                               // 7: sonic.SonicInterfaceIpFamily
	(*ywrapper.StringValue)(nil),                              // 8: ywrapper.StringValue
}
var file_sonic_interface_proto_depIdxs = []int32{
	2, // 0: sonic.SonicInterface.interface:type_name -> sonic.SonicInterface.Interface
	4, // 1: sonic.SonicInterface.Interface.interface_ipprefix_list:type_name -> sonic.SonicInterface.Interface.InterfaceIpprefixListKey
	6, // 2: sonic.SonicInterface.Interface.interface_list:type_name -> sonic.SonicInterface.Interface.InterfaceListKey
	7, // 3: sonic.SonicInterface.Interface.InterfaceIpprefixList.family:type_name -> sonic.SonicInterfaceIpFamily
	0, // 4: sonic.SonicInterface.Interface.InterfaceIpprefixList.scope:type_name -> sonic.SonicInterface.Interface.InterfaceIpprefixList.Scope
	3, // 5: sonic.SonicInterface.Interface.InterfaceIpprefixListKey.interface_ipprefix_list:type_name -> sonic.SonicInterface.Interface.InterfaceIpprefixList
	8, // 6: sonic.SonicInterface.Interface.InterfaceList.vrf_name:type_name -> ywrapper.StringValue
	5, // 7: sonic.SonicInterface.Interface.InterfaceListKey.interface_list:type_name -> sonic.SonicInterface.Interface.InterfaceList
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_sonic_interface_proto_init() }
func file_sonic_interface_proto_init() {
	if File_sonic_interface_proto != nil {
		return
	}
	file_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sonic_interface_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicInterface); i {
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
		file_sonic_interface_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicInterface_Interface); i {
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
		file_sonic_interface_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicInterface_Interface_InterfaceIpprefixList); i {
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
		file_sonic_interface_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicInterface_Interface_InterfaceIpprefixListKey); i {
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
		file_sonic_interface_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicInterface_Interface_InterfaceList); i {
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
		file_sonic_interface_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicInterface_Interface_InterfaceListKey); i {
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
			RawDescriptor: file_sonic_interface_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sonic_interface_proto_goTypes,
		DependencyIndexes: file_sonic_interface_proto_depIdxs,
		EnumInfos:         file_sonic_interface_proto_enumTypes,
		MessageInfos:      file_sonic_interface_proto_msgTypes,
	}.Build()
	File_sonic_interface_proto = out.File
	file_sonic_interface_proto_rawDesc = nil
	file_sonic_interface_proto_goTypes = nil
	file_sonic_interface_proto_depIdxs = nil
}
