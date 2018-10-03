package main

import (
	"fmt"
	"log"
	"net/http"
)

type Page struct {
  Title string
  Body []byte
}

type

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>, p.Title, p.Body")
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
