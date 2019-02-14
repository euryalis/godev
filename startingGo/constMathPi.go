package trash

import (
	"fmt"
	"math"
)

func main() {
	f32 := float32(math.Pi)
	f64 := float64(math.Pi)

	fmt.Printf("type:%T value:%v\n", math.Pi, math.Pi)
	fmt.Println(f32)
	fmt.Println(f64)
}