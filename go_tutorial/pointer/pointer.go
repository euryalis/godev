package main

import "fmt"

//メモリのアドレスを渡されるので、型は*intになる
func one(x *int){
	//1の値をメモリの実際の値に入れたいので、*xに入れる
	*x = 1
}

func main() {
	var n int = 100
	//nのメモリのアドレスを渡す。
	one(&n)
	fmt.Println(n)
	//アドレス->値->アドレス->値といった感じで試してみる
	fmt.Println(&*&*&n)

	/*
	//変数に&をつけると、変数を格納しているメモリのアドレスを指す
	fmt.Println(&n)
	//メモリのアドレスを変数にいれる場合は、typeの前に*をつける
	var p *int = &n
	fmt.Println(p)
	//メモリのアドレスをいれた変数から実際の値を指す場合は、変数の前に*をつける
	fmt.Println(*p)
	*/
}
