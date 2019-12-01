package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// bitbankのベースurl
var bitbankBaseUrl = "https://public.bitbank.cc/"

// bitflyerのベースurlとエンドポイント
var bitflyerBaseUrl = "https://api.bitflyer.jp/v1/getticker?product_code="

// BitFlyerのBitcoinのAPIを取得する関数
func GetBitcoinApi() {
	// GetでWebAPIに対してアクセスする
	api, err := http.Get(bitflyerBaseUrl + "btc_jpy")

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

// BitFlyerのEthereumのAPIを取得する関数
func GetEthereumApi() {
	// GetでWebAPIに対してアクセスする
	api, err := http.Get(bitflyerBaseUrl + "eth_jpy")

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

// BitbankのXRPのAPIを取得する関数
func GetXrpApi() {
	// GetでWebAPIに対してアクセスする(bitbankの場合エンドポイントはこちらに記述)
	api, err := http.Get(bitbankBaseUrl + "xrp_jpy/ticker")

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

// BitbankのMonaのAPIを取得する関数
func GetMonaApi() {
	// GetでWebAPIに対してアクセスする(bitbankの場合エンドポイントはこちらに記述)
	api, err := http.Get(bitbankBaseUrl + "mona_jpy/ticker")

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

// BitbankのBccのAPIを取得する関数
func GetBccApi() {
	// GetでWebAPIに対してアクセスする(bitbankの場合エンドポイントはこちらに記述)
	api, err := http.Get(bitbankBaseUrl + "bcc_jpy/ticker")

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
