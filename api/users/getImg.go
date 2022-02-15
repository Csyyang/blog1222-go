package users

import (
	"blog1222-go/mysql"
	"blog1222-go/response"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type userType struct {
	Avatar string `db:"avatar" json:"avatar"`
}

func GetUserImg(c *gin.Context) {
	account := c.Query("account")
	var user userType

	db := mysql.Db
	sqlStr := "select avatar from users where account = ?"
	err := db.Get(&user, sqlStr, account)
	if err != nil {
		response.BadRes(c, err.Error())
		return
	}

	index := strings.LastIndex(user.Avatar, "/")
	fileName := user.Avatar[index:]
	fmt.Println(fileName)
	c.File("./images/" + fileName)
}
