package trash

import "fmt"

func main() {
	s := "ABCDE"
	fmt.Println(s[0:3])

	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	s2 := []int{1, 2, 3, 4}
	s3 := append(s1, s2...)
	fmt.Println(s3)
	fmt.Println(cap(s3))

	s4 := []int{1, 2, 3, 4, 5}
	s5 := []int{10, 11}
	n := copy(s4, s5)
	fmt.Println(n, s4)

	s6 := []string{"apple", "banana", "grape"}

	//無限ループになる
	//for i := 0; i < len(s6); i++ {
	//	fmt.Printf("[%d] => %s\n", i, s6[i])
	//	s6 = append(s6, "melon")
	//}

	for i, v := range s6 {
		fmt.Printf("[%d] => %s\n", i, v)
		s6 = append(s6, "melon")
	}
	fmt.Println(s6)
}
