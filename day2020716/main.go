package main

import (
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/tidwall/gjson"
)

func main() {
	quantity := decimal.NewFromInt(3)
	fmt.Println(quantity)

	var num1 float64 = 3.1
	var num2 float64 = 4.2
	fmt.Println(num1 + num2)

	const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	value := gjson.Get(json, "name.last")
	println(value.String())
}
