package main

import "fmt"

func main() {
	//byte型はbyte単位のサイズで整数値を表す
	//ascii codeとして変換する
	b := []byte{72, 73}
	fmt.Println(b) //[72 73]

	//byte型をstring型に変換
	fmt.Println(string(b)) //HI

	//文字列をbyte型に変換することも可能
	c := []byte("H")
	fmt.Println(c) //72
	fmt.Println(string(c)) //H
}