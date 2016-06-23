package main

import (
	"log"
)

//定义ByteSlice 类型
type  ByteSlice  []byte

//接收者为值类型，不会修改接收者的值
func (slice ByteSlice) Append1(data []byte) []byte {
	slice = append(slice, data...)
	log.Println("Append1: ", string(slice))
	return slice
}

//接收者为指针类型，函数执行完后，修改了接受者的值
func (slice *ByteSlice) Append2(data []byte){
	lenSlice := len(*slice)
	lenNewSlice := lenSlice + len(data)
	newSlice := make( []byte, lenNewSlice)
	copy(newSlice[0:lenSlice], *slice)
	copy(newSlice[lenSlice:lenNewSlice], data)
	*slice = newSlice
}
func main() {
	var a ByteSlice //声明一个ByteSlice类型的变量
	a = []byte("hello") //赋值hello
	b := a.Append1( []byte("world1"))
	log.Println("a: ", string(a))
	log.Println("b: ", string(b))
	a.Append2([]byte("world2"))
	log.Println("a: ", string(a))
}
