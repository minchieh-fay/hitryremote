# HiTryRemote 客户端

基于 Wails 和 QUIC 协议的高性能代理客户端。

## 功能特性

- 🚀 基于 QUIC 协议的高性能传输
- 🎨 现代化的 Vue 3 + Element Plus 界面
- ⚙️ 灵活的连接管理
- 📊 实时状态监控
- 📝 详细的日志记录
- 🔧 丰富的配置选项

## 技术栈

### 后端
- Go 1.21+
- Wails v2
- QUIC-Go

### 前端
- Vue 3
- Element Plus
- Pinia
- Vue Router
- Vite

## 开发环境要求

- Go 1.21 或更高版本
- Node.js 16 或更高版本
- Wails v2

## 安装依赖

### 安装 Wails
```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### 安装前端依赖
```bash
cd frontend
npm install
```

## 开发

### 启动开发模式
```bash
# 使用脚本
./scripts/dev.sh

# 或直接使用 wails
wails dev
```

### 构建应用
```bash
# 使用脚本
./scripts/build.sh

# 或直接使用 wails
wails build
```

## 项目结构

```
client/
├── app.go                 # 主应用逻辑
├── main.go               # 应用入口
├── wails.json            # Wails 配置
├── go.mod               # Go 模块文件
├── package.json         # 项目配置
├── frontend/            # 前端代码
│   ├── src/
│   │   ├── views/       # 页面组件
│   │   ├── router/      # 路由配置
│   │   ├── App.vue      # 根组件
│   │   └── main.js      # 前端入口
│   ├── package.json     # 前端依赖
│   └── vite.config.js   # Vite 配置
├── internal/            # 内部包
│   ├── config/          # 配置管理
│   ├── logger/          # 日志管理
│   └── quic/            # QUIC 客户端
└── scripts/             # 构建脚本
    ├── build.sh         # 构建脚本
    └── dev.sh           # 开发脚本
```

## 配置

应用配置文件位置：
- Windows: `%USERPROFILE%\.hitryremote\client.json`
- macOS/Linux: `~/.hitryremote/client.json`

## 日志

日志文件位置：
- Windows: `%USERPROFILE%\.hitryremote\logs\`
- macOS/Linux: `~/.hitryremote/logs/`

## 许可证

MIT License
