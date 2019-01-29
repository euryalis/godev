package main

import "fmt"

func main() {
	//配列は以下のような形式で記載する
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a) //[100 200]

	//配列は最初に配列の数も定義する
	//したがって、あとから値を追加することはできない
	var b [2]int = [2]int{100, 200} //[100, 200]
	fmt.Println(b)

	//値を追加したい場合は、スライスを使う
	var c []int = []int{100, 200}
	c = append(c, 300)
	fmt.Println(c) //[100, 200, 300]
}