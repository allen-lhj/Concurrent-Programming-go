package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

const AllLetters = "abcdefghijklmnopqrstuvwxyz"

func main() {
	mutex := sync.Mutex{}
	var frequency = make([]int, 26)
	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go CountLettersBetter(url, frequency, &mutex)
	}
	time.Sleep(60 * time.Second)
	mutex.Lock()
	for i, c := range AllLetters {
		fmt.Printf("%c-%d ", c, frequency[i])
	}
	mutex.Unlock()
}

// CountLetters
// Note: this program us locking the entire goroutine with mutex on purpose to demonstrate
// bad placement of the lock and unlock. We fix this in the next listing
func CountLetters(url string, frequency []int, mutex *sync.Mutex) {
	mutex.Lock()
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("Server returning error status code: " + resp.Status)
	}
	body, _ := io.ReadAll(resp.Body)
	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(AllLetters, c)
		if cIndex >= 0 {
			frequency[cIndex] += 1
		}
	}
	fmt.Println("Completed:", url, time.Now().Format("15:04:05"))
	mutex.Unlock()
}

// 该锁定方式的简化调度图，仅使用三个goroutine。“我们的 goroutine 大部分时间用于下载文档，
// 而处理文档的时间只有一秒钟的微小部分。从性能角度来看，没有必要阻塞整个执行过程。文档下载步骤与其他 goroutine
// 没有共享内容，因此在此过程中不会发生竞争条件。”

// “过度锁定代码将我们的字母频率并发程序转变为顺序程序”

// 改进后的代码
// “并发执行函数的慢速部分（下载）”
// 仅锁定函数中的快速处理部分
func CountLettersBetter(url string, frequency []int, mutex *sync.Mutex) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("Server returning error code: " + resp.Status)
	}
	body, _ := io.ReadAll(resp.Body)
	mutex.Lock()
	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(AllLetters, c)
		if cIndex >= 0 {
			frequency[cIndex] += 1
		}
	}
	mutex.Unlock()
	fmt.Println("Completed:", url, time.Now().Format("15:04:05"))
}

// “下载部分，即我们函数中的耗时部分，将并发执行。然后，快速的字母计数处理将顺序执行。
// 我们基本上通过只对运行速度非常快且与整体相比比例很小的代码部分使用锁，来最大化我们程序的可扩展性。
// 我们可以运行前面的代码示例，正如预期的那样，它运行得更快，并给出了一致的正确结果：”
