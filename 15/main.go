package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

func createHugeString(size int) string {
	// используем билдер для эффективной конкатенации строк
	var b strings.Builder

	for i := 0; i < size; i++ {
		fmt.Fprint(&b, "永")
	}

	return b.String()
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	fmt.Println(v)
	// руна может занимать не один байт, это и есть проблема
	fmt.Println(utf8.RuneLen('永')) // 3

	// срезаем по количеству рун
	justString = v[:100]

	// преобразовываем строку в слайс рун
	r := []rune(v)
	// в даннам случае мы срезаем по количеству рун
	justString = string(r[:100])
}
func main() {
	someFunc()
}