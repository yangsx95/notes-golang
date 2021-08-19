package main

import (
	"flag"
	"fmt"
	"strings"
)

// 参数1： 命令参数名称
// 参数2： 命令参数默认值
// 参数3： 命令描述，输入错误时，会自动显示命令描述
// 使用--help或者-h会自动展示各个参数的描述信息
var n = flag.Bool("n", false, "省略行尾换行符")
var sep = flag.String("s", " ", "间隔符")

func main() {
	flag.Parse() // 解析用户闯过来的标准输入
	fmt.Print(strings.Join(flag.Args(), *sep)) // 使用Join打印参数数组
	if !*n {
		fmt.Println()
	}

}
