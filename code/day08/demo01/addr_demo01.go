package main

import "fmt"

func main() {
	a := 10
	b := &a // 取变量a的指针
	fmt.Printf("type:%T\n",b)
	c := *b
	fmt.Printf("type:%T\n",c)
	fmt.Printf("ptr:%v", c)
}
