package main

import "time"

func Sum(a, b int) int {
	return a + b
}

func GetMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

type Person struct {
	Code      string
	Name      string
	BirthYear int
}

type Employee struct {
	Id       string
	Position string
}

type FullTimeEmployee struct {
	Employee
	Person
}

var GetPersonByCode = func(code string) (Person, error) {
	time.Sleep(60 * time.Second)
	// SELECT * FROM Person WHERE Code = code
	return Person{Code: "1", Name: "John", BirthYear: 1980}, nil
}

var GetEmployeeById = func(id string) (Employee, error) {
	time.Sleep(60 * time.Second)
	// SELECT * FROM Employee WHERE Id = id
	return Employee{Id: "1", Position: "Software Developer"}, nil
}

func GetFullTimeEmpoyeeById(code string, id string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee

	e, error := GetEmployeeById(id)

	if error != nil {
		return ftEmployee, error
	}

	ftEmployee.Employee = e

	p, error := GetPersonByCode(code)

	if error != nil {
		return ftEmployee, error
	}

	ftEmployee.Person = p

	return ftEmployee, nil
}
