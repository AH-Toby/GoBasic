/*1.编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型。*/
package main

import "fmt"



func main() {
	int1 := 5
	float1 := 3.1415
	bool1 := true
	str1 := "这是一个字符串"

	fmt.Println("int1:",int1)
	fmt.Printf("int1:%T",int1)
	fmt.Println()
	fmt.Println("float1:",float1)
	fmt.Printf("float1:%T",float1)
	fmt.Println()
	fmt.Println("bool1:",bool1)
	fmt.Printf("bool1:%T",bool1)
	fmt.Println()
	fmt.Println("str1:",str1)
	fmt.Printf("str1:%T",str1)
}
