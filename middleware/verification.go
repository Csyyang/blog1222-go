package middleware

import (
	"blog1222-go/jwt"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/*
*  jwt验证
**/
func JwtVer() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		fmt.Println("token")

		if token == "" {
			c.Redirect(302, "http://yangyangcsy.cn/#/login")
			c.Abort()
			return
		}

		j := jwt.NewJwt()

		claim, err := j.ParseToken(token)

		if err != nil {
			fmt.Print(err)
			c.JSON(500, gin.H{"message": err})
			c.Abort()
		}

		c.Set("token", claim)
		c.Next()
	}
}

/*
* session验证
**/
func SessVer() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		fmt.Println(session.Get("account"))
		if session.Get("account") == nil {
			c.JSON(302, gin.H{
				"code": "03",
				"url":  "/?date=" + strconv.FormatInt(time.Now().Unix(), 10) + "#/login",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
