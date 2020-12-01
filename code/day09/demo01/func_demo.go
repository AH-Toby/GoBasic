package main

import "fmt"

func sum(x int,y int)(ret int){
	ret = x+y
	return
}
func main() {
	ret := sum(1,3)
	fmt.Println(ret)
}
