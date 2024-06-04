package main

import "fmt"

// записывает числа в канал из nums
func write(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// читает числа из канала, возводит в квадрат и передает в другой канал
func read(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	c := write(1, 2, 3, 4, 5, 6, 7)
	out := read(c)

	// выводим результат из выходного канала
	for i := range out {
		fmt.Println(i)
	}
}