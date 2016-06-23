package main

import (
	"fmt"
)
// 定义一个结构体，结构体中有两个方法
type MyType interface {
	Equal(a int, b interface{}) bool
	Len() int
}

// 定义index函数，返回b在slice中的索引位置
func index(slice MyType, b interface{}) int {
	for i := 0; i < slice.Len(); i++ {
		if slice.Equal(i, b) {
			return i
		}
	}
	return -1
}

// 定义int类型的结构体IntType
type IntType []int

// 实现Equal方法
func (slice IntType) Equal(i int, b interface{}) bool {
	return slice[i] == b.(int)
}

// 实现Len方法
func (slice IntType) Len() int {
	return len(slice)
}

func main(){

}