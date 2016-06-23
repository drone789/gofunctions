package main

import (
	"fmt"
	"runtime"
)

func main() {
	where := func() {
		// Caller(1) 跳过栈深度，跳过闭包的函数栈，输出调用位置的调试信息
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s: %d\n", file, line)
	}

	where()
	where()
}