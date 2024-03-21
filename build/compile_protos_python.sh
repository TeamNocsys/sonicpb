#!/bin/sh

export GOPATH=`go env GOPATH`
proto_imports="../api/protobuf/sonic:${GOPATH}/src:${GOPATH}/src/github.com/gogo/protobuf"
python_output="wheel"
if [ $# -eq 1 ];then
  python_output=$1
fi

protoc -I=$proto_imports --python_out=$python_output `find ../api/protobuf/sonic -name "*.proto" | xargs echo`

