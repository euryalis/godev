package main

import "fmt"

func goroutine(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		//一回ごとにsumの結果をchannelに送る
		c <- sum
	}
	close(c)

}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, len(s))
	go goroutine(s, c)
	//channelの値をfor文で都度取得して表示する

	for v := range c {
		fmt.Println(v)
	}
}