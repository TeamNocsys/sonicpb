// openconfig.sonic_lldp is generated by proto_generator as a protobuf
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
// source: sonic_lldp.proto

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

type SonicLldp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lldp *SonicLldp_Lldp `protobuf:"bytes,433335412,opt,name=lldp,proto3" json:"lldp,omitempty"`
}

func (x *SonicLldp) Reset() {
	*x = SonicLldp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_lldp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicLldp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicLldp) ProtoMessage() {}

func (x *SonicLldp) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_lldp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicLldp.ProtoReflect.Descriptor instead.
func (*SonicLldp) Descriptor() ([]byte, []int) {
	return file_sonic_lldp_proto_rawDescGZIP(), []int{0}
}

func (x *SonicLldp) GetLldp() *SonicLldp_Lldp {
	if x != nil {
		return x.Lldp
	}
	return nil
}

type SonicLldp_Lldp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceList []*SonicLldp_Lldp_DeviceListKey `protobuf:"bytes,134452366,rep,name=device_list,json=deviceList,proto3" json:"device_list,omitempty"`
}

func (x *SonicLldp_Lldp) Reset() {
	*x = SonicLldp_Lldp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_lldp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicLldp_Lldp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicLldp_Lldp) ProtoMessage() {}

func (x *SonicLldp_Lldp) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_lldp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicLldp_Lldp.ProtoReflect.Descriptor instead.
func (*SonicLldp_Lldp) Descriptor() ([]byte, []int) {
	return file_sonic_lldp_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SonicLldp_Lldp) GetDeviceList() []*SonicLldp_Lldp_DeviceListKey {
	if x != nil {
		return x.DeviceList
	}
	return nil
}

type SonicLldp_Lldp_DeviceList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *SonicLldp_Lldp_DeviceList_Config `protobuf:"bytes,264343743,opt,name=config,proto3" json:"config,omitempty"`
	State  *SonicLldp_Lldp_DeviceList_State  `protobuf:"bytes,160798448,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *SonicLldp_Lldp_DeviceList) Reset() {
	*x = SonicLldp_Lldp_DeviceList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_lldp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicLldp_Lldp_DeviceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicLldp_Lldp_DeviceList) ProtoMessage() {}

func (x *SonicLldp_Lldp_DeviceList) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_lldp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicLldp_Lldp_DeviceList.ProtoReflect.Descriptor instead.
func (*SonicLldp_Lldp_DeviceList) Descriptor() ([]byte, []int) {
	return file_sonic_lldp_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *SonicLldp_Lldp_DeviceList) GetConfig() *SonicLldp_Lldp_DeviceList_Config {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *SonicLldp_Lldp_DeviceList) GetState() *SonicLldp_Lldp_DeviceList_State {
	if x != nil {
		return x.State
	}
	return nil
}

type SonicLldp_Lldp_DeviceListKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceName string                     `protobuf:"bytes,1,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	DeviceList *SonicLldp_Lldp_DeviceList `protobuf:"bytes,2,opt,name=device_list,json=deviceList,proto3" json:"device_list,omitempty"`
}

func (x *SonicLldp_Lldp_DeviceListKey) Reset() {
	*x = SonicLldp_Lldp_DeviceListKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_lldp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicLldp_Lldp_DeviceListKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicLldp_Lldp_DeviceListKey) ProtoMessage() {}

func (x *SonicLldp_Lldp_DeviceListKey) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_lldp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicLldp_Lldp_DeviceListKey.ProtoReflect.Descriptor instead.
func (*SonicLldp_Lldp_DeviceListKey) Descriptor() ([]byte, []int) {
	return file_sonic_lldp_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *SonicLldp_Lldp_DeviceListKey) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *SonicLldp_Lldp_DeviceListKey) GetDeviceList() *SonicLldp_Lldp_DeviceList {
	if x != nil {
		return x.DeviceList
	}
	return nil
}

type SonicLldp_Lldp_DeviceList_Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SonicLldp_Lldp_DeviceList_Config) Reset() {
	*x = SonicLldp_Lldp_DeviceList_Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_lldp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicLldp_Lldp_DeviceList_Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicLldp_Lldp_DeviceList_Config) ProtoMessage() {}

