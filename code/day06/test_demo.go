//求数组[1, 3, 5, 7, 8]所有元素的和
package main

import "fmt"

func main() {
	array1 := [...]int{1, 3, 5, 7, 8}
	result := 0
	for i := 0; i < len(array1); i++ {
		result += array1[i]
	}
	fmt.Printf("结果为：%d", result)
}
