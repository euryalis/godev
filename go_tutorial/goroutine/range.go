package main

import "fmt"

//channelの中にはstringが入る
func goroutine(s []string, c chan string) {
	//処理が終わったらchannelをcloseする
	defer close(c)
	sum := ""
	for _, v := range s {
		sum += v
		c <- sum
	}
}

func main() {
	words := []string{"test1!", "test2!", "test3!", "test4!"}
	　　　//channelの中にはstringが入る
	c := make(chan string)
	go goroutine(words, c)
	for w := range c {
		fmt.Println(w)
	}

}