package main

import "fmt"

func incrementGenerator() (func() int) {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius *radius
	}
}

//関数を値として扱うクロージャー
//関数の返り値として関数を返す
//debug機能のconsole->step intoで動きをおうとわかりやすい
func main() {
	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	//円周率が3の場合と3.14の場合とで使い分けることができる
	c1 := circleArea(3.14)
	fmt.Println(c1(2))

	c2 := circleArea(3)
	fmt.Println(c2(2))
}