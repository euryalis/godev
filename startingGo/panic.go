package trash

import "fmt"

func main() {
	defer func() {
		if x := recover(); x != nil {
			//変数xはpanicに渡されたinterface{}
			fmt.Println(x)
		}
	}()
	panic("panic!!")
	//以下は実行されない
	fmt.Println("hello world")
}