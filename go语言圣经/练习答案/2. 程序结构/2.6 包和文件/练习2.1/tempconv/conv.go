package tempconv

// CToF converts a Celsius temperature to Fahrenheit. 摄氏度转换华氏度
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius. 华氏度转换摄氏度
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
