package main

import "fmt"

//interface型はどんな型の値も入れることができる
//値をinterface型にいれた場合、もとの方に戻すにはタイプアサーションする必要がある
//同じような仕組みとしてキャストがあるが、interfaceではキャストが使えない点に注意
func do(i interface{}){
	/*
	//int型にtype assertionする
	ii := i.(int)
	ii *= 2
	fmt.Println(ii)

	//string型にtype assertionする
	ss := i.(string)
	fmt.Println(ss + "!")
	*/

	//(type)によるタイプアサーションは、int,string,bool等のすべての型を意味する
	//このタイプアサーションはswitch文の中でしか使えない
	//switch type文としてセットで覚える
	//各caseにtype別の処理を記載して処理を分けることができる
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T\n", v)
	}


}

func main() {
	//var i interface{} = 10
	//do(i)
	do(10)
	do("mike")
	do(true)


	//以下の処理はキャスト(type conversion)
	var i int = 10
	ii := float64(i)
	fmt.Println(i, ii)
}