package mysql

import (
	"blog1222-go/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	dns := config.Configs.GetDNS()

	database, err := sqlx.Connect("mysql", dns)
	// database, err := sqlx.Open("mysql", "root:CSY19961222@tcp(101.34.66.232:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	database.SetMaxOpenConns(20) // 最大连接数量
	database.SetMaxIdleConns(10) // 最大空闲数

	Db = database
}
