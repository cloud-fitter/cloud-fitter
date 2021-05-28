# cloud-fitter 云适配

[![Go Report Card](https://goreportcard.com/badge/github.com/cloud-fitter/cloud-fitter?style=flat-square)](https://goreportcard.com/report/github.com/cloud-fitter/cloud-fitter)
[![Build Status](https://travis-ci.com/cloud-fitter/cloud-fitter.svg?branch=master)](https://travis-ci.com/cloud-fitter/cloud-fitter)
[![codecov](https://codecov.io/gh/cloud-fitter/cloud-fitter/branch/master/graph/badge.svg?token=OJJG8KF8A3)](https://codecov.io/gh/cloud-fitter/cloud-fitter)
[![LICENSE](https://img.shields.io/github/license/cloud-fitter/cloud-fitter.svg?style=flat-square)](https://github.com/cloud-fitter/cloud-fitter/blob/master/LICENSE)

Communicate with public and private clouds conveniently by a set of apis.

**用一套接口，便捷地访问各类公有云和私有云**

## 快速使用

```shell
# 下载Cloud-Fitter镜像
docker pull cloudfitter/cloud-fitter:latest

# 配置账户信息
# 参考样例 https://github.com/cloud-fitter/cloud-fitter/blob/master/config_template.yaml 

# 启动容器 注意必须为绝对路径
docker run -v {config_file}:/app/config/config.yaml -p 8081:8081 -p 9090:9090 cloudfitter/cloud-fitter:latest

# 界面展示 (上面适用于linux，下面为指定ip的通用方式)
docker run --add-host localnode:$(ifconfig en0 | grep inet | grep -v inet6 | awk '{print $2}') -p 8080:8080 junedayday/cloud-fitter-web:latest
docker run --add-host localnode:{本机ip} -p 8080:8080 cloudfitters/cloud-fitter-web:latest
```

## 对接计划

[对接进度](https://shimo.im/sheets/5bqnr9jnopfbOOqy)

更多内容正在内部筹备中，后续开放，有需求欢迎联系。

## 开发者社区

[开发者社区文档](doc/develop.md)

## 联系我们

[沟通群](https://shimo.im/docs/KrkEVnB7NRcwpmAJ)
