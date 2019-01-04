package main

import "fmt"

func main() {
	t, f := true, false
	fmt.Printf("%T %v %t\n", t, t, t) // bool true
	fmt.Printf("%T %v %t\n", f, f, f) // bool false

	//論理演算
	fmt.Println(true && true) // true
	fmt.Println(true && false) // false
	fmt.Println(false && false) // false

	fmt.Println(true || true) // true
	fmt.Println(true || false) // true
	fmt.Println(false || false) // false

	fmt.Println(!true) //false
	fmt.Println(!false) //true
}