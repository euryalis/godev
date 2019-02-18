package trash

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func Swap(p *Point) {
	x, y := p.Y, p.X
	p.X = x
	p.Y = y
}

func (p *Point) Distance(dp *Point) float64 {
	x,y := p.X - dp.X, p.Y - dp.Y
	return math.Sqrt(float64(x*x + y*y))
}

type Strings []string

func (s Strings) Join(d string) string {
	sum := ""
	for _, v := range s{
		if sum != "" {
			sum += d
		}
		sum += v
	}
	return sum
}

func main() {
	p := &Point{
		X: 1,
		Y: 10,
	}

	//メモリのアドレスを値に変換（デリファレンス）するには *
	//値をメモリのアドレスに変換するには &
	Swap(p)
	fmt.Println(*p)

	p1 := new(Point)

	fmt.Printf("%T %v\n", p1, p1)
	p1.X = 1
	p1.Y = 2
	fmt.Println(p1.X, p1.Y)

	fmt.Println(p.Distance(&Point{X: 1, Y:1}))

	s := Strings{"A", "B", "C"}
	result := s.Join(",")
	fmt.Println(result)
}