// openconfig.sonic_portchannel_enums is generated by proto_generator as a protobuf
// representation of a YANG schema.
//
// Input schema modules:
//  - ../api/yang/sonic/sonic-portchannel.yang
// Include paths:
//   - ../api/yang/sonic:../api/yang/public/third_party/...

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.3
// source: sonic_portchannel_enums.proto

package openconfig

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/openconfig/ygot/proto/yext"
	_ "github.com/openconfig/ygot/proto/ywrapper"
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

// SonicPortAdminStatus represents an enumerated type generated for the YANG enumerated type admin-status.
type SonicPortAdminStatus int32

const (
	SonicPortAdminStatus_SONICPORTADMINSTATUS_UNSET SonicPortAdminStatus = 0
	SonicPortAdminStatus_SONICPORTADMINSTATUS_up    SonicPortAdminStatus = 1
	SonicPortAdminStatus_SONICPORTADMINSTATUS_down  SonicPortAdminStatus = 2
)

// Enum value maps for SonicPortAdminStatus.
var (
	SonicPortAdminStatus_name = map[int32]string{
		0: "SONICPORTADMINSTATUS_UNSET",
		1: "SONICPORTADMINSTATUS_up",
		2: "SONICPORTADMINSTATUS_down",
	}
	SonicPortAdminStatus_value = map[string]int32{
		"SONICPORTADMINSTATUS_UNSET": 0,
		"SONICPORTADMINSTATUS_up":    1,
		"SONICPORTADMINSTATUS_down":  2,
	}
)

func (x SonicPortAdminStatus) Enum() *SonicPortAdminStatus {
	p := new(SonicPortAdminStatus)
	*p = x
	return p
}

func (x SonicPortAdminStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SonicPortAdminStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_sonic_portchannel_enums_proto_enumTypes[0].Descriptor()
}

func (SonicPortAdminStatus) Type() protoreflect.EnumType {
	return &file_sonic_portchannel_enums_proto_enumTypes[0]
}

