package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetConnect() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := "http://www.baidu.com" // 这里使用百度作为测试URL
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("服务器网络连接失败")
			c.Abort()
		}
		if resp.StatusCode == http.StatusOK {
			fmt.Println("服务器网络连接正常")
		} else {
			fmt.Println("服务器网络连接异常")
			c.Abort()
		}
	}
}
