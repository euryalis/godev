package trash

import "fmt"

func main() {
	LOOP:
		for {
			for {
				for {
					fmt.Println("開始")
					break LOOP
				}
				fmt.Println("ここは通らない")
			}
			fmt.Println("ここは通らない")
		}
	fmt.Println("完了")

	L:
		for i := 1; i <= 3; i++ {
			for j := 1; j <= 3; j++ {
				if j > 1 {
					continue L
				}
					fmt.Printf("%d * %d = %d\n", i, j, i*j)
			}
			fmt.Println("ここは処理されない")
		}
}
