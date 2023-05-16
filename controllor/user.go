package controllor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"web/models"
	"web/utility"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	sex := c.PostForm("sex")
	if username == "" || password == "" || email == "" || sex == "" {
		//c.JSON(http.StatusOK, gin.H{
		//	"code": 101,
		//	"msg":  "参数不能为空！",
		//})
		panic(&utility.ResponseError{Code: 201, Msg: "参数不能为空！"})
		return
	}
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		panic(&utility.ResponseError{Code: 201, Msg: err.Error()})
		return
	}
	t := models.GetUsername(username)
	if t {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "用户名已经存在！",
		})
		return
	}
	e := models.GetEmail(email)
	if e {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "邮箱已经注册！",
		})
		return
	}
	if err := models.InsertUser(&models.User{Indently: utility.SetUuid(), Username: username, Password: utility.GetMa5(password), Email: email, Status: 0, Sex: sex, Create_time: time.Now().Format("2006-01-02 15:04:05")}); err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功！",
	})

}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 101,
			"msg":  "参数不能为空！",
		})
		return
	}

}
