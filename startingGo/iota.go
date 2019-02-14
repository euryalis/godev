package trash

import "fmt"

const (
	_ = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
)

const (
	A = iota
	B
	C
	D = 17
	E
	F = iota
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(A, B, C, D, E, F)
}


