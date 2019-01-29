package main

import (
	"fmt"
)

type UserNotFound struct {
	Username string
}

//ポイントレシーバを使う
//複数のエラーが発生したときに変数ベースだと、異なるエラーが同じエラー扱いになる場合があるため
func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not Found : %v\n", e.Username)
}

func myFunc() error {
	//something went wrong
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{"Mike"}
}

//error型が生成されると、go内部のerror interfaceによって Error()が実行される？
func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}

}