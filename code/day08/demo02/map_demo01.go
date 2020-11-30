package main

import "fmt"

func main() {
	// 声明变量
	var mapA map[string]int
	fmt.Println(mapA == nil)
	mapA = make(map[string]int, 10)  //妖股算好内存空间避免在程序运行过程中动态扩容
	fmt.Println(mapA == nil)
	mapA["张三"] = 1
	mapA["李四"] = 2
	fmt.Println(mapA)
}
