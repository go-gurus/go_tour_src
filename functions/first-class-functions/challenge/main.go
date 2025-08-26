package main

import "fmt"

// Challenge: Use a type alias here!

func process(sample int, fn func(int) int) int {
	return fn(sample)
}

func adder(term int) func(a int) int {
	return func(a int) int { return a + term }
}

// Challenge: Refactor the code so this function has an empty argument list
func increment(term int) int {
	return adder(1)(term)
}

func main() {
	sample := 23
	fmt.Println(process(sample, increment),
		process(sample, adder(100)))

}
