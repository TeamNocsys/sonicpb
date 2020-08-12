// openconfig.sonic_portchannel is generated by proto_generator as a protobuf
// representation of a YANG schema.
//
// Input schema modules:
//  - ../api/yang/sonic/sonic-acl.yang
//  - ../api/yang/sonic/sonic-port.yang
//  - ../api/yang/sonic/sonic-types.yang
//  - ../api/yang/sonic/sonic-portchannel.yang
//  - ../api/yang/sonic/sonic-lldp-types.yang
//  - ../api/yang/sonic/third_party/ietf/iana-if-type.yang
//  - ../api/yang/sonic/third_party/ietf/ietf-inet-types.yang
//  - ../api/yang/sonic/third_party/ietf/ietf-interfaces.yang
//  - ../api/yang/sonic/third_party/ietf/ietf-yang-types.yang
//  - ../api/yang/sonic/sonic-vlan.yang
//  - ../api/yang/sonic/sonic-lldp.yang
//  - ../api/yang/sonic/sonic-interface.yang
//  - ../api/yang/sonic/sonic-loopback-interface.yang
//  - ../api/yang/sonic/sonic-platform.yang
//  - ../api/yang/sonic/sonic-extension.yang

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: sonic_portchannel.proto

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

type SonicPortchannel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Portchannel *SonicPortchannel_Portchannel `protobuf:"bytes,36582390,opt,name=portchannel,proto3" json:"portchannel,omitempty"`
}

func (x *SonicPortchannel) Reset() {
	*x = SonicPortchannel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_portchannel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicPortchannel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicPortchannel) ProtoMessage() {}

func (x *SonicPortchannel) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_portchannel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicPortchannel.ProtoReflect.Descriptor instead.
func (*SonicPortchannel) Descriptor() ([]byte, []int) {
	return file_sonic_portchannel_proto_rawDescGZIP(), []int{0}
}

func (x *SonicPortchannel) GetPortchannel() *SonicPortchannel_Portchannel {
	if x != nil {
		return x.Portchannel
	}
	return nil
}

type SonicPortchannel_Portchannel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PortchannelList []*SonicPortchannel_Portchannel_PortchannelListKey `protobuf:"bytes,118674730,rep,name=portchannel_list,json=portchannelList,proto3" json:"portchannel_list,omitempty"`
}

func (x *SonicPortchannel_Portchannel) Reset() {
	*x = SonicPortchannel_Portchannel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_portchannel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicPortchannel_Portchannel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicPortchannel_Portchannel) ProtoMessage() {}

func (x *SonicPortchannel_Portchannel) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_portchannel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicPortchannel_Portchannel.ProtoReflect.Descriptor instead.
func (*SonicPortchannel_Portchannel) Descriptor() ([]byte, []int) {
	return file_sonic_portchannel_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SonicPortchannel_Portchannel) GetPortchannelList() []*SonicPortchannel_Portchannel_PortchannelListKey {
	if x != nil {
		return x.PortchannelList
	}
	return nil
}

type SonicPortchannel_Portchannel_PortchannelList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminStatus SonicPortchannelAdminStatus `protobuf:"varint,289246861,opt,name=admin_status,json=adminStatus,proto3,enum=sonic.SonicPortchannelAdminStatus" json:"admin_status,omitempty"`
	Description *ywrapper.StringValue       `protobuf:"bytes,489610797,opt,name=description,proto3" json:"description,omitempty"`
	Members     []*ywrapper.StringValue     `protobuf:"bytes,390606290,rep,name=members,proto3" json:"members,omitempty"`
	MinLinks    *ywrapper.UintValue         `protobuf:"bytes,106858059,opt,name=min_links,json=minLinks,proto3" json:"min_links,omitempty"`
	Mtu         *ywrapper.UintValue         `protobuf:"bytes,527911279,opt,name=mtu,proto3" json:"mtu,omitempty"`
}

func (x *SonicPortchannel_Portchannel_PortchannelList) Reset() {
	*x = SonicPortchannel_Portchannel_PortchannelList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_portchannel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicPortchannel_Portchannel_PortchannelList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicPortchannel_Portchannel_PortchannelList) ProtoMessage() {}

func (x *SonicPortchannel_Portchannel_PortchannelList) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_portchannel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicPortchannel_Portchannel_PortchannelList.ProtoReflect.Descriptor instead.
func (*SonicPortchannel_Portchannel_PortchannelList) Descriptor() ([]byte, []int) {
	return file_sonic_portchannel_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *SonicPortchannel_Portchannel_PortchannelList) GetAdminStatus() SonicPortchannelAdminStatus {
	if x != nil {
		return x.AdminStatus
	}
	return SonicPortchannelAdminStatus_SONICPORTCHANNELADMINSTATUS_UNSET
}

