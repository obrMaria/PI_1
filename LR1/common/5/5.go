package main

import "fmt"

func main() {
	var a int

	fmt.Print("Enter num:")
	fmt.Scan(&a)

	result := a * a
	fmt.Println(result)
}
