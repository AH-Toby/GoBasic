package main

import "fmt"

func main() {
	// 切片再切片
	array1 := [...]int{0,1,2,3,4,5,6}
	slice1 := array1[0:3]  // s := a[low:high]  [0 1 2 3]
	fmt.Printf("len(slice1):%d,cap(slice1):%d", len(slice1), cap(slice1))
	fmt.Println()

	// 切片再切片
	slice2 := slice1[2:]    // [3]
	fmt.Printf("len(slice2):%d,cap(slice2):%d", len(slice2), cap(slice2))
}