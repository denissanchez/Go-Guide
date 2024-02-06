package main

import "fmt"

func main() {
	var number int = 1

	number++
	fmt.Println("<var>++: ", number)

	number--
	fmt.Println("<var>--: ", number)
}
