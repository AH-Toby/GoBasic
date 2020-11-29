package main

import "fmt"

func main() {
	var slice1 []int
	fmt.Println(slice1)
	slice2 := []int{1,2,3,4,5}
	slice1 = append(slice1,slice2...)
	fmt.Println(slice1)
}
