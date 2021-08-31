// 7.1 接口是合约
// 练习 7.1： 使用来自ByteCounter的思路，实现一个针对单词和行数的计数器。你会发现bufio.ScanWords非常的有用
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type ByteCounter int
type WordCounter int
type LineCounter int

// 输入一个字节数组，将会返回字节数
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

// 输入一个直接数组，将会返回单词数
func (c *WordCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewReader(p))
	s.Split(bufio.ScanWords) // 按单词扫描
	for s.Scan() {
		*c += 1
	}
	if e := s.Err(); e != nil {
		fmt.Fprintln(os.Stderr, "reading input:", e)
	}
	return int(*c), nil
}

// 输入一个字节数组，计算行数
func (c *LineCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewReader(p))
	s.Split(bufio.ScanLines) // 按行扫描
	for s.Scan() {
		*c += 1
	}
	if e := s.Err(); e != nil {
		fmt.Fprintln(os.Stderr, "reading input:", e)
	}
	return int(*c), nil
}

func main() {
	var c ByteCounter
	fmt.Println(c.Write([]byte("hello")))

	var w WordCounter
	fmt.Println(w.Write([]byte("hello world, boy")))

	var l LineCounter
	fmt.Println(l.Write([]byte(`张三
李四
都爱Golang`)))
}
