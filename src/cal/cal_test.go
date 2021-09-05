package cal

import (
	"testing"
)

func TestPlus(t *testing.T) {
	res := Plus(1, 2)
	e := 3 // 期望值

	if res != 3 { // 失败
		t.Fatalf("执行失败：期望值 %v， 目标值 %v ", e, res) // 失败并打印信息
	}
	t.Logf("执行成功：期望值 %v， 目标值 %v", e, res) // 打印日志
}

func TestDivision(t *testing.T) {
	res := Division(1, 2)
	t.Logf("执行成功：目标值 %v", res)
}

func TestDivisionByZero(t *testing.T) {
	res := Division(1, 0)
	t.Logf("执行成功：目标值 %v", res) // 打印日志
}
