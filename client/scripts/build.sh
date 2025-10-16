#!/bin/bash

# HiTryRemote 客户端构建脚本

set -e

echo "开始构建 HiTryRemote 客户端..."

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

# 安装前端依赖
echo "安装前端依赖..."
cd frontend
npm install
cd ..

# 构建前端
echo "构建前端..."
cd frontend
npm run build
cd ..

# 构建应用
echo "构建应用..."
wails build -clean

echo "构建完成！"
echo "输出目录: build/bin/"
