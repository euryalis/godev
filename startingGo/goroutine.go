package trash

import "fmt"

func main() {
	//go は無名関数でも実行可能
	go func() {
		for {
			fmt.Println("sub routine")
		}
	}()

	for {
		fmt.Println("main routine")
	}
}