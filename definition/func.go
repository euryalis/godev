package main

import "fmt"

func add(x, y int) (int, int) {
	return x + y, x - y
}

//関数の返り値で変数を宣言することができる
//この場合、関数内で変数を宣言する必要はない
//返り値が多くなってくると何を返すのかわかりづらいので、変数でわかるようにする
func calc(price, item int) (result int) {
	result = price * item
	return result
}

//関数から返り値が推測しやすい場合はわざわざ変数を宣言しなくてもよい
func convert(price int) float64 {
	return float64(price)
}

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := calc(100 , 2)
	fmt.Println(r3)

	//関数内で変数にfuncを定義することもできる
	f := func(x int){
		fmt.Println("inner func", x)
	}
	f(1)

	//宣言した関数のあとに値をすぐにいれて実行することもできる
	func(x int){
		fmt.Println("inner func", x)
	}(2)
}