package models

import (
	"log"
	"web/db"
)

type Title struct {
	Id          int    `json:"id"`
	Identifying string `json:"identifying"`
	Title       string `json:"title"`
	Create_time string `json:"create_time"`
	Option_a    string `json:"A"`
	Option_b    string `json:"B"`
	Option_c    string `json:"C"`
	Option_d    string `json:"D"`
	Option_e    string `json:"E"`
	Option_f    string `json:"F"`
	Check_type  string `json:"check_type"`
	Result      string `json:"result"`
	Book_number string `json:"book_number"`
}

func GetTitleList(page, titleNumber, bookNumber int) []Title {
	title := []Title{}
	err := db.DB.Select(&title, "select * from title where book_number=? limit ?,?", bookNumber, (page-1)*titleNumber, titleNumber)
	if err != nil {
		//panic(&utility.ResponseError{Code: 401, Msg: err.Error()})
		log.Panicln(err)
	}
	return title
}
