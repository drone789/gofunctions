package main

import (
	"fmt"
	"strings"
)

func double(data string) string {
	return data + data
}

func up(data string) string {
	return strings.ToUpper(data)
}

var command = map[string]func(data string) string {
	"double": double, "up":up}

// 定义一个函数，包装对map 的动态调用
func call(com, data string) string {
	if function, ok := command[com]; ok {
		return function(data)
	}
	return ""
}

func main() {
	var com string = "up"
	var data string = "abcd"

	fmt.Println(call(com, data))

	com = "double"
	fmt.Println(call(com, data))
	
}

