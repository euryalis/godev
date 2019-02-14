package trash

import "fmt"

const X = 1

func one() (int, int) {
	const Y = 2
	return X, Y
}

func main() {
	x, y := one()
	fmt.Println(x, y )
}