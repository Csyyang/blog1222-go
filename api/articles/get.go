package articles

import (
	"blog1222-go/mysql"
	"blog1222-go/response"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
	Account     string `db:"account" json:"account"`
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

	sqlStr := "SELECT article_title,article_create_date,article_brief,article_view,article_likes,article_comment,article_image, article_id, account FROM article"
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
