# cloud-fitter 公有云适配口

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
