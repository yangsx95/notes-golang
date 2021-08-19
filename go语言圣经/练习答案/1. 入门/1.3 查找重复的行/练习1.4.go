package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)       // 用于记录 行字符串:重复数量 的map
	lines := make(map[string][]*os.File) // 用于记录，行对应的文件
	files := os.Args[1:]
	if len(files) == 0 { // 如果没有传入参数，则打开输入流让用户输入
		countLines(os.Stdin, counts, lines)
	} else {
		for _, arg := range files { // 遍历打开文件
			f, err := os.Open(arg)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, lines) // 计算重复行
			_ = f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 { // 大于1则说明重复
			fileName := make([]string, 0)
			for _, file := range lines[line] {
				fileName = append(fileName, file.Name())
			}
			fmt.Printf("%d\t%s\t%v\n", n, line, fileName)
		}
	}
}

func countLines(f *os.File, counts map[string]int, lines map[string][]*os.File) {
	input := bufio.NewScanner(f) // 按行扫描代码扫描器
	for input.Scan() {           // 读入下一行，并移除行末的换行符，如果到达没有后面一行，函数返回false
		counts[input.Text()]++                     // 设置更新 kv对，第一次为1 第二次为2 ...
		var files []*os.File = lines[input.Text()] // 将行字符串对应的文件指针存入files中
		if files == nil {
			files = make([]*os.File, 0)
		}
		files = append(files, f)
		lines[input.Text()] = files
	}
}
