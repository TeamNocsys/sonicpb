// openconfig.accton_ntp is generated by proto_generator as a protobuf
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
// source: accton_ntp.proto

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

type AcctonNtp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ntp *AcctonNtp_Ntp `protobuf:"bytes,1,opt,name=ntp,proto3" json:"ntp,omitempty"`
}

func (x *AcctonNtp) Reset() {
	*x = AcctonNtp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accton_ntp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcctonNtp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcctonNtp) ProtoMessage() {}

func (x *AcctonNtp) ProtoReflect() protoreflect.Message {
	mi := &file_accton_ntp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcctonNtp.ProtoReflect.Descriptor instead.
func (*AcctonNtp) Descriptor() ([]byte, []int) {
	return file_accton_ntp_proto_rawDescGZIP(), []int{0}
}

func (x *AcctonNtp) GetNtp() *AcctonNtp_Ntp {
	if x != nil {
		return x.Ntp
	}
	return nil
}

type AcctonNtp_Ntp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NtpList []*AcctonNtp_Ntp_NtpListKey `protobuf:"bytes,1,rep,name=ntp_list,json=ntpList,proto3" json:"ntp_list,omitempty"`
}

func (x *AcctonNtp_Ntp) Reset() {
	*x = AcctonNtp_Ntp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accton_ntp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcctonNtp_Ntp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcctonNtp_Ntp) ProtoMessage() {}

func (x *AcctonNtp_Ntp) ProtoReflect() protoreflect.Message {
	mi := &file_accton_ntp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcctonNtp_Ntp.ProtoReflect.Descriptor instead.
func (*AcctonNtp_Ntp) Descriptor() ([]byte, []int) {
	return file_accton_ntp_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AcctonNtp_Ntp) GetNtpList() []*AcctonNtp_Ntp_NtpListKey {
	if x != nil {
		return x.NtpList
	}
	return nil
}

type AcctonNtp_Ntp_NtpList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State *AcctonNtp_Ntp_NtpList_State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *AcctonNtp_Ntp_NtpList) Reset() {
	*x = AcctonNtp_Ntp_NtpList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accton_ntp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcctonNtp_Ntp_NtpList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcctonNtp_Ntp_NtpList) ProtoMessage() {}

