package scripts

////go:generate go run ${GOPATH}/src/github.com/openconfig/ygot/proto_generator/protogenerator.go -base_import_path=github.com/TeamNocsys/sonicpb/api/protobuf -path=../api/yang/public/release/models:../api/yang/public/release/models/types:../api/yang/public/release/models/interfaces:../api/yang/public/release/models/acl:../api/yang/public/third_party -output_dir=../api/protobuf -enum_package_name=openconfig_acl_enums -exclude_modules=ietf-interfaces openconfig-acl.yang

//go:generate go run ${GOPATH}/src/github.com/openconfig/ygot/proto_generator/protogenerator.go -path=../api/yang/sonic:../api/yang/public/third_party -output_dir=.tmp -exclude_modules=ietf-interfaces ../api/yang/sonic/sonic-extension.yang ../api/yang/sonic/sonic-types.yang ../api/yang/sonic/sonic-port.yang ../api/yang/sonic/sonic-portchannel.yang ../api/yang/sonic/sonic-loopback-interface.yang ../api/yang/sonic/sonic-interface.yang ../api/yang/sonic/sonic-vlan.yang ../api/yang/sonic/sonic-acl.yang
//go:generate ./modify.sh

func init() {}