package models

import (
	"fmt"
	"log"
	"web/db"
)

type Book struct {
	Id          int    `json:"id"`
	Book_name   string `json:"book_name"`
	Book_number string `json:"book_number"`
	Book_status int    `json:"book_status"`
	Book_time   string `json:"book_time"`
}

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
