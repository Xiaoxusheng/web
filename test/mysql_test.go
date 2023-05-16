package test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"testing"
	"time"
)

var DB *sqlx.DB

func Test_mysql(t *testing.T) {

	database, err := sqlx.Open("mysql", "root:admin123@tcp(127.0.0.1:3306)/project")
	if err != nil {
		log.Println("open mysql failed,", err)
		return
	}
	err = database.Ping()
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		log.Panicln("mysql数据库连接失败" + err.Error())
	}
	/*
		database.SetConnMaxLifetime()需要确保在 MySQL 服务器、操作系统或其他中间件关闭连接之前，驱动程序安全地关闭连接。
		由于某些中间件将空闲连接关闭 5 分钟，因此我们建议超时时间短于 5 分钟。此设置还有助于负载平衡和更改系统变量
	*/
	database.SetConnMaxIdleTime(time.Minute * 3)
	// 最大连接数
	database.SetMaxOpenConns(100)
	// 闲置连接数
	database.SetMaxIdleConns(10)

	DB = database

	log.Println("Mysql服务启动成功")
	type User struct {
		Id          string `json:"id"`
		Indently    string `json:"indently" `
		Username    string `json:"username"  `
		Password    string `json:"password" `
		Status      int    `json:"status" `
		Create_time string `json:"create_time" `
		Email       string `json:"email" `
		Sex         string `json:"sex"`
	}
	user := []User{}

	err = DB.Select(&user, "SELECT * FROM userinfo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)

}
