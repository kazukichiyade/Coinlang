package handler

import (
	// pongo2テンプレートエンジン(Djangoライクな文法を利用できる)
	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
	"net/http"
)

// 相対パスを定数として宣言
const tmplPath = "src/template/"

func htmlBlob(file string, data map[string]interface{}) ([]byte, error) {
	return pongo2.Must(pongo2.FromCache(tmplPath + file)).ExecuteBytes(data)
}

func render(c echo.Context, file string, data map[string]interface{}) error {
	// 定義した htmlBlob() 関数を呼び出し、生成された HTML をバイトデータとして受け取る
	b, err := htmlBlob(file, data)

	// エラーチェック
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	// ステータスコード 200 で HTML データをレスポンス
	return c.HTMLBlob(http.StatusOK, b)
}
