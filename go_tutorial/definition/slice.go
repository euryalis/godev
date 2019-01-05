package main

import "fmt"

func main() {
	//スライスの基本
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n) //[1 2 3 4 5 6]
	fmt.Println(n[2]) //3
	fmt.Println(n[2:4]) //[3 4]
	fmt.Println(n[:4]) //[1 2 3 4]
	fmt.Println(n[2:]) //[3 4 5 6]
	fmt.Println(n[2:3]) //[3]

	//スライス内の要素を指定して値を入れることも可能
	n[2] = 100
	fmt.Println(n) //[1 2 100 4 5 6]

	board := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},

	}
	fmt.Println(board) // [[1 2 3] [4 5 6] [7 8 9]]
}