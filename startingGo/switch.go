package trash

import "fmt"

func main() {
	s := "A"
	switch s {
	case "A":
		s += "B"
		fallthrough
	case "B":
		s += "C"
		fallthrough
	case "C":
		s += "D"
		fallthrough
	default:
		s += "E"
	}
	fmt.Println(s)

	n := 4
	switch true {
	case n > 0 && n < 3:
		fmt.Println("0 < n < 3")
	case n > 3 && n < 6:
		fmt.Println("3 < n < 6")
	}

	var x interface{} = 3
	i := x.(float64)
	fmt.Printf("%T %v", i, i)
}
