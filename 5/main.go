package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	fmt.Scanln(&n)
	timer, _ := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)

	ch := make(chan int, 2)
	go func() {
		for {
			ch <- rand.Intn(5)
		}
	}()

	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()

	<-timer.Done()
}
