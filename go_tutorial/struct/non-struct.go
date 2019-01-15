main

import "fmt"

//struct型ではなく、自身で型を決めることができる
//こういった方をnon-structという
type MyInt int

//struct指向でMyInt型に紐づけたDouble関数。
//MyInt型からInt型にキャストして、値を２倍にして返す
func (i MyInt) Double() int {
return int(i * 2)
}

func main() {
//変数myIntに値10をMyInt型にキャストして入れる
myInt := MyInt(10)
//MyInt型には関数Doubleが含まれるため、myInt.Doubleで呼び出すことができる
fmt.Printf("%T %v", myInt.Double(), myInt.Double()) //int 20
}