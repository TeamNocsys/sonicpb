// openconfig.sonic_port is generated by proto_generator as a protobuf
// representation of a YANG schema.
//
// Input schema modules:
//  - ../api/yang/sonic/sonic-acl.yang
//  - ../api/yang/sonic/sonic-port.yang
//  - ../api/yang/sonic/sonic-types.yang
//  - ../api/yang/sonic/sonic-portchannel.yang
//  - ../api/yang/sonic/sonic-platform-types.yang
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
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.0
// source: sonic_port.proto

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

type SonicPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port *SonicPort_Port `protobuf:"bytes,491297659,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *SonicPort) Reset() {
	*x = SonicPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_port_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicPort) ProtoMessage() {}

func (x *SonicPort) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_port_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicPort.ProtoReflect.Descriptor instead.
func (*SonicPort) Descriptor() ([]byte, []int) {
	return file_sonic_port_proto_rawDescGZIP(), []int{0}
}

func (x *SonicPort) GetPort() *SonicPort_Port {
	if x != nil {
		return x.Port
	}
	return nil
}

type SonicPort_Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PortList []*SonicPort_Port_PortListKey `protobuf:"bytes,203138784,rep,name=port_list,json=portList,proto3" json:"port_list,omitempty"`
}

func (x *SonicPort_Port) Reset() {
	*x = SonicPort_Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_port_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicPort_Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicPort_Port) ProtoMessage() {}

func (x *SonicPort_Port) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_port_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicPort_Port.ProtoReflect.Descriptor instead.
func (*SonicPort_Port) Descriptor() ([]byte, []int) {
	return file_sonic_port_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SonicPort_Port) GetPortList() []*SonicPort_Port_PortListKey {
	if x != nil {
		return x.PortList
	}
	return nil
}

type SonicPort_Port_PortList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminStatus SonicPortAdminStatus  `protobuf:"varint,204119143,opt,name=admin_status,json=adminStatus,proto3,enum=sonic.SonicPortAdminStatus" json:"admin_status,omitempty"`
	Alias       *ywrapper.StringValue `protobuf:"bytes,451217765,opt,name=alias,proto3" json:"alias,omitempty"`
	Description *ywrapper.StringValue `protobuf:"bytes,372446763,opt,name=description,proto3" json:"description,omitempty"`
	Fec         *ywrapper.StringValue `protobuf:"bytes,352220911,opt,name=fec,proto3" json:"fec,omitempty"`
	Index       *ywrapper.UintValue   `protobuf:"bytes,280841649,opt,name=index,proto3" json:"index,omitempty"`
	Lanes       *ywrapper.StringValue `protobuf:"bytes,379225588,opt,name=lanes,proto3" json:"lanes,omitempty"`
	Mtu         *ywrapper.UintValue   `protobuf:"bytes,65189849,opt,name=mtu,proto3" json:"mtu,omitempty"`
	Speed       *ywrapper.UintValue   `protobuf:"bytes,112572536,opt,name=speed,proto3" json:"speed,omitempty"`
}

func (x *SonicPort_Port_PortList) Reset() {
	*x = SonicPort_Port_PortList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_port_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicPort_Port_PortList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicPort_Port_PortList) ProtoMessage() {}

func (x *SonicPort_Port_PortList) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_port_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicPort_Port_PortList.ProtoReflect.Descriptor instead.
func (*SonicPort_Port_PortList) Descriptor() ([]byte, []int) {
	return file_sonic_port_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *SonicPort_Port_PortList) GetAdminStatus() SonicPortAdminStatus {
	if x != nil {
		return x.AdminStatus
	}
	return SonicPortAdminStatus_SONICPORTADMINSTATUS_UNSET
}

func (x *SonicPort_Port_PortList) GetAlias() *ywrapper.StringValue {
	if x != nil {
		return x.Alias
	}
	return nil
}

func (x *SonicPort_Port_PortList) GetDescription() *ywrapper.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *SonicPort_Port_PortList) GetFec() *ywrapper.StringValue {
	if x != nil {
		return x.Fec
	}
	return nil
}

func (x *SonicPort_Port_PortList) GetIndex() *ywrapper.UintValue {
	if x != nil {
		return x.Index
	}
	return nil
}

func (x *SonicPort_Port_PortList) GetLanes() *ywrapper.StringValue {
	if x != nil {
		return x.Lanes
	}
	return nil
}

func (x *SonicPort_Port_PortList) GetMtu() *ywrapper.UintValue {
	if x != nil {
		return x.Mtu
	}
	return nil
}

func (x *SonicPort_Port_PortList) GetSpeed() *ywrapper.UintValue {
	if x != nil {
		return x.Speed
	}
	return nil
}

type SonicPort_Port_PortListKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PortName string                   `protobuf:"bytes,1,opt,name=port_name,json=portName,proto3" json:"port_name,omitempty"`
	PortList *SonicPort_Port_PortList `protobuf:"bytes,2,opt,name=port_list,json=portList,proto3" json:"port_list,omitempty"`
}

func (x *SonicPort_Port_PortListKey) Reset() {
	*x = SonicPort_Port_PortListKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_port_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicPort_Port_PortListKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicPort_Port_PortListKey) ProtoMessage() {}

func (x *SonicPort_Port_PortListKey) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_port_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicPort_Port_PortListKey.ProtoReflect.Descriptor instead.
func (*SonicPort_Port_PortListKey) Descriptor() ([]byte, []int) {
	return file_sonic_port_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *SonicPort_Port_PortListKey) GetPortName() string {
	if x != nil {
		return x.PortName
	}
	return ""
}

func (x *SonicPort_Port_PortListKey) GetPortList() *SonicPort_Port_PortList {
	if x != nil {
		return x.PortList
	}
	return nil
}

var File_sonic_port_proto protoreflect.FileDescriptor

var file_sonic_port_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x1a, 0x0e, 0x79, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x79, 0x65, 0x78, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa1, 0x08, 0x0a, 0x09, 0x53, 0x6f, 0x6e, 0x69, 0x63, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x42, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0xfb, 0xb6, 0xa2, 0xea, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x53, 0x6f, 0x6e, 0x69, 0x63,
	0x50, 0x6f, 0x72, 0x74, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x42, 0x13, 0x82, 0x41, 0x10, 0x2f, 0x73,
	0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x1a, 0xcf, 0x07, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x60, 0x0a,
	0x09, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0xe0, 0xcd, 0xee, 0x60, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x53, 0x6f, 0x6e, 0x69,
	0x63, 0x50, 0x6f, 0x72, 0x74, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x42, 0x1d, 0x82, 0x41, 0x1a, 0x2f, 0x73, 0x6f, 0x6e, 0x69,
	0x63, 0x2d, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74,
	0x2d, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x08, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x1a,
	0xd1, 0x05, 0x0a, 0x08, 0x50, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x6d, 0x0a, 0x0c,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0xe7, 0xb8, 0xaa,
	0x61, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x53, 0x6f,
	0x6e, 0x69, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x42, 0x2a, 0x82, 0x41, 0x27, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x70, 0x6f,
	0x72, 0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x6c, 0x69, 0x73,
	0x74, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x54, 0x0a, 0x05, 0x61,
	0x6c, 0x69, 0x61, 0x73, 0x18, 0xe5, 0x92, 0x94, 0xd7, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x79, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x23, 0x82, 0x41, 0x20, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63,
	0x2d, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2d,
	0x6c, 0x69, 0x73, 0x74, 0x2f, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61,
	0x73, 0x12, 0x66, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0xab, 0xac, 0xcc, 0xb1, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x79, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x29, 0x82, 0x41, 0x26, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x70, 0x6f, 0x72,
	0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x6c, 0x69, 0x73, 0x74,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x03, 0x66, 0x65, 0x63,
	0x18, 0xef, 0xed, 0xf9, 0xa7, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x79, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x21, 0x82, 0x41, 0x1e, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x70, 0x6f, 0x72,
	0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x6c, 0x69, 0x73, 0x74,
	0x2f, 0x66, 0x65, 0x63, 0x52, 0x03, 0x66, 0x65, 0x63, 0x12, 0x52, 0x0a, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0xb1, 0x9b, 0xf5, 0x85, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x79,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x23, 0x82, 0x41, 0x20, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x70, 0x6f, 0x72,
	0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x6c, 0x69, 0x73, 0x74,
	0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x54, 0x0a,
	0x05, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x18, 0xf4, 0x8b, 0xea, 0xb4, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x79, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x23, 0x82, 0x41, 0x20, 0x2f, 0x73, 0x6f, 0x6e,
	0x69, 0x63, 0x2d, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x6f, 0x72,
	0x74, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x52, 0x05, 0x6c, 0x61,
	0x6e, 0x65, 0x73, 0x12, 0x4b, 0x0a, 0x03, 0x6d, 0x74, 0x75, 0x18, 0xd9, 0xef, 0x8a, 0x1f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x79, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x55,
	0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x21, 0x82, 0x41, 0x1e, 0x2f, 0x73, 0x6f,
	0x6e, 0x69, 0x63, 0x2d, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x6f,
	0x72, 0x74, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x6d, 0x74, 0x75, 0x52, 0x03, 0x6d, 0x74, 0x75,
	0x12, 0x51, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0xf8, 0xf0, 0xd6, 0x35, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x79, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x55, 0x69,
	0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x23, 0x82, 0x41, 0x20, 0x2f, 0x73, 0x6f, 0x6e,
	0x69, 0x63, 0x2d, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x6f, 0x72,
	0x74, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x52, 0x05, 0x73, 0x70,
	0x65, 0x65, 0x64, 0x1a, 0x90, 0x01, 0x0a, 0x0b, 0x50, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x4b, 0x65, 0x79, 0x12, 0x44, 0x0a, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0x82, 0x41, 0x24, 0x2f, 0x73, 0x6f, 0x6e, 0x69,
	0x63, 0x2d, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74,
	0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x52,
	0x08, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73,
	0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x53, 0x6f, 0x6e, 0x69, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x2e, 0x50,
	0x6f, 0x72, 0x74, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x08, 0x70, 0x6f,
	0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x1b, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x6f,
	0x63, 0x73, 0x79, 0x73, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x5a, 0x07, 0x2e, 0x3b, 0x73, 0x6f,
	0x6e, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sonic_port_proto_rawDescOnce sync.Once
	file_sonic_port_proto_rawDescData = file_sonic_port_proto_rawDesc
)

func file_sonic_port_proto_rawDescGZIP() []byte {
	file_sonic_port_proto_rawDescOnce.Do(func() {
		file_sonic_port_proto_rawDescData = protoimpl.X.CompressGZIP(file_sonic_port_proto_rawDescData)
	})
	return file_sonic_port_proto_rawDescData
}

var file_sonic_port_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sonic_port_proto_goTypes = []interface{}{
	(*SonicPort)(nil),                  // 0: sonic.SonicPort
	(*SonicPort_Port)(nil),             // 1: sonic.SonicPort.Port
	(*SonicPort_Port_PortList)(nil),    // 2: sonic.SonicPort.Port.PortList
	(*SonicPort_Port_PortListKey)(nil), // 3: sonic.SonicPort.Port.PortListKey
	(SonicPortAdminStatus)(0),          // 4: sonic.SonicPortAdminStatus
	(*ywrapper.StringValue)(nil),       // 5: ywrapper.StringValue
	(*ywrapper.UintValue)(nil),         // 6: ywrapper.UintValue
}
var file_sonic_port_proto_depIdxs = []int32{
	1,  // 0: sonic.SonicPort.port:type_name -> sonic.SonicPort.Port
	3,  // 1: sonic.SonicPort.Port.port_list:type_name -> sonic.SonicPort.Port.PortListKey
	4,  // 2: sonic.SonicPort.Port.PortList.admin_status:type_name -> sonic.SonicPortAdminStatus
	5,  // 3: sonic.SonicPort.Port.PortList.alias:type_name -> ywrapper.StringValue
	5,  // 4: sonic.SonicPort.Port.PortList.description:type_name -> ywrapper.StringValue
	5,  // 5: sonic.SonicPort.Port.PortList.fec:type_name -> ywrapper.StringValue
	6,  // 6: sonic.SonicPort.Port.PortList.index:type_name -> ywrapper.UintValue
	5,  // 7: sonic.SonicPort.Port.PortList.lanes:type_name -> ywrapper.StringValue
	6,  // 8: sonic.SonicPort.Port.PortList.mtu:type_name -> ywrapper.UintValue
	6,  // 9: sonic.SonicPort.Port.PortList.speed:type_name -> ywrapper.UintValue
	2,  // 10: sonic.SonicPort.Port.PortListKey.port_list:type_name -> sonic.SonicPort.Port.PortList
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_sonic_port_proto_init() }
func file_sonic_port_proto_init() {
	if File_sonic_port_proto != nil {
		return
	}
	file_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sonic_port_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicPort); i {
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
		file_sonic_port_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicPort_Port); i {
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
		file_sonic_port_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicPort_Port_PortList); i {
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
		file_sonic_port_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicPort_Port_PortListKey); i {
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
			RawDescriptor: file_sonic_port_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sonic_port_proto_goTypes,
		DependencyIndexes: file_sonic_port_proto_depIdxs,
		MessageInfos:      file_sonic_port_proto_msgTypes,
	}.Build()
	File_sonic_port_proto = out.File
	file_sonic_port_proto_rawDesc = nil
	file_sonic_port_proto_goTypes = nil
	file_sonic_port_proto_depIdxs = nil
}
