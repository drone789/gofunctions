package main

import (
	"fmt"
)

//定义可选参数
type Options struct {
	one string
	two int
	three bool
}

//使用struct来聚合可变参数
func f1(a int, options Options) {
	fmt.Println(options.one)
	fmt.Println(options.two)
	fmt.Println(options.three)
}

//使用interface{}接口来统一变量类型
func f2(a int, values ...interface{}) {
	for _, v := range values {
		switch v.(type) {
		case string:
			fmt.Println("string: ", v.(string))
		case int:
			fmt.Println("int: ", v.(int))
		case bool:
			fmt.Println("bool: ", v.(bool))
		}
	}
}

func main() {
	fmt.Println("f1:")
	f1(3, Options{one:"haha", two:10, three:true})
	fmt.Println("f2:")
	f2(3,"haha",10,true)
}