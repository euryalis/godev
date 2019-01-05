package main

import "fmt"

func main() {
	//make関数でスライスを作成できる
	//スライス内の値のlen(長さ)とcap(容量)を指定することができる
	n := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n) //len=3 cap=5 value=[0 0 0]
	n = append(n, 0, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n) //len=6 cap=10 value=[0 0 0 0 0 0]

	//値を１つにすると、長さと容量の２つがその値になる
	a := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a) //len=3 cap=3 value=[0 0 0]

	//値を0にするとスライスの中身はからになる
	//makeの場合はメモリの領域を確保するので、厳密にいうと動作は異なる
	b := make([]int, 0)
	var c []int
	fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b) //len=0 cap=0 value=[]
	fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c) //len=0 cap=0 value=[]

	//for文で繰り返す
	//この場合、スライスの中身がないのでcapが５におさまる
	d := make([]int, 0, 5)
	for i := 0; i < 5; i++{
		d = append(d, i)
		fmt.Printf("len=%d cap=%d value=%v\n", len(d), cap(d), d)
	}

	//こちらの場合はすでに値が存在するので、capが１０になる
	e := make([]int, 5)
	for i := 0; i < 5; i++{
		e = append(e, i)
		fmt.Printf("len=%d cap=%d value=%v\n", len(e), cap(e), e)
	}
}