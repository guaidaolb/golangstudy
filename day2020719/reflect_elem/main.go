package main

import (
	"fmt"
	"reflect"
)

func reflectValueOf(x interface{}) {
	var v = reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(123)
	} else if v.Elem().Kind() == reflect.String {
		v.Elem().SetString("hello world!")
	}
}

func main() {

	var a int64 = 100
	var b string = "hello golang!"

	reflectValueOf(&a)
	fmt.Println(a)

	reflectValueOf(&b)
	fmt.Println(b)

}
