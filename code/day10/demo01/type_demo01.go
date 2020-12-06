package main

import "fmt"

type smallInt = int  //将MyInt定义为int类型
func main() {
	var a myInt
	fmt.Printf("value:%v,type:%T \n",a,a)
}
