package controllor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web/models"
	"web/utility"
)

// TODO 获取题目
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

// TODO 删除题目
func DeleteTitle(c *gin.Context) {
	identifying := c.Query("identifying")
	if identifying == "" {
		panic(&utility.ResponseError{Code: 101, Msg: "参数不能为空！"})
	}
	f := models.GetTitle(identifying)
	if !f {
		panic(&utility.ResponseError{Code: 501, Msg: "数据不存在！"})
	}
	b := models.DeleteTitle(identifying)
	if b {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "删除成功！",
		})
	}
}

// TODO  修改题目
func UpdateTitle(c *gin.Context) {
	bookNumber, _ := strconv.Atoi(c.PostForm("book_number"))
	identifying := c.PostForm("identifying")
	title := c.PostForm("title")
	option_a := c.PostForm("option_a")
	option_b := c.PostForm("option_b")
	option_c := c.PostForm("option_c")
	option_d := c.PostForm("option_d")
	result := c.PostForm("result")
	option_e := c.DefaultPostForm("option_e", "")
	option_f := c.DefaultPostForm("option_f", "")
	if identifying == "" || (bookNumber/1000000 < 1 && bookNumber/1000000 > 10) || title == "" || option_a == "" || option_b == "" || option_c == "" || option_d == "" || result == "" {
		panic(&utility.ResponseError{Code: 101, Msg: "参数不能为空！"})
	}
	f := models.GetBookByNumber(bookNumber)
	if !f {
		panic(&utility.ResponseError{Code: 501, Msg: "数据不存在！"})
	}
	b := models.UpdateTitle(title, option_a, option_b, option_c, option_d, option_e, option_f, result, identifying, bookNumber)
	if b {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "修改成功！",
		})
	}
}

//增加题目
