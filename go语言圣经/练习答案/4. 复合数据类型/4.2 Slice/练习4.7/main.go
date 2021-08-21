package main

import "fmt"

func main() {
	var s = "ab cde ff"
	var b = []byte(s)
	reverse(b)
	fmt.Printf("%s", b)
}

func reverse(targetSlice []byte) {
	for i := 0; i < len(targetSlice)/2; i++ {
		targetSlice[i], targetSlice[len(targetSlice)-i-1] = targetSlice[len(targetSlice)-i-1], targetSlice[i]
	}
}
