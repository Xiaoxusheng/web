package controllor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web/models"
	"web/utility"
)

// TODO 获取书籍列表
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

// TODO 修改书籍数据
func BookUpdate(c *gin.Context) {
	bookName := c.PostForm("book_name")
	bookStatus := c.PostForm("book_status")
	bookNumber := c.PostForm("book_number")
	if bookName == "" || bookStatus == "" || bookNumber == "" {
		panic(&utility.ResponseError{Code: 101, Msg: "参数不能为空！"})
	}
	BookStatus, _ := strconv.Atoi(bookStatus)
	BookNumber, _ := strconv.Atoi(bookNumber)
	b := models.GetBookByNumber(BookNumber)
	if !b {
		c.JSON(http.StatusOK, gin.H{
			"code": 501,
			"msg":  "书籍号不存在！",
		})
		return
	}
	f := models.UpdateBook(bookName, BookStatus, BookNumber)
	if f && b {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "修改成功！",
		})
	}
}
