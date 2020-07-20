package main

import (
	"fmt"
)

//Usber Usb接口
type Usber interface {
	start()
	stop()
}

//Computer 电脑
type Computer struct {
}

func (c Computer) work(usb Usber) {
	usb.start()
	usb.stop()
}

//Phone 手机
type Phone struct {
	name string
}

func (p Phone) start() {
	fmt.Println(p.name, "run")
}
func (p Phone) stop() {
	fmt.Println(p.name, "stop")
}

//Camera 照相机
type Camera struct {
}

func (c Camera) start() {
	fmt.Println("camera run")
}
func (c Camera) stop() {
	fmt.Println("camera stop")
}

func main() {
	var computer = Computer{}
	var p = Phone{
		name: "MI",
	}
	var c = Camera{}

	computer.work(p)
	computer.work(c)
}
