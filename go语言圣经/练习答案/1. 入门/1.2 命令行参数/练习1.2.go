// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for index, arg := range os.Args {
		s += sep + strconv.FormatInt(int64(index), 10) + arg
		sep = " "
	}
	fmt.Println(s)
}
