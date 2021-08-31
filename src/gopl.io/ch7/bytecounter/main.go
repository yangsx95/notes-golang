package main

import "fmt"

type ByteCounter int

// 输入一个字节数组，将会返回字节数
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // 5

	// 将字节计算器置为0
	c = 0
	// 因为 *ByteCounter类型满足 io.Writer 的约定，所以可以将该类型传递给Fprintf函数中
	fmt.Fprintf(&c, "hello, %s", "Dolly")
	fmt.Println(c)
}
