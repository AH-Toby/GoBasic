package main

import "fmt"

func main() {
	var num1 int // 101
	num1 = 5
	// 位运算符
	//=
	fmt.Printf("num1=5:%d", num1) // 5
	fmt.Println()
	//+=
	num1 += 1
	fmt.Printf("num1+=1:%d", num1) // 6
	fmt.Println()
	//-=
	num1 -= 1
	fmt.Printf("num1-=1:%d", num1) // 5
	fmt.Println()
	// *=
	num1 *= 2
	fmt.Printf("num1*=2:%d", num1) // 10
	fmt.Println()
	// /=
	num1 /= 2
	fmt.Printf("num1/=2:%d", num1) // 5
	fmt.Println()
	// %=
	num1 %= 2
	fmt.Printf("num1o/o=2:%d", num1) // 1
	fmt.Println()
	// <<=
	num1 <<= 1
	fmt.Printf("num1<<=1:%b", num1) // 10
	fmt.Println()
	// >>=
	num1 >>= 1
	fmt.Printf("num1>>=1:%b", num1) // 1
	fmt.Println()
	//&=
	num1 &= 1
	fmt.Printf("num1&=1:%b", num1) // 1
	fmt.Println()
	//|=
	num1 |= 1
	fmt.Printf("num1|=1:%b", num1) // 1
	fmt.Println()
	//^=
	num1 ^= 1
	fmt.Printf("num1^=1:%b", num1) // 0
	fmt.Println()
}
