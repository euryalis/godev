package trash

import "fmt"

func integer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	int := integer()

	fmt.Println(int())
	fmt.Println(int())
	fmt.Println(int())

	otherInt := integer()
	fmt.Println(otherInt())

}
