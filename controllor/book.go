package controllor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web/models"
	"web/utility"
)

// 获取题目
func TitleList(c *gin.Context) {
	bookNumber := c.Query("bookNumber")
	if bookNumber == "" {
		panic(&utility.ResponseError{Code: 101, Msg: "参数不能为空"})
	}
	bookNumbers, _ := strconv.Atoi(bookNumber)
	//TODO 判断 bookNumber的是否有效
	f := models.GetBookByNumber(bookNumbers)
	if !f {
		c.JSON(http.StatusOK, gin.H{
			"code": 501,
			"msg":  "书籍不存在！",
		})
		return
	}
	page := c.DefaultQuery("page", "1")
	pages, _ := strconv.Atoi(page)
	titleNumber := 20
	//TODO 获取题目
	title := models.GetTitleList(pages, titleNumber, bookNumbers)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取数据成功！",
		"data": title,
	})
}

//删除题目

//修改题目

//增加题目
