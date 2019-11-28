package repository

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

// sqlx定義
var db *sqlx.DB

func ConnectDB() *sqlx.DB {
	dsn := os.Getenv("DSN")

	db, err := sqlx.Open("mysql", dsn)

	if err != nil {
		// エラーメッセージ出力
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		// エラーに入るとDBをcloseする
		defer db.Close()
		log.Fatal(err)
	}

	// エラーが無ければメッセージとdbを返す
	log.Println("DB CONNECT SUCCESS!!!")
	return db
}
