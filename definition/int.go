package main

import "fmt"

func main() {
	var (
		u8  uint8           = 255
		i8      int8                =127
		f32 float32     = 0.2
		c64 complex64 = -5 + 12i
	)
	fmt.Println(u8, i8, f32, c64) // 255 127 0.2 5 + 12i
	fmt.Printf("type=%T value=%v", u8, u8) // type=uint8 value=255

	x := 1 + 1
	fmt.Println(x)
	fmt.Println(1+1, 2+2)
	fmt.Println("1 + 1 =", 1+1) // 2
	fmt.Println("10 - 1=", 10-1) // 9
	fmt.Println("10 / 2 =", 10/2) // 5
	fmt.Println("10 / 3 =", 10/3) // 3
	fmt.Println("10.0 / 3 =", 10.0/3) // 3.3333333333333335
	fmt.Println("10 / 3.0 =", 10/3.0) // 3.3333333333333335
	fmt.Println("10 % 2 =", 10%2) // 0
	fmt.Println("10 % 3 =", 10%3) // 1

	y := 0
	fmt.Println(y) // 0
	//y = y + 1
	y++
	fmt.Println(y) // 1
	//y = y - 1
	y--
	fmt.Println(y) // 0

	fmt.Println(1 << 0) // 1
	fmt.Println(1 << 1) // 2
	fmt.Println(1 << 2) // 4
	fmt.Println(1 << 3) // 8
}