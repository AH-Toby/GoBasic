package main

import "fmt"

//for 循环基本格式

func main() {
	a := "我是一个好人"
	for i, chrA := range a {
		fmt.Printf("%d:%c", i, chrA)
		fmt.Println()
	}
}
