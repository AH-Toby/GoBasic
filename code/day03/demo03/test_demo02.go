/*2.编写代码统计出字符串"hello go语言"中汉字的数量。*/
package main

import "fmt"

func main() {
	str1 := "hello go语言"
	//// 强制转换成rune切片
	strRune := []rune(str1)
	count := 0
	for _,c:= range strRune {
		if c > 255 {
			count += 1
		}
	}
	fmt.Printf("汉字的数量为：%d", count)
}
