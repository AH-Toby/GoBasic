package main

import "fmt"

func main() {
	// 声明变量
	mapA := map[string]int{
		"张三": 23,
		"李四": 30,
		"王五": 40,
	}
	fmt.Println(mapA)
	delete(mapA,"张三")
	fmt.Println(mapA)
}
