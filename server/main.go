package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var S *Server

func main() {

	fmt.Println("HiTryRemote 服务器启动")

	// 设置 Gin 模式
	gin.SetMode(gin.ReleaseMode)

	S = NewServer()
	S.Run()

	S.RunWeb(8080)

	log.Println("服务器已关闭")
}
