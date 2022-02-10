package users

import (
	"blog1222-go/email"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/*
* 发送邮件
**/
func Reset(c *gin.Context) {
	cookie, _ := c.Cookie("mysession")
	session := sessions.Default(c)

	eml := session.Get("email")
	// 发送邮件
	errs := email.SendGoMail([]string{eml.(string)}, "重置邮件", "点击重置密码，重置后密码为111111：http://yangyangcsy.cn/#/reset?token="+cookie)
	if errs != nil {
		fmt.Println(errs.Error())
	}

	c.JSON(200, gin.H{
		"code":    "00",
		"message": "邮件发送成功",
	})
}
