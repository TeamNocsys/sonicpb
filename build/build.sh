#!/bin/sh

export GOPATH=`go env GOPATH`
go run ${GOPATH}/src/github.com/openconfig/ygot/proto_generator/protogenerator.go -output_dir=.tmp -exclude_modules=ietf-interfaces `find ../yang/ -path '../yang/sonic/*.yang' -o -path '../yang/sonic/third_party/ietf/*.yang' | xargs echo`
find .tmp -name "*.proto" | xargs sed -i 's/^package openconfig.*;$/package openconfig;/g'
find .tmp -name "*.proto" | xargs sed -i 's/^import "github.com.*\//import "/g'
find .tmp -name "*.proto" | xargs sed -i 's/^import "openconfig.*\//import "/g'
find .tmp -name "*.proto" | xargs sed -i 's/openconfig\.\w*\./openconfig./g'
find .tmp -name "*.proto" | xargs sed -i '/^package openconfig;$/a\option java_package = "com.nocsys.openconfig";'
find .tmp -name "*.proto" | xargs sed -i 's/^import "enums.proto";/import "sonic_enums.proto";/g'
for file in `find .tmp -name "enums.proto"`
do
  newfile=`echo $file | sed s/enums.proto/sonic_enums.proto/`
  mv $file $newfile
done
find .tmp -name "*.proto" | xargs -i mv {} ../protobuf/sonic
cp ../protobuf/yext/yext.proto ../protobuf/sonic
cp ../protobuf/ywrapper/ywrapper.proto ../protobuf/sonic
rm -rf .tmp
