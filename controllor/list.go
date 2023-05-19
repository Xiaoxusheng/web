package controllor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web/models"
	"web/utility"
)

func BooKList(c *gin.Context) {
	book := 20
	page := c.DefaultQuery("page", "1")
	if page == "" {
		panic(&utility.ResponseError{Code: 101, Msg: "参数不能为空"})
	}
	p, err := strconv.Atoi(page)
	if err != nil {
		panic(&utility.ResponseError{Code: 500, Msg: "系统错误！"})
	}
	list := models.GetBookList(p, book)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取数据成功！",
		"data": list,
	})
}
