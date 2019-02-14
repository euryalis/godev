package trash

import "fmt"

func AorB(a int) (b string) {
	a = 8
	b = "test"
	{
		a := 7
		const b = "god"
		fmt.Println(a, b)
	}
	fmt.Println(a, b)
	return b
}

func main() {
	AorB(7)
}