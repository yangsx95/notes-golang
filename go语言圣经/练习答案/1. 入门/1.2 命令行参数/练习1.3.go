// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var startTime = time.Now()
	var ss []string = make([]string, 0)
	var sep string
	for index, arg := range os.Args {
		ss = append(ss, sep, strconv.FormatInt(int64(index), 10), arg)
		sep = " "
	}
	fmt.Println(strings.Join(ss, " "))
	log.Print(time.Since(startTime))
}
