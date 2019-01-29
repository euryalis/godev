package main

import "fmt"

//iotaとはギリシア文字のΙ（イオタ）を指す。
//index generator が、色々あって ι になり、それをアルファベットで読むと iota となった


const (
	//iotaで連番をふれる
	c1 = iota
	c2
	c3
)

const (
	//iotaの最初の0は使わないので捨てる
	_ = iota
	//シフト演算。1を左に10ビットずらすと1024になる
	KB int = 1 << (10 * iota)
	//1を左に20ビットずらすと 1024 * 1024 = 1MB(1048576)
	MB
	GB
)

func main() {
	fmt.Println(c1, c2, c3)
	fmt.Println(KB, MB, GB)
}