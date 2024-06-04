// Удалить i-ый элемент из слайса.
package main

import "fmt"

func main() {
	slice := []any{1, 2, 3, 4, 6, 7, 8, 9}
	var i int
	fmt.Scanln(&i)
	fmt.Println(deleteAp(slice, i))
	
	slice = []any{1, 2, 3, 4, 6, 7, 8, 9}
	fmt.Println(deleteCopy(slice, i))
}

func deleteAp(slice []any, i int) []any {
	if i < 0 || i >= len(slice) {
		fmt.Println("Индекс вне диапазона")
		return nil
	}

	slice = append(slice[:i], slice[i+1:]...)
	return slice
}

func deleteCopy(slice []any, i int) []any {
	if i < 0 || i >= len(slice) {
		fmt.Println("Индекс вне диапазона")
		return nil
	}

	res := make([]any, len(slice) - 1) 
	copy(res, slice[:i])
	copy(res[i:], slice[i+1:])	

	return res
}