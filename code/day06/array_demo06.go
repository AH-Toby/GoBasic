package main

import "fmt"

func main() {
	array1 := []int{1,2,3,4,5,6,7,8,9}
	for i:=0;i<len(array1);i++{
		fmt.Println(array1[i])
	}
}
