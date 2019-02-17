package trash

import "fmt"

var S string

func init() {
	S += "A"
}

func init() {
	S += "B"
}

func init() {
	S += "C"
}

func main() {
	//init()はソースコードに出現した順に実行される
	fmt.Println(S) //"ABC"
}