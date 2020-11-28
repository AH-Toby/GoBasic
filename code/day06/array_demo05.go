package main

import "fmt"

func main() {
	var array1 [3]int
	array1 = [...]int{0: 3, 1: 2, 2: 4}
	fmt.Println("array1：", array1)

	var array2 = []int{0:1, 2:3}
	fmt.Println("array2：", array2)

	array3 := []int{0:1, 2:3}
	fmt.Println("array3：", array3)
}
