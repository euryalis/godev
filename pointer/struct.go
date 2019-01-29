package main

import "fmt"

//構造体の中の変数が小文字だとprivate扱いになって、外部からはアクセスできなくなる点に注意
type Vertex struct{
	X int
	Y int
	S string
}

func changeVertex(v Vertex){
	v.X = 1000
}
//アドレスを渡しているので、参照先の値が変わることがわかる
func changeVertex2(v *Vertex){
	//ここで渡しているvはアドレスである。なので、本来であれば値に戻す必要がある ex. (*v).X
	//しかし、構造体ではここの値を自動で解釈してくれるので、vのままでよい
	v.X = 1000
}

func main() {
	v := Vertex{X: 1, Y: 2, S: "test"}
	//changeVertexではvの値をコピーして渡しているだけなので、main内では値は変わらない
	changeVertex(v)
	fmt.Println(v) //{1 2 test}

	v2 := &Vertex{X: 1, Y: 2, S: "test"}
	//&(アンパサンド)でv2のアドレスを渡しているので、main内の数値も変わる
	changeVertex2(v2)
	fmt.Println(v2) //&{1000 2 test}
	fmt.Println(*v2) //{1000 2 test}

	/*
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	//個別にアクセスすることも可能
	fmt.Println(v.X, v.Y)
	//値を差し替えることも可能
	v.X = 100
	fmt.Println(v)

	//変数に値をいれないと、初期値が入れられる
	v2 := Vertex{X: 1}
	fmt.Println(v2)

	v3 := Vertex{X: 1, Y: 2, S: "test"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Printf("%T %v\n", v4, v4)

	//構造体は型として指定することもできる
	var v5 Vertex
	fmt.Printf("%T %v\n", v5, v5)

	//newを使うと構造体のポインタになる
	v6 := new(Vertex)
	fmt.Printf("%T %v\n", v6, v6)

	//&を頭につけることでもポインタにできる
	//&の方がアドレスであることが明快なので、現状は&で表現する人の方が多めかも
	v7 := &Vertex{}
	fmt.Printf("%T %v\n", v7, v7)
	*/
}