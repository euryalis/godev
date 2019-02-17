package trash

import "fmt"

func receiver(ch chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func main() {
	ch := make(chan int)

	go receiver(ch)

	for i := 0; i < 10000; i++ {
		fmt.Println(len(ch))
		ch <- i
		fmt.Println(len(ch))
	}
}