func (x *SonicLldp_Lldp_DeviceList_Config) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_lldp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicLldp_Lldp_DeviceList_Config.ProtoReflect.Descriptor instead.
func (*SonicLldp_Lldp_DeviceList_Config) Descriptor() ([]byte, []int) {
	return file_sonic_lldp_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

type SonicLldp_Lldp_DeviceList_State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Age               *ywrapper.UintValue    `protobuf:"bytes,146671196,opt,name=age,proto3" json:"age,omitempty"`
	ChassisId         *ywrapper.StringValue  `protobuf:"bytes,104698433,opt,name=chassis_id,json=chassisId,proto3" json:"chassis_id,omitempty"`
	ChassisIdType     SonicLldpChassisIdType `protobuf:"varint,265100734,opt,name=chassis_id_type,json=chassisIdType,proto3,enum=sonic.SonicLldpChassisIdType" json:"chassis_id_type,omitempty"`
	Index             *ywrapper.UintValue    `protobuf:"bytes,258738913,opt,name=index,proto3" json:"index,omitempty"`
	LocalPort         *ywrapper.StringValue  `protobuf:"bytes,30130656,opt,name=local_port,json=localPort,proto3" json:"local_port,omitempty"`
	PortDescription   *ywrapper.StringValue  `protobuf:"bytes,158789391,opt,name=port_description,json=portDescription,proto3" json:"port_description,omitempty"`
	PortId            *ywrapper.StringValue  `protobuf:"bytes,374919748,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	PortIdType        SonicLldpPortIdType    `protobuf:"varint,264867561,opt,name=port_id_type,json=portIdType,proto3,enum=sonic.SonicLldpPortIdType" json:"port_id_type,omitempty"`
	SystemDescription *ywrapper.StringValue  `protobuf:"bytes,480308725,opt,name=system_description,json=systemDescription,proto3" json:"system_description,omitempty"`
	SystemName        *ywrapper.StringValue  `protobuf:"bytes,169239952,opt,name=system_name,json=systemName,proto3" json:"system_name,omitempty"`
}

func (x *SonicLldp_Lldp_DeviceList_State) Reset() {
	*x = SonicLldp_Lldp_DeviceList_State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_lldp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicLldp_Lldp_DeviceList_State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicLldp_Lldp_DeviceList_State) ProtoMessage() {}

func (x *SonicLldp_Lldp_DeviceList_State) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_lldp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicLldp_Lldp_DeviceList_State.ProtoReflect.Descriptor instead.
func (*SonicLldp_Lldp_DeviceList_State) Descriptor() ([]byte, []int) {
	return file_sonic_lldp_proto_rawDescGZIP(), []int{0, 0, 0, 1}
}

func (x *SonicLldp_Lldp_DeviceList_State) GetAge() *ywrapper.UintValue {
	if x != nil {
		return x.Age
	}
	return nil
}

func (x *SonicLldp_Lldp_DeviceList_State) GetChassisId() *ywrapper.StringValue {
	if x != nil {
		return x.ChassisId
	}
	return nil
}

func (x *SonicLldp_Lldp_DeviceList_State) GetChassisIdType() SonicLldpChassisIdType {
	if x != nil {
		return x.ChassisIdType
	}
	return SonicLldpChassisIdType_SONICLLDPCHASSISIDTYPE_UNSET
}

func (x *SonicLldp_Lldp_DeviceList_State) GetIndex() *ywrapper.UintValue {
	if x != nil {
		return x.Index
	}
	return nil
}

func (x *SonicLldp_Lldp_DeviceList_State) GetLocalPort() *ywrapper.StringValue {
	if x != nil {
		return x.LocalPort
	}
	return nil
}

func (x *SonicLldp_Lldp_DeviceList_State) GetPortDescription() *ywrapper.StringValue {
	if x != nil {
		return x.PortDescription
	}
	return nil
}

func (x *SonicLldp_Lldp_DeviceList_State) GetPortId() *ywrapper.StringValue {
	if x != nil {
		return x.PortId
	}
	return nil
}

func (x *SonicLldp_Lldp_DeviceList_State) GetPortIdType() SonicLldpPortIdType {
	if x != nil {
		return x.PortIdType
	}
	return SonicLldpPortIdType_SONICLLDPPORTIDTYPE_UNSET
}

func (x *SonicLldp_Lldp_DeviceList_State) GetSystemDescription() *ywrapper.StringValue {
	if x != nil {
		return x.SystemDescription
	}
	return nil
}

func (x *SonicLldp_Lldp_DeviceList_State) GetSystemName() *ywrapper.StringValue {
	if x != nil {
		return x.SystemName
	}
	return nil
}

var File_sonic_lldp_proto protoreflect.FileDescriptor

var file_sonic_lldp_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x5f, 0x6c, 0x6c, 0x64, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x1a, 0x0e, 0x79, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x79, 0x65, 0x78, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xab, 0x0d, 0x0a, 0x09, 0x53, 0x6f, 0x6e, 0x69, 0x63, 0x4c, 0x6c, 0x64, 0x70,
	0x12, 0x42, 0x0a, 0x04, 0x6c, 0x6c, 0x64, 0x70, 0x18, 0xf4, 0xd8, 0xd0, 0xce, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x53, 0x6f, 0x6e, 0x69, 0x63,
	0x4c, 0x6c, 0x64, 0x70, 0x2e, 0x4c, 0x6c, 0x64, 0x70, 0x42, 0x13, 0x82, 0x41, 0x10, 0x2f, 0x73,
	0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x6c, 0x6c, 0x64, 0x70, 0x2f, 0x6c, 0x6c, 0x64, 0x70, 0x52, 0x04,
	0x6c, 0x6c, 0x64, 0x70, 0x1a, 0xd9, 0x0c, 0x0a, 0x04, 0x4c, 0x6c, 0x64, 0x70, 0x12, 0x68, 0x0a,
	0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x8e, 0xa9, 0x8e,
	0x40, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x53, 0x6f,
	0x6e, 0x69, 0x63, 0x4c, 0x6c, 0x64, 0x70, 0x2e, 0x4c, 0x6c, 0x64, 0x70, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x42, 0x1f, 0x82, 0x41, 0x1c, 0x2f,
	0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x6c, 0x6c, 0x64, 0x70, 0x2f, 0x6c, 0x6c, 0x64, 0x70, 0x2f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0xc3, 0x0a, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x6a, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0xbf, 0xa1, 0x86, 0x7e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x6f, 0x6e, 0x69,
	0x63, 0x2e, 0x53, 0x6f, 0x6e, 0x69, 0x63, 0x4c, 0x6c, 0x64, 0x70, 0x2e, 0x4c, 0x6c, 0x64, 0x70,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x42, 0x26, 0x82, 0x41, 0x23, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x6c, 0x6c,
	0x64, 0x70, 0x2f, 0x6c, 0x6c, 0x64, 0x70, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6c,
	0x69, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x66, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0xf0, 0xad, 0xd6, 0x4c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x53, 0x6f, 0x6e,
	0x69, 0x63, 0x4c, 0x6c, 0x64, 0x70, 0x2e, 0x4c, 0x6c, 0x64, 0x70, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x25, 0x82, 0x41,
	0x22, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x6c, 0x6c, 0x64, 0x70, 0x2f, 0x6c, 0x6c, 0x64,
	0x70, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x08, 0x0a, 0x06, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x1a, 0xd6, 0x08, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x53,
	0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0xdc, 0x8c, 0xf8, 0x45, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x79, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x29, 0x82, 0x41, 0x26, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x6c,
	0x6c, 0x64, 0x70, 0x2f, 0x6c, 0x6c, 0x64, 0x70, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d,
	0x6c, 0x69, 0x73, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x67, 0x65, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x12, 0x69, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x73, 0x73, 0x69, 0x73, 0x5f, 0x69,
	0x64, 0x18, 0xc1, 0xa4, 0xf6, 0x31, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x79, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x30, 0x82, 0x41, 0x2d, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x6c, 0x6c, 0x64,
	0x70, 0x2f, 0x6c, 0x6c, 0x64, 0x70, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6c, 0x69,
	0x73, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x73, 0x73, 0x69, 0x73,
	0x2d, 0x69, 0x64, 0x52, 0x09, 0x63, 0x68, 0x61, 0x73, 0x73, 0x69, 0x73, 0x49, 0x64, 0x12, 0x7f,
	0x0a, 0x0f, 0x63, 0x68, 0x61, 0x73, 0x73, 0x69, 0x73, 0x5f, 0x69, 0x64, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0xbe, 0xbb, 0xb4, 0x7e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x73, 0x6f, 0x6e,
	0x69, 0x63, 0x2e, 0x53, 0x6f, 0x6e, 0x69, 0x63, 0x4c, 0x6c, 0x64, 0x70, 0x43, 0x68, 0x61, 0x73,
	0x73, 0x69, 0x73, 0x49, 0x64, 0x54, 0x79, 0x70, 0x65, 0x42, 0x35, 0x82, 0x41, 0x32, 0x2f, 0x73,
	0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x6c, 0x6c, 0x64, 0x70, 0x2f, 0x6c, 0x6c, 0x64, 0x70, 0x2f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2f, 0x63, 0x68, 0x61, 0x73, 0x73, 0x69, 0x73, 0x2d, 0x69, 0x64, 0x2d, 0x74, 0x79, 0x70, 0x65,
	0x52, 0x0d, 0x63, 0x68, 0x61, 0x73, 0x73, 0x69, 0x73, 0x49, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x59, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0xe1, 0x95, 0xb0, 0x7b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x79, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x55, 0x69, 0x6e,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x2b, 0x82, 0x41, 0x28, 0x2f, 0x73, 0x6f, 0x6e, 0x69,
	0x63, 0x2d, 0x6c, 0x6c, 0x64, 0x70, 0x2f, 0x6c, 0x6c, 0x64, 0x70, 0x2f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x69, 0x0a, 0x0a, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0xe0, 0x83, 0xaf, 0x0e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x79, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x30, 0x82, 0x41, 0x2d, 0x2f, 0x73, 0x6f,
	0x6e, 0x69, 0x63, 0x2d, 0x6c, 0x6c, 0x64, 0x70, 0x2f, 0x6c, 0x6c, 0x64, 0x70, 0x2f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x7b, 0x0a, 0x10, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x8f, 0xde, 0xdb, 0x4b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x79, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x36, 0x82, 0x41, 0x33, 0x2f, 0x73,
	0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x6c, 0x6c, 0x64, 0x70, 0x2f, 0x6c, 0x6c, 0x64, 0x70, 0x2f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0f, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x61, 0x0a, 0x07, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0xc4, 0xa4,
	0xe3, 0xb2, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x79, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x2d,
	0x82, 0x41, 0x2a, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x6c, 0x6c, 0x64, 0x70, 0x2f, 0x6c,
	0x6c, 0x64, 0x70, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x69, 0x64, 0x52, 0x06, 0x70,
	0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x73, 0x0a, 0x0c, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0xe9, 0x9d, 0xa6, 0x7e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a,
	0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x53, 0x6f, 0x6e, 0x69, 0x63, 0x4c, 0x6c, 0x64, 0x70,
	0x50, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x54, 0x79, 0x70, 0x65, 0x42, 0x32, 0x82, 0x41, 0x2f, 0x2f,
	0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x6c, 0x6c, 0x64, 0x70, 0x2f, 0x6c, 0x6c, 0x64, 0x70, 0x2f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x69, 0x64, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x52, 0x0a,
	0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x12, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0xf5, 0xdb, 0x83, 0xe5, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x79, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x38, 0x82, 0x41, 0x35, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x6c, 0x6c,
	0x64, 0x70, 0x2f, 0x6c, 0x6c, 0x64, 0x70, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6c,
	0x69, 0x73, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x6c, 0x0a, 0x0b, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x90,
	0xcb, 0xd9, 0x50, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x79, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x31,
	0x82, 0x41, 0x2e, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x6c, 0x6c, 0x64, 0x70, 0x2f, 0x6c,
	0x6c, 0x64, 0x70, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2d, 0x6e, 0x61, 0x6d,
	0x65, 0x52, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0xa0, 0x01,
	0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x12,
	0x4c, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0x82, 0x41, 0x28, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d,
	0x6c, 0x6c, 0x64, 0x70, 0x2f, 0x6c, 0x6c, 0x64, 0x70, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6e, 0x61, 0x6d,
	0x65, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a,
	0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x53, 0x6f, 0x6e, 0x69, 0x63,
	0x4c, 0x6c, 0x64, 0x70, 0x2e, 0x4c, 0x6c, 0x64, 0x70, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x12, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x6f, 0x63, 0x73, 0x79, 0x73, 0x2e, 0x73,
	0x6f, 0x6e, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sonic_lldp_proto_rawDescOnce sync.Once
	file_sonic_lldp_proto_rawDescData = file_sonic_lldp_proto_rawDesc
)

func file_sonic_lldp_proto_rawDescGZIP() []byte {
	file_sonic_lldp_proto_rawDescOnce.Do(func() {
		file_sonic_lldp_proto_rawDescData = protoimpl.X.CompressGZIP(file_sonic_lldp_proto_rawDescData)
	})
	return file_sonic_lldp_proto_rawDescData
}

var file_sonic_lldp_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sonic_lldp_proto_goTypes = []interface{}{
	(*SonicLldp)(nil),                        // 0: sonic.SonicLldp
	(*SonicLldp_Lldp)(nil),                   // 1: sonic.SonicLldp.Lldp
	(*SonicLldp_Lldp_DeviceList)(nil),        // 2: sonic.SonicLldp.Lldp.DeviceList
	(*SonicLldp_Lldp_DeviceListKey)(nil),     // 3: sonic.SonicLldp.Lldp.DeviceListKey
	(*SonicLldp_Lldp_DeviceList_Config)(nil), // 4: sonic.SonicLldp.Lldp.DeviceList.Config
	(*SonicLldp_Lldp_DeviceList_State)(nil),  // 5: sonic.SonicLldp.Lldp.DeviceList.State
	(*ywrapper.UintValue)(nil),               // 6: ywrapper.UintValue
	(*ywrapper.StringValue)(nil),             // 7: ywrapper.StringValue
	(SonicLldpChassisIdType)(0),              // 8: sonic.SonicLldpChassisIdType
	(SonicLldpPortIdType)(0),                 // 9: sonic.SonicLldpPortIdType
}
var file_sonic_lldp_proto_depIdxs = []int32{
	1,  // 0: sonic.SonicLldp.lldp:type_name -> sonic.SonicLldp.Lldp
	3,  // 1: sonic.SonicLldp.Lldp.device_list:type_name -> sonic.SonicLldp.Lldp.DeviceListKey
	4,  // 2: sonic.SonicLldp.Lldp.DeviceList.config:type_name -> sonic.SonicLldp.Lldp.DeviceList.Config
	5,  // 3: sonic.SonicLldp.Lldp.DeviceList.state:type_name -> sonic.SonicLldp.Lldp.DeviceList.State
	2,  // 4: sonic.SonicLldp.Lldp.DeviceListKey.device_list:type_name -> sonic.SonicLldp.Lldp.DeviceList
	6,  // 5: sonic.SonicLldp.Lldp.DeviceList.State.age:type_name -> ywrapper.UintValue
	7,  // 6: sonic.SonicLldp.Lldp.DeviceList.State.chassis_id:type_name -> ywrapper.StringValue
	8,  // 7: sonic.SonicLldp.Lldp.DeviceList.State.chassis_id_type:type_name -> sonic.SonicLldpChassisIdType
	6,  // 8: sonic.SonicLldp.Lldp.DeviceList.State.index:type_name -> ywrapper.UintValue
	7,  // 9: sonic.SonicLldp.Lldp.DeviceList.State.local_port:type_name -> ywrapper.StringValue
	7,  // 10: sonic.SonicLldp.Lldp.DeviceList.State.port_description:type_name -> ywrapper.StringValue
	7,  // 11: sonic.SonicLldp.Lldp.DeviceList.State.port_id:type_name -> ywrapper.StringValue
	9,  // 12: sonic.SonicLldp.Lldp.DeviceList.State.port_id_type:type_name -> sonic.SonicLldpPortIdType
	7,  // 13: sonic.SonicLldp.Lldp.DeviceList.State.system_description:type_name -> ywrapper.StringValue
	7,  // 14: sonic.SonicLldp.Lldp.DeviceList.State.system_name:type_name -> ywrapper.StringValue
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_sonic_lldp_proto_init() }
func file_sonic_lldp_proto_init() {
	if File_sonic_lldp_proto != nil {
		return
	}
	file_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sonic_lldp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicLldp); i {
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
		file_sonic_lldp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicLldp_Lldp); i {
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
		file_sonic_lldp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicLldp_Lldp_DeviceList); i {
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
		file_sonic_lldp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicLldp_Lldp_DeviceListKey); i {
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
		file_sonic_lldp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicLldp_Lldp_DeviceList_Config); i {
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
		file_sonic_lldp_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicLldp_Lldp_DeviceList_State); i {
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
			RawDescriptor: file_sonic_lldp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sonic_lldp_proto_goTypes,
		DependencyIndexes: file_sonic_lldp_proto_depIdxs,
		MessageInfos:      file_sonic_lldp_proto_msgTypes,
	}.Build()
	File_sonic_lldp_proto = out.File
	file_sonic_lldp_proto_rawDesc = nil
	file_sonic_lldp_proto_goTypes = nil
	file_sonic_lldp_proto_depIdxs = nil
}