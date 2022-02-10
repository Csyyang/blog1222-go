package users

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LogOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.JSON(200, gin.H{
		"code":    "00",
		"message": "登出成功",
	})
}
