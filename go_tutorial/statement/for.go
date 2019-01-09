package main

import "fmt"

func main(){
	//for文はループのこと
	//変数の宣言; 条件; 処理 といった形で記載する
	for i := 0; i < 10; i++ {
		if i == 3 {
			//continueは、それ以下の記載を飛ばす
			fmt.Println("continue")
			continue
		}
		if i > 5 {
			//breakは処理を中断する
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	} //0 1 2 continue 4 5 break

	//以下のような形で変数の宣言と処理を分けることができる
	//値を0にしてしまうと無限ループになってしまうので注意
	sum := 1
	for ; sum < 10; { //このセミコロンは省略可能 ex. for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
	//この場合は外でも変数を使用することができる
	fmt.Println(sum)

	//無限ループ
	for {
		fmt.Println("hello")
	}
}