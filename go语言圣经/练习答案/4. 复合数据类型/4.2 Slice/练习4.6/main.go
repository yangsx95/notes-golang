package main

import (
	"fmt"
	"unicode"
)

func main() {

	str := "a  bc def   a"
	strbyte := []byte(str)
	fmt.Println(string(removeSpace(strbyte)))
}

func removeSpace(strbyte []byte) []byte {
	for i := 0; i < len(strbyte); i++ { // 没有变量安全问题，在for循环内可随意修改变量
		if unicode.IsSpace(rune(strbyte[i])) && strbyte[i] == strbyte[i+1] {
			strbyte = append(strbyte[:i], strbyte[i+1:]...)
			i--
		}
	}
	return strbyte
}
