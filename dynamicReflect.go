package main

import (
	"fmt"
	"reflect"
)

func print(data string) {
	fmt.Println("print:", data)
}

func add(a, b int) {
	fmt.Println("data:", a+b)
}

var command = map[string]interface{} {
	"print":print, "add":add}

func call(name string, params ...interface{}) {
	f := reflect.ValueOf(command[name])
	if len(params) != f.Type().NumIn(){
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	f.Call(in)
}

func main() {
	call("print", "abcdefg")
	call("add", 7, 9)
}
