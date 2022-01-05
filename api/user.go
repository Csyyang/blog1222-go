package api

import (
	"blog1222-go/mysql"
	"fmt"

	"github.com/gin-gonic/gin"
)

type user struct {
	Acc      string `json:"account" db:"account"`
	Password string `json:"password" db:"password"`
}

func Login(c *gin.Context) {
	var userG user
	if err := c.ShouldBindJSON(&userG); err != nil {
		c.JSON(500, err)
		return
	}

	db := mysql.Db

	var userS user

	err := db.Get(&userS, "SELECT account,password FROM users WHERE account = ?", userG.Acc)
	if err != nil {

		if err.Error() == "sql: no rows in result set" {
			c.JSON(200, gin.H{"message": "账号或密码错误"})
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

	c.JSON(200, gin.H{"code": "00", "message": "登录成功"})
}

type registerS struct {
	Acc      string `json:"account" db:"account"`   // 账号
	Pass     string `json:"password" db:"password"` // 密码
	UserName string `json:"username" db:"username"` // 用户名
	Avatar   string `json:"avatar" db:"avatar"`     // 图像
	Email    string `json:"email" db:"email"`       // 邮箱
}

func Register(c *gin.Context) {
	var reg registerS

	if err := c.ShouldBindJSON(&reg); err != nil {
		fmt.Println(err)
		c.JSON(500, "serve error")
		return
	}
	fmt.Print(reg.Email)
	c.JSON(200, "ok")
}
