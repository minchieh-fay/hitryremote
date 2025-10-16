package main

import (
	"fmt"
	"hitryremote/util"
	"log"

	"github.com/quic-go/quic-go"
)

func (s *SQUIC) handlePing(stream *quic.Stream, msg *util.QuicMessage) {
	fmt.Printf("收到 PING 消息\n")

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

func (s *SQUIC) handleDataChannel(stream *quic.Stream) {
	defer stream.Close()

	// 从流中读取消息
	msg := util.QuicMessageFromStream(stream)
	if msg == nil {
		log.Printf("读取消息失败")
		return
	}

	// 处理不同类型的消息
	switch msg.Type {
	case util.MSG_TYPE_TCP_CHANNEL:
		s.handleTCPChannel(stream, msg)
	}
}

func (s *SQUIC) handleTCPChannel(stream *quic.Stream, msg *util.QuicMessage) {
	fmt.Printf("收到 TCP_CHANNEL 消息: %+v\n", msg.Data)
}
