package test

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"testing"
	"web/db"
	"web/utility"
)

func Test_ParseToken(t *testing.T) {
	user := utility.MyCustomClaims{}
	// Token from another example.  This token is expired
	tokens := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpbmRlbnRseSI6ImJhciIsImlzcyI6InRlc3QiLCJleHAiOjE2ODQ0MTUzNzB9.EwpKZ4YnJbRS13CjVR3hAEkpBSmg77jUtmECjdDKQM4"
	ctx := context.Background()
	i, err := db.Rdb.Exists(ctx, "bar").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("i", i)

	result, err := db.Rdb.Get(ctx, "bar").Result()
	if err != nil {
		log.Panicln(err)
		return
	}
	fmt.Println("tokenString", result)
	if result != tokens {
		fmt.Println(" 账号在其他地方登录，只允许一台设备登录！")
	}

	token, err := jwt.ParseWithClaims(tokens, &user, func(token *jwt.Token) (interface{}, error) {
		return utility.MySigningKey, nil
	})

	fmt.Println(user)
	if token.Valid {
		fmt.Println("You look nice today")
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		fmt.Println("That's not even a token")
		panic("")
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		// Invalid signature
		fmt.Println("Invalid signature")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		// Token is either expired or not active yet
		fmt.Println("Timing is everything")
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}

}
