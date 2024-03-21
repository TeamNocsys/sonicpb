#!/bin/sh

export GOPATH=`go env GOPATH`
proto_imports="../api/protobuf/sonic:${GOPATH}/src:${GOPATH}/src/github.com/gogo/protobuf"
java_output="jar/src/main/java"

if [ ! -d "$java_output" ];then
  mkdir -p "$java_output"
else 
  path="$java_output""/[^.]*"
  rm -rf $path
fi

if [ $# -eq 1 ];then
  java_output=$1
fi

protoc -I=$proto_imports --java_out=$java_output `find ../api/protobuf/sonic -name "*.proto" | xargs echo`
