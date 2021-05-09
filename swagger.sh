#!/bin/bash

# 请先保证swagger的安装，可参考 https://goswagger.io/install.html

swagger mixin gen/openapiv2/idl/*/*.json -o gen/swagger.json

rm -rf gen/openapiv2

# 将机器的sshkey保存到公有云机器121.41.88.120上
scp -P 22 gen/swagger.json root@121.41.88.120:/root/doc

# 服务端运行的docker
# docker run --name swagger -d --rm -p 80:80 -v /root/doc:/usr/share/nginx/html/swagger  -e SPEC_URL=swagger/swagger.json redocly/redoc
