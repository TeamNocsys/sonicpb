> 获取代码

```
# go get -u github.com/TeamNocsys/sonicpb
```

> 下载YANG编译工具

```
# go get -u github.com/openconfig/ygot
```

> 将YANG模型转换成Protobuf文件

```
# cd tools
# go generate 
```

> 根据Protobuf生成Java、GO和Python版本代码

```
# ./compile_protos.sh
```

> 打包Jar包

```
# cd build
### Jar包位于target目录
# mvn package
```