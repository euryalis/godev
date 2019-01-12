package main

import "fmt"

func main() {
	//newはポインタを介するものをつくるときにつかう
	var p *int = new(int)
	fmt.Printf("%T\n", p)

	var st = new(struct{})
	fmt.Printf("%T\n", st)

	//makeはそれ以外の場合に使う
	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	/*
	//メモリのアドレスだけ先に確保したい場合はnew(type)でメモリを確保する
	var p *int = new(int)
	fmt.Println(p) // アドレスが表示される
	fmt.Println(*p) //初期値が表示される　0
	*p++
	fmt.Println(*p)// 1

	//以下はnewを使わない例
	//変数に対してメモリが確保されていないので、値はnilになる
	var p2 *int
	fmt.Println(p2) //nil
	//メモリが割り当てられていないポインタに値をいれようとすると…
	*p2++
	//メモリが割り当てられていないことからpanicを起こす
	fmt.Println(*p2)
	*/
}