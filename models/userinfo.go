package models

import (
	"fmt"
	"log"
	"web/db"
	"web/utility"
)

type User struct {
	Id          string `json:"id" form:"id"`
	Indently    string `json:"indently" form:"indently"`
	Username    string `json:"username" form:"username"  binding:"required,min=3,max=10" `
	Password    string `json:"password" form:"password"  binding:"required,min=5,max=8" `
	Status      int    `json:"status" form:"status"`
	Create_time string `json:"create_time" form:"create_Time"`
	Email       string `json:"email" form:"email" binding:"required,email"`
	Sex         string `json:"sex" form:"sex" binding:"required"`
}

func Get() {
	user := []User{}
	err := db.DB.Select(&user, "SELECT * FROM userinfo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}

func GetUsername(username string) bool {
	user := User{}
	if err := db.DB.Get(&user, "select * from userinfo where username=?", username); err != nil {
		return false
	}
	return true
}

func GetEmail(email string) bool {
	user := User{}
	if err := db.DB.Get(&user, "select * from userinfo where email=?", email); err != nil {
		return false
	}
	fmt.Println(user)
	return true
}

func InsertUser(user *User) error {
	if _, err := db.DB.Exec("insert into userinfo(indently,username,password,email,sex,create_time,status) value(?,?,?,?,?,?,?)", user.Indently, user.Username, user.Password, user.Email, user.Status, user.Create_time, user.Status); err != nil {
		panic(&utility.ResponseError{Code: 401, Msg: err.Error()})
		return err
	}
	return nil
}
