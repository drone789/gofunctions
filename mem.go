package main

import (
	"fmt"
)

type memFunction func(int) int

var fib memFunction

var count int  = 0

func memory(f memFunction) memFunction {
	cache := make(map[int]int)
	return func(n int) int {
		// 有变量n，则返回v
		if v, ok := cache[n]; ok {
			return v
		}
		// 没有则计算
		v := f(n)
		cache[n] = v
		return v
	}


}

func init() {
	fib = memory(func(n int) int {
		count++
		if n < 2 {
			return 1
		}
		return fib(n-1) + fib(n-2)
		})
}

func main() {
	fmt.Println(fib(43))
	fmt.Println(count)

}