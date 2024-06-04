// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseSen("snow dog sun"))
	fmt.Println(reverseSen("снег собака солнце"))
}

func reverseSen(str string) string {
	s := strings.Split(str, " ")

	var b strings.Builder //Преимущество builder в оптимизации составление строки, в отличие от string из за аллокации памяти при изменении строки
	b.WriteString(s[len(s)-1])

	for i := len(s) - 2; i >= 0; i-- {
		b.WriteString(" ")
		b.WriteString(s[i])
	}

	return b.String()
}
