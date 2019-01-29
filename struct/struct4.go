package main

import "fmt"

type Vertex struct{
	x, y int
}

func (v Vertex) Area() int {
	return v.x * v.y
}

func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

//func Area(v Vertex) int {
// return v.x * v.y
//}

//構造体をつくる関数をコンストラクタという
//構造体を新たに作りたい場合はNew関数をつくって他のパッケージでも構造体を作れるようにすることができる
//このNewはGOにおけるデザインパターンようなもので慣習として覚える
//コンストラクタで返すの構造体のポインタ
func New(x, y int) *Vertex{
	return &Vertex{x, y}
}

func main() {
	v := New(3, 4)
	//v := Vertex{x: 3, y: 4}
	//fmt.Println(Area(v)) //12
	v.Scale(10)
	fmt.Println(v.Area()) //1200
}