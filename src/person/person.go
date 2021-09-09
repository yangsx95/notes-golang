// Package person
// 包 person下，person.go
package person

import "fmt"

// 小写开头，结构体person不导出
type person struct {
	Name string
	age uint8
	sal float32
}

// SetAge 大写字母，方法SetAge导出
func (p *person) SetAge(age uint8) {
	if age > 200 {
		fmt.Println("年龄范围在0-200之间")
	} else {
		p.age = age
	}
}

func (p *person) GetAge() uint8 {
	return p.age
}

func (p *person) SetSal(sal float32) {
	if sal <= 0 {
		fmt.Println("薪水不可小于零")
	} else {
		p.sal = sal
	}
}

func NewPerson(name string) *person{
	return &person{
		Name: name,
	}
}