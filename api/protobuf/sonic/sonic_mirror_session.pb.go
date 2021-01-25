// openconfig.sonic_mirror_session is generated by proto_generator as a protobuf
// representation of a YANG schema.
//
// Input schema modules:
//  - ../api/yang/sonic/sonic-platform.yang
//  - ../api/yang/sonic/sonic-interface.yang
//  - ../api/yang/sonic/sonic-port.yang
//  - ../api/yang/sonic/sonic-ntp.yang
//  - ../api/yang/sonic/sonic-vrf.yang
//  - ../api/yang/sonic/third_party/ietf/ietf-interfaces.yang
//  - ../api/yang/sonic/third_party/ietf/ietf-yang-types.yang
//  - ../api/yang/sonic/third_party/ietf/ietf-inet-types.yang
//  - ../api/yang/sonic/third_party/ietf/iana-if-type.yang
//  - ../api/yang/sonic/sonic-loopback-interface.yang
//  - ../api/yang/sonic/sonic-platform-types.yang
//  - ../api/yang/sonic/sonic-neighbor.yang
//  - ../api/yang/sonic/sonic-route.yang
//  - ../api/yang/sonic/sonic-lldp.yang
//  - ../api/yang/sonic/sonic-vlan.yang
//  - ../api/yang/sonic/sonic-portchannel.yang
//  - ../api/yang/sonic/sonic-extension.yang
//  - ../api/yang/sonic/sonic-mirror-session.yang
//  - ../api/yang/sonic/sonic-vxlan.yang
//  - ../api/yang/sonic/sonic-fdb.yang
//  - ../api/yang/sonic/sonic-acl.yang
//  - ../api/yang/sonic/sonic-types.yang

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: sonic_mirror_session.proto

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

type SonicMirrorSession_MirrorSession_MirrorSessionList_Direction int32

const (
	SonicMirrorSession_MirrorSession_MirrorSessionList_DIRECTION_UNSET SonicMirrorSession_MirrorSession_MirrorSessionList_Direction = 0
	SonicMirrorSession_MirrorSession_MirrorSessionList_DIRECTION_RX    SonicMirrorSession_MirrorSession_MirrorSessionList_Direction = 1
	SonicMirrorSession_MirrorSession_MirrorSessionList_DIRECTION_TX    SonicMirrorSession_MirrorSession_MirrorSessionList_Direction = 2
	SonicMirrorSession_MirrorSession_MirrorSessionList_DIRECTION_BOTH  SonicMirrorSession_MirrorSession_MirrorSessionList_Direction = 3
)

// Enum value maps for SonicMirrorSession_MirrorSession_MirrorSessionList_Direction.
var (
	SonicMirrorSession_MirrorSession_MirrorSessionList_Direction_name = map[int32]string{
		0: "DIRECTION_UNSET",
		1: "DIRECTION_RX",
		2: "DIRECTION_TX",
		3: "DIRECTION_BOTH",
	}
	SonicMirrorSession_MirrorSession_MirrorSessionList_Direction_value = map[string]int32{
		"DIRECTION_UNSET": 0,
		"DIRECTION_RX":    1,
		"DIRECTION_TX":    2,
		"DIRECTION_BOTH":  3,
	}
)

func (x SonicMirrorSession_MirrorSession_MirrorSessionList_Direction) Enum() *SonicMirrorSession_MirrorSession_MirrorSessionList_Direction {
	p := new(SonicMirrorSession_MirrorSession_MirrorSessionList_Direction)
	*p = x
	return p
}

func (x SonicMirrorSession_MirrorSession_MirrorSessionList_Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SonicMirrorSession_MirrorSession_MirrorSessionList_Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_sonic_mirror_session_proto_enumTypes[0].Descriptor()
}

func (SonicMirrorSession_MirrorSession_MirrorSessionList_Direction) Type() protoreflect.EnumType {
	return &file_sonic_mirror_session_proto_enumTypes[0]
}

