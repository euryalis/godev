package main

import "fmt"

type Stringer interface {
	String()
}

type T struct {
	X int
	Y string
}

func (t *T) String() string {
	return fmt.Sprintf("<<%d, %s>>", t.X, t.Y)
}

func main() {
	t := &T{X: 1, Y: "a"}
	fmt.Println(t)
}