package main

import "fmt"

func main() {
	// 十进制
	a1 := 101
	fmt.Printf("a1:%d\n", a1) // 十进制进
	fmt.Printf("a1:%b\n", a1) // 把十进制转成二进制
	fmt.Printf("a1:%o\n", a1) // 把十进制转成八进制
	fmt.Printf("a1:%x\n", a1) // 把十进制转成十六进制

	// 八进制 传统写法
	a2 := 077
	fmt.Printf("a2:%d\n", a2)  // 八进制转十进制
	// 新形式
	a3 := 0o77
	fmt.Printf("a3:%d\n", a3)

	// 十六进制
	a4 := 0xfff
	fmt.Printf("a4:%d\n", a4) // 十六进制转十进制
}
