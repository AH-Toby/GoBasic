package main

import "fmt"

func main() {
	var (
		a int
		b int
	)
	a = 5
	b = 2
	// 关系运算符
	//&&
	fmt.Println("a>b&&a>0", a>b&&a>0)
	//||
	fmt.Println("a>b||b>2", a>b||b>2)
	//!
	fmt.Println("!(a>b)", !(a>b))
}
