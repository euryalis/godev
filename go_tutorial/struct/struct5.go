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

//構造体Vertexを組み込んで新たな構造体を作る
type Vertex3D struct{
	//構造体の値として、Vertexをそのまま指定できる
	Vertex
	z int
}

func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

//Goにはclassの継承は存在しない
//しかしembeddedを用いることで同様の仕組みを実現できる
//Vertex構造体をx軸y軸の構造体として、Vertex3D構造体をx軸y軸z軸の構造体に拡張してみる
func New(x, y, z int) *Vertex3D{
	return &Vertex3D{Vertex{x, y}, z}
}

func main() {
	v := New(3, 4, 5)
	v.Scale3D(10)
	fmt.Println(v.Area3D()) //60000
}