func (x SonicMirrorSession_MirrorSession_MirrorSessionList_Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SonicMirrorSession_MirrorSession_MirrorSessionList_Direction.Descriptor instead.
func (SonicMirrorSession_MirrorSession_MirrorSessionList_Direction) EnumDescriptor() ([]byte, []int) {
	return file_sonic_mirror_session_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

type SonicMirrorSession_MirrorSession_MirrorSessionList_Status int32

const (
	SonicMirrorSession_MirrorSession_MirrorSessionList_STATUS_UNSET    SonicMirrorSession_MirrorSession_MirrorSessionList_Status = 0
	SonicMirrorSession_MirrorSession_MirrorSessionList_STATUS_active   SonicMirrorSession_MirrorSession_MirrorSessionList_Status = 1
	SonicMirrorSession_MirrorSession_MirrorSessionList_STATUS_inactive SonicMirrorSession_MirrorSession_MirrorSessionList_Status = 2
)

// Enum value maps for SonicMirrorSession_MirrorSession_MirrorSessionList_Status.
var (
	SonicMirrorSession_MirrorSession_MirrorSessionList_Status_name = map[int32]string{
		0: "STATUS_UNSET",
		1: "STATUS_active",
		2: "STATUS_inactive",
	}
	SonicMirrorSession_MirrorSession_MirrorSessionList_Status_value = map[string]int32{
		"STATUS_UNSET":    0,
		"STATUS_active":   1,
		"STATUS_inactive": 2,
	}
)

func (x SonicMirrorSession_MirrorSession_MirrorSessionList_Status) Enum() *SonicMirrorSession_MirrorSession_MirrorSessionList_Status {
	p := new(SonicMirrorSession_MirrorSession_MirrorSessionList_Status)
	*p = x
	return p
}

func (x SonicMirrorSession_MirrorSession_MirrorSessionList_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SonicMirrorSession_MirrorSession_MirrorSessionList_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_sonic_mirror_session_proto_enumTypes[1].Descriptor()
}

func (SonicMirrorSession_MirrorSession_MirrorSessionList_Status) Type() protoreflect.EnumType {
	return &file_sonic_mirror_session_proto_enumTypes[1]
}

func (x SonicMirrorSession_MirrorSession_MirrorSessionList_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SonicMirrorSession_MirrorSession_MirrorSessionList_Status.Descriptor instead.
func (SonicMirrorSession_MirrorSession_MirrorSessionList_Status) EnumDescriptor() ([]byte, []int) {
	return file_sonic_mirror_session_proto_rawDescGZIP(), []int{0, 0, 0, 1}
}

type SonicMirrorSession_MirrorSession_MirrorSessionList_Type int32

const (
	SonicMirrorSession_MirrorSession_MirrorSessionList_TYPE_UNSET  SonicMirrorSession_MirrorSession_MirrorSessionList_Type = 0
	SonicMirrorSession_MirrorSession_MirrorSessionList_TYPE_SPAN   SonicMirrorSession_MirrorSession_MirrorSessionList_Type = 1
	SonicMirrorSession_MirrorSession_MirrorSessionList_TYPE_ERSPAN SonicMirrorSession_MirrorSession_MirrorSessionList_Type = 2
)

// Enum value maps for SonicMirrorSession_MirrorSession_MirrorSessionList_Type.
var (
	SonicMirrorSession_MirrorSession_MirrorSessionList_Type_name = map[int32]string{
		0: "TYPE_UNSET",
		1: "TYPE_SPAN",
		2: "TYPE_ERSPAN",
	}
	SonicMirrorSession_MirrorSession_MirrorSessionList_Type_value = map[string]int32{
		"TYPE_UNSET":  0,
		"TYPE_SPAN":   1,
		"TYPE_ERSPAN": 2,
	}
)

func (x SonicMirrorSession_MirrorSession_MirrorSessionList_Type) Enum() *SonicMirrorSession_MirrorSession_MirrorSessionList_Type {
	p := new(SonicMirrorSession_MirrorSession_MirrorSessionList_Type)
	*p = x
	return p
}

func (x SonicMirrorSession_MirrorSession_MirrorSessionList_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SonicMirrorSession_MirrorSession_MirrorSessionList_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_sonic_mirror_session_proto_enumTypes[2].Descriptor()
}

func (SonicMirrorSession_MirrorSession_MirrorSessionList_Type) Type() protoreflect.EnumType {
	return &file_sonic_mirror_session_proto_enumTypes[2]
}

func (x SonicMirrorSession_MirrorSession_MirrorSessionList_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SonicMirrorSession_MirrorSession_MirrorSessionList_Type.Descriptor instead.
func (SonicMirrorSession_MirrorSession_MirrorSessionList_Type) EnumDescriptor() ([]byte, []int) {
	return file_sonic_mirror_session_proto_rawDescGZIP(), []int{0, 0, 0, 2}
}

type SonicMirrorSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MirrorSession *SonicMirrorSession_MirrorSession `protobuf:"bytes,1,opt,name=mirror_session,json=mirrorSession,proto3" json:"mirror_session,omitempty"`
}

func (x *SonicMirrorSession) Reset() {
	*x = SonicMirrorSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_mirror_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicMirrorSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicMirrorSession) ProtoMessage() {}

func (x *SonicMirrorSession) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_mirror_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicMirrorSession.ProtoReflect.Descriptor instead.
func (*SonicMirrorSession) Descriptor() ([]byte, []int) {
	return file_sonic_mirror_session_proto_rawDescGZIP(), []int{0}
}

func (x *SonicMirrorSession) GetMirrorSession() *SonicMirrorSession_MirrorSession {
	if x != nil {
		return x.MirrorSession
	}
	return nil
}

type SonicMirrorSession_MirrorSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MirrorSessionList []*SonicMirrorSession_MirrorSession_MirrorSessionListKey `protobuf:"bytes,1,rep,name=mirror_session_list,json=mirrorSessionList,proto3" json:"mirror_session_list,omitempty"`
}

func (x *SonicMirrorSession_MirrorSession) Reset() {
	*x = SonicMirrorSession_MirrorSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_mirror_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicMirrorSession_MirrorSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicMirrorSession_MirrorSession) ProtoMessage() {}

func (x *SonicMirrorSession_MirrorSession) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_mirror_session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicMirrorSession_MirrorSession.ProtoReflect.Descriptor instead.
func (*SonicMirrorSession_MirrorSession) Descriptor() ([]byte, []int) {
	return file_sonic_mirror_session_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SonicMirrorSession_MirrorSession) GetMirrorSessionList() []*SonicMirrorSession_MirrorSession_MirrorSessionListKey {
	if x != nil {
		return x.MirrorSessionList
	}
	return nil
}

type SonicMirrorSession_MirrorSession_MirrorSessionList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Direction SonicMirrorSession_MirrorSession_MirrorSessionList_Direction `protobuf:"varint,1,opt,name=direction,proto3,enum=sonic.SonicMirrorSession_MirrorSession_MirrorSessionList_Direction" json:"direction,omitempty"`
	Dscp      *ywrapper.UintValue                                          `protobuf:"bytes,2,opt,name=dscp,proto3" json:"dscp,omitempty"`
	DstIp     *ywrapper.StringValue                                        `protobuf:"bytes,3,opt,name=dst_ip,json=dstIp,proto3" json:"dst_ip,omitempty"`
	DstPort   *ywrapper.StringValue                                        `protobuf:"bytes,4,opt,name=dst_port,json=dstPort,proto3" json:"dst_port,omitempty"`
	GreType   *ywrapper.UintValue                                          `protobuf:"bytes,5,opt,name=gre_type,json=greType,proto3" json:"gre_type,omitempty"`
	Policer   *ywrapper.StringValue                                        `protobuf:"bytes,6,opt,name=policer,proto3" json:"policer,omitempty"`
	Queue     *ywrapper.UintValue                                          `protobuf:"bytes,7,opt,name=queue,proto3" json:"queue,omitempty"`
	SrcIp     *ywrapper.StringValue                                        `protobuf:"bytes,8,opt,name=src_ip,json=srcIp,proto3" json:"src_ip,omitempty"`
	SrcPort   *ywrapper.StringValue                                        `protobuf:"bytes,9,opt,name=src_port,json=srcPort,proto3" json:"src_port,omitempty"`
	Status    SonicMirrorSession_MirrorSession_MirrorSessionList_Status    `protobuf:"varint,10,opt,name=status,proto3,enum=sonic.SonicMirrorSession_MirrorSession_MirrorSessionList_Status" json:"status,omitempty"`
	Ttl       *ywrapper.UintValue                                          `protobuf:"bytes,11,opt,name=ttl,proto3" json:"ttl,omitempty"`
	Type      SonicMirrorSession_MirrorSession_MirrorSessionList_Type      `protobuf:"varint,12,opt,name=type,proto3,enum=sonic.SonicMirrorSession_MirrorSession_MirrorSessionList_Type" json:"type,omitempty"`
}

func (x *SonicMirrorSession_MirrorSession_MirrorSessionList) Reset() {
	*x = SonicMirrorSession_MirrorSession_MirrorSessionList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_mirror_session_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicMirrorSession_MirrorSession_MirrorSessionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicMirrorSession_MirrorSession_MirrorSessionList) ProtoMessage() {}

func (x *SonicMirrorSession_MirrorSession_MirrorSessionList) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_mirror_session_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicMirrorSession_MirrorSession_MirrorSessionList.ProtoReflect.Descriptor instead.
func (*SonicMirrorSession_MirrorSession_MirrorSessionList) Descriptor() ([]byte, []int) {
	return file_sonic_mirror_session_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *SonicMirrorSession_MirrorSession_MirrorSessionList) GetDirection() SonicMirrorSession_MirrorSession_MirrorSessionList_Direction {
	if x != nil {
		return x.Direction
	}
	return SonicMirrorSession_MirrorSession_MirrorSessionList_DIRECTION_UNSET
}

func (x *SonicMirrorSession_MirrorSession_MirrorSessionList) GetDscp() *ywrapper.UintValue {
	if x != nil {
		return x.Dscp
	}
	return nil
}

func (x *SonicMirrorSession_MirrorSession_MirrorSessionList) GetDstIp() *ywrapper.StringValue {
	if x != nil {
		return x.DstIp
	}
	return nil
}

func (x *SonicMirrorSession_MirrorSession_MirrorSessionList) GetDstPort() *ywrapper.StringValue {
	if x != nil {
		return x.DstPort
	}
	return nil
}

func (x *SonicMirrorSession_MirrorSession_MirrorSessionList) GetGreType() *ywrapper.UintValue {
	if x != nil {
		return x.GreType
	}
	return nil
}

func (x *SonicMirrorSession_MirrorSession_MirrorSessionList) GetPolicer() *ywrapper.StringValue {
	if x != nil {
		return x.Policer
	}
	return nil
}

func (x *SonicMirrorSession_MirrorSession_MirrorSessionList) GetQueue() *ywrapper.UintValue {
	if x != nil {
		return x.Queue
	}
	return nil
}

func (x *SonicMirrorSession_MirrorSession_MirrorSessionList) GetSrcIp() *ywrapper.StringValue {
	if x != nil {
		return x.SrcIp
	}
	return nil
}

func (x *SonicMirrorSession_MirrorSession_MirrorSessionList) GetSrcPort() *ywrapper.StringValue {
	if x != nil {
		return x.SrcPort
	}
	return nil
}

func (x *SonicMirrorSession_MirrorSession_MirrorSessionList) GetStatus() SonicMirrorSession_MirrorSession_MirrorSessionList_Status {
	if x != nil {
		return x.Status
	}
	return SonicMirrorSession_MirrorSession_MirrorSessionList_STATUS_UNSET
}

func (x *SonicMirrorSession_MirrorSession_MirrorSessionList) GetTtl() *ywrapper.UintValue {
	if x != nil {
		return x.Ttl
	}
	return nil
}

func (x *SonicMirrorSession_MirrorSession_MirrorSessionList) GetType() SonicMirrorSession_MirrorSession_MirrorSessionList_Type {
	if x != nil {
		return x.Type
	}
	return SonicMirrorSession_MirrorSession_MirrorSessionList_TYPE_UNSET
}

type SonicMirrorSession_MirrorSession_MirrorSessionListKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string                                              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MirrorSessionList *SonicMirrorSession_MirrorSession_MirrorSessionList `protobuf:"bytes,2,opt,name=mirror_session_list,json=mirrorSessionList,proto3" json:"mirror_session_list,omitempty"`
}

