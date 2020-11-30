package main

import "fmt"

func main() {
	// 声明变量
	mapA := map[string]int{
		"张三": 23,
		"李四": 30,
		"王五": 40,
	}
	// 循环键值对
	for key, value := range mapA {
		fmt.Println(key, value)
	}

	// 循环键
	for key := range mapA {
		fmt.Println(key)
	}

	// 循环值
	for _,value := range mapA {
		fmt.Println(value)
	}
}
