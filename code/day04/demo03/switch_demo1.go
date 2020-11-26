package main

import "fmt"

func main() {
	switch phone := "苹果12"; phone {
	case "苹果12", "苹果12mini", "苹果12pro", "苹果12promax":
		fmt.Println("苹果手机")
	case "华为mate40","华为mate40pro","华为mate40max","华为mate40promax":
		fmt.Println("华为手机")
	default:
		fmt.Println("其他家手机")
	}
}
