package main

import (
	"context"
	"fmt"
	"hitryremote/util"
	"log"
	"net"
	"sync"

	quic "github.com/quic-go/quic-go"
)

type ClientInfo struct {
	ClientID string
	Conn     *quic.Conn
	Stream   *quic.Stream
}

type CQUIC struct {
	listener *quic.Listener // 监听器
	clients  map[string]*ClientInfo
	mu       sync.Mutex
}

func NewCQUIC() *CQUIC {
	return &CQUIC{
		clients: make(map[string]*ClientInfo),
	}
}

func (c *CQUIC) handleConnection(conn *quic.Conn) {
	defer conn.CloseWithError(0, "cquic handleConnection 连接关闭")

	fmt.Printf("新的 CQUIC 连接: %s\n", conn.RemoteAddr())

	// client上来的quic只有一个stream,用于传输信令
	controlStream, err := conn.AcceptStream(context.Background())
	if err != nil {
		log.Printf("cquic handleConnection 接受控制流失败: %v", err)
		return
	}

	// 处理客户端注册
	c.handleControlStream(controlStream, conn)
}

func (c *CQUIC) handleControlStream(stream *quic.Stream, conn *quic.Conn) {
	defer stream.Close()

	// 从流中读取消息:客户端注册
	msg := util.QuicMessageFromStream(stream)
	if msg == nil {
		log.Printf("cquic handleControlStream 读取消息失败")
		return
	}

	if msg.Type != util.MSG_TYPE_CLIENTREG {
		log.Printf("cquic handleControlStream 消息类型错误: %d", msg.Type)
		return
	}

	clientRegMsg := msg.Data.(*util.ClientRegMessage)
	if clientRegMsg.Version != util.VERSION {
		// 发送回复
		ackMsg := &util.QuicMessage{
			Type: util.MSG_TYPE_CLIENTREG_ACK,
			Data: &util.ClientRegAckMessage{
				Version: util.VERSION,
				Result:  false,
				Reason:  "版本错误,请升级到最新版本",
			},
		}
		stream.Write(ackMsg.ToBuffer())
		log.Printf("cquic handleControlStream 版本错误: %d", clientRegMsg.Version)
		return
	}

	// 发送回复:客户端注册成功
	ackMsg := &util.QuicMessage{
		Type: util.MSG_TYPE_CLIENTREG_ACK,
		Data: &util.ClientRegAckMessage{
			Version: util.VERSION,
			Result:  true,
			Reason:  "",
		},
	}
	stream.Write(ackMsg.ToBuffer())

	log.Printf("客户端注册成功: %s", clientRegMsg.ClientID)

	// 保存客户端信息
	c.mu.Lock()
	c.clients[clientRegMsg.ClientID] = &ClientInfo{
		ClientID: clientRegMsg.ClientID,
		Conn:     conn,
		Stream:   stream,
	}
	c.mu.Unlock()
	defer func() {
		c.mu.Lock()
		delete(c.clients, clientRegMsg.ClientID)
		c.mu.Unlock()
	}()

	// 循环读取来自客户端的控制流消息
	for {
		msg := util.QuicMessageFromStream(stream)
		if msg == nil {
			log.Printf("cquic handleControlStream 读取消息失败")
			return
		}

		switch msg.Type {
		case util.MSG_TYPE_PING:
			c.handlePing(stream, msg)
		default:
			log.Printf("cquic handleControlStream 未知消息类型: %d", msg.Type)
		}
	}
}

func (c *CQUIC) handlePing(stream *quic.Stream, msg *util.QuicMessage) {
	log.Printf("收到客户端 PING 消息")

	// 创建 PONG 回复
	pongMsg := &util.QuicMessage{
		Type: util.MSG_TYPE_PONG,
		Data: &util.PongMessage{},
	}

	// 发送回复
	_, err := stream.Write(pongMsg.ToBuffer())
	if err != nil {
		log.Printf("发送 PONG 失败: %v", err)
	}
}

func (c *CQUIC) Run() error {
	// 监听quic端口
	// 创建 QUIC 监听器, 监听在0.0.0.0
	addr := net.JoinHostPort("0.0.0.0", util.RELAY_CQUIC_PORT)
	listener, err := quic.ListenAddr(addr, nil, nil)
	if err != nil {
		fmt.Printf("启动 QUIC 监听器失败: %v", err)
		return err
	}

	c.listener = listener
	fmt.Printf("CQUIC 服务器启动，监听地址: %s\n", addr)

	// 开始接受连接

	for {
		conn, err := listener.Accept(context.Background())
		if err != nil {
			fmt.Printf("接受 QUIC 连接失败: %v", err)
			continue
		}

		// 为每个连接启动处理协程
		go c.handleConnection(conn)
	}
}

func (c *CQUIC) CreateStreamByClientID(clientID string) (*quic.Stream, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if client, ok := c.clients[clientID]; ok {
		// open一个stream
		stream, err := client.Conn.OpenStreamSync(context.Background())
		if err != nil {
			return nil, fmt.Errorf("open stream failed: %v", err)
		}
		return stream, nil
	}

	return nil, fmt.Errorf("clientID %s not found", clientID)
}

func (c *CQUIC) Close() error {
	if c.listener != nil {
		return c.listener.Close()
	}
	return nil
}
