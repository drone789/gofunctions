package main

import (
	"fmt"
)

// check 传入一个interface类型，返回bool值，另一个是interface类型的可变参数
func print(check func(interface{}) bool, data...interface{}) {
	for _, v := range data {
		if check(v) {
			fmt.Printf("%v\n", v)
		}
	}
}

func main() {
	f := func(v interface{}) bool {
		switch v.(type) {
		case string:
			return true
		default:
			return false
		}
	}
	print(f, 1, "data", 233.22)
}
