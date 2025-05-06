package main

import (
	"fmt"
	"time"
)

func doWork(id int) {
	fmt.Printf("Work %d started at %s\n", id, time.Now().Format("15:04:05"))
	time.Sleep(1 * time.Second)
	fmt.Printf("Work %d finished at %s\n", id, time.Now().Format("15:04:05"))
}

func main() {
	for i := 0; i < 5; i++ {
		go doWork(i)
	}
	time.Sleep(2 * time.Second)
}

// 1、Starts a new goroutine that calls the doWork function 启动一个新的goroutine，并调用doWork函数
// 2、Waits for all of the work to finish using a longer sleep 使用一个更长的睡眠来等待所有的工作完成

// We can also refer to this manner of calling functions as an asynchronous call, meaning that
// we don't have to wait for the function to finish before continuing execution.

// 我们也可以将这种调用函数的方式称为异步调用，这意味着我们不必等待函数完成就可以继续执行。

// So the sleep instruction needs to be there because in Go, when the main execution runs out
// of the instructions to run, the process terminates.“Without this sleep, the process would terminate without giving the goroutines a chance to run.”

//由于goroutines 是异步执行的，所以不会阻塞主进程。这个 sleep 指令是必须的，因为在 Go 中，当主执行没有可运行的指令时，进程会终止.”
