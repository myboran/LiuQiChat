#!/bin/bash
PATH=/bin:/sbin:/usr/bin:/usr/local/bin:/usr/sbin:/usr/local/sbin:~/bin
export PATH
docker rm -f nginxhtml
# 删除原有容器
docker rmi -f nginx:v
# 删除原有镜像

docker build -t nginx:v .
# 构建新的镜像
docker run -d -v /root/nginx/nginx.conf:/etc/nginx/nginx.conf \
        --name nginxhtml \
        -p 80:80 \
        nginx:v