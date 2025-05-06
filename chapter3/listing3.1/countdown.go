package main

import (
	"fmt"
	"time"
)

func countdown(seconds *int) {
	for *seconds > 0 {
		time.Sleep(1 * time.Second)
		*seconds -= 1
	}
}

// we get two goroutines to share memory.
// we craeted one goroutine that will share a variable in memory with the main() goroutine(executing the main() function)
// The variable will actlike every second, and another goroutine will read the variable more frequently.
// and output it on the console.

// 我们如何让两个 goroutine 共享内存？在第一个例子中，
// 我们将创建一个 goroutine，与执行 main() 函数的主 goroutine 共享内存中的一个变量。
// 该变量将充当倒计时计时器。一个 goroutine 每秒减少该变量的值，另一个 goroutine 则更频繁地读取该变量并将其输出到控制台
func main() {
	count := 5
	go countdown(&count)
	for count > 0 {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(count)
	}
}
