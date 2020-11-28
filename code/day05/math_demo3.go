package main

import "fmt"

func main() {
	var num1 int = 5  // 101
	var num2 int = 4 //100

	// 位运算符
	//&
	fmt.Printf("num1&num2=%b", num1&num2)  // 100
	fmt.Println()
	//|
	fmt.Printf("num1|num2=%b", num1|num2)  // 101
	fmt.Println()
	//^
	fmt.Printf("num1^num2=%b", num1^num2)  // 001
	fmt.Println()
	// <<
	fmt.Printf("num1<<1 %d", num1<<1)  // 10
	fmt.Println()
	// >>
	fmt.Printf("num2>>1 %d", num1>>1)  // 2
	fmt.Println()
}
