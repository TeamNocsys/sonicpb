#!/bin/sh

set -euo pipefail

proto_imports="api/protobuf/openconfig"
java_output="build/jar/src/main/java"
if [ $# -eq 1 ];then
  java_output=$1
fi

protoc -I=$proto_imports --java_out=$java_output api/protobuf/openconfig/yext.proto
protoc -I=$proto_imports --java_out=$java_output api/protobuf/openconfig/ywrapper.proto
protoc -I=$proto_imports --java_out=$java_output api/protobuf/openconfig/enums.proto
protoc -I=$proto_imports --java_out=$java_output api/protobuf/openconfig/sonic_acl.proto
protoc -I=$proto_imports --java_out=$java_output api/protobuf/openconfig/sonic_interface.proto
protoc -I=$proto_imports --java_out=$java_output api/protobuf/openconfig/sonic_loopback_interface.proto
protoc -I=$proto_imports --java_out=$java_output api/protobuf/openconfig/sonic_port.proto
protoc -I=$proto_imports --java_out=$java_output api/protobuf/openconfig/sonic_portchannel.proto
protoc -I=$proto_imports --java_out=$java_output api/protobuf/openconfig/sonic_vlan.proto