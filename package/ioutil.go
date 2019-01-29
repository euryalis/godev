package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main(){
	//ioutilパッケージのreadfileでファイルを読み込む
	content, err := ioutil.ReadFile("main.go")
	//errが出力された場合は表示
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(content))

	//ioutilパッケージのWritefileで読み込んだ内容をファイルに出力する
	//Writefileの場合は返り値がerrだけなので１行で書ける
	if err := ioutil.WriteFile("ioutil_temp.go", content, 0666); err != nil {
		log.Fatalln(err)
	}
}