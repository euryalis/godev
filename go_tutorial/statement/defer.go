package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("world foo")
	fmt.Println("hello foo")
}

func main(){
	//deferは遅延実行のこと
	//具体的には関数内の処理が完了したあとに実行する
	/*
	foo()
	defer fmt.Println("world")
	fmt.Println("hello")
	*/

	//deferは最初に実行したものからスタックされる
	//したがって最初に宣言されたdeferが一番最後に実行される

	/*
	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")
	*/

	//goではfileを開いた場合必ず閉じる必要がある
	//はじめにcloseを用意しておくことで閉じ忘れを防げる
	file, _ := os.Open("./lesson.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}