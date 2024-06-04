package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := []any{"", 1, true, make(chan struct{})}
	//Используем рефлексию для определния типа
	for _, t := range t {
		v := reflect.TypeOf(t)
		fmt.Print(v, " ")
	}
	//Еще можно мапать типы switch v.(type)
}
