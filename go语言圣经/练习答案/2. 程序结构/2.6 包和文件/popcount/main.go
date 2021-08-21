package main

import "fmt"

var pc [256]byte

func init() {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1) // &是按位与运算 |是按位或运算
	}
}

func main() {
	fmt.Printf("%v\n", pc)
	count := PopCount(3)
	fmt.Printf("%d", count)
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
