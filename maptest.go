package main

import (
	"fmt"
)

func main() {
	// 声明方式1
	m1 := make(map[string]string)
	m1["a"] = "aa"
	m1["b"] = "bb"
	fmt.Println(m1)
	fmt.Println("===================")

	// 声明方式2
	m2 := map[string]string{
		"a":"a2",
		"b":"b2",
	}
	fmt.Println(m2)

	// 修改键值
	m1["a"] = "xx"
	// 添加键值
	m1["x"] = "xx"
	fmt.Println(m1)

	// 检查键值是否存在，如果存在，则打印
	if v, ok := m1["a"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key Not Found")
	}

	// 删除键值
	delete(m1, "a")

	// 打印所有的键值
	for k, v := range m2 {
		fmt.Printf("%s === %s\n",k, v)
	}

	fmt.Println(len(m1), len(m2))
}