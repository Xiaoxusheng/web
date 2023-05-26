package test

import (
	"fmt"
	"testing"
	"time"
)

func Test_time(t *testing.T) {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Hour())

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	fmt.Println(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), 0, 0, 0, time.Now().Location()))
}
