package main

import "fmt"

// Export 変数名の先頭が大文字なら外部から見えるよ
const Export = true

const export = false

func main() {
	const Z = 123
	Bar()
}

func Bar() {
	fmt.Println("Bar desuyo")
}