func (x *SonicPortchannel_Portchannel_PortchannelList) GetDescription() *ywrapper.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *SonicPortchannel_Portchannel_PortchannelList) GetMembers() []*ywrapper.StringValue {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *SonicPortchannel_Portchannel_PortchannelList) GetMinLinks() *ywrapper.UintValue {
	if x != nil {
		return x.MinLinks
	}
	return nil
}

func (x *SonicPortchannel_Portchannel_PortchannelList) GetMtu() *ywrapper.UintValue {
	if x != nil {
		return x.Mtu
	}
	return nil
}

type SonicPortchannel_Portchannel_PortchannelListKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PortchannelName string                                        `protobuf:"bytes,1,opt,name=portchannel_name,json=portchannelName,proto3" json:"portchannel_name,omitempty"`
	PortchannelList *SonicPortchannel_Portchannel_PortchannelList `protobuf:"bytes,2,opt,name=portchannel_list,json=portchannelList,proto3" json:"portchannel_list,omitempty"`
}

func (x *SonicPortchannel_Portchannel_PortchannelListKey) Reset() {
	*x = SonicPortchannel_Portchannel_PortchannelListKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_portchannel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicPortchannel_Portchannel_PortchannelListKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicPortchannel_Portchannel_PortchannelListKey) ProtoMessage() {}

func (x *SonicPortchannel_Portchannel_PortchannelListKey) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_portchannel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicPortchannel_Portchannel_PortchannelListKey.ProtoReflect.Descriptor instead.
func (*SonicPortchannel_Portchannel_PortchannelListKey) Descriptor() ([]byte, []int) {
	return file_sonic_portchannel_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *SonicPortchannel_Portchannel_PortchannelListKey) GetPortchannelName() string {
	if x != nil {
		return x.PortchannelName
	}
	return ""
}

func (x *SonicPortchannel_Portchannel_PortchannelListKey) GetPortchannelList() *SonicPortchannel_Portchannel_PortchannelList {
	if x != nil {
		return x.PortchannelList
	}
	return nil
}

var File_sonic_portchannel_proto protoreflect.FileDescriptor

var file_sonic_portchannel_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x6f, 0x6e, 0x69, 0x63,
	0x1a, 0x0e, 0x79, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0a, 0x79, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x08, 0x0a, 0x10, 0x53, 0x6f,
	0x6e, 0x69, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x6b,
	0x0a, 0x0b, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0xf6, 0xe7,
	0xb8, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x53,
	0x6f, 0x6e, 0x69, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e,
	0x50, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x42, 0x21, 0x82, 0x41, 0x1e,
	0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x0b,
	0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x1a, 0xf4, 0x07, 0x0a, 0x0b,
	0x50, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x98, 0x01, 0x0a, 0x10,
	0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0xaa, 0xaa, 0xcb, 0x38, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x73, 0x6f, 0x6e, 0x69,
	0x63, 0x2e, 0x53, 0x6f, 0x6e, 0x69, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x50,
	0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65,
	0x79, 0x42, 0x32, 0x82, 0x41, 0x2f, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x70, 0x6f, 0x72,
	0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2d, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x0f, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0xe2, 0x04, 0x0a, 0x0f, 0x50, 0x6f, 0x72, 0x74, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x8a, 0x01, 0x0a, 0x0c, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x8d, 0x9d, 0xf6, 0x89,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x53, 0x6f,
	0x6e, 0x69, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x3f, 0x82, 0x41, 0x3c, 0x2f, 0x73,
	0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2f, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x70, 0x6f, 0x72,
	0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x7b, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0xad, 0xbc, 0xbb, 0xe9, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x79, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x3e, 0x82, 0x41, 0x3b, 0x2f, 0x73, 0x6f, 0x6e,
	0x69, 0x63, 0x2d, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x70,
	0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6f, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18,
	0xd2, 0xdb, 0xa0, 0xba, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x79, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x3a, 0x82, 0x41, 0x37, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x70, 0x6f, 0x72, 0x74,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2d,
	0x6c, 0x69, 0x73, 0x74, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x07, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x71, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x69, 0x6e,
	0x6b, 0x73, 0x18, 0xcb, 0x8c, 0xfa, 0x32, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x79, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x3c, 0x82, 0x41, 0x39, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x70, 0x6f, 0x72, 0x74,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2d,
	0x6c, 0x69, 0x73, 0x74, 0x2f, 0x6d, 0x69, 0x6e, 0x2d, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x08,
	0x6d, 0x69, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x61, 0x0a, 0x03, 0x6d, 0x74, 0x75, 0x18,
	0xef, 0x92, 0xdd, 0xfb, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x79, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x36,
	0x82, 0x41, 0x33, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2d, 0x6c, 0x69,
	0x73, 0x74, 0x2f, 0x6d, 0x74, 0x75, 0x52, 0x03, 0x6d, 0x74, 0x75, 0x1a, 0xe4, 0x01, 0x0a, 0x12,
	0x50, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x4b,
	0x65, 0x79, 0x12, 0x6e, 0x0a, 0x10, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x43, 0x82, 0x41,
	0x40, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f,
	0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2d, 0x6c, 0x69, 0x73, 0x74,
	0x2f, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2d, 0x6e, 0x61, 0x6d,
	0x65, 0x52, 0x0f, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x5e, 0x0a, 0x10, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x73,
	0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x53, 0x6f, 0x6e, 0x69, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x0f, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x12, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x6f, 0x63, 0x73, 0x79, 0x73,
	0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sonic_portchannel_proto_rawDescOnce sync.Once
	file_sonic_portchannel_proto_rawDescData = file_sonic_portchannel_proto_rawDesc
)

