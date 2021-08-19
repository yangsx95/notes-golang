// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) // 创建channel通道，是用来在goroutine协程之间进行参数传递
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine   使用go关键字启动一个goroutine，并将通道传给协程
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch   将响应的数据放入到通道中
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // 不需要响应，只要数据大小，将返回流拷贝到垃圾桶ioutil.DisCard中
	resp.Body.Close()                                 // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds() // 获取执行时间
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
	// 当一个goroutine尝试在一个channel上做send或者receive操作时，这个goroutine会阻塞在调用处，
	// 直到另一个goroutine从这个channel里接收或者写入值，这样两个goroutine才会继续执行channel操作之后的逻辑
	// 这里main就是接收方，如果main一直不接受通道内的内容，此协程将会一直等待
}
