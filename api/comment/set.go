package comment

import (
	"blog1222-go/mysql"
	"blog1222-go/response"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func SetComment(c *gin.Context) {
	var comment commontType
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		response.BadRes(c, err.Error())
		return
	}

	db := mysql.Db
	conn, err := db.Beginx()
	if err != nil {
		response.BadRes(c, err.Error())
		return
	}

	insertStr := "insert into comment (account,comment_context,article_id,comment_date,parent_comment_id,comment_level,replay_account) values (?,?,?,?,?,?,?)"
	updateStr := "update article set article_comment = article_comment + 1 where article_id = ?"

	timeObj := time.Now()
	str := timeObj.Format("2006-01-02 03:04:05")

	_, errInsert := conn.Exec(insertStr, comment.Account, comment.Comment_context, comment.Article_id, str, comment.Parent_comment_id, comment.Comment_level, comment.Replay_account)
	if errInsert != nil {
		fmt.Println(errInsert.Error())
		conn.Rollback()
		response.BadRes(c, "提交失败")
		return
	}
	_, errUpdate := conn.Exec(updateStr, comment.Article_id)
	if errUpdate != nil {
		fmt.Println(errUpdate.Error())
		conn.Rollback()
		response.BadRes(c, "提交失败")
		return
	}

	conn.Commit() // 提交事物
	response.SuccessRes(c, "ok")
}
