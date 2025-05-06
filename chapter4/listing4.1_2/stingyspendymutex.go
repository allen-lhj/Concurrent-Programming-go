package main

import (
	"fmt"
	"sync"
	"time"
)

func stingy(money *int, mutex *sync.Mutex) { //“接受指向共享互斥锁结构的指针”

	for i := 0; i < 1000000; i++ {
		mutex.Lock() //“在进入临界区之前锁定互斥锁”
		*money += 10
		mutex.Unlock() //“在离开临界区之后解锁互斥锁”
	}
	fmt.Println("Stingy Done")
}

func spendy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 1000000; i++ {
		mutex.Lock()
		*money -= 10
		mutex.Unlock()
	}
	fmt.Println("Spendy Done")
}

func main() {
	money := 100
	mutex := sync.Mutex{}
	go stingy(&money, &mutex)
	go spendy(&money, &mutex)
	time.Sleep(2 * time.Second)
	mutex.Lock()
	fmt.Println("Money in bank account: ", money)
	mutex.Unlock()
}

// “如果 Stingy 和 Spendy 的 goroutine 同时尝试锁定互斥锁，则互斥锁保证只有一个 goroutine 能够锁定它。
// 其他 goroutine 的执行将被挂起，直到互斥锁再次可用。例如，Stingy 将不得不等待 Spendy 减去金额并释放互斥锁。
// 当互斥锁再次可用时，挂起的 Stingy 的 goroutine 将被恢复，获取临界区的锁。”
