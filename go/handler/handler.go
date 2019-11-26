package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandlerIndex(c echo.Context) error {
	return c.String(http.StatusOK, "HelloWorld")
}

func HandlerSignup(c echo.Context) error {
	return c.String(http.StatusOK, "Signup")
}

func HandlerLogin(c echo.Context) error {
	return c.String(http.StatusOK, "Login")
}

func HandlerCoinList(c echo.Context) error {
	return c.String(http.StatusOK, "CoinList")
}

func HandlerMyPage(c echo.Context) error {
	return c.String(http.StatusOK, "MyPage")
}
