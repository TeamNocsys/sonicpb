> 配置代理

```
# go env -w GO111MODULE=on
# go env -w GOPROXY=https://goproxy.io,direct
```

> 获取代码

```
# GO111MODULE=off go get -u github.com/TeamNocsys/sonicpb
```

> 下载YANG编译工具

```
# GO111MODULE=off go get -u github.com/openconfig/ygot
```

> 将YANG模型转换成Protobuf文件

```
# cd build
# ./build.sh
```

> 配置编译环境

```
# sudo apt install -y protobuf-compiler
# GO111MODULE=off go get -u github.com/google/protobuf
```

> 根据Protobuf生成Java、GO和Python版本代码

```
# ./compile_protos_go.sh
# ./compile_protos_java.sh
# ./compile_protos_python.sh
```

> 打包Jar包

```
# cd build
### Jar包位于target目录
# mvn package
```
