package main

import (
	"fmt"
	"math"
)

func main() {
	// Задание значений a и b
	a := 5.0
	b := 3.0

	// Вычисление площади поверхности
	S := 2 * math.Pi * a * b * math.Sqrt(1+math.Pow((b/a), 2))

	// Вычисление объема
	V := (4 / 3) * math.Pi * a * b * b

	// Вывод результата
	fmt.Println("Площадь поверхности:", S)
	fmt.Println("Объем:", V)
}
