#!/bin/sh

export GOPATH=`go env GOPATH`
proto_imports="../protobuf/sonic:${GOPATH}/src:${GOPATH}/src/github.com/google/protobuf/src"
protoc -I=$proto_imports --go_out=go `find ../protobuf/sonic -name "*.proto" ! -name "yext.proto" ! -name "ywrapper.proto" | xargs echo`
