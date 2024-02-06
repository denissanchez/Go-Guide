package main

import "fmt"

func main() {
	var helloMsg string = "Hello"
	var worldMsg string = "World"

	fmt.Println(helloMsg, worldMsg)

	fmt.Printf("HelloMsg: %s WorldMsg: %s\n", helloMsg, worldMsg)

	var tinyInt int8 = 25
	fmt.Printf("uint8: %d\n", tinyInt)

	savedString := fmt.Sprintf("SAVED STRING: HelloMsg: %s WorldMsg: %s\n", helloMsg, worldMsg)
	fmt.Println(savedString)

	// Show the type of the variable

	fmt.Printf("Type of helloMsg: %T\n", helloMsg)
	fmt.Printf("Type of tinyInt: %T\n", tinyInt)
}
