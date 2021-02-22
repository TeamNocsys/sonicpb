// openconfig.nocsys_todo is generated by proto_generator as a protobuf
// representation of a YANG schema.
//
// Input schema modules:
//  - ../api/yang/sonic/nocsys-portchannel.yang
//  - ../api/yang/sonic/nocsys-port.yang
//  - ../api/yang/sonic/nocsys-todo.yang
//  - ../api/yang/sonic/nocsys-route.yang
//  - ../api/yang/sonic/nocsys-platform.yang
//  - ../api/yang/sonic/nocsys-fdb.yang
//  - ../api/yang/sonic/nocsys-vrf.yang
//  - ../api/yang/sonic/nocsys-neighbor.yang
//  - ../api/yang/sonic/nocsys-lldp.yang
//  - ../api/yang/sonic/nocsys-vxlan.yang
//  - ../api/yang/sonic/third_party/ietf/ietf-interfaces.yang
//  - ../api/yang/sonic/third_party/ietf/ietf-yang-types.yang
//  - ../api/yang/sonic/third_party/ietf/ietf-inet-types.yang
//  - ../api/yang/sonic/third_party/ietf/iana-if-type.yang
//  - ../api/yang/sonic/nocsys-mirror-session.yang
//  - ../api/yang/sonic/nocsys-interface.yang
//  - ../api/yang/sonic/nocsys-extension.yang
//  - ../api/yang/sonic/nocsys-platform-types.yang
//  - ../api/yang/sonic/nocsys-types.yang
//  - ../api/yang/sonic/nocsys-vlan.yang
//  - ../api/yang/sonic/nocsys-acl.yang
//  - ../api/yang/sonic/nocsys-ntp.yang
//  - ../api/yang/sonic/nocsys-loopback-interface.yang

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: nocsys_todo.proto

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

type NocsysTodo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *NocsysTodo_Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *NocsysTodo) Reset() {
	*x = NocsysTodo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nocsys_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NocsysTodo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NocsysTodo) ProtoMessage() {}

func (x *NocsysTodo) ProtoReflect() protoreflect.Message {
	mi := &file_nocsys_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NocsysTodo.ProtoReflect.Descriptor instead.
func (*NocsysTodo) Descriptor() ([]byte, []int) {
	return file_nocsys_todo_proto_rawDescGZIP(), []int{0}
}

func (x *NocsysTodo) GetTodo() *NocsysTodo_Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type NocsysTodo_Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TodoList []*NocsysTodo_Todo_TodoListKey `protobuf:"bytes,1,rep,name=todo_list,json=todoList,proto3" json:"todo_list,omitempty"`
}

func (x *NocsysTodo_Todo) Reset() {
	*x = NocsysTodo_Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nocsys_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NocsysTodo_Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NocsysTodo_Todo) ProtoMessage() {}

