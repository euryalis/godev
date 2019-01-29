package main

import "fmt"

//panicはプログラムの実行ができなくなったときに強制終了する関数
//goではpanicを非推奨としている
//理由としては、panicにすると何が起きたかログをたどることができないため
func thirdPartyConnectDB() {
	panic("Unable to connect database")
}

//recover関数によってpanic状態から復帰することができる
//deferの実行をthirdParty~よりあとにすると、先にpanicが実行されるため、recoverが動作しなくなる点に注意
func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}

func main(){
	save()
	fmt.Println("OK?")
}