package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ErrorHandler 捕获错误中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 处理错误
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err,
					"code":  http.StatusInternalServerError,
				})
			}
		}()
		c.Next()
	}
}
