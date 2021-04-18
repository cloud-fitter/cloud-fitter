#!/bin/bash

# 废弃

# 安装到 $GOPATH/bin 目录中
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

# 采用buf的方式生成
wget https://github.com/bufbuild/buf/releases/download/v0.41.0/buf-Darwin-x86_64.tar.gz
tar xvzf buf-Darwin-x86_64.tar.gz
mv buf/bin/* ${GOPATH}/bin/


# 弃用：protoc使用起来不方便
## 下载protoc并安装对应的依赖，默认都放到 $GOPATH 中
#wget https://github.com/protocolbuffers/protobuf/releases/download/v3.15.8/protoc-3.15.8-osx-x86_64.zip
#unzip protoc-3.15.8-osx-x86_64.zip
#mv bin/protoc $GOPATH/bin/
#mv include/ $GOPATH/
#rm -rf include/ bin/ readme.txt protoc-3.15.8-osx-x86_64.zip
#
#mkdir $GOPATH/include/google/api
#cd $GOPATH/include/google/api
#wget https://github.com/googleapis/googleapis/blob/master/google/api/annotations.proto
#wget https://github.com/googleapis/googleapis/blob/master/google/api/http.proto
#wget https://github.com/googleapis/googleapis/blob/master/google/api/httpbody.proto
#wget https://github.com/googleapis/googleapis/blob/master/google/api/field_behavior.proto
