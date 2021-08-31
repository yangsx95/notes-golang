package main

import (
	"fmt"
	"sort"
)

// 每个课程的前置课程
// 找出一组课程，这组课程必须确保按顺序学习时，能全部被完成
// 拓扑排序:前置条件可以构成有向图。图中的顶点表示课程，边表示课程间的依赖关系
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

// 深度优先搜索整张图
func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool) // 记录课程是否已经被检查过
	var visitAll func(items []string) // 必须分为两步骤定义，如果一个步骤定义，函数字面量无法与visitAll绑定. compile error: undefined: visitAll
	visitAll = func(items []string) {
		for _, item := range items { // 遍历所有的key
			if !seen[item] { // 如果没有被检查，则进行检查
				seen[item] = true // 标记检查过
				visitAll(m[item]) // 检查该元素的依赖课程
				order = append(order, item) // 检查完毕，将该课程放到排序结果集中
			}
		}
	}
	// 遍历map，获取所有的key，放到变量keys中
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 检查所有图
	visitAll(keys)
	return order
}
