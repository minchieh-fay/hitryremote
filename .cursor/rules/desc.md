# 这是一个华创公司自用的一个远程端口工具

## 协议
client和server之间使用quic

## 分3个程序
1. server  
公司内网服务器的一台机器运行该程序

功能:
a.他主动注册到relay中继的1号端口
b.他开放本地的http查看端口
c.根据客户的发过来的json串.映射本地端口

2.client
用wails开发的一个桌面app
功能:
a.主动注册到relay中继的2号端口
b.上报用户填写的网络信息(例如:192.168.1.1 22 master机器的ssh地址)
c.接受来自relay的stream读取信息后,代理到本地的网络地址

3.relay
代理client和server
他有2个端口分别用quic监听
然后来自client的连接全部对接到server上

### **请严格尊重我的提问或者请求,不要额外加东西**
**我让你加什么你就加什么,千万不要扩展**

### 目录结构
├── util  // 通用库
│   └── proto // 协议
│   └── define // 通用定义
├── server // 服务端
├── relay // 中继
├── client // wails开发的客户端
