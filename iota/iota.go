package main

import "fmt"

func main() {
	const (
		ZERO  = iota
		ONE   = iota
		TWO   = iota
		THREE = iota
		FOUR  = iota
		FIVE  = iota
	)

	const (
		A = iota
		B
		C
	)

	const (
		BIT0 = 1 << iota
		BIT1
		BIT2
	)

	fmt.Println("ZERO = ", ZERO)
	fmt.Println("ONE = ", ONE)
	fmt.Println("TWO = ", TWO)
	fmt.Println("THREE = ", THREE)
	fmt.Println("FOUR = ", FOUR)
	fmt.Println("FIVE = ", FIVE)

	fmt.Println("A = ", A)
	fmt.Println("B = ", B)
	fmt.Println("C = ", C)

	fmt.Println("BIT0 = ", BIT0)
	fmt.Println("BIT1 = ", BIT1)
	fmt.Println("BIT2 = ", BIT2)

}
