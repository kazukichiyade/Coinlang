package handler

import (
	"time"

	"github.com/labstack/echo/v4"
)

// echo.Context(標準パッケージのresponseWriterやRequestと同じ)
func HandlerIndex(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Index",
		"Now":     time.Now(),
	}
	return render(c, "html/index.html", data)
}

func HandlerSignup(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Signup",
		"Now":     time.Now(),
	}
	return render(c, "html/signup.html", data)
}

func HandlerLogin(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Login",
		"Now":     time.Now(),
	}
	return render(c, "html/login.html", data)
}

func HandlerCoinList(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "CoinList",
		"Now":     time.Now(),
	}
	return render(c, "html/coinlist.html", data)
}

func HandlerMyPage(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "MyPage",
		"Now":     time.Now(),
	}
	return render(c, "html/mypage.html", data)
}
