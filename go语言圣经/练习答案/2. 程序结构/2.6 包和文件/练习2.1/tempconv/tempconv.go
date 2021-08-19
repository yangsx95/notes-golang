package tempconv

import "fmt"

type Celsius float64    // 定义类型摄氏度
type Fahrenheit float64 // 定义类型华氏度
type Kelvin float64     // 定义类型绝对温度

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gk", k) }
