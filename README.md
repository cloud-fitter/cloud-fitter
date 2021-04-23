# cloud-fitter 公有云适配口

[![Go Report Card](https://goreportcard.com/badge/github.com/cloud-fitter/cloud-fitter?style=flat-square)](https://goreportcard.com/report/github.com/cloud-fitter/cloud-fitter)
[![Build Status](https://travis-ci.com/cloud-fitter/cloud-fitter.svg?branch=master)](https://travis-ci.com/cloud-fitter/cloud-fitter)
[![codecov](https://codecov.io/gh/cloud-fitter/cloud-fitter/branch/master/graph/badge.svg?token=OJJG8KF8A3)](https://codecov.io/gh/cloud-fitter/cloud-fitter)
[![LICENSE](https://img.shields.io/github/license/cloud-fitter/cloud-fitter.svg?style=flat-square)](https://github.com/cloud-fitter/cloud-fitter/blob/master/LICENSE)

## 开发环境准备

```shell script
# 安装依赖的库，如果无法下载请参考 GORPOXY 设置
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

# 采用 buf 工具的方式生成proto相关的代码，弃用protoc
wget https://github.com/bufbuild/buf/releases/download/v0.41.0/buf-Darwin-x86_64.tar.gz
tar xvzf buf-Darwin-x86_64.tar.gz
mv buf/bin/ ${GOPATH}/bin/
```

## 验证环境搭建

```shell script
# 编译通过
go build 

## 运行二进制文件
./cloud-fitter

## 通过HTTP请求访问到结果
curl --location --request POST 'http://localhost:8081/v1/demo'
```
