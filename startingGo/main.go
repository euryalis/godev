package trash

import "fmt"

func callFunction(f func()) {
	f()
}

func main() {
	callFunction(func() {
		fmt.Println("called function")
	})


}
