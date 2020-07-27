#!/bin/sh

set -euo pipefail

proto_imports="api/protobuf/openconfig"

protoc -I=$proto_imports --go_out=api/protobuf/openconfig api/protobuf/openconfig/sonic_acl_enums.proto
protoc -I=$proto_imports --go_out=api/protobuf/openconfig api/protobuf/openconfig/sonic_acl_enums.proto
protoc -I=$proto_imports --go_out=api/protobuf/openconfig api/protobuf/openconfig/sonic_acl.proto
protoc -I=$proto_imports --go_out=api/protobuf/openconfig api/protobuf/openconfig/sonic_interface_enums.proto
protoc -I=$proto_imports --go_out=api/protobuf/openconfig api/protobuf/openconfig/sonic_interface.proto
protoc -I=$proto_imports --go_out=api/protobuf/openconfig api/protobuf/openconfig/sonic_loopback_interface_enums.proto
protoc -I=$proto_imports --go_out=api/protobuf/openconfig api/protobuf/openconfig/sonic_loopback_interface.proto
protoc -I=$proto_imports --go_out=api/protobuf/openconfig api/protobuf/openconfig/sonic_port_enums.proto
protoc -I=$proto_imports --go_out=api/protobuf/openconfig api/protobuf/openconfig/sonic_port.proto
protoc -I=$proto_imports --go_out=api/protobuf/openconfig api/protobuf/openconfig/sonic_portchannel_enums.proto
protoc -I=$proto_imports --go_out=api/protobuf/openconfig api/protobuf/openconfig/sonic_portchannel.proto
protoc -I=$proto_imports --go_out=api/protobuf/openconfig api/protobuf/openconfig/sonic_vlan_enums.proto
protoc -I=$proto_imports --go_out=api/protobuf/openconfig api/protobuf/openconfig/sonic_vlan.proto