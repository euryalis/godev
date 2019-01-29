package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"time"
)

var s *semaphore.Weighted = semaphore.NewWeighted(1)

func longprocess(ctx context.Context) {
	// longprocessの処理を１つずつ順番に実行する
	if err := s.Acquire(ctx, 1); err != nil {
		fmt.Println(err)
		return
	}
	defer s.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func main() {
	ctx := context.TODO()
	go longprocess(ctx)
	go longprocess(ctx)
	go longprocess(ctx)
	time.Sleep(5 * time.Second)


}