func (x *SonicMirrorSession_MirrorSession_MirrorSessionListKey) Reset() {
	*x = SonicMirrorSession_MirrorSession_MirrorSessionListKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonic_mirror_session_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SonicMirrorSession_MirrorSession_MirrorSessionListKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SonicMirrorSession_MirrorSession_MirrorSessionListKey) ProtoMessage() {}

func (x *SonicMirrorSession_MirrorSession_MirrorSessionListKey) ProtoReflect() protoreflect.Message {
	mi := &file_sonic_mirror_session_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SonicMirrorSession_MirrorSession_MirrorSessionListKey.ProtoReflect.Descriptor instead.
func (*SonicMirrorSession_MirrorSession_MirrorSessionListKey) Descriptor() ([]byte, []int) {
	return file_sonic_mirror_session_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *SonicMirrorSession_MirrorSession_MirrorSessionListKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SonicMirrorSession_MirrorSession_MirrorSessionListKey) GetMirrorSessionList() *SonicMirrorSession_MirrorSession_MirrorSessionList {
	if x != nil {
		return x.MirrorSessionList
	}
	return nil
}

var File_sonic_mirror_session_proto protoreflect.FileDescriptor

var file_sonic_mirror_session_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x5f, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x6f,
	0x6e, 0x69, 0x63, 0x1a, 0x0e, 0x79, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x79, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb6, 0x12, 0x0a, 0x12, 0x53, 0x6f, 0x6e, 0x69, 0x63, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x77, 0x0a, 0x0e, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x53, 0x6f, 0x6e, 0x69, 0x63, 0x4d, 0x69, 0x72, 0x72,
	0x6f, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x27, 0x82, 0x41, 0x24, 0x2f, 0x73, 0x6f, 0x6e,
	0x69, 0x63, 0x2d, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2f, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x0d, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a,
	0xa6, 0x11, 0x0a, 0x0d, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0xa9, 0x01, 0x0a, 0x13, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x3c, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x53, 0x6f, 0x6e, 0x69, 0x63, 0x4d, 0x69, 0x72,
	0x72, 0x6f, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x69, 0x72, 0x72, 0x6f,
	0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x42, 0x3b, 0x82,
	0x41, 0x38, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x11, 0x6d, 0x69, 0x72, 0x72,
	0x6f, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x8e, 0x0e,
	0x0a, 0x11, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0xa8, 0x01, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x43, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e,
	0x53, 0x6f, 0x6e, 0x69, 0x63, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x45, 0x82, 0x41,
	0x42, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x69,
	0x0a, 0x04, 0x64, 0x73, 0x63, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x79,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x40, 0x82, 0x41, 0x3d, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x6d, 0x69, 0x72,
	0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x69, 0x72, 0x72,
	0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x69, 0x72, 0x72, 0x6f,
	0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x64,
	0x73, 0x63, 0x70, 0x52, 0x04, 0x64, 0x73, 0x63, 0x70, 0x12, 0x70, 0x0a, 0x06, 0x64, 0x73, 0x74,
	0x5f, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x79, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x42, 0x82, 0x41, 0x3f, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x6d, 0x69, 0x72, 0x72,
	0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x69, 0x72, 0x72, 0x6f,
	0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72,
	0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x64, 0x73,
	0x74, 0x2d, 0x69, 0x70, 0x52, 0x05, 0x64, 0x73, 0x74, 0x49, 0x70, 0x12, 0x76, 0x0a, 0x08, 0x64,
	0x73, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x79, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x44, 0x82, 0x41, 0x41, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d,
	0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d,
	0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x69,
	0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x69, 0x73,
	0x74, 0x2f, 0x64, 0x73, 0x74, 0x2d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x07, 0x64, 0x73, 0x74, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x74, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x79, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x2e, 0x55, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x44, 0x82, 0x41, 0x41, 0x2f,
	0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x72, 0x65, 0x2d, 0x74, 0x79, 0x70, 0x65,
	0x52, 0x07, 0x67, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x74, 0x0a, 0x07, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x79, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x43, 0x82, 0x41, 0x40, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x6d, 0x69, 0x72,
	0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x69, 0x72, 0x72,
	0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x69, 0x72, 0x72, 0x6f,
	0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x72, 0x12,
	0x6c, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x79, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x41, 0x82, 0x41, 0x3e, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x6d,
	0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x69,
	0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x69, 0x72,
	0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x69, 0x73, 0x74,
	0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x52, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x70, 0x0a,
	0x06, 0x73, 0x72, 0x63, 0x5f, 0x69, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x79, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x42, 0x82, 0x41, 0x3f, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d,
	0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d,
	0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x69,
	0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x69, 0x73,
	0x74, 0x2f, 0x73, 0x72, 0x63, 0x2d, 0x69, 0x70, 0x52, 0x05, 0x73, 0x72, 0x63, 0x49, 0x70, 0x12,
	0x76, 0x0a, 0x08, 0x73, 0x72, 0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x79, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x44, 0x82, 0x41, 0x41, 0x2f, 0x73, 0x6f,
	0x6e, 0x69, 0x63, 0x2d, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2f, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2f, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x73, 0x72, 0x63, 0x2d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x07,
	0x73, 0x72, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x9c, 0x01, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x40, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63,
	0x2e, 0x53, 0x6f, 0x6e, 0x69, 0x63, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x42, 0x82, 0x41, 0x3f, 0x2f,
	0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x66, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x79, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x55,
	0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x3f, 0x82, 0x41, 0x3c, 0x2f, 0x73, 0x6f,
	0x6e, 0x69, 0x63, 0x2d, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2f, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2f, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x74, 0x74, 0x6c, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x94,
	0x01, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3e, 0x2e,
	0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x53, 0x6f, 0x6e, 0x69, 0x63, 0x4d, 0x69, 0x72, 0x72, 0x6f,
	0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x40, 0x82,
	0x41, 0x3d, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x6f, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x0c, 0x44, 0x49, 0x52, 0x45, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x58, 0x10, 0x01, 0x1a, 0x05, 0x82, 0x41, 0x02, 0x52, 0x58,
	0x12, 0x17, 0x0a, 0x0c, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x58,
	0x10, 0x02, 0x1a, 0x05, 0x82, 0x41, 0x02, 0x54, 0x58, 0x12, 0x1b, 0x0a, 0x0e, 0x44, 0x49, 0x52,
	0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x42, 0x4f, 0x54, 0x48, 0x10, 0x03, 0x1a, 0x07, 0x82,
	0x41, 0x04, 0x42, 0x4f, 0x54, 0x48, 0x22, 0x5a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54,
	0x10, 0x00, 0x12, 0x1c, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x10, 0x01, 0x1a, 0x09, 0x82, 0x41, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x20, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x69, 0x6e, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x10, 0x02, 0x1a, 0x0b, 0x82, 0x41, 0x08, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x22, 0x4a, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x09, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x53, 0x50, 0x41, 0x4e, 0x10, 0x01, 0x1a, 0x07, 0x82, 0x41, 0x04, 0x53, 0x50,
	0x41, 0x4e, 0x12, 0x1a, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x52, 0x53, 0x50, 0x41,
	0x4e, 0x10, 0x02, 0x1a, 0x09, 0x82, 0x41, 0x06, 0x45, 0x52, 0x53, 0x50, 0x41, 0x4e, 0x1a, 0xd7,
	0x01, 0x0a, 0x14, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x54, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x40, 0x82, 0x41, 0x3d, 0x2f, 0x73, 0x6f, 0x6e, 0x69, 0x63,
	0x2d, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f,
	0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d,
	0x69, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x69,
	0x73, 0x74, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x69, 0x0a,
	0x13, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x73, 0x6f, 0x6e,
	0x69, 0x63, 0x2e, 0x53, 0x6f, 0x6e, 0x69, 0x63, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x11, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x1b, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e,
	0x6e, 0x6f, 0x63, 0x73, 0x79, 0x73, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x5a, 0x07, 0x2e, 0x3b,
	0x73, 0x6f, 0x6e, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sonic_mirror_session_proto_rawDescOnce sync.Once
	file_sonic_mirror_session_proto_rawDescData = file_sonic_mirror_session_proto_rawDesc
)

