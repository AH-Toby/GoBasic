package main

import "fmt"

func f5(x,y int ) (sum int){
	sum = x+y
	return
}

func ff(x,y int,z func(int,int) int) int{
	ret := z(x,y)
	return ret
}

func main() {
	a := ff(10,20,f5)
	fmt.Println(a)
}
