package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err) // 将字符串写入到 Writer中，并返回字节数以及错误
			os.Exit(1)                                        // 以错误状态退出命令行
		}

		//body, err := ioutil.ReadAll(resp.Body)
		_, err = io.Copy(os.Stdout, resp.Body) // 返回写入的字符数
		_ = resp.Body.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err) // 将字符串写入到 Writer中，并返回字节数以及错误
			os.Exit(1)
		}
		// 打印字节数组
		//fmt.Printf("%s", body)
	}
}