func (x *AcctonNtp_Ntp_NtpList) ProtoReflect() protoreflect.Message {
	mi := &file_accton_ntp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcctonNtp_Ntp_NtpList.ProtoReflect.Descriptor instead.
func (*AcctonNtp_Ntp_NtpList) Descriptor() ([]byte, []int) {
	return file_accton_ntp_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *AcctonNtp_Ntp_NtpList) GetState() *AcctonNtp_Ntp_NtpList_State {
	if x != nil {
		return x.State
	}
	return nil
}

type AcctonNtp_Ntp_NtpListKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip      string                 `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	NtpList *AcctonNtp_Ntp_NtpList `protobuf:"bytes,2,opt,name=ntp_list,json=ntpList,proto3" json:"ntp_list,omitempty"`
}

func (x *AcctonNtp_Ntp_NtpListKey) Reset() {
	*x = AcctonNtp_Ntp_NtpListKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accton_ntp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcctonNtp_Ntp_NtpListKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcctonNtp_Ntp_NtpListKey) ProtoMessage() {}

func (x *AcctonNtp_Ntp_NtpListKey) ProtoReflect() protoreflect.Message {
	mi := &file_accton_ntp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcctonNtp_Ntp_NtpListKey.ProtoReflect.Descriptor instead.
func (*AcctonNtp_Ntp_NtpListKey) Descriptor() ([]byte, []int) {
	return file_accton_ntp_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *AcctonNtp_Ntp_NtpListKey) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *AcctonNtp_Ntp_NtpListKey) GetNtpList() *AcctonNtp_Ntp_NtpList {
	if x != nil {
		return x.NtpList
	}
	return nil
}

type AcctonNtp_Ntp_NtpList_State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Current *ywrapper.StringValue `protobuf:"bytes,1,opt,name=current,proto3" json:"current,omitempty"`
	Poll    *ywrapper.UintValue   `protobuf:"bytes,2,opt,name=poll,proto3" json:"poll,omitempty"`
}

func (x *AcctonNtp_Ntp_NtpList_State) Reset() {
	*x = AcctonNtp_Ntp_NtpList_State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accton_ntp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcctonNtp_Ntp_NtpList_State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcctonNtp_Ntp_NtpList_State) ProtoMessage() {}

func (x *AcctonNtp_Ntp_NtpList_State) ProtoReflect() protoreflect.Message {
	mi := &file_accton_ntp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcctonNtp_Ntp_NtpList_State.ProtoReflect.Descriptor instead.
func (*AcctonNtp_Ntp_NtpList_State) Descriptor() ([]byte, []int) {
	return file_accton_ntp_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

func (x *AcctonNtp_Ntp_NtpList_State) GetCurrent() *ywrapper.StringValue {
	if x != nil {
		return x.Current
	}
	return nil
}

func (x *AcctonNtp_Ntp_NtpList_State) GetPoll() *ywrapper.UintValue {
	if x != nil {
		return x.Poll
	}
	return nil
}

var File_accton_ntp_proto protoreflect.FileDescriptor

var file_accton_ntp_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x63, 0x63, 0x74, 0x6f, 0x6e, 0x5f, 0x6e, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x1a, 0x0e, 0x79, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x79, 0x65, 0x78, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x04, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x74, 0x6f, 0x6e,
	0x4e, 0x74, 0x70, 0x12, 0x3a, 0x0a, 0x03, 0x6e, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x74, 0x6f, 0x6e, 0x4e,
	0x74, 0x70, 0x2e, 0x4e, 0x74, 0x70, 0x42, 0x12, 0x82, 0x41, 0x0f, 0x2f, 0x61, 0x63, 0x63, 0x74,
	0x6f, 0x6e, 0x2d, 0x6e, 0x74, 0x70, 0x2f, 0x6e, 0x74, 0x70, 0x52, 0x03, 0x6e, 0x74, 0x70, 0x1a,
	0xf5, 0x03, 0x0a, 0x03, 0x4e, 0x74, 0x70, 0x12, 0x57, 0x0a, 0x08, 0x6e, 0x74, 0x70, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x6f, 0x6e, 0x69,
	0x63, 0x2e, 0x41, 0x63, 0x63, 0x74, 0x6f, 0x6e, 0x4e, 0x74, 0x70, 0x2e, 0x4e, 0x74, 0x70, 0x2e,
	0x4e, 0x74, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x42, 0x1b, 0x82, 0x41, 0x18, 0x2f,
	0x61, 0x63, 0x63, 0x74, 0x6f, 0x6e, 0x2d, 0x6e, 0x74, 0x70, 0x2f, 0x6e, 0x74, 0x70, 0x2f, 0x6e,
	0x74, 0x70, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x07, 0x6e, 0x74, 0x70, 0x4c, 0x69, 0x73, 0x74,
	0x1a, 0x9d, 0x02, 0x0a, 0x07, 0x4e, 0x74, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x5b, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x6f,
	0x6e, 0x69, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x74, 0x6f, 0x6e, 0x4e, 0x74, 0x70, 0x2e, 0x4e, 0x74,
	0x70, 0x2e, 0x4e, 0x74, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42,
	0x21, 0x82, 0x41, 0x1e, 0x2f, 0x61, 0x63, 0x63, 0x74, 0x6f, 0x6e, 0x2d, 0x6e, 0x74, 0x70, 0x2f,
	0x6e, 0x74, 0x70, 0x2f, 0x6e, 0x74, 0x70, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x1a, 0xb4, 0x01, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x5a, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x79, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x29, 0x82, 0x41, 0x26,
	0x2f, 0x61, 0x63, 0x63, 0x74, 0x6f, 0x6e, 0x2d, 0x6e, 0x74, 0x70, 0x2f, 0x6e, 0x74, 0x70, 0x2f,
	0x6e, 0x74, 0x70, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12,
	0x4f, 0x0a, 0x04, 0x70, 0x6f, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x79, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x26, 0x82, 0x41, 0x23, 0x2f, 0x61, 0x63, 0x63, 0x74, 0x6f, 0x6e, 0x2d, 0x6e,
	0x74, 0x70, 0x2f, 0x6e, 0x74, 0x70, 0x2f, 0x6e, 0x74, 0x70, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x6f, 0x6c, 0x6c, 0x52, 0x04, 0x70, 0x6f, 0x6c, 0x6c,
	0x1a, 0x75, 0x0a, 0x0a, 0x4e, 0x74, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x2e,
	0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0x82, 0x41, 0x1b, 0x2f,
	0x61, 0x63, 0x63, 0x74, 0x6f, 0x6e, 0x2d, 0x6e, 0x74, 0x70, 0x2f, 0x6e, 0x74, 0x70, 0x2f, 0x6e,
	0x74, 0x70, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x69, 0x70, 0x52, 0x02, 0x69, 0x70, 0x12, 0x37,
	0x0a, 0x08, 0x6e, 0x74, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x74, 0x6f, 0x6e, 0x4e,
	0x74, 0x70, 0x2e, 0x4e, 0x74, 0x70, 0x2e, 0x4e, 0x74, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x07,
	0x6e, 0x74, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x1b, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x63, 0x63, 0x74, 0x6f, 0x6e, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x5a, 0x07, 0x2e, 0x3b, 0x73,
	0x6f, 0x6e, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accton_ntp_proto_rawDescOnce sync.Once
	file_accton_ntp_proto_rawDescData = file_accton_ntp_proto_rawDesc
)

func file_accton_ntp_proto_rawDescGZIP() []byte {
	file_accton_ntp_proto_rawDescOnce.Do(func() {
		file_accton_ntp_proto_rawDescData = protoimpl.X.CompressGZIP(file_accton_ntp_proto_rawDescData)
	})
	return file_accton_ntp_proto_rawDescData
}

var file_accton_ntp_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_accton_ntp_proto_goTypes = []interface{}{
	(*AcctonNtp)(nil),                   // 0: sonic.AcctonNtp
	(*AcctonNtp_Ntp)(nil),               // 1: sonic.AcctonNtp.Ntp
	(*AcctonNtp_Ntp_NtpList)(nil),       // 2: sonic.AcctonNtp.Ntp.NtpList
	(*AcctonNtp_Ntp_NtpListKey)(nil),    // 3: sonic.AcctonNtp.Ntp.NtpListKey
	(*AcctonNtp_Ntp_NtpList_State)(nil), // 4: sonic.AcctonNtp.Ntp.NtpList.State
	(*ywrapper.StringValue)(nil),        // 5: ywrapper.StringValue
	(*ywrapper.UintValue)(nil),          // 6: ywrapper.UintValue
}
var file_accton_ntp_proto_depIdxs = []int32{
	1, // 0: sonic.AcctonNtp.ntp:type_name -> sonic.AcctonNtp.Ntp
	3, // 1: sonic.AcctonNtp.Ntp.ntp_list:type_name -> sonic.AcctonNtp.Ntp.NtpListKey
	4, // 2: sonic.AcctonNtp.Ntp.NtpList.state:type_name -> sonic.AcctonNtp.Ntp.NtpList.State
	2, // 3: sonic.AcctonNtp.Ntp.NtpListKey.ntp_list:type_name -> sonic.AcctonNtp.Ntp.NtpList
	5, // 4: sonic.AcctonNtp.Ntp.NtpList.State.current:type_name -> ywrapper.StringValue
	6, // 5: sonic.AcctonNtp.Ntp.NtpList.State.poll:type_name -> ywrapper.UintValue
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_accton_ntp_proto_init() }
func file_accton_ntp_proto_init() {
	if File_accton_ntp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_accton_ntp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcctonNtp); i {
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
		file_accton_ntp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcctonNtp_Ntp); i {
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
		file_accton_ntp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcctonNtp_Ntp_NtpList); i {
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
		file_accton_ntp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcctonNtp_Ntp_NtpListKey); i {
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
		file_accton_ntp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcctonNtp_Ntp_NtpList_State); i {
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
			RawDescriptor: file_accton_ntp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_accton_ntp_proto_goTypes,
		DependencyIndexes: file_accton_ntp_proto_depIdxs,
		MessageInfos:      file_accton_ntp_proto_msgTypes,
	}.Build()
	File_accton_ntp_proto = out.File
	file_accton_ntp_proto_rawDesc = nil
	file_accton_ntp_proto_goTypes = nil
	file_accton_ntp_proto_depIdxs = nil
}