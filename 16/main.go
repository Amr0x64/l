// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
package main

import "fmt"

//Рекурсивный код, плюсы: реализуется просто, минусы: возможный stackoverflow, по причине большого размера массива. 
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0] //Первый элемент массива
	var less, more []int
	for _, num := range arr[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			more = append(more, num)
		}
	}
	result := append(quickSort(less), pivot)
	result = append(result, quickSort(more)...)
	return result
}

func main() {
	fmt.Print(quickSort([]int{10, 0, -9, 15, 120}))
}
