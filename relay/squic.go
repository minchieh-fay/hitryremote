package main

import (
	"context"
	"fmt"
	"hitryremote/util"
	"log"
	"net"
	"time"

	quic "github.com/quic-go/quic-go"
)

type SQUIC struct {
	listener *quic.Listener // 监听器
	conn     *quic.Conn     // 连接
	stream   *quic.Stream   // 1号连接的流 用于发送信令/心跳等信息
}

func NewSQUIC() *SQUIC {
	return &SQUIC{}
}

func (s *SQUIC) Run() error {
	// 创建 QUIC 监听器, 监听在0.0.0.0
	addr := net.JoinHostPort("0.0.0.0", util.RELAY_SQUIC_PORT)
	listener, err := quic.ListenAddr(addr, nil, nil)
	if err != nil {
		return fmt.Errorf("启动 QUIC 监听器失败: %v", err)
	}

	s.listener = listener
	fmt.Printf("SQUIC 服务器启动，监听地址: %s\n", addr)

	// 开始接受连接
	go func() {
		for {
			conn, err := listener.Accept(context.Background())
			if err != nil {
				log.Printf("接受 QUIC 连接失败: %v", err)
				continue
			}

			// 为每个连接启动处理协程
			go s.handleConnection(conn)
		}
	}()
	return nil
}

func (s *SQUIC) handleConnection(conn *quic.Conn) {
	defer conn.CloseWithError(0, "squic handleConnection 连接关闭")

	fmt.Printf("新的 SQUIC 连接: %s\n", conn.RemoteAddr())

	// 第一个stream是信令流,独立处理

	// 创建一个bool的chan,等控制流返回成功或者失败
	successChan := make(chan bool, 1)
	controlStream, err := conn.AcceptStream(context.Background())
	if err != nil {
		log.Printf("squic handleConnection 接受控制流失败: %v", err)
		return
	}
	go s.handleControlStream(controlStream, successChan)

	// 如果successChan true=成功, false=失败
	select {
	case success := <-successChan:
		if success {
			log.Printf("squic handleConnection 控制流返回成功")
			oldConn := s.conn
			s.conn = conn
			s.stream = controlStream
			if oldConn != nil {
				oldConn.CloseWithError(0, "新的连接建立,关闭旧的连接")
			}
		} else {
			log.Printf("squic handleConnection 控制流返回失败")
			return
		}
	case <-time.After(5 * time.Second):
		log.Printf("squic handleConnection 控制流返回失败,超时")
		return
	}

	// 接受代理数据流
	for {
		stream, err := conn.AcceptStream(context.Background())
		if err != nil {
			log.Printf("接受流失败: %v", err)
			break
		}

		// 为每个流启动处理协程
		go s.handleDataChannel(stream)
	}
}

func (s *SQUIC) handleControlStream(stream *quic.Stream, successChan chan bool) {
	defer stream.Close()

	// 从流中读取消息:服务端注册
	msg := util.QuicMessageFromStream(stream)
	if msg == nil {
		log.Printf("squic handleControlStream 读取消息失败")
		successChan <- false
		return
	}
	if msg.Type != util.MSG_TYPE_SERVERREG {
		log.Printf("squic handleControlStream 控制流消息类型错误: %d", msg.Type)
		successChan <- false
		return
	}
	serverRegMsg := msg.Data.(*util.ServerRegMessage)
	if serverRegMsg.Version != util.VERSION {
		// 发送回复
		ackMsg := &util.QuicMessage{
			Type: util.MSG_TYPE_SERVERREG_ACK,
			Data: &util.ServerRegAckMessage{
				Version: util.VERSION,
				Result:  false,
				Reason:  "版本错误,请升级到最新版本",
			},
		}
		stream.Write(ackMsg.ToBuffer())

		log.Printf("squic handleControlStream 控制流消息版本错误: %d", serverRegMsg.Version)
		successChan <- false
		return
	}
	// 发送回复:服务端注册成功
	ackMsg := &util.QuicMessage{
		Type: util.MSG_TYPE_SERVERREG_ACK,
		Data: &util.ServerRegAckMessage{
			Version: util.VERSION,
			Result:  true,
			Reason:  "",
		},
	}
	stream.Write(ackMsg.ToBuffer())
	successChan <- true

	// 循环读取来自服务端的控制流消息
	for {
		msg := util.QuicMessageFromStream(stream)
		if msg == nil {
			log.Printf("squic handleControlStream 读取消息失败")
			s.conn.CloseWithError(0, "squic handleControlStream 读取消息失败")
			s.stream = nil
			s.conn = nil
			return
		}
		switch msg.Type {
		case util.MSG_TYPE_PING:
			s.handlePing(stream, msg)
		default:
			log.Printf("squic handleControlStream 未知消息类型: %d", msg.Type)
		}
	}
}

func (s *SQUIC) Close() error {
	if s.listener != nil {
		return s.listener.Close()
	}
	return nil
}
