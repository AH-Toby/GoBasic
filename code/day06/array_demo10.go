package main

import "fmt"

func main() {
	array1 := [][]int{{0,1},{2,3},{4,5}}
	fmt.Println(array1)
	for _,v1 :=range array1{
		for _,v2 := range v1{
			fmt.Println(v2)
		}
	}
}
