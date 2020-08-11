#!/bin/sh

export GOPATH=`go env GOPATH`

go_output="go"

if [ ! -d  "$go_output" ];then
mkdir "$go_output"
else
path="$go_output""/sonic*"
rm -f $path 
fi

proto_imports="../protobuf/sonic:${GOPATH}/src:${GOPATH}/src/github.com/google/protobuf/src"
protoc -I=$proto_imports --go_out=$go_output `find ../protobuf/sonic -name "*.proto" ! -name "yext.proto" ! -name "ywrapper.proto" | xargs echo`
