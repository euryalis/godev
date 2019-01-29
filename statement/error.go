package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	file, err := os.Open("./lesson.go")
	if err != nil {
		log.Fatalln("Error!!")
	}
	defer file.Close()
	data := make([]byte, 100)
	//initializeにおいて、errが複数回でてきてもエラーにはならない
	//initializeでは最低ひとつ宣言できていればエラーにはならないため
	//したがって、err自体は上書きされていることに注意すること
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("file read error")
	}
	fmt.Println(count, string(data))

	//すでに上記でerrを定義しているため、再度定義しようとするとエラーになる
	//なのでこの場合は変数に値をいれるだけでOK
	if err = os.Chdir("test"); err != nil {
		log.Fatalln("Error")
	}

}