func file_sonic_portchannel_proto_rawDescGZIP() []byte {
	file_sonic_portchannel_proto_rawDescOnce.Do(func() {
		file_sonic_portchannel_proto_rawDescData = protoimpl.X.CompressGZIP(file_sonic_portchannel_proto_rawDescData)
	})
	return file_sonic_portchannel_proto_rawDescData
}

var file_sonic_portchannel_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sonic_portchannel_proto_goTypes = []interface{}{
	(*SonicPortchannel)(nil),                                // 0: sonic.SonicPortchannel
	(*SonicPortchannel_Portchannel)(nil),                    // 1: sonic.SonicPortchannel.Portchannel
	(*SonicPortchannel_Portchannel_PortchannelList)(nil),    // 2: sonic.SonicPortchannel.Portchannel.PortchannelList
	(*SonicPortchannel_Portchannel_PortchannelListKey)(nil), // 3: sonic.SonicPortchannel.Portchannel.PortchannelListKey
	(SonicPortchannelAdminStatus)(0),                        // 4: sonic.SonicPortchannelAdminStatus
	(*ywrapper.StringValue)(nil),                            // 5: ywrapper.StringValue
	(*ywrapper.UintValue)(nil),                              // 6: ywrapper.UintValue
}
var file_sonic_portchannel_proto_depIdxs = []int32{
	1, // 0: sonic.SonicPortchannel.portchannel:type_name -> sonic.SonicPortchannel.Portchannel
	3, // 1: sonic.SonicPortchannel.Portchannel.portchannel_list:type_name -> sonic.SonicPortchannel.Portchannel.PortchannelListKey
	4, // 2: sonic.SonicPortchannel.Portchannel.PortchannelList.admin_status:type_name -> sonic.SonicPortchannelAdminStatus
	5, // 3: sonic.SonicPortchannel.Portchannel.PortchannelList.description:type_name -> ywrapper.StringValue
	5, // 4: sonic.SonicPortchannel.Portchannel.PortchannelList.members:type_name -> ywrapper.StringValue
	6, // 5: sonic.SonicPortchannel.Portchannel.PortchannelList.min_links:type_name -> ywrapper.UintValue
	6, // 6: sonic.SonicPortchannel.Portchannel.PortchannelList.mtu:type_name -> ywrapper.UintValue
	2, // 7: sonic.SonicPortchannel.Portchannel.PortchannelListKey.portchannel_list:type_name -> sonic.SonicPortchannel.Portchannel.PortchannelList
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_sonic_portchannel_proto_init() }
func file_sonic_portchannel_proto_init() {
	if File_sonic_portchannel_proto != nil {
		return
	}
	file_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sonic_portchannel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicPortchannel); i {
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
		file_sonic_portchannel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicPortchannel_Portchannel); i {
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
		file_sonic_portchannel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicPortchannel_Portchannel_PortchannelList); i {
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
		file_sonic_portchannel_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicPortchannel_Portchannel_PortchannelListKey); i {
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
			RawDescriptor: file_sonic_portchannel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sonic_portchannel_proto_goTypes,
		DependencyIndexes: file_sonic_portchannel_proto_depIdxs,
		MessageInfos:      file_sonic_portchannel_proto_msgTypes,
	}.Build()
	File_sonic_portchannel_proto = out.File
	file_sonic_portchannel_proto_rawDesc = nil
	file_sonic_portchannel_proto_goTypes = nil
	file_sonic_portchannel_proto_depIdxs = nil
}