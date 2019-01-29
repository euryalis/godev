package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx) // float64 1 1.000000

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy) // int 1 1

	//string -> int の場合はキャストを利用できない
	//なので、strconvパッケージを利用する
	var s string = "14"
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v\n", i, i) // int 14

	//バイト表示 ->　文字列はキャスト可能
	h := "Hello World"
	fmt.Println(h[0]) // 72
	fmt.Println(string(h[0])) // H
}