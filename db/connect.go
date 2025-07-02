package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Conn *sql.DB

func init() {
	var err error
	Conn, err = sql.Open("mysql", "user:password@tcp(localhost:3306)/yourdbname")
	if err != nil {
		log.Fatal("DB接続失敗:", err)
	}
}
