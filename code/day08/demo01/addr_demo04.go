package main

import "fmt"

func main() {
	var  a *int  // 声明一个int类型指针变量a
	a = new(int)  //开辟一个int类型的内存空间
	*a =100
	fmt.Println(*a)
}
