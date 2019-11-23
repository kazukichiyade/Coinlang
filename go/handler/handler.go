package handler

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func HandlerIndex(c echo.Context) error {
	return c.String(http.StatusOK, "HelloWorld")
}