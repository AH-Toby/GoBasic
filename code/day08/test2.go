//观察下面代码，写出最终的打印结果。
package main

import (
	"fmt"
)

func main() {
	type Map map[string][]int

	m := make(Map)
	s := []int{1, 2} //s:[1 2]

	s = append(s, 3) // s:[1 2 3]
	fmt.Printf("%+v\n", s)  // [1 2 3]

	s = append(s, 4) // s:[1 2 3]
	fmt.Printf("%+v\n", s)  // [1 2 3 4]

	m["q1mi"] = s // {"q1mi":[1 2 3，4]}

	s = append(s[:1], s[2:]...)  // [1 3]

	fmt.Printf("%+v\n", s)   // [1 3]

	fmt.Printf("%+v\n", m["q1mi"])  // [1 3 3 4] 注意这个列表保存的长度在上一步赋值处固定
}
