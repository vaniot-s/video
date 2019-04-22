#!/usr/bin/env bash

#build web ui
cd ./web
#编译
env GOOS=linux GOARCH=amd64 go build .

if [ ! -d "web_ui" ]; then
  mkdir web_ui
fi
#复制编译后的二进制文件
cp  web ./web_ui/web
#复制资源文件
cp -R  ./templates ./web_ui/