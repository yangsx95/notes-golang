package main

import "fmt"

func main() {
	var m1 map[string]int = make(map[string]int, 10) // map的存储空间为可以存储10个键值对，默认为1个
	m1["语文"] = 88
	m1["数学"] = 98
	fmt.Println(m1) // map[数学:98 语文:88]

	m2 := map[string]int{
		"语文": 88,
		"数学": 98,
	}
	fmt.Println(m2) // map[数学:98 语文:88]
}