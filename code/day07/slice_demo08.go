package main

import "fmt"

func main() {
	// 切片的赋值
	slice1 := []int{1,2,3,4}

	slice2 := make([]int,4)  // 注意目的切片的长度容量必须和源切片一样
	copy(slice2,slice1)  //使用copy()函数将切片slice1中的元素复制到切片slice2

	fmt.Printf("slice1=%v,slice2=%v", slice1,slice2)
	fmt.Println()

	slice2[0] = 1000
	fmt.Printf("slice1=%v,slice2=%v", slice1,slice2)
}
