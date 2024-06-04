package main

import (
	"fmt"
	"sync"
)

func main() {
	wg()
	// g()
	// ch()
}

// Реализация через WaitGroup, вообщем позволяет задать счетчик, который информирует функцию Wait о блокировке вызывающей функции.
func wg() {
	nums := [5]int{2, 4, 6, 8, 10}
	wg := new(sync.WaitGroup)
	wg.Add(len(nums))
	for _, i := range nums {
		go func(n int) {
			defer wg.Done()
			fmt.Println(n * n)
		}(i)
	}
	wg.Wait()
}

//Реализация посредством каналов.
func ch() {
	nums := [5]int{2, 4, 6, 8, 10}
	ch := make(chan int)
	for _, i := range nums {
		go func(n int) {
			v := <- ch
			fmt.Println(v * v)
		}(i)
		ch <- i
	}
	fmt.Scanln()
}

// Минус данного способа, в ожидании пользовательского ввода,
// который должен совершится после всех окончаний горутин.
func g() {
	nums := [5]int{2, 4, 6, 8, 10}
	for _, i := range nums {
		go func(n int) {
			fmt.Println(n * n)
		}(i)
	}
	fmt.Scanln()
}
