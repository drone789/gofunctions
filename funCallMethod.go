package main

import ()

func f() {
	// todo
}

type Mytype int
// 接收者类型为MyType
func (myType MyType) f() {
	// todo
}

func main() {

	// 1 普通函数调用
	f()

	// 2 闭包调用
	x := func() {
		// todo
	}
	x()

	// 3 匿名函数调用
	func(){
		// todo
	}()

	// 4 方法调用
	var myType MyType
	myType.f()
}
