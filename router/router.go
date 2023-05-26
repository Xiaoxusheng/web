package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"web/controllor"
	"web/middleware"
)

func Router() *gin.Engine {
	r := gin.Default()
	//错误处理中间件
	r.Use(middleware.ErrorHandler())

	//限速中间件
	r.Use(middleware.LimitIP())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	r.Use(cors.New(config))
	r.Use(cors.Default())

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
		user.GET("/titleDelete", middleware.ParseToken(), controllor.DeleteTitle)
		user.POST("/titleUpdate", middleware.ParseToken(), controllor.UpdateTitle)
	}

	return r
}
