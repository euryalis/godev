package trash

import "fmt"

func main() {
	x, y := 3, 5

	if n := x * y; n % 2 == 0 {
		fmt.Println("n(%d) is even\n", n)
	} else {
		fmt.Printf("n(%d) is odd\n", n)
	}

	//fmt.Println(n) if文内で定義されていて、この領域ではブロック外にあたるため、未定義によるコンパイルエラー
}