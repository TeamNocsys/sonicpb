#!/bin/sh

export GOPATH=`go env GOPATH`
proto_imports="../protobuf/sonic:${GOPATH}/src:${GOPATH}/src/github.com/google/protobuf/src"
java_output="jar/src/main/java"
if [ $# -eq 1 ];then
  java_output=$1
fi

protoc -I=$proto_imports --java_out=$java_output `find ../protobuf/sonic -name "*.proto" | xargs echo`
