package main

import "fmt"

func sum6(x int, y int) (int, string) {
	ret := x + y
	retName := "综和是"
	return ret,retName
}
func main() {
	ret,name := sum6(1, 2)
	fmt.Println(ret)
	fmt.Println(name)
}
