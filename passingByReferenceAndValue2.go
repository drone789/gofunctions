package main

import (
	"fmt"
)



// 值传递,参数为int类型的切片
func modify3(a []int) {
	a = []int{3,4}
}

// 引用传递
func modify4(a []int) {
	a[0] = 3
}

func main() {
	var arr []int = []int{1,2}
	fmt.Println(arr)

	modify3(arr)
	fmt.Println(arr)

	modify4(arr)
	fmt.Println(arr)

}