// Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

func main() {
	pointA := New(36, 7)
	pointB := New(1, 9)
	fmt.Println(pointA.distance(*pointB))
}

type Point struct {
	x int //сделаем поля приватными для этого пакета, так как нет необходимости делать публичными
	y int
}

func New(x, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p Point) distance(point Point) float64 {
	return math.Sqrt(math.Pow(float64(p.x-point.x), 2) + math.Pow(float64(p.y-point.y), 2))
}
