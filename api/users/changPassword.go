package users

import (
	"blog1222-go/mysql"
	"blog1222-go/response"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/*
* 更换密码
**/
type users struct {
	Acc      string `json:"account"`
	Pass     string `json:"password"`
	NewPass  string `json:"newpassword"`
	NewPass2 string `json:"newpassword2"`
}

type dbPass struct {
	Pass string `db:"password"`
}

func ChangPassword(c *gin.Context) {
	// session := sessions.Default(c)
	// account := session.Get("account")

	var usermessage users
	if err := c.ShouldBindJSON(&usermessage); err != nil {
		response.BadRes(c, err.Error())
	}

	fmt.Println(usermessage)

	db := mysql.Db
	sqlPas := "SELECT password FROM users WHERE account = ?"

	var dbPassworld dbPass

	err := db.Get(&dbPassworld, sqlPas, usermessage.Acc)
	if err != nil {
		response.BadRes(c, err.Error())
		return
	}

	if dbPassworld.Pass != usermessage.Pass {
		response.BadRes(c, "密码错误")
		return
	}

	if usermessage.NewPass != usermessage.NewPass2 {
		response.BadRes(c, "两次输入密码不一致")
		return
	}

	sqlStr := "update users set password= ? where account = ?"

	_, err2 := db.Exec(sqlStr, usermessage.NewPass, usermessage.Acc)
	if err2 != nil {
		response.BadRes(c, err2.Error())
		return
	}

	session := sessions.Default(c)
	session.Clear()
	session.Save()

	response.SuccessRes(c, "修改密码成功")
}
