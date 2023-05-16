package router

import (
	"github.com/gin-gonic/gin"
	"web/controllor"
	"web/middleware"
)

func Router() *gin.Engine {
	r := gin.Default()
	//r.Use(middleware.GetConnect())
	r.Use(middleware.ErrorHandler())

	user := r.Group("/user")
	{
		user.POST("/login", controllor.Login)
		user.POST("/register", controllor.Register)

	}

	return r
}
