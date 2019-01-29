package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{5, 3, 2, 8, 7}
	s := []string{"d", "a", "f"}
	p := []struct{
		Name string
		Age int
	}{
		{"Nancy", 20},
		{"Vera", 40},
		{"Mike", 30},
		{"Bob", 50},
	}
	fmt.Println(i, s, p)
	//数値の少ない順にソート
	sort.Ints(i)
	//アルファベット順にソート
	sort.Strings(s)
	//Nameの大小を比較する関数を作ってソート
	sort.Slice(p, func(i, j int) bool {
		return p[i].Name < p[j].Name
	})
	fmt.Println(i, s, p)
}