package main

import "fmt"

func main() {
	// 声明变量
	mapA := map[string]int{
		"张三": 23,
		"李四": 30,
		"王五": 40,
	}
	fmt.Println(mapA["张三"])
	fmt.Println(mapA["李四"])
	fmt.Println(mapA["王五"])
	v, ok := mapA["王六"]
	if !ok{
		fmt.Println("查无此人")
	}else {
		fmt.Println(v)
	}
}
