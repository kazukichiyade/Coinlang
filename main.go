package main

import (
	"coinlang/go/router"
)

// パッケージの読み込み グローバル定数 グローバル変数 init() main() の順に実行

func main() {
	// ルーティングを変数へ
	route := router.Router()
	route.Logger.Fatal(route.Start(":8080"))
}
