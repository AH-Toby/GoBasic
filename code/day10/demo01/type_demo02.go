package main

import "fmt"

type bigInt int
type lowInt = int

func main() {
	var a bigInt
	var b lowInt
	fmt.Printf("value:%v,type:%T \n",a,a) //value:0,type:main.bigInt
	fmt.Printf("value:%v,type:%T \n",b,b) //value:0,type:int
}
