package main

import "fmt"

func main() {

	slice := make([]string, 0)
	slice = append(slice, "a", "b", "b", "c", "c", "c", "d", "a", "a")

	fmt.Println(distinct(slice))

}

func distinct(sa []string) []string {
	for i := 0; ; {
		if i == len(sa)-1 { // 到达切片尾部
			break
		}
		if sa[i] == sa[i+1] {
			sa = append(sa[:i], sa[i+1:]...) // 重复删除
			//fmt.Println(sa)
		} else {
			i++
		}
	}
	return sa
}
