#!/usr/bin/env bash

#静态文件
cp -R ./templates ./bin/

#videos暂存文件
if [ ! -d "/bin/videos" ]; then
mkdir ./bin/videos
fi


# 启动
cd bin
nohup ./api &
nohup ./scheduler &
nohup ./streamserver &
nohup ./web &

echo "deploy finished"