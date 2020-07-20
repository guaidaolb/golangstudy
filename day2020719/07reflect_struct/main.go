package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name" form:"username"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s Student) GetInfo() string {
	var str = fmt.Sprintf("姓名:%v 年龄：%v 得分：%v", s.Name, s.Age, s.Score)
	return str
}

func (s *Student) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score

}

func (s Student) PrintFn() {
	fmt.Println("这是一个打印方法...")

}

func PrintStructField(s interface{}) {
	//判断s是不是一个结构体
	v := reflect.TypeOf(s)
	val := reflect.ValueOf(s)
	if v.Kind() != reflect.Struct && v.Elem().Kind() != reflect.Struct {
		fmt.Println("这不是一个学生类型。。。")
		return
	}

	method0 := v.Method(0)
	fmt.Println(method0.Name)
	fmt.Println(method0.Type)
	fmt.Println("-----------------------------------------")

	method1, ok := v.MethodByName("PrintFn")
	if ok {
		fmt.Println(method1.Name)
		fmt.Println(method1.Type)
	}
	fmt.Println("-----------------------------------------")

	val.Method(1).Call(nil)
	fmt.Println("-----------------------------------------")

	val.MethodByName("PrintFn").Call(nil)
	var str = val.MethodByName("GetInfo").Call(nil)
	fmt.Println(str)
	fmt.Println("-----------------------------------------")

	var params []reflect.Value
	params = append(params, reflect.ValueOf("李四"))
	params = append(params, reflect.ValueOf(23))
	params = append(params, reflect.ValueOf(99))
	val.MethodByName("SetInfo").Call(params)
	var str1 = val.MethodByName("GetInfo").Call(nil)
	fmt.Println(str1)
	fmt.Println("-----------------------------------------")
	num := v.NumMethod()
	fmt.Println(num)
	// //通过变量里面的field方法获得结构体的字段
	// field0 := v.Field(0)
	// fmt.Println(field0)
	// fmt.Printf("%#v \n", field0)

	// fmt.Println("字段名称", field0.Name)
	// fmt.Println("字段类型", field0.Type)
	// fmt.Println("字段标签", field0.Tag.Get("json"))
	// fmt.Println("字段标签", field0.Tag.Get("form"))

	// fmt.Println("-----------------------------------------")
	// //通过类型变量里面的FieldByName可以获得结构体里面的字段
	// field1, ok := v.FieldByName("Age")
	// if ok {
	// 	fmt.Println("字段名称", field1.Name)
	// 	fmt.Println("字段类型", field1.Type)
	// 	fmt.Println("字段标签", field1.Tag.Get("json"))
	// }

	// fmt.Println("-----------------------------------------")
	// //通过类型变量里面的NumField获取到该结构体里面有几个字段
	// num := v.NumField()
	// fmt.Println(num)

	// fmt.Println(val.FieldByName("Name"))
	// fmt.Println(val.FieldByName("Age"))
	// fmt.Println(val.FieldByName("Score"))

	// for i := 0; i < num; i++ {
	// 	fmt.Printf("字段名称:%v 字段的值:%v 字段类型:%v 字段标签:%v \n", v.Field(i).Name, val.Field(i), v.Field(i).Type, v.Field(i).Tag.Get("json"))
	// }

}

func main() {

	var s1 Student
	s1.Name = "张三"
	s1.Age = 20
	s1.Score = 98

	// var s2 = Student{
	// 	Name:  "李四",
	// 	Age:   20,
	// 	Score: 89,
	// }

	PrintStructField(&s1)
	fmt.Println(s1)
	fmt.Printf("%#v", s1)
}
