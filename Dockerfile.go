FROM hyperjiang/golang:1.16.3

WORKDIR $GOPATH/src/github.com/TeamNocsys

RUN go env -w GO111MODULE=on; \
    go env -w GOPROXY=https://goproxy.io,direct; \
    git clone https://github.com/TeamNocsys/ygot.git; \
    cd ygot/proto_generator; \
    go install .

RUN apt-get update && apt-get install -y protobuf-compiler

RUN go get -u google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0

RUN GO111MODULE=off go get -u github.com/gogo/protobuf/proto


