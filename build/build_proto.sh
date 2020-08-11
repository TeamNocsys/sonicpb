#!/bin/sh

export GOPATH=`go env GOPATH`
output="tmp"

if [ ! -d "$output" ];then
rm -rf "$output"
fi

go run ${GOPATH}/src/github.com/openconfig/ygot/proto_generator/protogenerator.go -output_dir=$output -exclude_modules=ietf-interfaces `find ../yang/ -path '../yang/sonic/*.yang' -o -path '../yang/sonic/third_party/ietf/*.yang' | xargs echo`
find $output -name "*.proto" | xargs sed -i 's/^package openconfig.*;$/package openconfig;/g'
find $output -name "*.proto" | xargs sed -i 's/^import "github.com.*\//import "/g'
find $output -name "*.proto" | xargs sed -i 's/^import "openconfig.*\//import "/g'
find $output -name "*.proto" | xargs sed -i 's/openconfig\.\w*\./openconfig./g'
find $output -name "*.proto" | xargs sed -i '/^package openconfig;$/a\option java_package = "com.nocsys.openconfig";'
find $output -name "*.proto" | xargs sed -i 's/^import "enums.proto";/import "sonic_enums.proto";/g'
for file in `find $output -name "enums.proto"`
do
  newfile=`echo $file | sed s/enums.proto/sonic_enums.proto/`
  mv $file $newfile
done
find $output -name "*.proto" | xargs -i mv {} ../protobuf/sonic
cp ../protobuf/yext/yext.proto ../protobuf/sonic
cp ../protobuf/ywrapper/ywrapper.proto ../protobuf/sonic
rm -rf $output
