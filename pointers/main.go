package main

import (
	"fmt"
)

type Computer struct {
	brand   string
	memory  int
	storage int
}

func (c Computer) assignDefault() {
	c.brand = "Unknown"
	c.memory = 16
	c.storage = 512
}

func (c Computer) String() string {
	return fmt.Sprintf("Brand: %s, Memory: %d, Storage: %d", c.brand, c.memory, c.storage)
}

func (c *Computer) setMemory(newCapacity int) {
	c.memory = newCapacity
}

func (c *Computer) duplicateMemory() {
	c.memory = c.memory * 2
}

func main() {
	a := 30
	b := &a

	println(fmt.Sprintf("a: %d, b: %x, *b: %d", a, b, *b))

	*b = 10

	println(fmt.Sprintf("a: %d, b: %x, *b: %d", a, b, *b))

	var myPC Computer

	myPC.assignDefault()

	fmt.Println(myPC)

	myPC.setMemory(8)

	fmt.Println(myPC)

	myPC.duplicateMemory()

	fmt.Println(myPC)
}
