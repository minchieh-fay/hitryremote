package quic

import (
	"context"
	"hitryremote-client/internal/config"
	"hitryremote-client/internal/logger"
	"time"

	quic "github.com/quic-go/quic-go"
)

// Client QUIC 客户端
type Client struct {
	serverAddr string
	config     *config.Config
	logger     *logger.Logger
	conn       *quic.Conn
	stream     *quic.Stream
	status     string
}

// NewClient 创建新的 QUIC 客户端
func NewClient(serverAddr string, cfg *config.Config) *Client {
	return &Client{
		serverAddr: serverAddr,
		config:     cfg,
		logger:     logger.New(),
		status:     "disconnected",
	}
}

// Connect 连接到服务器
func (c *Client) Connect() error {
	c.logger.Info("正在连接到服务器: " + c.serverAddr)

	// 创建 TLS 配置
	tlsConfig := &quic.Config{
		// 这里可以添加 QUIC 配置
	}

	// 连接到服务器
	conn, err := quic.DialAddr(context.Background(), c.serverAddr, nil, tlsConfig)
	if err != nil {
		c.status = "error"
		c.logger.Error("连接失败: " + err.Error())
		return err
	}

	c.conn = conn
	c.status = "connected"
	c.logger.Info("连接成功")

	// 启动心跳
	go c.heartbeat()

	return nil
}

// Close 关闭连接
func (c *Client) Close() error {
	c.status = "disconnected"
	if c.conn != nil {
		return c.conn.CloseWithError(0, "客户端关闭")
	}
	return nil
}

// GetStatus 获取连接状态
func (c *Client) GetStatus() string {
	return c.status
}

// heartbeat 心跳检测
func (c *Client) heartbeat() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if c.status == "connected" {
				c.logger.Debug("发送心跳")
				// TODO: 实现心跳逻辑
			}
		}
	}
}
