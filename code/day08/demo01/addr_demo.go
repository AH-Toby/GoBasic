package main

import "fmt"

func main() {
	a := 10
	b := &a // 取变量a的指针
	fmt.Printf("a:%d, ptr%p\n", a, &a)

	fmt.Printf("b%p, type:%T\n", b, b)

	fmt.Printf("ptr:%p", &b)
}
