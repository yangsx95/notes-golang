// 7.1 接口是合约
// 练习 7.2： 写一个带有如下函数签名的函数CountingWriter，传入一个io.Writer接口类型，
//		返回一个把原来的Writer封装在里面的新的Writer类型和一个表示新的写入字节数的int64类型指针。
// func CountingWriter(w io.Writer) (io.Writer, *int64)
package main

import (
	"fmt"
	"io"
	"os"
)

type ByteWriter struct {
	w io.Writer // 被包装的Writer
	C int64     // 新写入字节数的int64
}

// 新的类型ByteWriter的写入方法，实际上是记录了字节数之后，
// 再调用内部的w成员的Write方法，也就是包装了一下
func (bw *ByteWriter) Write(p []byte) (int, error) {
	bw.C = int64(len(p))
	n, err := bw.w.Write(p)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	bw := &ByteWriter{w: w}
	return bw, &bw.C
}

func main() {
	rw, l := CountingWriter(os.Stdout) // 写出到屏幕上
	_, _ = rw.Write([]byte("hello,world\n"))
	fmt.Println(*l)
}
