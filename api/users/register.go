package users

import (
	"blog1222-go/mysql"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

/*
* 注册
**/
type registerS struct {
	Acc      string `json:"account"`   // 账号
	Pass     string `json:"password"`  // 密码
	Pass2    string `json:"password2"` // 重复密码
	UserName string `json:"username" ` // 用户名
	// Avatar   string `json:"avatar"`    // 头像
	Email string `json:"email"` // 邮箱
}

func Register(c *gin.Context) {
	var reg registerS

	if err := c.ShouldBindJSON(&reg); err != nil {
		fmt.Println(err)
		c.JSON(500, "serve error")
		return
	}
	fmt.Println(reg)

	// 表单校验
	t := reflect.TypeOf(reg)
	v := reflect.ValueOf(reg)

	for k := 0; k < t.NumField(); k++ {
		if v.Field(k).Interface() == "" {
			c.JSON(200, gin.H{"code": "01", "message": "表单不完整"})
			return
		}
	}

	if reg.Pass != reg.Pass2 {
		c.JSON(200, gin.H{"code": "01", "message": "密码不一致"})
		return
	}

	// 入库
	db := mysql.Db
	sqlStr := "insert into users (account,password,username,email) values (?,?,?,?)"

	_, err := db.Exec(sqlStr, reg.Acc, reg.Pass, reg.UserName, reg.Email)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, "bad")
		return
	}

	fmt.Print(reg.Email)
	c.JSON(200, gin.H{
		"code":    "00",
		"message": "success",
	})
}
