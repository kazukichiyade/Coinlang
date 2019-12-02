package main

import (
	"coinlang/go/api"
	"coinlang/go/repository"
	"coinlang/go/router"

	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/jmoiron/sqlx"
)

// パッケージの読み込み グローバル定数 グローバル変数 init() main() の順に実行

var db *sqlx.DB

func main() {
	// DB接続(ConnectDBの戻り値をグローバル変数に格納)
	db = repository.ConnectDB()
	repository.SetDB(db)

	// BitflyerビットコインのAPIを取得
	api.GetBitcoinApi()
	// BitflyerイーサリアムのAPIを取得
	api.GetEthereumApi()
	// BitbankリップルのAPIを取得
	api.GetXrpApi()
	// BitbankモナコインのAPIを取得
	api.GetMonaApi()
	// BitbankビットコインキャッシュのAPIを取得
	api.GetBccApi()

	// ルーティングを変数へ
	route := router.Router()
	route.Logger.Fatal(route.Start(":8080"))
}
