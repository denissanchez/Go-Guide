package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

func Speak(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	myCat := Cat{Name: "Fido"}
	myDog := Dog{Name: "Fido"}

	Speak(myCat)
	Speak(myDog)
}
