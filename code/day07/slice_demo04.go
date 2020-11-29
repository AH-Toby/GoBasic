package main

import "fmt"

func main() {
	// 简单切片表达式
	array1 := [...]int{0,1,2,3,4,5,6}
	slice1 := array1[0:3]  // s := a[low:high]
	fmt.Printf("len(slice1):%d,cap(slice1):%d", len(slice1), cap(slice1))
}