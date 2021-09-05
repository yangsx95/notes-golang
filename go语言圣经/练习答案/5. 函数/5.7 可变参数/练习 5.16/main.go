package main

import "fmt"

func main() {
	fmt.Println(Join("a", "b", "c"))
}

func Join(s ...string) (str string) {
	for _, ss := range s {
		str += ss
	}
	return
}
