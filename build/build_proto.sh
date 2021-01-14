#!/bin/sh

export GOPATH=`go env GOPATH`
output="tmp"

if [ -d "$output" ];then
  rm -rf "$output"
fi

proto_generator -output_dir=$output -typedef_enum_with_defmod -consistent_union_enum_names -exclude_modules=ietf-interfaces `find ../api/yang/ -path '../api/yang/sonic/*.yang' -o -path '../api/yang/sonic/third_party/ietf/*.yang' | xargs echo`
find $output -name "*.proto" | xargs sed -i 's/^package openconfig.*;$/package sonic;/g'
find $output -name "*.proto" | xargs sed -i 's/^import "github.com.*\//import "/g'
find $output -name "*.proto" | xargs sed -i 's/^import "openconfig.*\//import "/g'
find $output -name "*.proto" | xargs sed -i 's/openconfig\.\w*\./sonic./g'
find $output -name "*.proto" | xargs sed -i '/^package sonic;$/a\option java_package = "com.nocsys.sonic";\noption go_package = ".;sonic";'

dest="../api/protobuf/sonic"
if [ -d "$dest" ];then
  path="$dest""/*.proto"
  rm -rf $path
fi
find $output -name "*.proto" | xargs -i mv {} ../api/protobuf/sonic

cp ${GOPATH}/src/github.com/openconfig/ygot/proto/yext/yext.proto ../api/protobuf/sonic
cp ${GOPATH}/src/github.com/openconfig/ygot/proto/ywrapper/ywrapper.proto ../api/protobuf/sonic

rm -rf $output
