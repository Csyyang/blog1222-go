package users

import (
	"blog1222-go/mysql"
	"blog1222-go/response"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/*
* 重置密码
**/
func ResetPassword(c *gin.Context) {
	session := sessions.Default(c)
	account := session.Get("account")

	db := mysql.Db
	sqlStr := "update users set password= ? where account = ?"

	_, err := db.Exec(sqlStr, "111111", account)
	if err != nil {
		response.BadRes(c, err.Error())
		return
	}

	// c.JSON(200, gin.H{
	// 	"code":    "00",
	// 	"message": "密码重置成功",
	// })
	response.SuccessRes(c, "密码重置成功")
}
