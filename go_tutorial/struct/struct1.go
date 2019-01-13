package main

import "fmt"

type Vertex struct {
	X, Y int
}

func (v Vertex) Plus() int {
	return v.X + v.Y
}

func main() {
	v := Vertex{X: 3, Y: 4}
	fmt.Println(v.Plus())
}