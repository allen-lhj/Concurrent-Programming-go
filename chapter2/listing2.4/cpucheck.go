package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())

	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
}
