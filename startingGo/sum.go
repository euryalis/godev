package trash

import "fmt"

func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func pow(s []int) {
	for i, v := range s {
		s[i] = v * v
	}
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5))

	s := []int{10, 11, 12}
	fmt.Println(sum(s...))

	pow(s)
	fmt.Println(s)

}
