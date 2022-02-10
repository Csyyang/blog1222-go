package users

import (
	"blog1222-go/mysql"
	"blog1222-go/response"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type user struct {
	Acc      string `json:"account" db:"account"`
	Password string `json:"password" db:"password"`
}

/*
* 登录
**/
func Login(c *gin.Context) {
	session := sessions.Default(c)

	var userG user
	if err := c.ShouldBindJSON(&userG); err != nil {
		response.BadRes(c, err.Error())
		// c.JSON(500, err)
		return
	}

	// 数据库操作
	db := mysql.Db
	var userS struct {
		Account  string `json:"account"`
		Password string `json:"password"`
		Username string `json:"username"`
		Avatar   string `json:"avatar"`
	}
	err := db.Get(&userS, "SELECT account,password,username,avatar FROM users WHERE account = ?", userG.Acc)
	if err != nil {

		if err.Error() == "sql: no rows in result set" {
			// c.JSON(200, gin.H{"message": "账号或密码错误"})
			response.BadRes(c, "账号或密码错误")
			return
		}

		fmt.Printf("%v", err)
		c.JSON(500, "bad")
		return
	}

	if userG.Password != userS.Password {
		c.JSON(200, gin.H{"message": "账号或密码错误"})
		return
	}
	// 生成jwt
	// ccc, err := jwt.NewJwt().GenerateToken(userG.Acc)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	c.JSON(500, err.Error())
	// 	return
	// }

	// 生成session，存入cookie
	session.Set("account", userS.Account)
	session.Save()

	c.JSON(200, gin.H{
		"code":    "00",
		"message": "登录成功",
		"context": gin.H{
			"userData": gin.H{
				"account":  userS.Account,
				"username": userS.Username,
				"avatar":   userS.Avatar,
			},
		},
	})
}
