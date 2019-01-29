package mylib

import "fmt"

type Person struct {
	Name string
	Age int
}

var Public string = "Public"
var private string = "private"

func Say() {
	fmt.Println("human!")
}