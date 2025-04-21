#!/bin/bash

set -e

CONTAINER_NAME=fake-server
TARGET_PATH=/usr/local/bin/go-demo

echo "📦 检查容器是否存在..."
if ! docker ps -a --format '{{.Names}}' | grep -q "^$CONTAINER_NAME$"; then
    echo "🚀 创建模拟容器 $CONTAINER_NAME..."
    docker run -dit --name $CONTAINER_NAME ubuntu:20.04
fi

echo "📤 上传 go-demo 到容器 $CONTAINER_NAME:$TARGET_PATH ..."
docker cp bin/go-demo $CONTAINER_NAME:$TARGET_PATH

echo "🔁 在容器中重启 go-demo（先杀进程）..."
docker exec $CONTAINER_NAME pkill go-demo || true
docker exec -d $CONTAINER_NAME $TARGET_PATH

echo "✅ 部署完成！你可以执行下面命令进入容器查看运行情况："
echo "docker exec -it $CONTAINER_NAME bash"
