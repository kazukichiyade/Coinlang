package main

import (
	"coinlang/go/repository"
	"coinlang/go/router"

	"github.com/jmoiron/sqlx"
)

// パッケージの読み込み グローバル定数 グローバル変数 init() main() の順に実行

var db *sqlx.DB

func main() {
	db = repository.ConnectDB()

	// ルーティングを変数へ
	route := router.Router()
	route.Logger.Fatal(route.Start(":8080"))
}
