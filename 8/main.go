package main

import (
	"fmt"
)

func setBit(n int64, pos uint, val uint) int64 {
	mask := int64(1 << pos)
	n &^= mask // сбросывает бит в позиции pos
	if val == 1 {
		n |= mask // устанавливает бит в позиции pos
	}
	return n
}

func main() {
	var num int64 = 10
	var pos uint = 1
	var val uint = 0

	result := setBit(num, pos, val)

	fmt.Printf("Число после установки %d-го бита в %d: %d\n", pos, val, result)
}
