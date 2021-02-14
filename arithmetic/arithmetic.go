package main

import "fmt"

func main() {
	fmt.Println("1 + 2 = ", 1+2)

	fmt.Println("\"abc\" + \"XYZ\" =", "abc"+"XYZ")

	fmt.Println("3 - 2 =", 3-2)

	fmt.Println("2 * 3 =", 2*3)

	fmt.Println("5 / 2 =", 5/2)

	fmt.Println("-5 / 2 =", -5/2)

	fmt.Println("5 % 2 =", 5%2)

	// ビット演算(and)
	// 3 = 011
	// 6 = 110
	// 2 = 010 (答えと同じになった！)
	// OUTPUT: 3 & 6 = 2
	fmt.Println("3 & 6 =", 3&6)

	// ビット演算(or)
	// 3 = 011
	// 6 = 110
	// 7 = 111
	// OUTPUT: 3 | 6 = 7
	fmt.Println("3 | 6 =", 3|6)

	// 左シフト演算
	// 2 = 010
	// 4 = 100
	// OUTPUT: 2 << 1 = 4
	fmt.Println("2 << 1 =", 2<<1)

	// 右シフト演算
	// 2 = 010
	// 1 = 001
	// OUTPUT: 2 >> 1 = 1
	fmt.Println("2 >> 1 =", 2>>1)
}
