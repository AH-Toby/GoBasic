package main

import "fmt"

func main() {
	// 定义一个长度为3元素类型为int的数组array1
	var array1 [3]int
	var array2 [4]int
	//array1=array2  //不可以这样做，因为此时二者是不同的类型
	fmt.Printf("%T",array1)
	fmt.Printf("%T",array2)
}
