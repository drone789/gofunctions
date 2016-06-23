package main

import (
	"fmt"
)

func main(){
	a := []string{"a", "b", "c"}
	// 所有的协程共享一个变量
	for _, v := range a {
		// 闭包以协程的方式运行

		// 定义一个局部变量，供协程中的闭包使用
		// v := v

		go func(v string) {
			fmt.Println(v)
		}(v)
	}
	// 阻塞 main 协程
	var input string
	fmt.Scan(&input)
}
