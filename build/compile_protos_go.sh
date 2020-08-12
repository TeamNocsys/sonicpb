#!/bin/sh

export GOPATH=`go env GOPATH`

go_output="../api/protobuf/sonic"

if [ -d  "$go_output" ];then
path="$go_output""/*.pb.go"
rm -f $path 
fi

proto_imports="../api/protobuf/sonic:${GOPATH}/src:${GOPATH}/src/github.com/google/protobuf/src"
protoc -I=$proto_imports --go_out=$go_output `find ../api/protobuf/sonic -name "*.proto" ! -name "yext.proto" ! -name "ywrapper.proto" | xargs echo`
