#!/bin/sh

export GOPATH=`go env GOPATH`
proto_imports="api/protobuf/openconfig:${GOPATH}/src:${GOPATH}/src/github.com/google/protobuf/src"
protoc -I=$proto_imports --go_out=api/protobuf/openconfig `find api/protobuf/openconfig -name "*.proto" ! -name "yext.proto" ! -name "ywrapper.proto" | xargs echo`
