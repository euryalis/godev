package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "jklsafdlkj"
}

func main(){
	//switch文は条件によって処理を切り替えられる
	//該当しないものはdefaultですべて処理する（オプショナル
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("mac!!")
	case "windows":
		fmt.Println("Windows!!")
	default:
		fmt.Println("Default!!", os)
	}

	//switchで変数を宣言せずにcase内に条件をかいてもOK
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("morning!")
	case t.Hour() < 17:
		fmt.Println("afternoon!")
	}
}