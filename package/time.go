package main

/*
see time package format.go
const (
   ANSIC       = "Mon Jan _2 15:04:05 2006"
   UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
   RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
   RFC822      = "02 Jan 06 15:04 MST"
   RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
   RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
   RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
   RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
   RFC3339     = "2006-01-02T15:04:05Z07:00"
   RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
   Kitchen     = "3:04PM"
   // Handy time stamps.
   Stamp      = "Jan _2 15:04:05"
   StampMilli = "Jan _2 15:04:05.000"
   StampMicro = "Jan _2 15:04:05.000000"
   StampNano  = "Jan _2 15:04:05.000000000"
)

この中でRFC3339はデータベース内の時間の型として使われることが多い
 */

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t) // 2018-12-29 18:48:10.168535 +0900 JST m=+0.000318434
	//時間をRFC3339型にフォーマットする
	fmt.Println(t.Format(time.RFC3339)) //2018-12-29T18:48:10+09:00
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()) //2018 December 29 18 48 10

}