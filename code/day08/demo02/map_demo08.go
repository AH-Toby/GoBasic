package main

import "fmt"

func main() {
	sliceMap := make([]map[string]int, 3) // 声明一个长度为3容量为3的切片

	// 初始化内部map
	sliceMap[0] = make(map[string]int, 1)
	sliceMap[0]["张三"]=23
	sliceMap[1] = make(map[string]int, 1)
	sliceMap[1]["李四"]=30
	sliceMap[2] = make(map[string]int, 1)
	sliceMap[2]["王五"]=40
	fmt.Println(sliceMap)

}
