// openconfig.nocsys_mdns is generated by proto_generator as a protobuf
// representation of a YANG schema.
//
// Input schema modules:
//  - ../api/yang/sonic/nocsys-platform-types.yang
//  - ../api/yang/sonic/nocsys-acl.yang
//  - ../api/yang/sonic/nocsys-todo.yang
//  - ../api/yang/sonic/nocsys-mirror-session.yang
//  - ../api/yang/sonic/nocsys-extension.yang
//  - ../api/yang/sonic/nocsys-vlan.yang
//  - ../api/yang/sonic/nocsys-lldp.yang
//  - ../api/yang/sonic/nocsys-port.yang
//  - ../api/yang/sonic/nocsys-ntp.yang
//  - ../api/yang/sonic/nocsys-vxlan.yang
//  - ../api/yang/sonic/third_party/openconfig/openconfig-types.yang
//  - ../api/yang/sonic/third_party/openconfig/openconfig-extensions.yang
//  - ../api/yang/sonic/third_party/ietf/ietf-yang-types.yang
//  - ../api/yang/sonic/third_party/ietf/iana-if-type.yang
//  - ../api/yang/sonic/third_party/ietf/ietf-inet-types.yang
//  - ../api/yang/sonic/third_party/ietf/ietf-interfaces.yang
//  - ../api/yang/sonic/nocsys-types.yang
//  - ../api/yang/sonic/nocsys-fdb.yang
//  - ../api/yang/sonic/nocsys-neighbor.yang
//  - ../api/yang/sonic/nocsys-vrf.yang
//  - ../api/yang/sonic/nocsys-route.yang
//  - ../api/yang/sonic/nocsys-loopback-interface.yang
//  - ../api/yang/sonic/nocsys-mdns.yang
//  - ../api/yang/sonic/nocsys-interface.yang
//  - ../api/yang/sonic/nocsys-system.yang
//  - ../api/yang/sonic/nocsys-platform.yang
//  - ../api/yang/sonic/nocsys-portchannel.yang

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: nocsys_mdns.proto

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

type NocsysMdns struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mdns *NocsysMdns_Mdns `protobuf:"bytes,1,opt,name=mdns,proto3" json:"mdns,omitempty"`
}

func (x *NocsysMdns) Reset() {
	*x = NocsysMdns{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nocsys_mdns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NocsysMdns) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NocsysMdns) ProtoMessage() {}

func (x *NocsysMdns) ProtoReflect() protoreflect.Message {
	mi := &file_nocsys_mdns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NocsysMdns.ProtoReflect.Descriptor instead.
func (*NocsysMdns) Descriptor() ([]byte, []int) {
	return file_nocsys_mdns_proto_rawDescGZIP(), []int{0}
}

func (x *NocsysMdns) GetMdns() *NocsysMdns_Mdns {
	if x != nil {
		return x.Mdns
	}
	return nil
}

type NocsysMdns_Mdns struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MdnsList []*NocsysMdns_Mdns_MdnsListKey `protobuf:"bytes,1,rep,name=mdns_list,json=mdnsList,proto3" json:"mdns_list,omitempty"`
}

func (x *NocsysMdns_Mdns) Reset() {
	*x = NocsysMdns_Mdns{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nocsys_mdns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NocsysMdns_Mdns) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NocsysMdns_Mdns) ProtoMessage() {}

