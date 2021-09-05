package main

import "fmt"

func main() {
	intChan := make(chan int, 5)
	strChan := make(chan string, 5)

	for i := 0; i < 5; i++ {
		intChan <- i
		strChan <- fmt.Sprintf("hello%d", i)
	}

	// 不关闭chan
over:
	for true { // 循环使用select读取管道，且按顺序读，如果第一个case没有读取到，则从下一个case管道读取
		select {
		case v := <-intChan:
			fmt.Println("从intChan读取数据", v)
		case v := <-strChan:
			fmt.Println("从strChan读取数据", v)
		default:
			fmt.Println("都读取完了，退出for")
			break over
		}
	}

}
