package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() { //读写方式打开文件
	file, err := os.OpenFile("/etc/network/interfaces", os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("open file filed.", err)
		return
	} //defer关闭文件
	defer file.Close() //获取文件大小
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stat.Size()
	fmt.Println("file size:", size) //读取文件内容到io中
	reader := bufio.NewReader(file)
	pos := int64(0)
	ip := "40.40.40.220"
	gateway := "40.40.40.1"
	netmask := "255.255.255.0"
	for { //读取每一行内容
		line, err := reader.ReadString('\n')
		if err != nil { //读到末尾 if err == io.EOF { fmt.Println("File read ok!") break } else { fmt.Println("Read file error!", err) return }
		}
		fmt.Println(line) //根据关键词覆盖当前行
		if strings.Contains(line, "address") {
			bytes := []byte("address " + ip + "\n")
			file.WriteAt(bytes, pos)
		} else if strings.Contains(line, "gateway") {
			bytes := []byte("gateway " + gateway + "\n")
			file.WriteAt(bytes, pos)
		} else if strings.Contains(line, "netmask") {
			bytes := []byte("netmask " + netmask + "\n")
			file.WriteAt(bytes, pos)
		} //每一行读取完后记录位置
		pos += int64(len(line))
	}
}
