package api

import (
	"blog1222-go/mysql"
	"blog1222-go/response"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/**
 *  新增文章
 */
type articleT struct {
	Aticle_title    string `json:"title"`
	Article_context string `json:"context"`
	Article_type    string `json:"lable"`
	Article_image   string `json:"image"`
	Article_brief   string `json:"brief"`
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
	sqlstr := "insert into article (article_title,article_create_date,article_modify_date,article_context,account,article_type,article_image,article_brief) values (?,?,?,?,?,?,?,?)"
	_, err := db.Exec(sqlstr, article.Aticle_title, str, str, article.Article_context, account, article.Article_type, article.Article_image, article.Article_brief)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, "bad")
		return
	}

	response.SuccessRes(c, "添加成功")
}

func FixArticle(c *gin.Context) {}

/**
 *  获取文章
 */
type articleType struct {
	Title       string `db:"article_title" json:"title"`
	Create_date string `db:"article_create_date" json:"create_date"`
	Brief       string `db:"article_brief" json:"brief"`
	View        string `db:"article_view" json:"view"`
	Likes       string `db:"article_likes" json:"likes"`
	Comment     string `db:"article_comment" json:"comment"`
	Image       string `db:"article_image" json:"image"`
	Id          string `db:"article_id" json:"id"`
	Context     string `db:"article_context" json:"context"`
}

func GetArticle(c *gin.Context) {
	// 解析JSON 判断是否有文章id 有范文具体文章 无返回列表
	articleId := make(map[string]string)
	err2 := c.ShouldBindJSON(&articleId)
	if err2 != nil {
		response.BadRes(c, err2.Error())
	}
	db := mysql.Db

	// 返回具体文章
	if len(articleId) != 0 {
		var article articleType
		err := db.Get(&article, "SELECT article_title,article_create_date,article_brief,article_view,article_likes,article_comment,article_image, article_id,article_context FROM article WHERE article_id = ?", articleId["id"])
		if err != nil {
			response.BadRes(c, err.Error())
			return
		}

		viweSql := "update article set article_view=? where article_id=?"
		oldId, _ := strconv.Atoi(article.View)
		fmt.Println(oldId, articleId["id"])
		_, err3 := db.Exec(viweSql, oldId+1, articleId["id"])
		if err3 != nil {
			fmt.Println("更新失败")
			c.JSON(500, "bad")
			return
		}

		c.JSON(200, gin.H{
			"code": "00",
			"data": article,
		})

		return
	}

	// 返回所有文章
	var articles []articleType

	sqlStr := "SELECT article_title,article_create_date,article_brief,article_view,article_likes,article_comment,article_image, article_id FROM article"
	err := db.Select(&articles, sqlStr)
	if err != nil {
		response.BadRes(c, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"code": "00",
		"data": articles,
	})
}
