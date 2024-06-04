// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	//Когда диапазона значений нативных типов не хватает, поможет пакет big
	num1 := big.NewInt(600000000000)
	num2 := big.NewInt(600000000000)

	res := big.NewInt(0)
	fmt.Println("multiplex:", res.Mul(num1, num2))
	fmt.Println("sum:", res.Add(num1, num2))
	fmt.Println("subtract:", res.Sub(num1, num2))
	fmt.Println("division:", res.Div(num1, num2))
}
