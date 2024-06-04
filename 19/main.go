// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import "fmt"

func main() {
	fmt.Println(reverse("главрыба"))
	fmt.Println(reverse("日本語 test"))
}

func reverse(str string) string {
	s := []rune(str)

	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}

	return string(s)
}