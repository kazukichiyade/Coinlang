package router

import (
	// GolangのWeb FWでAPIサーバーによく使われる(外部パッケージ)
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"coinlang/go/handler"
)

var e = Router()

func Router() *echo.Echo {
	e := echo.New()

	// ルーティング
	e.GET("/", handler.HandlerIndex)
	e.GET("/signup", handler.HandlerSignup)
	e.GET("/login", handler.HandlerLogin)
	e.GET("/coinlist", handler.HandlerCoinList)
	e.GET("/mypage", handler.HandlerMyPage)

	// ミドルウェア
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	// `src/css` ディレクトリ配下のファイルに `/css`, `/js`  のパスでアクセスできるようにする
	e.Static("/css", "src/css")
	e.Static("/js", "src/js")

	return e
}
