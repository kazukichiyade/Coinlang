package main

import (
	"coinlang/go/router"
)

// パッケージの読み込み グローバル定数 グローバル変数 init() main() の順に実行

func main() {
	router := router.Router()
	router.Logger.Fatal(router.Start(":8080"))
}
