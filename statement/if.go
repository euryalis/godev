package main

import "fmt"

func by2(num int) string {
	if num % 2 == 0{
		return "ok"
	} else {
		return "no"
	}
}

func main(){
	//変数を単体で宣言した場合は、if文の外でも変数を使い回せる
	result := by2(10)
	if result == "ok"{
		fmt.Println("great")
	}
	fmt.Println(result)

	if result2 := by2(10); result2 == "ok" {
		fmt.Println("great 2")
	}
	//fmt.Println(result2) //undefined扱いになる


	num := 10
	if num % 2 == 0 {
		fmt.Println("by 2") //by 2
	} else if num % 3 == 0 {
		fmt.Println("by 3")
	} else {
		fmt.Println("else")
	}

	//and は &&, or は ||
	x, y := 10, 11
	if x == 10 && y == 10{
		fmt.Println("&&")
	}

	if x == 10 || y == 10{
		fmt.Println("||") // ||
	}
}