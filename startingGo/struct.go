package trash

import "fmt"

type Feed struct {
	Name string
	Amount int
}

type Animal struct {
	Name string
	Feed
}

type T0 struct {
	Name1 string
}

type T1 struct {
	T0
	Name2 string
}

type T2 struct {
	T1
	Name3 string
}

type Base struct {
	Id int
	Owner string
}

type A struct {
	Base /*共通のフィールド*/
	Name string
	Area string
}

type B struct {
	Base /*共通のフィールド*/
	Title string
	Bodies []string
}





func main() {
	a := Animal{
		Name: "dog",
		Feed: Feed{
			Name: "dogFood",
			Amount: 10,
		},
	}

	fmt.Println(a.Feed.Amount)
	fmt.Println(a.Amount)
	fmt.Println(a.Name)

	t := T2{
		T1: T1{
			T0: T0{
				Name1: "X",
			},
			Name2: "Y",
		},
		Name3: "Z",
	}
	fmt.Println(t.Name3)
	fmt.Println(t.Name2)


	a1 := A{
		Base: Base{
			Id: 10,
			Owner: "Taro",
		},
		Name: "Taro",
		Area: "Tokyo",
	}

	b1 := B{
		Base: Base{
			Id: 81,
			Owner: "Hanako",
		},
		Title: "no title",
		Bodies: []string{"A", "B"},
	}

	fmt.Println(a1.Id)
	fmt.Println(b1.Id)
	fmt.Println(a1.Owner)
	fmt.Println(b1.Owner)
}
