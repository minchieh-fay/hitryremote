#!/bin/bash

# HiTryRemote 客户端开发脚本

set -e

echo "启动 HiTryRemote 客户端开发模式..."

# 检查是否安装了 wails
if ! command -v wails &> /dev/null; then
    echo "错误: 未找到 wails 命令，请先安装 wails"
    echo "安装命令: go install github.com/wailsapp/wails/v2/cmd/wails@latest"
    exit 1
fi

# 检查是否安装了 node
if ! command -v node &> /dev/null; then
    echo "错误: 未找到 node 命令，请先安装 Node.js"
    exit 1
fi

# 安装前端依赖（如果还没有安装）
if [ ! -d "frontend/node_modules" ]; then
    echo "安装前端依赖..."
    cd frontend
    npm install
    cd ..
fi

# 启动开发模式
echo "启动开发模式..."
wails dev
