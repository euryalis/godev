package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))

	/*
	バッファが2であるのに対し、値を3つ以上いれるとエラーになる
	ch <- 300
	fmt.Println(len(ch))
	*/

	//for rangeでバッファチャネルの中の値を１つずつとる
	close(ch)
	for c := range ch {
		fmt.Println(c)
	}
}