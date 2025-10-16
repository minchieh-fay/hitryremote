package util

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/quic-go/quic-go"
)

const VERSION = 1

// 这里会定义一些结构体。用于传输数据，json格式

// 消息类型常量
const (
	MSG_TYPE_SERVERREG         = 0 // 服务端注册
	MSG_TYPE_SERVERREG_ACK     = 1 // 服务端注册回复
	MSG_TYPE_CLIENTREG         = 2 // 客户端注册
	MSG_TYPE_CLIENTREG_ACK     = 3 // 客户端注册回复
	MSG_TYPE_PING              = 4 // ping消息
	MSG_TYPE_PONG              = 5 // pong回复
	MSG_TYPE_TCP_CHANNEL       = 6 // tcp隧道
	MSG_TYPE_CLIENT_TARGETLIST = 7 // 客户端上报的他那边开放给server可以远程的局域网地址列表

	MSG_TYPE_ERROR = 99 // 错误消息
)

// 基础消息结构
type QuicMessage struct {
	Type int         `json:"type"` // 消息类型
	Data interface{} `json:"data"` // 消息数据
}

func NewQuicMessage(typ int, data interface{}) *QuicMessage {
	return &QuicMessage{
		Type: typ,
		Data: data,
	}
}

func (m *QuicMessage) ToBuffer() []byte {
	buf, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	// 前面4个字节是len
	lenBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(lenBuf, uint32(len(buf)))
	buf = append(lenBuf, buf...)
	return buf
}

func QuicMessageFromStream(stream *quic.Stream) *QuicMessage {
	buf := make([]byte, 4)
	_, err := stream.Read(buf)
	if err != nil {
		return nil
	}
	len := binary.BigEndian.Uint32(buf)
	buf = make([]byte, len)
	_, err = stream.Read(buf)
	if err != nil {
		fmt.Printf("读取消息长度失败: %v\n", err)
		return nil
	}
	var msg QuicMessage
	err = json.Unmarshal(buf, &msg)
	if err != nil {
		fmt.Printf("解析消息失败: %v\n", err)
		return nil
	}

	// 根据消息类型反序列化Data字段
	switch msg.Type {
	case MSG_TYPE_SERVERREG:
		var serverRegMsg ServerRegMessage
		dataBytes, _ := json.Marshal(msg.Data)
		json.Unmarshal(dataBytes, &serverRegMsg)
		msg.Data = &serverRegMsg
	case MSG_TYPE_SERVERREG_ACK:
		var serverRegAckMsg ServerRegAckMessage
		dataBytes, _ := json.Marshal(msg.Data)
		json.Unmarshal(dataBytes, &serverRegAckMsg)
		msg.Data = &serverRegAckMsg
	case MSG_TYPE_PING:
		var pingMsg PingMessage
		dataBytes, _ := json.Marshal(msg.Data)
		json.Unmarshal(dataBytes, &pingMsg)
		msg.Data = &pingMsg
	case MSG_TYPE_PONG:
		var pongMsg PongMessage
		dataBytes, _ := json.Marshal(msg.Data)
		json.Unmarshal(dataBytes, &pongMsg)
		msg.Data = &pongMsg
	case MSG_TYPE_TCP_CHANNEL:
		var tcpChannelMsg TCPChannelMessage
		dataBytes, _ := json.Marshal(msg.Data)
		json.Unmarshal(dataBytes, &tcpChannelMsg)
		msg.Data = &tcpChannelMsg
	case MSG_TYPE_CLIENT_TARGETLIST:
		var clientTargetListMsg ClientTargetListMessage
		dataBytes, _ := json.Marshal(msg.Data)
		json.Unmarshal(dataBytes, &clientTargetListMsg)
		msg.Data = &clientTargetListMsg
	}

	return &msg
}

// MSG_TYPE_SERVERREG
type ServerRegMessage struct {
	Version int `json:"version"` // 版本号
}

// MSG_TYPE_SERVERREG_ACK
type ServerRegAckMessage struct {
	Version int    `json:"version"` // 版本号
	Result  bool   `json:"result"`  // 是否成功
	Reason  string `json:"reason"`  // 原因
}

// MSG_TYPE_CLIENTREG
type ClientRegMessage struct {
	Version  int    `json:"version"`   // 版本号
	ClientID string `json:"client_id"` // 客户端ID
}

// MSG_TYPE_CLIENTREG_ACK
type ClientRegAckMessage struct {
	Version int    `json:"version"` // 版本号
	Result  bool   `json:"result"`  // 是否成功
	Reason  string `json:"reason"`  // 原因
}

// MSG_TYPE_TCP_CHANNEL
type TCPChannelMessage struct {
	ClientID   string `json:"client_id"`   // 客户端ID
	TargetAddr string `json:"target_addr"` // 目标地址
	TargetPort int    `json:"target_port"` // 目标端口
}

// Ping消息结构
type PingMessage struct {
}

// Pong回复结构
type PongMessage struct {
}

type TargetInfo struct {
	IP          string `json:"ip"`
	Port        int    `json:"port"`
	Description string `json:"description"`
}

// MSG_TYPE_CLIENT_TARGETLIST
type ClientTargetListMessage struct {
	ClientID string       `json:"client_id"` // 客户端ID
	Name     string       `json:"name"`      // 客户端使用人的名字,如: 张三
	Phone    string       `json:"phone"`     // 客户端使用人的电话,如: 13800138000
	Targets  []TargetInfo `json:"targets"`   // 目标列表
}
