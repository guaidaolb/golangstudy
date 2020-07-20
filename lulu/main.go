package main

import (
	"fmt"
	"lulu/calc"
	"lulu/tools"
)

func main() {
	var sum = calc.Add(100, 200)
	fmt.Println(sum)

	var suw = calc.Sub(200, 100)
	fmt.Println(suw)

	var suu = tools.Mul(5, 5)
	fmt.Println(suu)

	tools.PrintInfo()

	tools.SortIntAsc()
}
