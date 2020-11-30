package main

import "fmt"

func main() {
	sliceMap := make(map[string][]int, 3) // 声明一个容量为3，key是string,值为切片的map

	// 初始化内部map
	sliceMap["张三"] = []int{1,2,3}
	sliceMap["李四"] = []int{1,2,3}
	sliceMap["王五"] = []int{1,2,3}
	fmt.Println(sliceMap)
}
