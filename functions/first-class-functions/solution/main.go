package main

import "fmt"

type IntFn = func(int) int

func process(sample int, fn IntFn) int {
	return fn(sample)
}

func adder(term int) IntFn {
	return func(a int) int { return a + term }
}

func increment() IntFn {
	return adder(1)
}

func main() {
	sample := 23
	fmt.Println(process(sample, increment()),
		process(sample, adder(100)))

}
