// 7.1. 接口是合约
// 练习 7.3： 为在gopl.io/ch4/treesort (§4.4)中的*tree类型实现一个String方法去展示tree类型的值序列。

// Package treesort
// 二叉树排序：将切片元素依次放到到二叉树中，如果大于当前值，则放在左树，否则放在右树
// 最后通过前置遍历即可得到从大到小的值
package treesort

import (
	"fmt"
)

// tree 二叉树
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values { // 将所有的数值添加到二叉树中
		root = add(root, v)
	}
	appendValues(values[:0], root) // 进行前序遍历（左, 中, 右），并将遍历结果放到 values切片中
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// 将一个值添加到二叉树中，并返回这个新的二叉树
func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	// 小的值放在左边，大的值放在右边
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}


func (t *tree) String() string {
	return fmt.Sprintf(" %s %d %s ", t.left, t.value, t.right)
}
