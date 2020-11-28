//找出数组中和为指定值的两个元素的下标，
//比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
package main

import "fmt"

func main() {
	array1 := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(array1); i++ {
		for j := i + 1; j < len(array1); j++ {
			if array1[i]+array1[j] == 8 {
				fmt.Printf("结果为：(%d,%d)", i, j)
			}
		}
	}
}
