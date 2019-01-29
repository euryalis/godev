package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)
	//上の例は条件と対象を一緒にしているので使い回しができない
	//正規表現を単独で定義することで、使い回すことができるようになる
	r1 := regexp.MustCompile("a([a-z]+)e")
	ms := r1.MatchString("apple")
	fmt.Println(ms)

	//正規表現にひっかかる文字列を表示する
	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test")
	fmt.Println(fs)
	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
}