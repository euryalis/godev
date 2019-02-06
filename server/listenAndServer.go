package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// http.Requestにはクライアントのリクエスト情報が格納されている
// 今回は *http.Request内のURLのPathを取得してtitleを取得した
// http.ResponseWriterにはレスポンスのための雛形がある
// したがって、 http.ResponseWriterにhtml情報を格納して返すとクライントにその情報が表示される
func viewHandler(w http.ResponseWriter, r *http.Request) {
	// /view/test
	// 37行目にデバッグポイントをつけてデバッグモードで処理の流れを追うと良い
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))


}
