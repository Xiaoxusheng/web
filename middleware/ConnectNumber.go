package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

func GetConnect() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录当前连接数
		var currentConnections int64
		// 打印当前连接数
		go func() {
			for {
				log.Printf("当前连接数 connections: %d\n", atomic.LoadInt64(&currentConnections))
				time.Sleep(time.Second)
			}
		}()
		// 在连接状态改变时更新连接数
		atomic.AddInt64(&currentConnections, 1)
		c.Writer.(http.CloseNotifier).CloseNotify()
		atomic.AddInt64(&currentConnections, -1)
	}
}
