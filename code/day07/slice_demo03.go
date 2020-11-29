package main

import "fmt"

func main() {
	// 长度和容量
	s1 := []int{1, 2, 3, 4, 5, 6}
	println(s1)
	fmt.Printf("len(s1):%d,cap(s1):%d", len(s1), cap(s1))

	// 由数组得到切片
	array1 := [...]int{0,1,2,3,4,5,6,7,8,9,10}
	s2 := array1[1:7]
	fmt.Println(s2)
	fmt.Printf("len(s2):%d,cap(s2):%d", len(s2), cap(s2))
}