func (x SonicPortAdminStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SonicPortAdminStatus.Descriptor instead.
func (SonicPortAdminStatus) EnumDescriptor() ([]byte, []int) {
	return file_sonic_portchannel_enums_proto_rawDescGZIP(), []int{0}
}

// SonicPortchannelAdminStatus represents an enumerated type generated for the YANG enumerated type admin-status.
type SonicPortchannelAdminStatus int32

const (
	SonicPortchannelAdminStatus_SONICPORTCHANNELADMINSTATUS_UNSET SonicPortchannelAdminStatus = 0
	SonicPortchannelAdminStatus_SONICPORTCHANNELADMINSTATUS_up    SonicPortchannelAdminStatus = 1
	SonicPortchannelAdminStatus_SONICPORTCHANNELADMINSTATUS_down  SonicPortchannelAdminStatus = 2
)

// Enum value maps for SonicPortchannelAdminStatus.
var (
	SonicPortchannelAdminStatus_name = map[int32]string{
		0: "SONICPORTCHANNELADMINSTATUS_UNSET",
		1: "SONICPORTCHANNELADMINSTATUS_up",
		2: "SONICPORTCHANNELADMINSTATUS_down",
	}
	SonicPortchannelAdminStatus_value = map[string]int32{
		"SONICPORTCHANNELADMINSTATUS_UNSET": 0,
		"SONICPORTCHANNELADMINSTATUS_up":    1,
		"SONICPORTCHANNELADMINSTATUS_down":  2,
	}
)

func (x SonicPortchannelAdminStatus) Enum() *SonicPortchannelAdminStatus {
	p := new(SonicPortchannelAdminStatus)
	*p = x
	return p
}

func (x SonicPortchannelAdminStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SonicPortchannelAdminStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_sonic_portchannel_enums_proto_enumTypes[1].Descriptor()
}

func (SonicPortchannelAdminStatus) Type() protoreflect.EnumType {
	return &file_sonic_portchannel_enums_proto_enumTypes[1]
}

func (x SonicPortchannelAdminStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SonicPortchannelAdminStatus.Descriptor instead.
func (SonicPortchannelAdminStatus) EnumDescriptor() ([]byte, []int) {
	return file_sonic_portchannel_enums_proto_rawDescGZIP(), []int{1}
}

var File_sonic_portchannel_enums_proto protoreflect.FileDescriptor

var file_sonic_portchannel_enums_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x0e, 0x79, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x79, 0x65, 0x78,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x82, 0x01, 0x0a, 0x14, 0x53, 0x6f, 0x6e, 0x69,
	0x63, 0x50, 0x6f, 0x72, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x4f, 0x4e, 0x49, 0x43, 0x50, 0x4f, 0x52, 0x54, 0x41, 0x44, 0x4d,
	0x49, 0x4e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00,
	0x12, 0x22, 0x0a, 0x17, 0x53, 0x4f, 0x4e, 0x49, 0x43, 0x50, 0x4f, 0x52, 0x54, 0x41, 0x44, 0x4d,
	0x49, 0x4e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x75, 0x70, 0x10, 0x01, 0x1a, 0x05, 0x82,
	0x41, 0x02, 0x75, 0x70, 0x12, 0x26, 0x0a, 0x19, 0x53, 0x4f, 0x4e, 0x49, 0x43, 0x50, 0x4f, 0x52,
	0x54, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x64, 0x6f, 0x77,
	0x6e, 0x10, 0x02, 0x1a, 0x07, 0x82, 0x41, 0x04, 0x64, 0x6f, 0x77, 0x6e, 0x2a, 0x9e, 0x01, 0x0a,
	0x1b, 0x53, 0x6f, 0x6e, 0x69, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x21,
	0x53, 0x4f, 0x4e, 0x49, 0x43, 0x50, 0x4f, 0x52, 0x54, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c,
	0x41, 0x44, 0x4d, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x45,
	0x54, 0x10, 0x00, 0x12, 0x29, 0x0a, 0x1e, 0x53, 0x4f, 0x4e, 0x49, 0x43, 0x50, 0x4f, 0x52, 0x54,
	0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x75, 0x70, 0x10, 0x01, 0x1a, 0x05, 0x82, 0x41, 0x02, 0x75, 0x70, 0x12, 0x2d,
	0x0a, 0x20, 0x53, 0x4f, 0x4e, 0x49, 0x43, 0x50, 0x4f, 0x52, 0x54, 0x43, 0x48, 0x41, 0x4e, 0x4e,
	0x45, 0x4c, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x64, 0x6f,
	0x77, 0x6e, 0x10, 0x02, 0x1a, 0x07, 0x82, 0x41, 0x04, 0x64, 0x6f, 0x77, 0x6e, 0x42, 0x17, 0x0a,
	0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x6f, 0x63, 0x73, 0x79, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sonic_portchannel_enums_proto_rawDescOnce sync.Once
	file_sonic_portchannel_enums_proto_rawDescData = file_sonic_portchannel_enums_proto_rawDesc
)

func file_sonic_portchannel_enums_proto_rawDescGZIP() []byte {
	file_sonic_portchannel_enums_proto_rawDescOnce.Do(func() {
		file_sonic_portchannel_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_sonic_portchannel_enums_proto_rawDescData)
	})
	return file_sonic_portchannel_enums_proto_rawDescData
}

var file_sonic_portchannel_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_sonic_portchannel_enums_proto_goTypes = []interface{}{
	(SonicPortAdminStatus)(0),        // 0: openconfig.SonicPortAdminStatus
	(SonicPortchannelAdminStatus)(0), // 1: openconfig.SonicPortchannelAdminStatus
}
var file_sonic_portchannel_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sonic_portchannel_enums_proto_init() }
func file_sonic_portchannel_enums_proto_init() {
	if File_sonic_portchannel_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sonic_portchannel_enums_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sonic_portchannel_enums_proto_goTypes,
		DependencyIndexes: file_sonic_portchannel_enums_proto_depIdxs,
		EnumInfos:         file_sonic_portchannel_enums_proto_enumTypes,
	}.Build()
	File_sonic_portchannel_enums_proto = out.File
	file_sonic_portchannel_enums_proto_rawDesc = nil
	file_sonic_portchannel_enums_proto_goTypes = nil
	file_sonic_portchannel_enums_proto_depIdxs = nil
}
