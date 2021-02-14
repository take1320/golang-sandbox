package main

import "fmt"

// C1 型を持たない定数
const C1 = 12345

// C2 型を持つ(int型)定数
const C2 int = 34567

// C3 型を持つ(string型)定数
// const C3 string = "34567"

func main() {

	var x int = C1 * C2

	fmt.Println("12345 * 34567 =", x)

	var a int32 = 123
	var b int64 = 234

	fmt.Println("123 + 234 =", a+int32(b))
}
