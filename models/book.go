package models

import (
	"fmt"
	"log"
	"web/db"
	"web/utility"
)

type Book struct {
	Id          int    `json:"id"`
	Book_name   string `json:"book_name"`
	Book_number string `json:"book_number"`
	Book_status int    `json:"book_status"`
	Book_time   string `json:"book_time"`
}

// GetBookList TODO 书籍列表
func GetBookList(page, number int) []Book {
	book := []Book{}
	err := db.DB.Select(&book, "select * from book limit ?,?", (page-1)*number, number)
	if err != nil {
		log.Panicln(err)
		return nil
	}
	fmt.Println(book)
	return book
}

// GetBookByNumber TODO  查询 book_number是否存在
func GetBookByNumber(bookNumber int) bool {
	BookNumber := Book{}
	err := db.DB.Get(&BookNumber, "select * from book where book_number=?", bookNumber)
	if err != nil {
		return false
	}
	return true
}

// UpdateBook TODO 修改书籍信息
func UpdateBook(bookName string, bookStatus, bookNumber int) bool {
	_, err := db.DB.Exec("update book set book_name=?,book_status=? where book_number=?", bookName, bookStatus, bookNumber)
	if err != nil {
		panic(&utility.ResponseError{Code: 401, Msg: err.Error()})
	}
	return true

}
