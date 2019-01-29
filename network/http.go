package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	//resp, _ := http.Get("https://example.com")
	//defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	base, _ := url.Parse("https://example.com")
	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Printf("%T %v\n", endpoint, endpoint)

	//リクエストを生成する
	req, _ := http.NewRequest("GET", endpoint, nil)
	//リクエストにheaderを追加する
	req.Header.Add("If-None-Match", `W/"wyzzyz"`)
	//クエリを生成する
	q := req.URL.Query()
	//クエリに値を追加する
	q.Add("c", "3&%")
	fmt.Println(q)
	//クエリを実際のurlにエンコードする
	fmt.Println(q.Encode())
	req.URL.RawQuery = q.Encode()

	//clientの構造体を作る
	var client *http.Client = &http.Client{}
	//clientの構造体にリクエストを実行させる
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	//bodyを読み込む
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}