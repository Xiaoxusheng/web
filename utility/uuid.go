package utility

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
	"time"
	"web/db"
)

// SetUuid 生成uuid
func SetUuid() string {
	fmt.Println(uuid.NewV4().String())
	return uuid.NewV4().String()
}

// GetMa5 生成MD5
func GetMa5(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GetToken 生成token
func GetToken(username string) string {
	ctx := context.Background()
	// Create claims with multiple fields populated
	claims := MyCustomClaims{
		username,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "mr.lei",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}
	//生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(MySigningKey)
	//i, err := db.Rdb.Exists(ctx, identif).Result()
	//if err != nil {
	//	fmt.Println(err)
	//	return ""
	//}
	//if i == 1 {
	//	panic(&ResponseError{
	//		Code: 1,
	//		Msg:  "账号已经登录，只允许一台设备登录！",
	//	})
	//}
	result, err := db.Rdb.Set(ctx, username, ss, time.Hour*24).Result()
	if err != nil {
		return ""
	}
	fmt.Println(result)

	fmt.Printf("%v ", ss)
	return ss
}
