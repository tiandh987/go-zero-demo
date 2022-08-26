#!/bin/bash

# 编译 server.go

# Docker alpine 部署 Go 项目失败分析
#   https://blog.csdn.net/hanziyuan08/article/details/105463930

GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" server.go

# docker
#   run    Run a command in a new container
#
# Usage:
#   docker run [OPTIONS] IMAGE [COMMAND] [ARG...]
#
#   --cpus decimal          Number of CPUs
#   --rm                    Automatically remove the container when it exits
#   -i, --interactive       Keep STDIN open even if not attached
#   -t, --tty               Allocate a pseudo-TTY
#   -p, --publish list      Publish a container's port(s) to the host
#   -v, --volume list       Bind mount a volume
#   -w, --workdir string    Working directory inside the container

# -v `pwd`:/app
#   将当前目录挂在到容器内 /app

# alpine
#   IMAGE

# /app/server
#   COMMAND

sudo docker run --rm -it --cpus=1 -p 8080:8080 -v `pwd`:/app -w /app alpine /app/server

rm -f server