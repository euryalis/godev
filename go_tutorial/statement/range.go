package main

import "fmt"

func main(){

	//for文でスライスのキーと値をとる
	l := []string{"go", "java", "python"}
	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i]) // 0 go, 1 java, 2 python
	}

	//rangeを使えば簡単にとれる
	//iにindex, vに値が入る
	for i, v := range l {
		fmt.Println(i, v) //0 go, 1 java, 2 python
	}

	//mapでも同様のことができる
	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v) //apple 100, banana 200
	}

	//第二要素はオプショナルなので明示しなくてもエラーにはならない
	for k := range m {
		fmt.Println(k) //apple, banana
	}

	//第一要素は必須なので_で明示する
	for _, v := range m {
		fmt.Println(v) //100, 200
	}
}