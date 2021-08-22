// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters
	abcCount := 0                   // count of abc

	in := bufio.NewReader(os.Stdin) // 读取输入
	for {
		r, n, err := in.ReadRune() // 从控制台读取一个rune
		if err == io.EOF {         // 如果遇到EOF，则跳出循环，
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 { // 判断字符是否是无效字符，也就是没有对应的Unicode码点表示，比如�
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++

		if unicode.IsLetter(r) { // 该字符是否是字母
			abcCount++
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
	fmt.Printf("共包含字母 %v 个", abcCount)
}
