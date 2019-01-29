package main

import "fmt"

type Person struct {
	Name string
	Age int
}

//Stringerはfmtパッケージで用いられるインターフェース
//StringerにはString()が含められていて、fmtパッケージの実行時に必ず実行される
//したがって、String()を自分で定義することで、Printの出力結果を変更することができる
func (p Person) String() string {
	return fmt.Sprintf("Mu name is %v.", p.Name)
}

func main() {
	mike := Person{"Mike", 22}
	fmt.Println(mike)
}