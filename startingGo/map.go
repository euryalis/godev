package trash

import "fmt"

func main() {
	m := map[int]string{
		1: "No1",
		2: "No2",
		3: "No3",
	}

	s, ok := m[1]
	fmt.Println(s, ok)
	s, ok = m[7]
	fmt.Println(s, ok)
	s, ok = m[3]
	fmt.Println(s, ok)

	if _, ok := m[1]; ok {
		fmt.Println("map[1] exists")
	}
	if _, ok := m[7]; ok {
		fmt.Println("map[7 exists")
	}

	delete(m, 2)
	fmt.Println(m)
}
