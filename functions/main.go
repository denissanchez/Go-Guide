package main

import (
	"fmt"
	"math"
)

func add(a, b int) int {
	return a + b
}

func diff(a, b int) int {
	return a - b
}

func doubles(numbers ...int) []int {
	for i, n := range numbers {
		numbers[i] = n * 2
	}
	return numbers
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(add(3, 4))
	fmt.Println(diff(3, 4))

	fmt.Println("*** Doubles ***")

	result := doubles(1, 2, 3, 4, 5)
	first := result[0]
	rest := result[1:]

	fmt.Println(first, rest)

	x := func(x int) int {
		return x * 2
	}(8)

	fmt.Println(x)
}
