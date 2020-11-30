//写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "how do you do"
	s1Split := strings.Split(str1, " ")
	// 字符串切片
	fmt.Println(s1Split)

	// map存储出现册数
	strMap := make(map[string]int, len(s1Split))

	for _, key := range s1Split {
		value,ok := strMap[key]
		if !ok{
			strMap[key] = 1
		}else {
			strMap[key] = value+1
		}
	}
	fmt.Println(strMap)
}
