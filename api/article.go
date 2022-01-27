package api

import (
	"blog1222-go/mysql"
	"blog1222-go/response"
	"fmt"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/**
 *  新增文章
 */
type articleT struct {
	Aticle_title    string `json:'title'`
	Article_context string `json:context`
	Article_type    string `json:lable`
	Article_image   string `json:image`
}

func NewArticle(c *gin.Context) {
	session := sessions.Default(c)
	account := session.Get("account")

	// 序列化json
	var article articleT
	if err := c.ShouldBindJSON(&article); err != nil {
		fmt.Println(err)
		c.JSON(500, "serve error")
		return
	}

	// 暂不开放
	if account != "693765678" {
		response.BadRes(c, "权限不足")
		return
	}

	// 当前时间
	timeObj := time.Now()
	str := timeObj.Format("2006-01-02 03:04:05")

	// 存入数据库
	db := mysql.Db
	sqlstr := "insert into article (article_title,article_create_date,article_modify_date,article_context,account,article_type,article_topping,article_image) values (?,?,?,?,?,?,?,?)"
	_, err := db.Exec(sqlstr, article.Aticle_title, str, str, article.Article_context, account, article.Article_type, '0', article.Article_image)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, "bad")
		return
	}

	response.SuccessRes(c, "添加成功")
}

func FixArticle(c *gin.Context) {}
