package api

import (
	"blog1222-go/email"
	"blog1222-go/mysql"
	"blog1222-go/response"
	"fmt"
	"reflect"

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

/*
* 登出
**/
func LogOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.JSON(200, gin.H{
		"code":    "00",
		"message": "登出成功",
	})
}

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

/*
* 重置密码
**/

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

func Test(c *gin.Context) {

	c.JSON(200, "ok的")
}

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
