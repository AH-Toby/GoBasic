package main

import "fmt"

func main() {
	a1 := 0b10101010
	fmt.Printf("a1:%b\n", a1) // 打印二进制

	a2 := 0o777
	fmt.Printf("a2:%o\n", a2) // 打印八进制

	a3 := 0xfff
	fmt.Printf("a3:%x\n", a3) // 打印十六进制

	a4 := 123_456
	fmt.Printf("a4:%d\n", a4)
}
