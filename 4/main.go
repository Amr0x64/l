package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var numW int
	fmt.Scanln(&numW)

	ch := make(chan int)
	for i := 1; i < numW+1; i++ {
		go worker(ch, i)
	}

	s := make(chan os.Signal, 1)
	//подписываемся на событие определенного сигнала, в данном случае поможет нам с обработкой ctrl+C.
	signal.Notify(s, syscall.SIGINT)
	for {
		//Работа select построена так, что если ни один из case не срабатывает, срабатывает default случай.
		select {
		case <-s:
			fmt.Println("Остсановка программы пользователем")
			close(ch)
			return
		default:
			ch <- rand.Intn(100)
		}
	}
}

func worker(ch chan int, num int) {
	for i := range ch { //Остановка воркеров произойдет после закрытия канала, данный способ является безопасным с точки зрения параллельной работы программы
		fmt.Println(i, " : ", num)
	}
}
