package trash

import (
	"fmt"
	"runtime"
)

func main() {
	go fmt.Println("go 1")
	go fmt.Println("go 2")
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoRoutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("Version: %s\n", runtime.Version())
}