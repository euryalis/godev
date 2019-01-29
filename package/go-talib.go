package main

import (
	"fmt"
	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
	//_ "github.com/markcheno/go-talib" パッケージの頭にアンダースコアを入れるとパッケージは無視される
	//a "github.com/markcheno/go-talib" パッケージの頭に値をいれるとパッケージ名を任意に設定できる
)

func main() {
	//spyとは米国の株価の指数のこと
	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
}