package main

import (
	"fmt"
)

func main() {
	var num1, num2 int

	fmt.Print("Введите первое число: ")
	fmt.Scan(&num1)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&num2)

	// Вычисление суммы, разности, произведения и частного
	sum := num1 + num2
	diff := num1 - num2
	prod := num1 * num2
	quotient := float64(num1) / float64(num2)

	// Вычисление среднего арифметического
	average := float64(sum) / 2.0

	// Вывод результатов
	fmt.Printf("%d+%d=%d\n", num1, num2, sum)
	fmt.Printf("%d-%d=%d\n", num1, num2, diff)
	fmt.Printf("%d*%d=%d\n", num1, num2, prod)
	fmt.Printf("%d/%d=%.16f\n", num1, num2, quotient)
	fmt.Printf("(%d+%d)/2=%.1f\n", num1, num2, average)
}
