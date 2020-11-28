/*有一堆数字，如果除了一个数字以外，其他数字都出现了两次，那么如何找到出现一次的数字？*/
package main

import "fmt"

func main() {
	nums := [...]int{1, 1, 2, 2, 3, 4, 4}
	res := 0
	for _, v := range nums {
		res ^= v
	}
	fmt.Println(res)
}
