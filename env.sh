#!/usr/bin/env bash

DBHOST="127.0.0.1"
DRIVERNAME="mysql"
USERNAME="root"
PASSWORD="root"
PORT="3307"
DATABASE="video"
OSSURL="http://pq5cmm2db.bkt.clouddn.com/"
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

echo "设置OSSURL"
if [ -n $OSSURL ]
then
    echo "已经设置过OSSURL,请检查"
      echo $OSSURL
    else
    export OSSURL
    echo $OSSURL
fi

