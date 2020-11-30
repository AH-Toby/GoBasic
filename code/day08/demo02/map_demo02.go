package main

import "fmt"

func main() {
	// 声明变量
	mapA := make(map[string]int, 5)
	mapA["张三"] = 23
	mapA["李四"] = 30
	mapA["王五"] = 40
	fmt.Println(mapA["张三"])
	fmt.Println(mapA["李四"])
	fmt.Println(mapA["王五"])
}
