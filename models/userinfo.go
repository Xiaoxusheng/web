package models

import (
	"fmt"
	"web/db"
	"web/utility"
)

type User struct {
	Id          int    `json:"id" form:"id"`
	Indently    string `json:"indently" form:"indently"`
	Username    string `json:"username" form:"username"  binding:"required,min=3,max=10" `
	Password    string `json:"password" form:"password"  binding:"required,min=5,max=8" `
	Status      int    `json:"status" form:"status"`
	Create_time string `json:"create_time" form:"create_Time"`
	Email       string `json:"email" form:"email" binding:"required,email"`
	Sex         string `json:"sex" form:"sex" binding:"required"`
}

// 注册
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

// 插入
func InsertUser(user *User) error {
	if _, err := db.DB.Exec("insert into userinfo(indently,username,password,email,sex,create_time,status) value(?,?,?,?,?,?,?)", user.Indently, user.Username, user.Password, user.Email, user.Status, user.Create_time, user.Status); err != nil {
		panic(&utility.ResponseError{Code: 401, Msg: err.Error()})
		return err
	}
	return nil
}

// 注销
func Logout(username string) bool {
	_, err := db.DB.Exec("update userinfo set states=4 where username=?", username)
	if err != nil {
		panic(&utility.ResponseError{Code: 401, Msg: err.Error()})
	}
	return true
}

// 登录
func GetUserByUserPwd(use, pwd string) error {
	user := User{}
	if err := db.DB.Get(&user, "select  * from userinfo where username=? and password=?", use, pwd); err != nil {
		return err
	}
	return nil
}
