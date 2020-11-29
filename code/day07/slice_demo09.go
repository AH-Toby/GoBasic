package main

import "fmt"

func main() {
	// 切片的赋值
	slice1 := []int{1,2,3,4}
	for i:=0;i<len(slice1);i++{
		fmt.Println(slice1[i])
	}
	fmt.Println()
	for _, i2 := range slice1 {
		fmt.Println(i2)
	}
}
