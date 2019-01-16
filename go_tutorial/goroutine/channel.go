package main

import "fmt"

func goroutine1(s []int, c chan int){
	sum := 0
	for _, v := range s {
		sum += v
	}
	//結果をchannelにわたす
	c <- sum
}

func goroutine2(s []int, c chan int){
	sum := 0
	for _, v := range s {
		sum -= v
	}
	//結果をchannelにわたす
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	//make関数でchannelを生成する
	c := make(chan int)
	//goroutineにsliceとchannelを渡す
	go goroutine1(s, c)
	go goroutine2(s, c)
	//channelの結果がくるまで待ち続ける　-> sync.WaitGroupは不要となる
	　　　//xにはgoroutineで最初に処理された方がはいる
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Println(y)
}