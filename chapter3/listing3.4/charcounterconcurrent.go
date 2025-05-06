package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

/*
Note: this program has a race condition for demonstration purposes
Additionally we have a timer at the end which you might need to adjust
depending on how fast your internet connection is.
In later chapters we cover how to wait for threads to complete their work
*/
func countLetters(url string, frequency []int) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("Server returning error status code: " + resp.Status)
	}
	body, _ := io.ReadAll(resp.Body)
	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(allLetters, c)
		if cIndex >= 0 {
			frequency[cIndex] += 1
		}
	}
	fmt.Println("Completed:", url)
}

func main() {
	var frequency = make([]int, 26)
	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, frequency)
	}
	time.Sleep(10 * time.Second)
	for i, c := range allLetters {
		fmt.Printf("%c-%d ", c, frequency[i])
	}
}

// 与上一个例子相比，我们是用31个goroutine并发读取和写入相同频率的切片, 首先会注意到下载速度比顺序版本快得多
// 其次输出消息不再按照顺序排列，因为文件大小不同，所以他们完成的顺序不同。
// 同时结果上出了问题，大多数结果在并发版本中的计数都较低，例如字母e在顺序运行中为181360，在并发版本中的计数为179936

// 这是所谓的竞态条件的结果，--当我们有多个线程（或进程）共享一个资源，“并且它们相互覆盖时，会给出意外的结果”
