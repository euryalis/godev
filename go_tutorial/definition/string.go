package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World") // Hello World
	fmt.Println("Hello " + "World") // Hello World
	// string(""[int])で表示したい文字列を指定できる
	fmt.Println(string("Hello World"[0])) // H

	var s string = "Hello World"
	// stringsパッケージのReplaceで文字の置き換えが可能
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s) // Xello World

	//stringパッケージのContainsで対象の文字列が含まれているか確認できる
	fmt.Println(strings.Contains(s, "World")) //true

	//　"を表示したい場合
	fmt.Println("\"") // "
	fmt.Println(`"`) // "
}