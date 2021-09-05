package main

import (
	"flag"
	"fmt"
)

// Celsius 摄氏度
type Celsius float64

// Fahrenheit 华氏度
type Fahrenheit float64

// CToF 摄氏度转换华氏度
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }

// FToC 华氏度转换摄氏度
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }

// String 摄氏度字符串表示
func (c Celsius) String() string {
	return fmt.Sprintf("%g℃", c) // %g : 根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出
}

// String 华氏度字符串表示
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g℉", f)
}

// celsiusFlag 摄氏度Flag类型，用与给Flag消费
type celsiusFlag struct{ Celsius }

// Set ，字符串转换为 celsiusFlag， 从 flag.Value 接口实现
func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // 把s解析成 "浮点数字符串"的格式
	switch unit {
	case "C", "°C", "℃":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F", "℉":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}