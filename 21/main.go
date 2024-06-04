// Реализовать паттерн «адаптер» на любом примере.

package main

import (
	"fmt"
)
//Паттерн адаптер позволяет использовать нам тип, нереализовывая  напрямую определенный интерфейс, а через типа-посредника(адаптера).
func main() {
	square := &Square{side: 10}
	adapter := &SquareAdapter{square: square}

	area := Check(adapter)
	fmt.Printf("Площадь квадрата: %f\n", area)
	
}



// Интерфейс фигуры
type Shape interface {
	CalculateArea() float64
}

type Square struct {
	side float64
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

// SquareAdapter адаптирует Квадрат к интерфейсу фигуры
type SquareAdapter struct {
	square *Square
}

//ТАк как наш Квадрат не реализует интерфейс Shape. Можем создать промежуточную структуру в которой инкапсулируем структуру square. 
//Она и будет вызываться в методе CalculateArea, метод который реализует интерфейс Shape(фигуры)
func (sa *SquareAdapter) CalculateArea() float64 {
	return sa.square.Area()
}

//Продемонстрируем реализацию
func Check(s Shape) float64 {
	return s.CalculateArea()
}