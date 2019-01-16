package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan int, i int) {
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println(i * 2000)
		wg.Done()
	}
	fmt.Println("FFFFFFFFFFFFFFf")
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	//Producer
	//producerの生成回数分addする必要がある
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	//Consumer
	//waitgroupは実態を渡す必要があるので&
	go consumer(ch, &wg)
	wg.Wait()
	//closeしないと、consumerがchを待ち続けてしまい、仮にchのあとの処理が実行されなくなってしまう
	close(ch)
	time.Sleep(2 * time.Second)

}
