package main

import (
	"context"
	"fmt"
	"time"
)

func goroutine(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func main(){
	ch := make(chan string)
	//context型を作る
	ctx := context.Background()
	//context型をもとにtimeout用の型とcancel funcを作る
	//バックグラウンドで1秒間待機する設定
	ctx, cancel := context.WithTimeout(ctx, 1 * time.Second)
	go goroutine(ctx, ch)
	defer cancel()

	for {
		select {
		//chから値が帰ってきた場合
		case <-ch:
			fmt.Println("success")
			return
		//ctx型の値が帰ってきて、Doneが実行された場合
		case <-ctx.Done():
			fmt.Println(ctx.Err())　//context deafline exceeded
			return
		}
	}
}