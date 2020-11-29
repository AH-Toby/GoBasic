package main

import "fmt"

func main() {
	slice1 := []int{0,1,2,3,4,5}
	fmt.Println(slice1)
	slice1 = append(slice1,6)
	fmt.Println(slice1)
}
