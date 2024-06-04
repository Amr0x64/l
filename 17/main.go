// Реализовать бинарный поиск встроенными методами языка.

package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	arrInt := []int{1, 2, 3, 4, 5, 6}
	arrFloat := []float64{-1.1, 2.1, 3.7, 4.4, 5.2, 6.1}

	fmt.Println(binarySearch(arrInt, 3))
	fmt.Println(binarySearch(arrFloat, 4.4))
	fmt.Println(binarySearch(arrInt, -10))
}

// ограничение T constraints.Ordered означает, что тип T должен реализовывать интерфейс Ordered, который предпологает, что тип T должен быть сравнимым
// например поддреживать операции ==, >, <
func binarySearch[T constraints.Ordered](list []T, target T) int {
	var s, e, mid int
	e = len(list) - 1

	for s <= e {
		mid = (s + e) / 2

		if list[mid] == target {
			return mid
		} else if target > list[mid] {
			s = mid + 1
		} else {
			e = mid - 1
		}

	}

	return -1
}
