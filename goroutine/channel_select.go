package main

import (
	"fmt"
	"time"
)

//1secごとにstringを送信する
func goroutine1(ch chan string) {
	for {
		ch <- "sending packet 1"
		time.Sleep(1 * time.Second)
	}
}

//1secごとにstringを送信する
func goroutine2(ch chan string) {
	for {
		ch <- "sending packet 2"
		time.Sleep(1 * time.Second)
	}
}


func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go goroutine1(c1)
	go goroutine2(c2)

	//forでループを作り、msgを受信する仕組みをつくる
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}