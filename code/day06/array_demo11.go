package main

import "fmt"

func main() {
	array1 := [...]int{0,1,2,3,4}
	array2 := array1
	array2[0] = 6
	fmt.Println(array1,array2)
}
