[预览](http://video.vaniot.net/)

文件目录
---

```shell
video
|__Api #api模块 用户相关的操作
|  |__defs #配置定义,返回错误信息,校验
|  |  |__apidefs.go #handlers 的数据结构
|  |  |__errs.go #错误信息数据结构
|  |__main.go #文件入口,路由规划
|  |__dbops #数据库
|  |  |__api.go #数据库操作
|  |  |__conn.go # 数据库连接
|  |  |__api_test.go #测试
|  |  |__internal.go #session的相关操作
|  |__utils
|  |  |__uuid.go #生成唯一的uuid
|  |__handlers.go #路由的具体文件处理
|  |__response.go #响应消息
|  |__session #会话信息
|  |   |__ops.go #生成session
|  |__auth.go #身份校验
|
|__streamserver #视频播放模块 上传/播放
|  |__main.go #入口函数
|  |__halders.go # 处理逻辑
|  |__defs.go # 定义
|  |__limiter.go #流控(控制connection):长链接
|  |__response.go #响应
|  |__config #配置文件夹
|     |__config.go #配置文件
|__Scheduler #对于无法及时返回的任务定时,延时触发
|  |__taskrunner #
|  |  |__defs.go # 预定义
|  |  |__trmain.go #dispatch execute 调用
|  |  |__runner.go # 逻辑生产消费者模型的实现
|  |  |__task.go #dispatch  coumser具体实现
|  |__dbops #数据库相关
|  |  |__internal.go #数据库具体操作 删除,读取软删除记录
|  |  |__conn.go #数据库连接
|  |  |__api.go # api
|  |__handlers.go # handlers
|  |__main.go # api
|  |__response.go  #响应
|__template #前端模板
|__web #webserver
|  |__main.go #
|  |__defs.go #
|  |__client.go #代理转发
|  |__handlers.go #
|__vendor #公共配置的包
|   |__video
|      |__config
|         |__config.go
|__bin #编译后的二进制文件
|  |__scheduler # 任务有调度器
|  |__api # 数据库服务
|  |__streamserver # 视频文件服务
|  |__web # web服务
|  |__conf.json
|__build.sh #内部调试部署脚本
|__duildprod.sh # 将各个部分服务编译并复制到bin
|__deploy.sh # 启动服务
```
## web前端相关
### api用户
- 创建用户 /user post 201 400 500
- 用户登录 /user/:usename post 200 4000 500
- 获取用户的登录信息: /user/:username get 200 400 401 403 500
- 用户注销 /user/:username delete 204 400 401 403 500
- 用户video /user/:usename/videos Get 200 400 500
- 获取video /user/:useranme/videos/:vid-id GET 200 400 500
- 删除video /user/:username/cideos/:cid-id delete 204 400 401 403 500
- 评论列表 /videos/:vid-id/comments GET 200 400 500
- 发布评论 /videos/:vid-id/comments POST 201 400 500
- 删除评论 /videos/:vid-id/comment/:comment-id delete 204 400 401 403 500
## streamserver获取video
- 
## scheduler


## 运行
### docker-mysql
docker下拉区mysql8.0镜像

```shell
docker pull daocloud.io/library/mysql:latest
```

创建container参数说明

```shell
docker run -d  -p 3306:3306 \
--rm --name mysql \ ## 测试时使用 --rm 正式使用 --restart=always 自动重启
–privileged=true \ # 提升容器内权限,错误日权限
###环境变量##
-e MYSQL_ROOT_PASSWORD=passsword \ #环境变量用户密码
-e LANG=C.UTF-8 \ #支持中文
-e MYSQL_USER=vaniot \ #创建用户
-e MYSQL_PASSWORD=123456 \ #设置密码
-e MYSQL_DATABASE=test \ #创建数据库
-e MYSQL_ROOT_HOST=% \
##文件挂载必须是目录##
-v /usr/local/src/mysql/data:/var/lib/mysql \ #mysql 文件目录
-v /usr/local/src/mysql/conf:/etc/mysql \ #配置文件目录
daocloud.io/library/mysql:latest \
--character-set-server=utf8 \
--collation-server=utf8_general_ci
```
可执行的命令:
```shell
docker run -d  -p 3306:3306 \
--restart=always --name mysqls \
-e MYSQL_ROOT_PASSWORD=password \
-e LANG=C.UTF-8 \
mysql \
--character-set-server=utf8  \
--collation-server=utf8_general_ci
```

#### 数据库备份

```shell
 docker exec mysql sh -c 'exec mysqldump --all-databases -uroot -p"$MYSQL_ROOT_PASSWORD"' > /data/mysql/all-databases.sql
```
### 创建数据库
```sql
create database ssss;
```
### 创建表
执行initdb.sql
### 启动部署
### 设置环境变量
在项目下创建env.sh将配需要的配置的环境变量写入到env.sh
```bash
#!/usr/bin/env bash

DBHOST="172.4.0.1"
DRIVERNAME="mysql"
USERNAME="user"
PASSWORD="passwordt"
PORT="3306"
DATABASE="database"
STORAGETYPE="OSS"
OSSURL="http://perewrwerb.bkt.clouddn.com/"
echo "设置数据类型"
if [ -n $DRIVERNAME ]
then
   echo "已经设置过DRIVERNAME,请检查"
  echo $DRIVERNAME
else
  export DRIVERNAME
  echo $DRIVERNAME
fi


echo "设置database Host"
if [ -n $DBHOST ]
then
    echo "已经设置过DBHOST,请检查"
    echo $DBHOST
    else
    export DBHOST
    echo $DBHOST
fi

echo "设置USERNAME"
if [ -n $USERNAME ]
then
     echo "已经设置过USERNAME,请检查"
     echo $PORT
     echo $USERNAME
    else
     export USERNAME
     echo $USERNAME
fi

echo "设置PASSWORD"
if [ -n $PASSWORD ]
then
    echo "已经设置过PASSWORD,请检查"
    echo $PASSWORD
    else
    export PASSWORD
    echo $PASSWORD
fi

echo "设置PORT"
if [ -n $PORT ]
then
    echo "已经设置过PORT,请检查"
    echo $PORT
    else
    export PORT
    echo $PORT
fi

echo "设置DATABASE"
if [ -n $DATABASE ]
then
    echo "已经设置过DATABASE,请检查"
      echo $DATABASE
    else
    export DATABASE
    echo $DATABASE
fi

echo "设置设置存储类型"
if [ -n $OSS ]
then
    echo "已经设置过OSS,请检查"
      echo $OSS
    else
    export OSS
    echo $OSS
fi

echo "设置OSSURL"
if [ -n $OSSURL ]
then
    echo "已经设置过OSSURL,请检查"
      echo $OSSURL
    else
    export OSSURL
    echo $OSSURL
fi

```
### 编译
需要go环境变量的支持
```sql
sh ./buildprod.sh
```
## 启动
```bash
sh ./deploy.sh
```
> 单独部署前端
```bash
sh ./deployFE.sh
```