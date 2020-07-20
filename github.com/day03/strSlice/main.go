package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "what is your name ha ha what is your name ha ha"
	var strSlice = strings.Split(str, " ")
	var strMap = make(map[string]int)
	for _, v := range strSlice {
		strMap[v]++

	}
	fmt.Println(strMap)
}
