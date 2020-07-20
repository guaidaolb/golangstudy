package main

import (
	"fmt"
	"reflect"
)

func reflectValue(x interface{}) {

	v := reflect.ValueOf(x)

	kind := v.Kind()

	switch kind {
	case reflect.Int:
		fmt.Printf("Int的原始类型值%v \n", v.Int()+11)
	case reflect.Float32:
		fmt.Printf("Float32的原始类型值%v \n", v.Float()+22.2)
	case reflect.Float64:
		fmt.Printf("float64的原始类型值%v \n", v.Float()+22.2)
	case reflect.String:
		fmt.Printf("String的原始类型值%v \n", v.String())
	default:
		fmt.Println("输入错误...")

	}

}

func main() {

	var a = 100
	var b = 10.4
	var c = 11.11
	var d = "hello golang!"

	reflectValue(a)
	reflectValue(b)
	reflectValue(c)
	reflectValue(d)

}
