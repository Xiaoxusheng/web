package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/net/context"
	"net/http"
	"strings"
	"web/db"
	"web/utility"
)

// ParseToken TODO 解析token中间件
func ParseToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Token from another example.  This token is expired
		tokenString := strings.Split(c.GetHeader("Authorization"), "Bearer ")[1]
		//fmt.Println(tokenString)
		ctx := context.Background()
		user := utility.MyCustomClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &user, func(token *jwt.Token) (interface{}, error) {
			return utility.MySigningKey, nil
		})
		if user.Username != "" {
			c.Set("username", user.Username)
		}
		fmt.Println("username", user.Username)
		result, err := db.Rdb.Get(ctx, user.Username).Result()
		if err != nil {
			c.Abort()
			panic(&utility.ResponseError{Code: 301, Msg: "token失效或过期！"})
		}
		fmt.Println(result)
		if result != tokenString {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "账号在其他地方登录，只允许一台设备登录！",
			})
			c.Abort()
			return
		}

		if token.Valid {
			fmt.Println("验证通过")
		} else if errors.Is(err, jwt.ErrTokenMalformed) {
			fmt.Println("That's not even a token")
			c.Abort()
			panic(&utility.ResponseError{Code: 301, Msg: "这不是一个token"})
		} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			// Invalid signature
			fmt.Println("Invalid signature")
			c.Abort()
			panic(&utility.ResponseError{Code: 301, Msg: "无效的签名"})
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
			c.Abort()
			panic(&utility.ResponseError{Code: 301, Msg: "token失效或过期！"})
		} else {
			fmt.Println("Couldn't handle this token:", err)
			c.Abort()
			panic(&utility.ResponseError{Code: 301, Msg: "未知错误！"})
		}

	}
}
