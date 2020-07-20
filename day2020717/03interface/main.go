package main

import (
	"fmt"
)

//Animaler 定义Animaler接口
type Animaler interface {
	SetName(string)
	GetName() string
}

//Dog 定义Dog结构体
type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}
func (d Dog) GetName() string {
	return d.Name
}

//Cat 定义Cat结构体
type Cat struct {
	Name string
}

func (c *Cat) SetName(name string) {
	c.Name = name
}
func (c Cat) GetName() string {
	return c.Name
}

func main() {

	var d = &Dog{
		Name: "大威天龙",
	}
	var d1 Animaler = d
	fmt.Println(d1.GetName())
	d1.SetName("大罗法咒")
	fmt.Println(d1.GetName())

	var c = &Cat{
		Name: "世尊地藏",
	}
	var c1 Animaler = c
	fmt.Println(c1.GetName())
	c1.SetName("啵业妈咪哄")
	fmt.Println(c1.GetName())

}
