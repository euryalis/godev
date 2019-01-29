package main

import "fmt"

func main() {
	//mapはkeyとvalueを格納できる箱
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m) //map[apple:100 banana:200]
	fmt.Println(m["apple"]) //100
	//値を上書きすることもできる
	m["banana"] = 300
	fmt.Println(m) //map[apple:100 banana:300]
	//存在しないkeyを指定すると、mapの中に追加される
	m["new"] = 500
	fmt.Println(m)

	//存在しないkeyを表示すると0が返される
	fmt.Println(m["nothing"]) //0

	//map内のkeyを変数にいれると、値とboolが帰ってくる
	//なおboolの返り値は必須ではないので、変数を宣言しなくてもエラーにはならない
	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	//mapもスライスと同様にmakeで初期化することができる
	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	//varの場合は宣言はできるが、メモリ上にmapが割り当てられないのでエラーになる
	//sliceでappendする場合は問題ない
	/*
	var m3 map[string]int
	m3["pc"] = 5000
	fmt.Println(m3)
	*/
}