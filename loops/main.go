package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ {
		println(i)
	}

	fmt.Println("**********")

	counter := 0

	for counter <= 10 {
		println(counter)
		counter++
	}

	fmt.Println("**********")

	counter = 0

	for {
		fmt.Println(counter)
		counter++
		if counter > 10_000 {
			break
		}
	}
}
