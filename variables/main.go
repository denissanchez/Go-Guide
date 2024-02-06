package main

import "fmt"

func main() {
	const pi float64 = 3.14159
	fmt.Println("PI: ", pi)

	const pi2 = 3.14159
	fmt.Println("PI 2: ", pi2)

	integer := 12
	fmt.Println("Integer: ", integer)

	var level int = 5
	fmt.Println("Level: ", level)

	var zeroInteger int
	fmt.Println("zeroInteger: ", zeroInteger)

	fmt.Println("Integer: ", integer, "| Level: ", level, "| zeroInteger: ", zeroInteger)

	myString := "Hello Go!"
	fmt.Println("String: ", myString)

	var zeroFloat float64
	var zeroString string
	var zeroBoolean bool

	fmt.Println("zeroFloat: ", zeroFloat, "| zeroString: ", zeroString, "| zeroBoolean: ", zeroBoolean)

	const squareSide float32 = 5
	squareArea := squareSide * squareSide

	fmt.Println("Square Area: ", squareArea)

	const highestScore uint64 = 99999 * 99999999999999

	fmt.Println("Highest Score: ", highestScore)
}
