package main

import "fmt"

func main() {
	var array1 [3][2]int
	array1 = [3][2]int{{0,1},{2,3},{4,5}}
	fmt.Println(array1)
	array2 := [][]int{{0,1},{2,3},{4,5}}
	fmt.Println(array2)
}
