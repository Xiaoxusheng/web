package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"web/router"
)

func main() {
	r := router.Router()
	//开启日志颜色
	gin.ForceConsoleColor()
	//// 记录到文件。
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r.Static("/index", "./view")

	err := r.Run(":8080")
	if err != nil {
		log.Panicln("启动服务出错" + err.Error())
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
