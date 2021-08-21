package main

import "fmt"

func main() {

	var abc []string = make([]string, 0)
	abc = append(abc, "a", "b", "c", "d", "e", "f")
	fmt.Println(abc)

	reverse(abc)
	fmt.Println(abc)

	abcArray := [...]string{"a", "b", "c", "d", "e", "f"}
	arrayReverse(&abcArray) // 数组作为参数长度必须是固定的
	fmt.Println(abcArray)

	// 数组可以使用切片替代，完成反转操作
	abcSlice := abcArray[:]
	reverse(abcSlice)
	fmt.Println(abcSlice)
}

func reverse(targetSlice []string) {
	for i := 0; i < len(targetSlice)/2; i++ {
		targetSlice[i], targetSlice[len(targetSlice)-i-1] = targetSlice[len(targetSlice)-i-1], targetSlice[i]
	}
}

func arrayReverse(targetArray *[6]string) {
	for i := 0; i < len(targetArray)/2; i++ {
		targetArray[i], targetArray[len(targetArray)-i-1] = targetArray[len(targetArray)-i-1], targetArray[i]
	}
}
