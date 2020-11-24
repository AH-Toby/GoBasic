package main

import "fmt"

func main() {
	n := 100
	// 查看类型
	fmt.Printf("%T\n",n)
	// 打印万物
	fmt.Printf("%v\n",n)
	// 打印二进制
	fmt.Printf("%b\n",n)
	// 打印整型十进制
	fmt.Printf("%d\n",n)
	// 打印八进制
	fmt.Printf("%o\n",n)
	// 打印十六进制
	fmt.Printf("%x\n",n)

	m := 1.233443
	// 打印浮点型
	fmt.Printf("%f\n",m)

	s := "hello 你好！"
	//打印字符串
	fmt.Printf("%s\n",s)

	// 加描述符打印
	fmt.Printf("%#v\n",s)
}