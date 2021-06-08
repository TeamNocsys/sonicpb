# sonicpb

> go最低版本1.16.3

## 编译前准备

> 配置代理

```shell
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.io,direct
```

> 下载YANG编译工具

```shell
### 记得导出环境变量 export PATH=$PATH:$HOME/go/bin
### 使用改版的ygot，不然java编译会报错
$ GO111MODULE=off go get -u github.com/TeamNocsys/ygot/proto_generator
```

> 配置编译环境

```shell
$ sudo apt install -y protobuf-compiler
$ go get -u google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0
$ GO111MODULE=off go get -u github.com/google/protobuf
```

> 获取代码

```shell
$ git clone  https://github.com/TeamNocsys/sonicpb.git
```

## 编译

编译本项目可以选择分布编译或一步编译，方法如下

### 分步编译

1. 进入编译目录

```shell
$ cd build
```

2. 将YANG模型转换成Protobuf文件

```shell
$ ./build_proto.sh
```

3. 根据Protobuf生成Java、GO和Python版本代码

```shell
$ ./compile_protos_go.sh
$ ./compile_protos_java.sh
$ ./compile_protos_python.sh
```

4. 打包Jar包

```shell
$ cd build/jar
$ mvn package
```

### 一步编译

在项目根目录执行
```shell
$ ./build.sh
```
