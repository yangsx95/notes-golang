package cal

import "testing"

func TestSub(t *testing.T) {
	res := Sub(1, 2)
	if res != -1 {
		t.Fatalf("执行失败：期望值 %v， 目标值 %v ", -1, res) // 失败并打印信息
	}
}
