package repository

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

// sqlx定義
var db *sqlx.DB

func ConnectDB() *sqlx.DB {
	// 接続に必要な情報をos.Getenvで環境変数から取得
	dsn := os.Getenv("DSN")

	db, err := sqlx.Open("mysql", dsn)

	if err != nil {
		// エラーメッセージ出力
		log.Fatal(err)
	}

	// Pingでコールして疎通確認
	if err := db.Ping(); err != nil {
		// エラーに入るとDBをcloseする
		log.Fatal(err)
	}
	// defer(一番最後に実行)
	defer db.Close()

	// エラーが無ければメッセージとdbを返す
	log.Println("DB CONNECT SUCCESS!!!")
	return db
}