func (x *NocsysTodo_Todo) ProtoReflect() protoreflect.Message {
	mi := &file_nocsys_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NocsysTodo_Todo.ProtoReflect.Descriptor instead.
func (*NocsysTodo_Todo) Descriptor() ([]byte, []int) {
	return file_nocsys_todo_proto_rawDescGZIP(), []int{0, 0}
}

func (x *NocsysTodo_Todo) GetTodoList() []*NocsysTodo_Todo_TodoListKey {
	if x != nil {
		return x.TodoList
	}
	return nil
}

type NocsysTodo_Todo_TodoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Json *ywrapper.StringValue `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
}

func (x *NocsysTodo_Todo_TodoList) Reset() {
	*x = NocsysTodo_Todo_TodoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nocsys_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NocsysTodo_Todo_TodoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NocsysTodo_Todo_TodoList) ProtoMessage() {}

func (x *NocsysTodo_Todo_TodoList) ProtoReflect() protoreflect.Message {
	mi := &file_nocsys_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NocsysTodo_Todo_TodoList.ProtoReflect.Descriptor instead.
func (*NocsysTodo_Todo_TodoList) Descriptor() ([]byte, []int) {
	return file_nocsys_todo_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *NocsysTodo_Todo_TodoList) GetJson() *ywrapper.StringValue {
	if x != nil {
		return x.Json
	}
	return nil
}

type NocsysTodo_Todo_TodoListKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TodoList *NocsysTodo_Todo_TodoList `protobuf:"bytes,2,opt,name=todo_list,json=todoList,proto3" json:"todo_list,omitempty"`
}

func (x *NocsysTodo_Todo_TodoListKey) Reset() {
	*x = NocsysTodo_Todo_TodoListKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nocsys_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NocsysTodo_Todo_TodoListKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NocsysTodo_Todo_TodoListKey) ProtoMessage() {}

func (x *NocsysTodo_Todo_TodoListKey) ProtoReflect() protoreflect.Message {
	mi := &file_nocsys_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NocsysTodo_Todo_TodoListKey.ProtoReflect.Descriptor instead.
func (*NocsysTodo_Todo_TodoListKey) Descriptor() ([]byte, []int) {
	return file_nocsys_todo_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *NocsysTodo_Todo_TodoListKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NocsysTodo_Todo_TodoListKey) GetTodoList() *NocsysTodo_Todo_TodoList {
	if x != nil {
		return x.TodoList
	}
	return nil
}

var File_nocsys_todo_proto protoreflect.FileDescriptor

var file_nocsys_todo_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6e, 0x6f, 0x63, 0x73, 0x79, 0x73, 0x5f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x1a, 0x0e, 0x79, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x79, 0x65, 0x78, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x03, 0x0a, 0x0a, 0x4e, 0x6f, 0x63, 0x73, 0x79,
	0x73, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x40, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x4e, 0x6f, 0x63, 0x73,
	0x79, 0x73, 0x54, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x42, 0x14, 0x82, 0x41, 0x11,
	0x2f, 0x6e, 0x6f, 0x63, 0x73, 0x79, 0x73, 0x2d, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x74, 0x6f, 0x64,
	0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x1a, 0xca, 0x02, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f,
	0x12, 0x5f, 0x0a, 0x09, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x4e, 0x6f, 0x63, 0x73,
	0x79, 0x73, 0x54, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x42, 0x1e, 0x82, 0x41, 0x1b, 0x2f, 0x6e, 0x6f, 0x63,
	0x73, 0x79, 0x73, 0x2d, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x74, 0x6f,
	0x64, 0x6f, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x08, 0x74, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x1a, 0x5a, 0x0a, 0x08, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4e, 0x0a,
	0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x79, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x23, 0x82, 0x41, 0x20, 0x2f, 0x6e, 0x6f, 0x63, 0x73, 0x79, 0x73, 0x2d, 0x74,
	0x6f, 0x64, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x6c, 0x69,
	0x73, 0x74, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x1a, 0x84, 0x01,
	0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x37, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0x82, 0x41, 0x20,
	0x2f, 0x6e, 0x6f, 0x63, 0x73, 0x79, 0x73, 0x2d, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x74, 0x6f, 0x64,
	0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x6e, 0x61, 0x6d, 0x65,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x6f, 0x6e, 0x69,
	0x63, 0x2e, 0x4e, 0x6f, 0x63, 0x73, 0x79, 0x73, 0x54, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64,
	0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x08, 0x74, 0x6f, 0x64, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x1b, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x6f, 0x63, 0x73,
	0x79, 0x73, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x5a, 0x07, 0x2e, 0x3b, 0x73, 0x6f, 0x6e, 0x69,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nocsys_todo_proto_rawDescOnce sync.Once
	file_nocsys_todo_proto_rawDescData = file_nocsys_todo_proto_rawDesc
)

func file_nocsys_todo_proto_rawDescGZIP() []byte {
	file_nocsys_todo_proto_rawDescOnce.Do(func() {
		file_nocsys_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_nocsys_todo_proto_rawDescData)
	})
	return file_nocsys_todo_proto_rawDescData
}

var file_nocsys_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nocsys_todo_proto_goTypes = []interface{}{
	(*NocsysTodo)(nil),                  // 0: sonic.NocsysTodo
	(*NocsysTodo_Todo)(nil),             // 1: sonic.NocsysTodo.Todo
	(*NocsysTodo_Todo_TodoList)(nil),    // 2: sonic.NocsysTodo.Todo.TodoList
	(*NocsysTodo_Todo_TodoListKey)(nil), // 3: sonic.NocsysTodo.Todo.TodoListKey
	(*ywrapper.StringValue)(nil),        // 4: ywrapper.StringValue
}
var file_nocsys_todo_proto_depIdxs = []int32{
	1, // 0: sonic.NocsysTodo.todo:type_name -> sonic.NocsysTodo.Todo
	3, // 1: sonic.NocsysTodo.Todo.todo_list:type_name -> sonic.NocsysTodo.Todo.TodoListKey
	4, // 2: sonic.NocsysTodo.Todo.TodoList.json:type_name -> ywrapper.StringValue
	2, // 3: sonic.NocsysTodo.Todo.TodoListKey.todo_list:type_name -> sonic.NocsysTodo.Todo.TodoList
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_nocsys_todo_proto_init() }
func file_nocsys_todo_proto_init() {
	if File_nocsys_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nocsys_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NocsysTodo); i {
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
		file_nocsys_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NocsysTodo_Todo); i {
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
		file_nocsys_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NocsysTodo_Todo_TodoList); i {
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
		file_nocsys_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NocsysTodo_Todo_TodoListKey); i {
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
			RawDescriptor: file_nocsys_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nocsys_todo_proto_goTypes,
		DependencyIndexes: file_nocsys_todo_proto_depIdxs,
		MessageInfos:      file_nocsys_todo_proto_msgTypes,
	}.Build()
	File_nocsys_todo_proto = out.File
	file_nocsys_todo_proto_rawDesc = nil
	file_nocsys_todo_proto_goTypes = nil
	file_nocsys_todo_proto_depIdxs = nil
}
