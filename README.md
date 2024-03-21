# sonicpb

## Build by docker container

1. Clone the source
```shell
$ git clone  https://github.com/TeamNocsys/sonicpb.git
$ cd sonicpb
```

2. Build Docker image
```shell
docker build --network=host -t test/sonicpb -f Dockerfile  .
```

3. Run the container to generate protobuf
```shell
$ docker run --net=host --rm -v "$(pwd)":/home/myapp:rw -w /home/myapp \
test/sonicpb:latest bash -x build.sh
```

NOTE: it's ok to ignore the error messages about the mvn.

4. Run another container to pacakage the jar
```shell
$ docker run --net=host --rm -v "$(pwd)":/opt/maven -w /opt/maven maven:3.3-jdk-8 \
mvn -f build/jar/pom.xml package
```

5. Copy build/jar/target/sonicpb-1.0.0.jar to use.

---

> go最低版本1.16.3

## 編譯前准備

> 配置代理

```shell
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.io,direct
```

> 下載YANG編譯工具

```shell
### 記得導出環境變量 export PATH=$PATH:$HOME/go/bin
### 使用改版的ygot，不然java編譯會報錯，不能直接使用go get的方式，否則生成的proto_generator還是官方的
$ GOPATH=`go env GOPATH`
$ mkdir -p $GOPATH/src/github.com/TeamNocsys
$ cd $GOPATH/src/github.com/TeamNocsys
$ git clone https://github.com/TeamNocsys/ygot.git
$ cd ygot/proto_generator
$ go install .
```

> 配置編譯環境

```shell
$ sudo apt install -y protobuf-compiler
$ go get -u google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0
$ GO111MODULE=off go get -u github.com/google/protobuf
```

> 獲取代碼

```shell
$ git clone  https://github.com/TeamNocsys/sonicpb.git
```

## 編譯

編譯本項目可以選擇分布編譯或一步編譯，方法如下

### 分步編譯

1. 進入編譯目錄

```shell
$ cd build
```

2. 將YANG模型轉換成Protobuf文件

```shell
$ ./build_proto.sh
```

3. 根據Protobuf生成Java、GO和Python版本代碼

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

### 一步編譯

在項目根目錄執行
```shell
$ ./build.sh
```
