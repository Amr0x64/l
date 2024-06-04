package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{5, 4, 3, 2, 2}
	for _, i := range Intersetcion(nums1, nums2) {
		fmt.Print(i, ", ") 
	}
}

//В этой реализации пересечения множества, отбрасываются повторяющиеся элементы,
//как и в других языках программирования(прим. C#)
func Intersetcion(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	list := make([]int, 0)

	for _, i := range nums1 {
		m[i] += 1
	}

	for _, i := range nums2 {
		m[i] += 1
	}

	for k, v := range m {
		if v > 1 {
			list = append(list, k)
		}
	}

	return list
}
