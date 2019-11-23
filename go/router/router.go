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
	e.GET("/", handler.HandlerIndex)

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}
