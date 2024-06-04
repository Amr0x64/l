// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(UniqueStr("abcd")) 
	fmt.Println(UniqueStr("abCdefAaf"))
	fmt.Println(UniqueStr("aabcd"))
	fmt.Println(UniqueStr("абвгд"))
	fmt.Println(UniqueStr("абвгдд"))
}

//Реализуем множество через map. Struct{} поможет сэкономить память

func UniqueStr(str string) bool {
	str = strings.ToLower(str)
	m := make(map[rune]struct{})

	for _, i := range str {
		if _, ok := m[i]; ok {
			return false
		}
		m[i] = struct{}{}

	}
	return true
}
