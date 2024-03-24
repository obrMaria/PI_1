package main

import (
	"fmt"
	"math"
)

func main() {
	var sideA, sideB, sideC float64

	// Ввод длин сторон с клавиатуры
	fmt.Println("Введите длины трех сторон сферического треугольника:")
	fmt.Print("Длина стороны A: ")
	fmt.Scan(&sideA)
	fmt.Print("Длина стороны B: ")
	fmt.Scan(&sideB)
	fmt.Print("Длина стороны C: ")
	fmt.Scan(&sideC)

	// Вычисление углов сферического треугольника
	angleA := math.Acos((math.Cos(sideA) - math.Cos(sideB)*math.Cos(sideC)) / (math.Sin(sideB) * math.Sin(sideC)))
	angleB := math.Acos((math.Cos(sideB) - math.Cos(sideA)*math.Cos(sideC)) / (math.Sin(sideA) * math.Sin(sideC)))
	angleC := math.Acos((math.Cos(sideC) - math.Cos(sideA)*math.Cos(sideB)) / (math.Sin(sideA) * math.Sin(sideB)))

	// Преобразование радиан в градусы для вывода
	angleA = angleA * (180 / math.Pi)
	angleB = angleB * (180 / math.Pi)
	angleC = angleC * (180 / math.Pi)

	// Вывод результатов
	fmt.Printf("Угол A: %.2f градусов\n", angleA)
	fmt.Printf("Угол B: %.2f градусов\n", angleB)
	fmt.Printf("Угол C: %.2f градусов\n", angleC)
}
