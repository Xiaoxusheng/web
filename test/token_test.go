package test

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"testing"
	"time"
	"web/db"
)

// GetToken 生成token
func Test_GetToken(T *testing.T) {
	mySigningKey := []byte("UFFij ji^&^*k*&&j")
	ctx := context.Background()
	type MyCustomClaims struct {
		Indently string `json:"indently"`
		jwt.RegisteredClaims
	}

	//// Create claims with multiple fields populated
	//claims := MyCustomClaims{
	//	"wer",
	//	jwt.RegisteredClaims{
	//		// A usual scenario is to set the expiration time relative to the current time
	//		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	//		IssuedAt:  jwt.NewNumericDate(time.Now()),
	//		NotBefore: jwt.NewNumericDate(time.Now()),
	//		Issuer:    "test",
	//		Subject:   "somebody",
	//		ID:        "1",
	//		Audience:  []string{"somebody_else"},
	//	},
	//}

	//fmt.Printf("foo: %v\n", claims.Indently)

	// Create claims while leaving out some of the optional fields
	claims := MyCustomClaims{
		"bar",
		jwt.RegisteredClaims{
			// Also fixed dates can be used for the NumericDate
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	_, err = db.Rdb.Set(ctx, "bar", ss, time.Second*60).Result()
	if err != nil {
		log.Panicln("redis", err)
	}
	fmt.Printf("%v %v", ss, err)

}
