package main

import "fmt"

func main() {
	var (
		a int
		b int
	)
	a = 5
	b = 2
	//算术运算符
	//+
	fmt.Println("a+b=", a+b)
	//-
	fmt.Println("a-b=", a-b)
	//*
	fmt.Println("a*b=", a*b)
	// /
	fmt.Println("a/b=", a/b)
	// %
	fmt.Println("a%b=", a%b)

	a++
	fmt.Println("a++", a)
	b--
	fmt.Println("b--", b)
}
