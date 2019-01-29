package main

import (
	"awesomeProject/mylib"
	"awesomeProject/mylib/under"
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s)) //3
	mylib.Say() //human!
	under.Hello() //hello from under

	person := mylib.Person{"Mike", 20}
	fmt.Println(person) //{Mike 20}
	fmt.Println(mylib.Public) //Public
	// fmt.Println(mylib.private) //privateなので呼び出せない
}