package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
OuterLoop:
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("Boom!")
			//return であればすべての処理から抜けるのでOK
			//break　の場合は条件判定の１つからしか抜けないため、selectからしか抜けない
			break OuterLoop
		default:
			time.Sleep(50 * time.Millisecond)
			fmt.Println(".")
		}
	}
}