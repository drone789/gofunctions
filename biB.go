package main

import (
	"fmt"
)

func main() {
	var a func()
	func() {
		var count int = 0
		a = func() {
			fmt.Println(count)
			count++
		}
	}()

	for i := 0; i < 10; i++ {
		a()
	}
}