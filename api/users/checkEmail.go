package users

import (
	"blog1222-go/mysql"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 邮箱验证
type emial struct {
	Email string `json:"email"`
}

func CheckEmail(c *gin.Context) {
	var eml emial

	if err := c.ShouldBindJSON(&eml); err != nil {
		fmt.Println(err)
		c.JSON(500, "serve error")
		return
	}

	fmt.Println(eml.Email)

	db := mysql.Db
	var user struct {
		Acc string `json:"account" db:"account"` // 账号
	}

	// 查询账号
	err := db.Get(&user, "SELECT account FROM users WHERE email = ?", eml.Email)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(200, gin.H{
				"code":    "01",
				"message": "账户不存在"})
			return
		}

		fmt.Printf("%v", err)
		c.JSON(500, "bad")
		return
	}

	// 生成jwt
	// ccc, err := jwt.NewJwt().GenerateToken(user.Acc)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	c.JSON(500, err.Error())
	// 	return
	// }

	session := sessions.Default(c)
	session.Set("account", user.Acc)
	session.Set("email", eml.Email)
	session.Save()

	// // 发送邮件
	// errs := email.SendGoMail([]string{eml.Email}, "重置邮件", "点击重置密码，重置后密码为111111：http://yangyangcsy.cn/reset?token=")
	// if errs != nil {
	// 	fmt.Println(err.Error())
	// }

	c.JSON(200, gin.H{
		"code":    "00",
		"message": "邮箱正确",
	})
}
