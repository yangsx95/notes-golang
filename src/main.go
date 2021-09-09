package main

import "fmt"

type A interface {
	sayA()
}

type B struct {
	Name string
}

// B 实现 A 接口
func (b B) sayA() {
	fmt.Println("你好 a")
}

func (b B) sayB() {
	fmt.Println("你好 b", b.Name)
}

func main() {
	// 上溯造型
	var a A = B{"小熊"}
	a.sayA()

	// 下塑造型 (类型断言)
	b, ok := a.(B)
	if ok {
		b.sayB()
	} else {
		fmt.Println("类型断言失败")
	}
}

func TypeJudge(items ...interface{}) {
	for i, item := range items {
		switch item.(type) { // type 是一个关键字，代表使用case语句中的类型对item进行类型断言
		case bool:
			fmt.Printf("param %d is bool, value is %t", i, item)
		case float32, float64:
			fmt.Printf("param %d is float, value is %v", i, item)
		case int8, int16, int32, int64, int:
			fmt.Printf("param %d is int, value is %v", i, item)
		case nil:
			fmt.Printf("param %d is nil, value is %v", i, item)
		case string:
			fmt.Printf("param %d is string, value is %v", i, item)
		}

	}
}
