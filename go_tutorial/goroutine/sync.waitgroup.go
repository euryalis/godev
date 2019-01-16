package main

import (
	"fmt"
	"sync"
	"time"
)

func normal(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroutine(s string, wg *sync.WaitGroup) {
	//1つのwg.Add()につき、1つのwg.Done()を実行する必要がある
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	/*
	//goによって、処理を並列実行することができる
	//goはmainの処理が完了する前に並列を開始しないといけないので、処理の前に置くこと
	//また、goroutineの処理が終わらなくとも、mainの処理が完了した時点で処理は終了する
	go goroutine("World")
	normal("Hello")
	*/

	//goroutineの処理を実行するために、並列処理を待つための仕組みがある
	var wg sync.WaitGroup
	wg.Add(1)
	normal("hello")
	go goroutine("world", &wg)
	wg.Wait()
}