# HiTryRemote 服务器

基于 Gin + QUIC 的客户端管理服务器，提供 Web 界面查看客户端注册信息。

## 功能特性

- 🚀 基于 QUIC 协议的高性能通信
- 🌐 现代化的 Web 管理界面
- 📊 实时客户端状态监控
- 📝 客户端地址信息管理
- 🔧 灵活的配置管理
- 💾 内存存储（可扩展为数据库存储）

## 技术栈

- **后端**: Go 1.21+
- **Web 框架**: Gin
- **QUIC**: quic-go
- **前端**: Bootstrap 5 + JavaScript
- **日志**: Logrus
- **配置**: Viper

## 项目结构

```
server/
├── cmd/server/          # 服务器入口
├── internal/            # 内部包
│   ├── api/            # API 处理器和路由
│   ├── config/         # 配置管理
│   ├── models/         # 数据模型
│   ├── quic/           # QUIC 中继服务器
│   └── storage/        # 存储层
├── pkg/                # 公共包
│   └── logger/         # 日志记录
├── web/                # Web 资源
│   ├── static/         # 静态文件
│   └── templates/      # HTML 模板
├── configs/            # 配置文件
├── scripts/            # 构建和启动脚本
└── build/              # 构建输出
```

## 快速开始

### 1. 安装依赖

```bash
go mod tidy
```

### 2. 配置服务器

编辑 `configs/config.yaml`:

```yaml
server:
  host: "0.0.0.0"
  port: 8080
  mode: "debug"

relay:
  host: "0.0.0.0"
  port: 10001

database:
  type: "memory"
  dsn: ""

log:
  level: "info"
  format: "text"
  file: "logs/server.log"
```

### 3. 启动服务器

```bash
# 使用脚本启动
./scripts/start.sh

# 或直接运行
go run cmd/server/main.go

# 或构建后运行
./scripts/build.sh
./build/hitryremote-server
```

### 4. 访问管理界面

打开浏览器访问: http://localhost:8080

## API 接口

### 客户端管理

- `GET /api/v1/clients` - 获取客户端列表
- `POST /api/v1/clients` - 创建客户端
- `GET /api/v1/clients/:id` - 获取单个客户端
- `PUT /api/v1/clients/:id/status` - 更新客户端状态

### 地址管理

- `GET /api/v1/clients/:id/addresses` - 获取客户端地址列表
- `POST /api/v1/clients/:id/addresses` - 为客户端添加地址
- `PUT /api/v1/clients/:id/addresses/:addressId` - 更新地址
- `DELETE /api/v1/clients/:id/addresses/:addressId` - 删除地址

### 健康检查

- `GET /health` - 服务器健康状态

## QUIC 协议

服务器监听 QUIC 连接，处理以下消息类型：

### 客户端注册
```json
{
  "type": "client_register",
  "data": {
    "client_id": "client-001",
    "name": "测试客户端"
  }
}
```

### 地址注册
```json
{
  "type": "address_register",
  "data": {
    "client_id": "client-001",
    "ip": "12.33.33.33",
    "port": 443,
    "description": "xxx管理平台"
  }
}
```

### 心跳
```json
{
  "type": "heartbeat",
  "data": {
    "client_id": "client-001"
  }
}
```

## 配置说明

### 服务器配置
- `server.host`: Web 服务器监听地址
- `server.port`: Web 服务器端口
- `server.mode`: Gin 运行模式 (debug/release/test)

### 中继配置
- `relay.host`: QUIC 中继服务器监听地址
- `relay.port`: QUIC 中继服务器端口

### 数据库配置
- `database.type`: 存储类型 (memory/sqlite/mysql)
- `database.dsn`: 数据库连接字符串

### 日志配置
- `log.level`: 日志级别 (debug/info/warn/error)
- `log.format`: 日志格式 (text/json)
- `log.file`: 日志文件路径

## 开发

### 添加新的 API 接口

1. 在 `internal/api/handlers.go` 中添加处理函数
2. 在 `internal/api/routes.go` 中注册路由
3. 在 `internal/models/` 中定义数据模型

### 添加新的存储后端

1. 实现 `internal/storage/interface.go` 中的 `Storage` 接口
2. 在 `cmd/server/main.go` 中注册新的存储类型

## 部署

### Docker 部署

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o server cmd/server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/server .
COPY --from=builder /app/web ./web
COPY --from=builder /app/configs ./configs
EXPOSE 8080 10001
CMD ["./server"]
```

### 系统服务

创建 systemd 服务文件 `/etc/systemd/system/hitryremote.service`:

```ini
[Unit]
Description=HiTryRemote Server
After=network.target

[Service]
Type=simple
User=hitryremote
WorkingDirectory=/opt/hitryremote
ExecStart=/opt/hitryremote/server
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
```

## 许可证

MIT License
