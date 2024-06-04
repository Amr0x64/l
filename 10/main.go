package main

import "fmt"

func main() {
	temp := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := make(map[int][]float32)
	for _, t := range temp {
		group := int(t/10) * 10 //ключ группы 
		m[group] = append(m[group], t)
	}
	fmt.Println(m)
}
