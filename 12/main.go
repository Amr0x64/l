package main

import "fmt"

func main() {
	//В go нет встроенного set, используем map
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})
	for _, v := range arr {
		set[v] = struct{}{}
	}
	fmt.Print("Подмножество: ")
	for v := range set {
		fmt.Print(v, " ")
	}
}
