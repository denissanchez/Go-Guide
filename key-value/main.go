package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["A"] = 1
	m["B"] = 2
	m["C"] = 3

	fmt.Println(m)

	for key, value := range m {
		fmt.Println(key, value)
	}

	valiue, ok := m["D"]
	fmt.Println(valiue, ok)
}
