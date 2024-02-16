package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
	Person
}

type Director struct {
	Person
	Employee
}

func Print(p Person) {
	fmt.Printf("%+v\n", p)
}

func main() {
	director := Director{}

	director.name = "Denis"
	director.id = 5

	fmt.Printf("%+v\n", director)

	Print(director.Person)
}
