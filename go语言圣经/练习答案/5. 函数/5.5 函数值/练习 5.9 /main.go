package main

import (
	"fmt"
	"strings"
)

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "foo", f(s) , 1)
}

func main() {
	s := "afoob"
	s = expand(s, func(s string) string {
		return "aaa"
	})
	fmt.Println(s)
}
