package main

import "fmt"

// 编写类似sum的可变参数函数max和min。考虑不传参时，max和min该如何处理，再编写至少接收1个参数的版本。
func main() {
	fmt.Println(max(3,1,5,0,4))
	fmt.Println(min(3,1,5,0,4))
}

func max (i...int)  int {
	max := i[0]
	for _, n := range i {
		if n > max {
			max = n
		}
	}
	return max
}

func min(i ...int) int {
	min := i[0]
	for _, n := range i {
		if n < min {
			min = n
		}
	}
	return min
}