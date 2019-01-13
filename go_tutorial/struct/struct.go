package main

import "fmt"

type Vertex struct{
	X, Y int
}

//メソッド。構造体指向
//構造体に対して関数が紐付けられている = これをメソッドと呼ぶ
//また、この関数では値を渡している。値レシーバ
func (v Vertex) Area() int {
	return v.X * v.Y
}

//ポインタレシーパ
//vの実態となる値を変更したい場合は、構造体のアドレスを指定する
//アドレスを指定しないと値をコピーするだけで、実態には影響しない
func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

//通常の関数
//構造体と関数に関連性はなく、入れらてた引数に応じて値を返すだけ
func Area(v Vertex) int {
	return v.X * v.Y
}

func main() {
	v := Vertex{X: 3, Y: 4}
	/*
	fmt.Println(Area(v)) //12
	//構造体と紐付けられたメソッドは構造体の変数からclassのように呼び出せる
	fmt.Println(v.Area()) //12
	*/
	v.Scale(10)
	fmt.Println(v.Area()) //1200
}