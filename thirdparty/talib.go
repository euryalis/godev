package main

import (
	"fmt"
	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

func main() {
	spy, _ := quote.NewQuoteFromYahoo(
		"spy", "2018-04-01", "2019-01-01", quote.Daily, true)
	fmt.Println(spy.CSV()) //yahooから取得した株価情報をcsv形式で出力する
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2) 　//取得した株価情報のrsiをしゅつりょくする　
	mva := talib.Ema(spy.Close, 14)
	fmt.Println(mva)
}