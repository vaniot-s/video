>[video](https://coding.imooc.com/lesson/227.html)

[![Build Status](https://travis-ci.com/vaniot-s/video.svg?branch=master)](https://travis-ci.com/vaniot-s/video)

命令
---
build

- 跨平台编译 `env GOOS=linux GOARCH=amd64  go build`

install

- 将编译后的文件打包成库放在pkg

get

- 获取package

fmt

clean

doc

env

bug

fix

generate

list

run

test

tool

version

vet

## docker-mysql

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
-e MYSQL_ROOT_PASSWORD=root \ #环境变量用户密码
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
--restart=always --name mysql \
-e MYSQL_ROOT_PASSWORD=root \
-e LANG=C.UTF-8 \
-v /home/vaniot/dev/dockerdata/mysql/data:/var/lib/mysql \
-v /home/vaniot/dev/dockerdata/mysql/conf:/etc/mysql/ \
daocloud.io/library/mysql:8.0\
--character-set-server=utf8  \
--collation-server=utf8_general_ci
```

docker run -d  -p 3307:3306 \
--restart=always --name mysqls \
-e MYSQL_ROOT_PASSWORD=root \
-e LANG=C.UTF-8 \
mysql \
--character-set-server=utf8  \
--collation-server=utf8_general_ci


### 数据库备份

```shell
 docker exec mysql sh -c 'exec mysqldump --all-databases -uroot -p"$MYSQL_ROOT_PASSWORD"' > /data/mysql/all-databases.sql
```

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

api流程:

handlers->validation{1.request,2.user}->business logic(data )->response

- handler:
  - apidefs.go 数据结构
  - errs.go error handler

scheduler(任务调度):处理异步任务
 
 - RESTful的HTTP server 
 - Timer(定时器)
 - 生产/消费者模型下的task runner(任务的读取)
## api设计

### 用户

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
## scheduler
delete video
 api->videoid-mysql->dispatcher->mysql-videoid->datachannel->executor->datachannel-videoid->delete videos
 
 
代理模式
  - api 
  - proxy:
# 云
### 云原生
特性:
 -松耦架构哦(SOA/Microservice)
 -无状态(stableless),伸缩性(Scalability),(Redundancy)冗余性
 - 平台无关性
##部署发布
 - 自动部署
 - 良好的迁移性
 - 多云共生
 
TODO:
 - [ ] 部署类
   - [ ] 一期:
     - [ ] 配置文件
         - yml
     - [ ] docker
     - [ ] CI/CD
 - [ ] 功能类
   - [ ] 一期:
     - 视频
       - [ ] 弹幕
       - [ ] 正在观看
       - [ ] 收藏
       - [ ] logo
     - [ ] 关注关系
     - [ ] 分片上传
     - [ ] login
        - sso
        - jwt
        - oauth
        - email
        - mobile phone
     - [ ] 站内信
     
 - [ ] 二期:
   - [ ] 推荐系统
   - [ ] 负载均衡
 - [ ] 模块
   - [ ] 一期:
     - [ ] admin 
     - [ ] 监控系统
        - grafna
        - snetry 
   - [ ] 二期: