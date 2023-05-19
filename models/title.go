package models

import "time"

type Title struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Create_time time.Time `json:"create_time"`
	OptionA     string    `json:"optionA"`
	OptionB     string    `json:"optionB"`
	OptionC     string    `json:"optionC"`
	OptionD     string    `json:"optionD"`
	OptionE     string    `json:"optionE"`
	OptionF     string    `json:"optionF"`
	Check_type  string    `json:"check_type"`
	Result      string    `json:"result"`
	Book_number string    `json:"book_number"`
}

func GetList() {
}
