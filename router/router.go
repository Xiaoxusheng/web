package router

import (
	"github.com/gin-gonic/gin"
	"web/controllor"
	"web/middleware"
)

func Router() *gin.Engine {
	r := gin.Default()

	//错误处理中间件
	r.Use(middleware.ErrorHandler())
	//验证网络是否连接
	r.Use(middleware.GetConnect())

	user := r.Group("/user")
	{
		user.POST("/login", controllor.Login)
		user.GET("/logout", controllor.Logout)
		user.POST("/register", controllor.Register)
		user.GET("/list", middleware.ParseToken(), controllor.BooKList)
		user.POST("/bookUpdate", controllor.BookUpdate)
		user.GET("/book", middleware.ParseToken(), controllor.TitleList)

	}

	administrator := r.Group("/administrator")
	{
		administrator.GET("/add")
	}
	return r
}
