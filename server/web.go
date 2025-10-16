package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 将../html弄到embed

//go:embed html
var templateFS embed.FS

func (s *Server) RunWeb(port int) {
	router := gin.Default()
	SetupRoutes(s, router)
	router.Run(fmt.Sprintf(":%d", port))
}

func SetupRoutes(s *Server, r *gin.Engine) {
	// 设置模板
	tmpl := template.Must(template.ParseFS(templateFS, "html/*"))
	r.SetHTMLTemplate(tmpl)

	// 往clientTargetList中模拟一些数据
	s.clientTargetList["123"] = &ClientTargetList{
		ClientID: "123",
		Name:     "张三",
		Phone:    "13800138000",
		Targets: []TargetInfo{
			{IP: "192.168.1.100", Port: 8080, Description: "本地测试服务器", LocalPort: 8080},
			{IP: "192.168.1.101", Port: 8081, Description: "本地测试服务器2", LocalPort: 8081},
			{IP: "192.168.1.102", Port: 8082, Description: "本地测试服务器3", LocalPort: 8082},
			{IP: "192.168.1.103", Port: 8083, Description: "本地测试服务器4", LocalPort: 8083},
		},
	}
	s.clientTargetList["456"] = &ClientTargetList{
		ClientID: "456",
		Name:     "李四",
		Phone:    "13800138001",
		Targets: []TargetInfo{
			{IP: "192.168.1.104", Port: 8084, Description: "本地测试服务器5", LocalPort: 8084},
		},
	}
	s.clientTargetList["789"] = &ClientTargetList{
		ClientID: "789",
		Name:     "王五",
		Phone:    "13800138002",
		Targets: []TargetInfo{
			{IP: "192.168.1.105", Port: 8085, Description: "本地测试服务器6", LocalPort: 8085},
		},
	}
	s.clientTargetList["101"] = &ClientTargetList{
		ClientID: "101",
		Name:     "赵六",
		Phone:    "13800138003",
		Targets: []TargetInfo{
			{IP: "192.168.1.106", Port: 8086, Description: "本地测试服务器7", LocalPort: 8086},
		},
	}
	s.clientTargetList["102"] = &ClientTargetList{
		ClientID: "102",
		Name:     "孙七",
		Phone:    "13800138004",
		Targets: []TargetInfo{
			{IP: "192.168.1.107", Port: 8087, Description: "本地测试服务器8", LocalPort: 8087},
		},
	}
	s.clientTargetList["103"] = &ClientTargetList{
		ClientID: "103",
		Name:     "周八",
		Phone:    "13800138005",
		Targets: []TargetInfo{
			{IP: "192.168.1.108", Port: 8088, Description: "本地测试服务器9", LocalPort: 8088},
		},
	}

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title":            "HiTryRemote 管理平台",
			"clientTargetList": s.clientTargetList,
		})
	})
}
