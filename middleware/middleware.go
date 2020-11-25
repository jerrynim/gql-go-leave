package middleware

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	database "github.com/jerrynim/gql-leave/db"
	"github.com/jerrynim/gql-leave/graph/model"
	"github.com/jerrynim/gql-leave/jwt"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
    name string
}


func Middleware() gin.HandlerFunc {
    
    return func(c *gin.Context) {
       token := c.Request.Header.Get("Authorization")
        if  token == "" {
            c.Next()
            return
        }

        db, err := database.GetDatabase()

        if err != nil {
            log.Println("데이터 베이스 연결 에러")
            c.Next()
            return
        }
        defer db.Close()

        var user model.User
        userId,parsingErr:=jwt.ParseToken(token)

        if parsingErr!=nil || userId == ""{
            log.Println("토큰 파싱 에러")
            c.Next()
            return
        }
        if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
            log.Println("유저 조회 에러")
            c.Next()
            return
        }
        ctx := context.WithValue(c.Request.Context(), "user", user)
		c.Request = c.Request.WithContext(ctx)
        c.Next()
}
}