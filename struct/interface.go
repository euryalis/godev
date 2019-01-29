package main

import "fmt"

//interface型に指定された変数は、必ずinterfaceに含まれるメソッドを持つ必要がある
//struct <- method <- interface的な？
type Human interface {
	Say() string
}

type Person struct {
	Name string
}

type Dog struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	return p.Name
}

//引数にHuman型を指定することで、Say()の実行を強制する
//また、Say()の実行にはPerson型が紐付いている
//そうすると、HumanInterfaceを型と指定したときPerson構造体によるSayの実行が強制される
func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get out")
	}
}

func main() {
	var mike Human = &Person{"Mike"}
	var x *Person = &Person{"X"}
	//var dog Dog = Dog{"Dog"} //DogにはHumanインターフェースがないため（Say()が構造体に紐付いていないため）、実行不可

	DriveCar(mike) //Run
	DriveCar(x) //Get out
	//DriveCar(dog)
}