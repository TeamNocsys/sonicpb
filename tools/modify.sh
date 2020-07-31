#!/bin/sh

find .tmp -name "*.proto" | xargs sed -i 's/^package openconfig.*;$/package openconfig;/g'
find .tmp -name "*.proto" | xargs sed -i 's/^import "github.com.*\//import "/g'
find .tmp -name "*.proto" | xargs sed -i 's/^import "openconfig.*\//import "/g'
find .tmp -name "*.proto" | xargs sed -i 's/openconfig\.\w*\./openconfig./g'
find .tmp -name "*.proto" | xargs sed -i '/^package openconfig;$/a\option java_package = "com.nocsys.openconfig";'

find .tmp -name "*.proto" | xargs -I '{}' cp {} ../api/protobuf/openconfig
cp ${GOPATH}/src/github.com/openconfig/ygot/proto/yext/yext.proto ../api/protobuf/openconfig
cp ${GOPATH}/src/github.com/openconfig/ygot/proto/ywrapper/ywrapper.proto ../api/protobuf/openconfig
rm -rf .tmp