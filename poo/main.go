package main

import "fmt"

type Employee struct {
	id        int
	name      string
	entryYear int
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) SetEntryYear(entryYear int) {
	e.entryYear = entryYear
}

func (e Employee) GetId() int {
	return e.id
}

func (e Employee) GetName() string {
	return e.name
}

func (e Employee) GetEntryYear() int {
	return e.entryYear
}

func main() {
	e := Employee{}
	fmt.Printf("%v\n", e)

	e.id = 0
	e.name = "Denis"
	e.entryYear = 2019

	fmt.Printf("%+v\n", e)

	e.SetId(2)
	e.SetName("Xiomi")
	e.SetEntryYear(2019)

	fmt.Printf("%+v\n", e)

	fmt.Printf("Id: %d\n", e.GetId())
	fmt.Printf("Name: %s\n", e.GetName())
	fmt.Printf("Entry year: %d\n", e.GetEntryYear())
}
