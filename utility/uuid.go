package utility

import (
	"crypto/md5"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

// SetUuid 生成uuid
func SetUuid() string {
	fmt.Println(uuid.NewV4().String())
	return uuid.NewV4().String()
}

func GetMa5(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}
