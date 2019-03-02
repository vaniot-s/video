#!/usr/bin/env bash

# 检查是否存在目录
if [ ! -d "bin" ]; then
  mkdir bin
fi

# build  web and other service
cd ./api
env GOOS=linux GOARCH=amd64 go build -o ../bin/api

cd ../scheduler
env GOOS=linux GOARCH=amd64 go build -o ../bin/scheduler

cd ../streamserver
env GOOS=linux GOARCH=amd64 go build -o ../bin/streamserver

cd ../web
env GOOS=linux GOARCH=amd64 go build -o ../bin/web

