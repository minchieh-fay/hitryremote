package main

type TargetInfo struct {
	IP          string `json:"ip"`
	Port        int    `json:"port"`
	Description string `json:"description"`
	LocalPort   int    `json:"local_port"` // 本地映射到目标的端口
}

type ClientTargetList struct {
	ClientID string       `json:"client_id"` // 客户端ID
	Name     string       `json:"name"`      // 客户端使用人的名字,如: 张三
	Phone    string       `json:"phone"`     // 客户端使用人的电话,如: 13800138000
	Targets  []TargetInfo `json:"targets"`   // 目标列表
}

type Server struct {
	clientTargetList map[string]*ClientTargetList
}

func NewServer() *Server {
	return &Server{
		clientTargetList: make(map[string]*ClientTargetList),
	}
}

func (s *Server) Run() {

}
