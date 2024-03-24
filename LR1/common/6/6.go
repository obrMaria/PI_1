package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64

	fmt.Print("Enter num:")
	fmt.Scan(&a)

	result := math.Mod(a, 10)
	fmt.Println(result)
}
