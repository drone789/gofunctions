package main

import (
	"fmt"
)

// 自定义类型
type MyType struct {
	a int
}

// 值传递
func modify1(myType MyType) {
	myType.a = 1
}

// 引用传递
func modify2(myType *MyType) {
	myType.a = 1
}

func main() {
	myType := MyType{0}
	fmt.Println(myType.a)
	modify1(myType)
	fmt.Println(myType.a)

	modify2(&myType)
	fmt.Println(myType.a)

}