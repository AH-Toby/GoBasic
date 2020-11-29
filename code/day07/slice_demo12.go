package main

import "fmt"

func main() {
	slice1 := []int{1,2,3,4,5}

	slice1 = append(slice1[:2],slice1[3:]...)
	fmt.Println(slice1)
}
