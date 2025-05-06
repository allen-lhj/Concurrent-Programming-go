package main

import (
	"fmt"
	"time"
)

/*
Note: this program has a race condition for demonstration purposes
In later chapters we cover how to wait for threads to complete their work
*/
func stingy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money += 10
	}
	fmt.Println("Stingy Done")
}

func spendy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money -= 10
	}
	fmt.Println("Spendy Done")
}

func main() {
	money := 100
	go stingy(&money)
	go spendy(&money)
	time.Sleep(2 * time.Second)
	fmt.Println("Money in bank account: ", money)
}

// “我们遇到这个问题是因为操作 *money += 10 和 *money -= 10 不是原子的；
// 在编译后，它们转换成了多个指令。执行过程中可能在这些指令之间发生中断。
//来自另一个 goroutine 的不同指令可能会干扰并导致竞态条件。当这种越界发生时，我们会得到不可预测的结果。”
