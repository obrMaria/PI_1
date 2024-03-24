package main

import "fmt"

func main() {
	var a int

	fmt.Print("Введите неотрицательное целое число:")
	fmt.Scan(&a)

	if a < 0 || a > 10000 {
		fmt.Println("Число не соответствует условиям")
		return
	}

	result := (a / 10) % 10
	fmt.Println(result)
}
