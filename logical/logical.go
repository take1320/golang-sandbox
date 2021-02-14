package main

import "fmt"

func main() {

	// AND条件
	fmt.Println("true && true =", true && true)
	fmt.Println("true && false =", true && false)
	fmt.Println("false && false =", false && false)

	// OR条件
	fmt.Println("true || true =", true || true)
	fmt.Println("true || false =", true || false)
	fmt.Println("false || false =", false || false)

	// 否定
	fmt.Println("!true =", !true)
	fmt.Println("!false =", !false)
	fmt.Println("!!true =", !!true)
	fmt.Println("!!false =", !!false)
}
