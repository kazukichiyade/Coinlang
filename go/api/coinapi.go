package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// bitflyerエンドポイント
var baseUrl = "https://api.bitflyer.jp/v1/getticker?product_code="

// BitFlyerのAPIを取得する関数
func GetCoinApi() {
	// GetでWebAPIに対してアクセスする
	api, err := http.Get(baseUrl + "btc_jpy")

	if err != nil {
		log.Fatal(err)
	}

	// 最後にapiをCloseする
	defer api.Body.Close()

	// ReadAllは、エラーまたはEOFに達するまでrから読み込み、読み込んだデータを返す
	byteArray, err := ioutil.ReadAll(api.Body)

	if err != nil {
		log.Fatal(err)
	}
	// byte配列をstring型へキャスト
	fmt.Println(string(byteArray))
}
