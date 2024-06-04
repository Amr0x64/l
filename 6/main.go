package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

func main() {
	goroutine_one()
	goroutine_two()
	goroutine_three()
	goroutine_four()
}

func goroutine_one() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan struct{})

	go func() {
		select {
		case <-ctx.Done(): //Внутри канал, который высвобождается при отмене контекста
			fmt.Println("Горутина останавливается")
			ch <- struct{}{} //Посылаем пустую структуру в канал, чтобы функция main продолжила выполнение
		}
	}()

	time.Sleep(2 * time.Second) //Предоставим горутине немного времени на жизнь
	cancel()
	<-ch
}

// Цикл в горутине прекратит свое выполнение после, того как закроется канал
// Также используется wait group, для ожидания в основной функции.
func goroutine_two() {
	ch := make(chan struct{})
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := range ch {
			fmt.Println(i)
		}
	}(wg)
	ch <- struct{}{} //для примера
	time.Sleep(2 * time.Second)
	close(ch)
	wg.Wait()
}

// Остановка горутины с помощью контекста с таймером
func goroutine_three() {
	timeout, _ := context.WithTimeout(context.Background(), 5*time.Second)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	func() {
		defer wg.Done()
		<-timeout.Done()
		fmt.Println("goroutine finished, context timeout")
	}()
	wg.Wait()
}

//WaitGroup
func goroutine_four() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		wg.Wait()
		ch <- 10
	}(&wg)
	
	wg.Done()
	fmt.Println(<-ch)
}