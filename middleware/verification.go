package middleware

import (
	"blog1222-go/jwt"
	"fmt"

	"github.com/gin-gonic/gin"
)

func JwtVer() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-tokne")

		if token == "" {
			c.Redirect(302, "/login")
			return
		}

		j := jwt.NewJwt()

		claim, err := j.ParseToken(token)

		if err != nil {
			fmt.Print(err)
			c.JSON(500, gin.H{"message": err})
		}

		c.Set("claim", claim)
	}
}
