package articles

import (
	"blog1222-go/mysql"
	"blog1222-go/response"

	"github.com/gin-gonic/gin"
)

type requestJSON struct {
	Account string `json:"account"`
	Id      string `json:"id"`
}

func Links(c *gin.Context) {
	var request requestJSON
	if err := c.ShouldBindJSON(&request); err != nil {
		response.BadRes(c, err.Error())
		return
	}

	db := mysql.Db
	// 判断是否点赞
	recordStr := "select likes_id from likes where article_id = ? and account = ?"
	var id int
	err := db.Get(&id, recordStr, request.Id, request.Account)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			// 点赞
			conn, err := db.Beginx()
			if err != nil {
				response.BadRes(c, err.Error())
				return
			}

			insertStr := "insert into likes (account,article_id) values (?,?)"
			updateStr := "update article set article_likes = article_likes + 1 where article_id = ?"
			_, errInsert := conn.Exec(insertStr, request.Account, request.Id)
			_, errUpdate := conn.Exec(updateStr, request.Id)
			if errInsert != nil || errUpdate != nil {
				conn.Rollback()
				response.BadRes(c, "提交失败")
				return
			}

			conn.Commit() // 提交事物
			response.SuccessRes(c, "add")
			return
		} else {
			response.BadRes(c, err.Error())
			return
		}
	}

	conn, err := db.Beginx()
	if err != nil {
		response.BadRes(c, err.Error())
		return
	}

	deleteStr := "delete from likes where likes_id = ?"
	updateStr := "update article set article_likes = article_likes - 1 where article_id = ?"
	_, errDele := conn.Exec(deleteStr, id)
	_, errUpdate := conn.Exec(updateStr, request.Id)
	if errDele != nil || errUpdate != nil {
		conn.Rollback()
		response.BadRes(c, "提交失败")
		return
	}

	conn.Commit() // 提交事物
	response.SuccessRes(c, "-1")
}
