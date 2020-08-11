#!/bin/sh

cd build

# 将yang model编译为proto文件
echo "编译protobuf文件"
./build_proto.sh

# 根据proto文件生成go代码
echo "生成go代码"
./compile_protos_go.sh

# 根据proto文件生成python代码
echo "生成python代码"
./compile_protos_python.sh

# 根据proto文件生成java代码
echo "生成java代码"
./compile_protos_java.sh

cd jar
echo "打包java文件"
mvn package

cd ../..
echo "编译成功"