package comment

import (
	"blog1222-go/mysql"
	"blog1222-go/response"

	"github.com/gin-gonic/gin"
)

type articleType struct {
	Id    string `json:"id" db:"article_id"`
	Level string `json:"level" db:"comment_level"`
}

type commontType struct {
	Comment_id        int    `json:"comment_id" db:"comment_id"`
	Account           string `json:"account" db:"account"`
	Comment_context   string `json:"comment_context" db:"comment_context"`
	Article_id        string `json:"article_id" db:"article_id"`
	Comment_date      string `json:"comment_date" db:"comment_date"`
	Parent_comment_id int    `json:"parent_comment_id" db:"parent_comment_id"`
	Comment_level     string `json:"comment_level" db:"comment_level"`
	Praise_num        int    `json:"praise_num" db:"praise_num"`
	Replay_account    string `json:"replay_account" db:"replay_account"`
}

func GetComment(c *gin.Context) {
	var article articleType
	err := c.ShouldBindJSON(&article)
	if err != nil {
		response.BadRes(c, err.Error())
		return
	}

	db := mysql.Db
	var conmmont []commontType
	getComStr := "SELECT comment_id,account,comment_context,article_id,comment_date,parent_comment_id,comment_level,praise_num,replay_account from comment WHERE article_id = ? and comment_level = ?"
	err2 := db.Select(&conmmont, getComStr, article.Id, article.Level)
	if err2 != nil {
		response.BadRes(c, err2.Error())
		return
	}

	response.SuccessData(c, conmmont)
}
