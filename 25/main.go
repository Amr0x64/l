package main

import (
	"fmt"
	"time"
)

func main() {
	sleep(3 * time.Second)
	fmt.Println("close")
}

// функция time.After после заданного промежутка времени возращает канал с текущим временем. Далее читаем из него и идем дальше по коду.
func sleep(d time.Duration) {
	<-time.After(d)
}

