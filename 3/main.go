package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	wg()
	atom()

}

// Используем mutex для безопасной операции суммирования.
func wg() {
	nums := [5]int{2, 4, 6, 8, 10}
	var sum int
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	wg.Add(len(nums))

	for _, i := range nums {
		go func(n int) {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			sum += n * n
		}(i)
	}
	wg.Wait()
	fmt.Println(sum)
}

// Использование атомиков для более низкоуровнего доступа, что дает скорсоть выполнения.
// atomic дает нам атомарность в прибавлении счетчика
func atom() {
	nums := []int64{2, 4, 6, 8, 10}
	var sum int64
	wg := new(sync.WaitGroup)
	wg.Add(len(nums))
	for _, n := range nums {
		go func(n int64) {
			atomic.AddInt64(&sum, n*n)
			wg.Done()
		}(n)
	}
	wg.Wait()
	fmt.Println(sum)
}

// func ch() {
// 	var sum int
// 	nums := []int{2, 4, 6, 8, 10}
// 	ch := make(chan int, 5)
// 	for _, n := range nums {
// 		go func(n int) {
// 			ch <- n * n
// 		}(n)
// 	}
// 	for i := 0; i < 5; i++ {
// 		sum += <-ch
// 	}
// 	fmt.Println(sum)
// }
