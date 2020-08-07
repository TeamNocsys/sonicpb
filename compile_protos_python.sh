#!/bin/sh

export GOPATH=`go env GOPATH`
proto_imports="api/protobuf/openconfig:${GOPATH}/src:${GOPATH}/src/github.com/google/protobuf/src"
python_output="build/wheel"
if [ $# -eq 1 ];then
  python_output=$1
fi

protoc -I=$proto_imports --python_out=$python_output `find api/protobuf/openconfig -name "*.proto" | xargs echo`

