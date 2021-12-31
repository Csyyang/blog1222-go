package main

import (
	"blog1222-go/mysql"
	"blog1222-go/router"
	"fmt"
)

type Ids struct {
	Id int `db:"book_id"`
}

func main() {
	var ids []Ids
	db := mysql.Db
	defer db.Close()

	err := db.Select(&ids, "select book_id from book")

	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	fmt.Println("select succ:", ids)

	router.CreateRouter().Run()
}
