# 开发者手册

欢迎各位Go语言爱好者或公有云用户加入到这个项目。无论是 **贡献代码** ，还是 **在自己的项目中使用**，都是对这个项目的巨大帮助。

如果希望能深度参与到这个项目中，请参考下面的文档：

## 如何加入到这个项目？

为了让新人能尽快地熟悉合作开发，申请加入项目的流程，也是模拟了一次Feature的开发

- 1 创建Issue
  - 新参与者：创建一个Issue，表示希望加入到本项目，如[示例](https://github.com/cloud-fitter/cloud-fitter/issues/8)
  - 常规开发：一个Issue为一个开发需求或者BugFix
- 2 被指定为Issue负责人
  - 新参与者：项目负责人会将这个Issue指派给你，表示将由你负责
  - 常规开发：表示这个Issue被项目接受，并由交给了对应的负责人，在非特殊情况下其他人不会参与
- 3 fork项目并进行修改
  - 新参与者：fork项目，并在 `idl/demo/demo.proto` 文件中的 `OurTeam` 添加github信息，表示希望加入到项目中
  - 参考下面 **IDL环境搭建** ，将proto文件生成对应的代码，可在 `gen/idl/demo/demo.pb.go`中查看到信息，表示开发者打通了开发流程
  - 常规开发：对 Issue 进行开发
  - 后续会对重度参与开发的朋友进一步开放权限
- 4 提交PR，合并到主项目的master中
  - 新参与者：提交PR，请求合入到master，并注明对应的Issue编号，表示已开发完，希望验收
  - 待管理员审核通过后，会将你的PR绑定到Issue上，邀请您加入开发者的微信群
  - 常规开发：表示需求完成开发，请求审核

## 开发者社区

目前，这个开发者社区虽然年轻，但有多位资深 Go语言开发专家 与 公有云领域专家 的参与，正在将各自原有的经验沉淀到这个新平台。

整体的开发难度并不高，关键有三点，：

1. 抽象极为复杂的 多租户、多区域、多产品 的场景，锻炼面向对象的抽象能力
2. 大量引入并发操作，实践高并发的场景
3. 对接各个公有云产品SDK，阅读相关产品与代码文档

总结下来，就是 **抽象能力、并发能力、以及阅读第三方文档与代码的能力**，对有自驱力的新手尤为有价值

我们开发的是一个平台，更是希望搭建一个健康、活跃的技术交流社区，目前重点会放在以下三块：

- 搭建一个技术领先的web平台，引入大量的优秀实践，供大家在开发自己的平台时有所参考
- 每月提供一次以上的技术分享，介绍平台中的优秀实践，解答用户常见的问题
- 对进行PR的新手朋友进行点对点的`code review`，持续提升每位参与者的项目实战能力

## Go语言环境准备

- **Go语言版本** : >= 1.15
- **环境变量配置**
  - GOPROXY=https://goproxy.io,direct
  - GO111MODULE=on

## IDL环境搭建

```shell script
# 安装grpc-gateway依赖的库，如果无法下载请参考上面的 GORPOXY 设置
# 详情可参考 https://github.com/grpc-ecosystem/grpc-gateway
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

# 采用 buf 工具的方式生成proto相关的代码（弃用protoc）
# 1. mac操作系统
wget https://github.com/bufbuild/buf/releases/download/v0.41.0/buf-Darwin-x86_64.tar.gz
tar xvzf buf-Darwin-x86_64.tar.gz
mv buf/bin/ ${GOPATH}/bin/

# 2. 其余操作系统
# 手动去 https://github.com/bufbuild/buf/releases 下载0.41.0版本的二进制文件，放到可执行的目录中

# 根据IDL生成Go代码（重要） - 也可直接运行 ./gen.sh
buf beta mod update
buf generate
```

## 验证环境搭建

```shell script
# 编译通过
go build 

# 创建属于自己的配置文件（由于涉及到账密信息，需要自行维护）,可参考 config_template.yaml
cp config_template.yaml config.yaml

## 运行二进制文件
./cloud-fitter -conf=config_template.yaml

## 通过HTTP请求访问到结果
curl --location --request POST 'http://localhost:8081/v1/demo'
```

## 容器环境验证(可选)

```shell script
# 拉取镜像(自行安装docker)
docker pull cloudfitter/cloud-fitter

# 确保本地端口的8081和9090没有被占用，配置文件放在{config_dir}里，名字必须为 config.yaml
# 需要定制化修改启动命令的，请参考Dockerfie
docker run -v {config_dir}:/app/config/ -p 8081:8081 -p 9090:9090 cloudfitter/cloud-fitter:latest

## 通过HTTP请求访问到结果
curl --location --request POST 'http://localhost:8081/v1/demo'
```
