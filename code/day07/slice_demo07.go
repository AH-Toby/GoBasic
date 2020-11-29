package main

import "fmt"

func main() {
	// 切片的赋值
	slice1 := []int{1,2,3,4}
	slice2 := slice1
	fmt.Printf("slice1=%v,slice2=%v", slice1,slice2)
	fmt.Println()
	slice2[0] = 1000
	fmt.Printf("slice1=%v,slice2=%v", slice1,slice2)
}
