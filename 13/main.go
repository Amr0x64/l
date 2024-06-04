package main

import "fmt"

func main() {
	swap1()
	swap2()
	swap3()
}

func swap1() {
	var a, b int = 1, 5
	fmt.Printf("before: a=%d, b=%d \n", a, b)
	a, b = b, a
	fmt.Printf("after: a=%d, b=%d \n\n", a, b)
}

func swap2() {
	var a, b int = 1, 5
	fmt.Printf("before: a=%d, b=%d \n", a, b)

	a += b
	b = a - b
	a -= b
	fmt.Printf("after: a=%d, b=%d \n\n", a, b)
}

func swap3() {
	var a, b int = 1, 5
	fmt.Printf("before: a=%d, b=%d \n", a, b)

	a ^= b
	b = a ^ b
	a ^= b
	fmt.Printf("after: a=%d, b=%d \n", a, b)
}
