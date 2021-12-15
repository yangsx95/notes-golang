package goconvey

import (
	. "github.com/smartystreets/goconvey/convey" // 导入convey包，使用Convey函数以及So等函数
	"testing"                                    // go testing 测试框架支持
)

func TestAdd(t *testing.T) {
	// 使用Convey方法声明一个测试的范围，主要包含 描述、以及一个用于执行测试逻辑的func()
	Convey("将两个数相加", t, func() {
		// 用于断言的通用方法，他可以接收如下三个参数
		// 1, 实际值
		// 2, 判断方式，这里是相同
		// 3, 期望值
		// 如果实际值不等于期望值，则代表测试未通过
		So(Add(1, 2), ShouldEqual, 3)
	})
}

func TestSubtract(t *testing.T) {
	Convey("将两个数相减", t, func() {
		So(Subtract(1, 2), ShouldEqual, -1)
	})
}

func TestMultiply(t *testing.T) {
	Convey("将两个数相乘", t, func() {
		So(Multiply(1, 2), ShouldEqual, 2)
	})
}

func TestDivision(t *testing.T) {
	Convey("将两个数相除", t, func() {
		// 不同的测试情况，请使用嵌套的Convey方法
		Convey("除数不为0", func() {
			r, err := Division(4, 2)
			So(r, ShouldEqual, 2)
			So(err, ShouldBeNil)
		})
		Convey("除数为0", func() {
			_, err := Division(4, 0)
			So(err, ShouldNotBeNil)
		})
	})
}
