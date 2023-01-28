package main

import (
	"fmt"
	"math"
)

func main() {
	var i int
	fmt.Printf("Enter a number to check: ")
	_, _ = fmt.Scanf("%d", &i)
	var result = IsPrime(i)
	fmt.Printf("result=%t ", result)
}

func IsPrime(value int) (result bool) {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