func (x *NocsysMdns_Mdns) ProtoReflect() protoreflect.Message {
	mi := &file_nocsys_mdns_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NocsysMdns_Mdns.ProtoReflect.Descriptor instead.
func (*NocsysMdns_Mdns) Descriptor() ([]byte, []int) {
	return file_nocsys_mdns_proto_rawDescGZIP(), []int{0, 0}
}

func (x *NocsysMdns_Mdns) GetMdnsList() []*NocsysMdns_Mdns_MdnsListKey {
	if x != nil {
		return x.MdnsList
	}
	return nil
}

type NocsysMdns_Mdns_MdnsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname *ywrapper.StringValue `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *NocsysMdns_Mdns_MdnsList) Reset() {
	*x = NocsysMdns_Mdns_MdnsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nocsys_mdns_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NocsysMdns_Mdns_MdnsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NocsysMdns_Mdns_MdnsList) ProtoMessage() {}

func (x *NocsysMdns_Mdns_MdnsList) ProtoReflect() protoreflect.Message {
	mi := &file_nocsys_mdns_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NocsysMdns_Mdns_MdnsList.ProtoReflect.Descriptor instead.
func (*NocsysMdns_Mdns_MdnsList) Descriptor() ([]byte, []int) {
	return file_nocsys_mdns_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *NocsysMdns_Mdns_MdnsList) GetHostname() *ywrapper.StringValue {
	if x != nil {
		return x.Hostname
	}
	return nil
}

type NocsysMdns_Mdns_MdnsListKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpPrefix string                    `protobuf:"bytes,1,opt,name=ip_prefix,json=ipPrefix,proto3" json:"ip_prefix,omitempty"`
	MdnsList *NocsysMdns_Mdns_MdnsList `protobuf:"bytes,2,opt,name=mdns_list,json=mdnsList,proto3" json:"mdns_list,omitempty"`
}

func (x *NocsysMdns_Mdns_MdnsListKey) Reset() {
	*x = NocsysMdns_Mdns_MdnsListKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nocsys_mdns_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NocsysMdns_Mdns_MdnsListKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NocsysMdns_Mdns_MdnsListKey) ProtoMessage() {}

func (x *NocsysMdns_Mdns_MdnsListKey) ProtoReflect() protoreflect.Message {
	mi := &file_nocsys_mdns_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NocsysMdns_Mdns_MdnsListKey.ProtoReflect.Descriptor instead.
func (*NocsysMdns_Mdns_MdnsListKey) Descriptor() ([]byte, []int) {
	return file_nocsys_mdns_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *NocsysMdns_Mdns_MdnsListKey) GetIpPrefix() string {
	if x != nil {
		return x.IpPrefix
	}
	return ""
}

func (x *NocsysMdns_Mdns_MdnsListKey) GetMdnsList() *NocsysMdns_Mdns_MdnsList {
	if x != nil {
		return x.MdnsList
	}
	return nil
}

var File_nocsys_mdns_proto protoreflect.FileDescriptor

var file_nocsys_mdns_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6e, 0x6f, 0x63, 0x73, 0x79, 0x73, 0x5f, 0x6d, 0x64, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x1a, 0x0e, 0x79, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x79, 0x65, 0x78, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x03, 0x0a, 0x0a, 0x4e, 0x6f, 0x63, 0x73, 0x79,
	0x73, 0x4d, 0x64, 0x6e, 0x73, 0x12, 0x40, 0x0a, 0x04, 0x6d, 0x64, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x4e, 0x6f, 0x63, 0x73,
	0x79, 0x73, 0x4d, 0x64, 0x6e, 0x73, 0x2e, 0x4d, 0x64, 0x6e, 0x73, 0x42, 0x14, 0x82, 0x41, 0x11,
	0x2f, 0x6e, 0x6f, 0x63, 0x73, 0x79, 0x73, 0x2d, 0x6d, 0x64, 0x6e, 0x73, 0x2f, 0x6d, 0x64, 0x6e,
	0x73, 0x52, 0x04, 0x6d, 0x64, 0x6e, 0x73, 0x1a, 0xe4, 0x02, 0x0a, 0x04, 0x4d, 0x64, 0x6e, 0x73,
	0x12, 0x5f, 0x0a, 0x09, 0x6d, 0x64, 0x6e, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x4e, 0x6f, 0x63, 0x73,
	0x79, 0x73, 0x4d, 0x64, 0x6e, 0x73, 0x2e, 0x4d, 0x64, 0x6e, 0x73, 0x2e, 0x4d, 0x64, 0x6e, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x42, 0x1e, 0x82, 0x41, 0x1b, 0x2f, 0x6e, 0x6f, 0x63,
	0x73, 0x79, 0x73, 0x2d, 0x6d, 0x64, 0x6e, 0x73, 0x2f, 0x6d, 0x64, 0x6e, 0x73, 0x2f, 0x6d, 0x64,
	0x6e, 0x73, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x08, 0x6d, 0x64, 0x6e, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x1a, 0x66, 0x0a, 0x08, 0x4d, 0x64, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x5a, 0x0a,
	0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x79, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x27, 0x82, 0x41, 0x24, 0x2f, 0x6e, 0x6f, 0x63, 0x73,
	0x79, 0x73, 0x2d, 0x6d, 0x64, 0x6e, 0x73, 0x2f, 0x6d, 0x64, 0x6e, 0x73, 0x2f, 0x6d, 0x64, 0x6e,
	0x73, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52,
	0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x92, 0x01, 0x0a, 0x0b, 0x4d, 0x64,
	0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x45, 0x0a, 0x09, 0x69, 0x70, 0x5f,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0x82, 0x41,
	0x25, 0x2f, 0x6e, 0x6f, 0x63, 0x73, 0x79, 0x73, 0x2d, 0x6d, 0x64, 0x6e, 0x73, 0x2f, 0x6d, 0x64,
	0x6e, 0x73, 0x2f, 0x6d, 0x64, 0x6e, 0x73, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x69, 0x70, 0x2d,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x08, 0x69, 0x70, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x12, 0x3c, 0x0a, 0x09, 0x6d, 0x64, 0x6e, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x4e, 0x6f, 0x63, 0x73,
	0x79, 0x73, 0x4d, 0x64, 0x6e, 0x73, 0x2e, 0x4d, 0x64, 0x6e, 0x73, 0x2e, 0x4d, 0x64, 0x6e, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x08, 0x6d, 0x64, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x1b,
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x6f, 0x63, 0x73, 0x79, 0x73, 0x2e, 0x73, 0x6f, 0x6e,
	0x69, 0x63, 0x5a, 0x07, 0x2e, 0x3b, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_nocsys_mdns_proto_rawDescOnce sync.Once
	file_nocsys_mdns_proto_rawDescData = file_nocsys_mdns_proto_rawDesc
)

func file_nocsys_mdns_proto_rawDescGZIP() []byte {
	file_nocsys_mdns_proto_rawDescOnce.Do(func() {
		file_nocsys_mdns_proto_rawDescData = protoimpl.X.CompressGZIP(file_nocsys_mdns_proto_rawDescData)
	})
	return file_nocsys_mdns_proto_rawDescData
}

var file_nocsys_mdns_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nocsys_mdns_proto_goTypes = []interface{}{
	(*NocsysMdns)(nil),                  // 0: sonic.NocsysMdns
	(*NocsysMdns_Mdns)(nil),             // 1: sonic.NocsysMdns.Mdns
	(*NocsysMdns_Mdns_MdnsList)(nil),    // 2: sonic.NocsysMdns.Mdns.MdnsList
	(*NocsysMdns_Mdns_MdnsListKey)(nil), // 3: sonic.NocsysMdns.Mdns.MdnsListKey
	(*ywrapper.StringValue)(nil),        // 4: ywrapper.StringValue
}
var file_nocsys_mdns_proto_depIdxs = []int32{
	1, // 0: sonic.NocsysMdns.mdns:type_name -> sonic.NocsysMdns.Mdns
	3, // 1: sonic.NocsysMdns.Mdns.mdns_list:type_name -> sonic.NocsysMdns.Mdns.MdnsListKey
	4, // 2: sonic.NocsysMdns.Mdns.MdnsList.hostname:type_name -> ywrapper.StringValue
	2, // 3: sonic.NocsysMdns.Mdns.MdnsListKey.mdns_list:type_name -> sonic.NocsysMdns.Mdns.MdnsList
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_nocsys_mdns_proto_init() }
func file_nocsys_mdns_proto_init() {
	if File_nocsys_mdns_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nocsys_mdns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NocsysMdns); i {
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
		file_nocsys_mdns_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NocsysMdns_Mdns); i {
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
		file_nocsys_mdns_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NocsysMdns_Mdns_MdnsList); i {
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
		file_nocsys_mdns_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NocsysMdns_Mdns_MdnsListKey); i {
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
			RawDescriptor: file_nocsys_mdns_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nocsys_mdns_proto_goTypes,
		DependencyIndexes: file_nocsys_mdns_proto_depIdxs,
		MessageInfos:      file_nocsys_mdns_proto_msgTypes,
	}.Build()
	File_nocsys_mdns_proto = out.File
	file_nocsys_mdns_proto_rawDesc = nil
	file_nocsys_mdns_proto_goTypes = nil
	file_nocsys_mdns_proto_depIdxs = nil
}