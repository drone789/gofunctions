package main

import (
	"fmt"
	"os"
)

// 包级函数
var f func() string

func init() {
	if len(os.Args) > 1{
		if os.Args[1] == "a" {
			f = func() string {
				return "bingo"
			}
		} else {
			f = func() string {
				return os.Args[1]
			}
		}
	}	
}

func main() {
	fmt.Println(f())
}

