package trash

import "fmt"

func main() {
	//i := 0
	//for {
	//	fmt.Println(i)
	//	i++
	//	if i == 100 {
	//		break
	//	}
	//
	//}

	//for i < 100 {
	//	fmt.Println(i)
	//	i++
	//}

	//for i := 0; i < 100; i++ {
	//	fmt.Println(i)
	//}

	fruits := [3]string{"apple", "banana", "grape"}
	for i, s := range fruits {
		fmt.Printf("key: %d, value: %s\n", i, s)
	}

	for s, r := range "あいうえお" {
		fmt.Printf("%d %d\n", s, r)
	}

}
