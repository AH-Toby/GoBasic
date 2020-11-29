package main

import "fmt"

func main() {
	//切片初始化
	//1.
	var s1 []int
	s1 = []int{1,2,3,4,5}
	//2.
	var s2 = []string{"张三","李四","王五","李二"}
	//3.
	s3 := []rune{'a','b','c','d'}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
