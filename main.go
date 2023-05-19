package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"web/router"
)

func main() {
	r := router.Router()

	gin.ForceConsoleColor()
	// 记录到文件。
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