func file_sonic_mirror_session_proto_rawDescGZIP() []byte {
	file_sonic_mirror_session_proto_rawDescOnce.Do(func() {
		file_sonic_mirror_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_sonic_mirror_session_proto_rawDescData)
	})
	return file_sonic_mirror_session_proto_rawDescData
}

var file_sonic_mirror_session_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_sonic_mirror_session_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sonic_mirror_session_proto_goTypes = []interface{}{
	(SonicMirrorSession_MirrorSession_MirrorSessionList_Direction)(0), // 0: sonic.SonicMirrorSession.MirrorSession.MirrorSessionList.Direction
	(SonicMirrorSession_MirrorSession_MirrorSessionList_Status)(0),    // 1: sonic.SonicMirrorSession.MirrorSession.MirrorSessionList.Status
	(SonicMirrorSession_MirrorSession_MirrorSessionList_Type)(0),      // 2: sonic.SonicMirrorSession.MirrorSession.MirrorSessionList.Type
	(*SonicMirrorSession)(nil),                                        // 3: sonic.SonicMirrorSession
	(*SonicMirrorSession_MirrorSession)(nil),                          // 4: sonic.SonicMirrorSession.MirrorSession
	(*SonicMirrorSession_MirrorSession_MirrorSessionList)(nil),        // 5: sonic.SonicMirrorSession.MirrorSession.MirrorSessionList
	(*SonicMirrorSession_MirrorSession_MirrorSessionListKey)(nil),     // 6: sonic.SonicMirrorSession.MirrorSession.MirrorSessionListKey
	(*ywrapper.UintValue)(nil),                                        // 7: ywrapper.UintValue
	(*ywrapper.StringValue)(nil),                                      // 8: ywrapper.StringValue
}
var file_sonic_mirror_session_proto_depIdxs = []int32{
	4,  // 0: sonic.SonicMirrorSession.mirror_session:type_name -> sonic.SonicMirrorSession.MirrorSession
	6,  // 1: sonic.SonicMirrorSession.MirrorSession.mirror_session_list:type_name -> sonic.SonicMirrorSession.MirrorSession.MirrorSessionListKey
	0,  // 2: sonic.SonicMirrorSession.MirrorSession.MirrorSessionList.direction:type_name -> sonic.SonicMirrorSession.MirrorSession.MirrorSessionList.Direction
	7,  // 3: sonic.SonicMirrorSession.MirrorSession.MirrorSessionList.dscp:type_name -> ywrapper.UintValue
	8,  // 4: sonic.SonicMirrorSession.MirrorSession.MirrorSessionList.dst_ip:type_name -> ywrapper.StringValue
	8,  // 5: sonic.SonicMirrorSession.MirrorSession.MirrorSessionList.dst_port:type_name -> ywrapper.StringValue
	7,  // 6: sonic.SonicMirrorSession.MirrorSession.MirrorSessionList.gre_type:type_name -> ywrapper.UintValue
	8,  // 7: sonic.SonicMirrorSession.MirrorSession.MirrorSessionList.policer:type_name -> ywrapper.StringValue
	7,  // 8: sonic.SonicMirrorSession.MirrorSession.MirrorSessionList.queue:type_name -> ywrapper.UintValue
	8,  // 9: sonic.SonicMirrorSession.MirrorSession.MirrorSessionList.src_ip:type_name -> ywrapper.StringValue
	8,  // 10: sonic.SonicMirrorSession.MirrorSession.MirrorSessionList.src_port:type_name -> ywrapper.StringValue
	1,  // 11: sonic.SonicMirrorSession.MirrorSession.MirrorSessionList.status:type_name -> sonic.SonicMirrorSession.MirrorSession.MirrorSessionList.Status
	7,  // 12: sonic.SonicMirrorSession.MirrorSession.MirrorSessionList.ttl:type_name -> ywrapper.UintValue
	2,  // 13: sonic.SonicMirrorSession.MirrorSession.MirrorSessionList.type:type_name -> sonic.SonicMirrorSession.MirrorSession.MirrorSessionList.Type
	5,  // 14: sonic.SonicMirrorSession.MirrorSession.MirrorSessionListKey.mirror_session_list:type_name -> sonic.SonicMirrorSession.MirrorSession.MirrorSessionList
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_sonic_mirror_session_proto_init() }
func file_sonic_mirror_session_proto_init() {
	if File_sonic_mirror_session_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sonic_mirror_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicMirrorSession); i {
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
		file_sonic_mirror_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicMirrorSession_MirrorSession); i {
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
		file_sonic_mirror_session_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicMirrorSession_MirrorSession_MirrorSessionList); i {
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
		file_sonic_mirror_session_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SonicMirrorSession_MirrorSession_MirrorSessionListKey); i {
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
			RawDescriptor: file_sonic_mirror_session_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sonic_mirror_session_proto_goTypes,
		DependencyIndexes: file_sonic_mirror_session_proto_depIdxs,
		EnumInfos:         file_sonic_mirror_session_proto_enumTypes,
		MessageInfos:      file_sonic_mirror_session_proto_msgTypes,
	}.Build()
	File_sonic_mirror_session_proto = out.File
	file_sonic_mirror_session_proto_rawDesc = nil
	file_sonic_mirror_session_proto_goTypes = nil
	file_sonic_mirror_session_proto_depIdxs = nil
}
