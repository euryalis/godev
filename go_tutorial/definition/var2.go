package main

import "fmt"

//可変長引数
//引数の数が場合によって変わる場合に使う
func foo(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params{
		fmt.Println(param)
	}
}

func main() {
	foo() //0 []
	foo(10, 20) //2 [10 20]
	foo(10, 20, 30) //3 [10 20 30]

	//sliceを引数として渡したい場合は、[変数]...という形で渡せる
	s := []int{1, 2, 3}
	fmt.Println(s)
	foo(s...)
}