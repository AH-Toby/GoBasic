package main

import "fmt"

func main() {
	// make函数创建切片
	slice1 := make([]int, 2, 10)

	fmt.Printf("slice1=%v,len(slice1):%d,cap(slice1):%d", slice1, len(slice1), cap(slice1))
}
