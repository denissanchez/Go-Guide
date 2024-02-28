package main

import "testing"

func TestSum(t *testing.T) {
	// total := Sum(5, 5)

	// if total != 10 {
	// 	t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	// }

	// total = Sum(5, 0)

	// if total != 5 {
	// 	t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 5)
	// }

	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{10000, 2, 10002},
		{20, 20, 40},
	}

	for _, table := range tables {
		total := Sum(table.a, table.b)
		if total != table.n {
			t.Errorf("Sum was incorrect, got: %d, want: %d.", total, table.n)
		}
	}
}

func TestGetMax(t *testing.T) {
	total := GetMax(5, 5)

	if total != 5 {
		t.Errorf("GetMax was incorrect, got: %d, want: %d.", total, 5)
	}

	total = GetMax(5, 0)

	if total != 5 {
		t.Errorf("GetMax was incorrect, got: %d, want: %d.", total, 5)
	}
}

func TestFibonacci(t *testing.T) {
	tables := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
	}

	for _, table := range tables {
		total := Fibonacci(table.a)
		if total != table.n {
			t.Errorf("Fibonacci was incorrect, got: %d, want: %d.", total, table.n)
		}
	}
}

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		code             string
		id               string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			code: "1",
			id:   "1",
			mockFunc: func() {
				GetEmployeeById = func(id string) (Employee, error) {
					return Employee{Id: "1", Position: "Software Developer"}, nil
				}

				GetPersonByCode = func(code string) (Person, error) {
					return Person{Code: "1", Name: "John", BirthYear: 1980}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Employee: Employee{Id: "1", Position: "Software Developer"},
				Person:   Person{Code: "1", Name: "John", BirthYear: 1980},
			},
		},
	}

	orginalGetEmployeeById := GetEmployeeById
	originalGetPersonByCode := GetPersonByCode

	for _, test := range table {
		test.mockFunc()

		employee, err := GetFullTimeEmpoyeeById(test.code, test.id)

		if err != nil {
			t.Errorf("GetFullTimeEmployeeById was incorrect, got: %v, want: %v.", err, nil)
		}

		if employee.Id != test.expectedEmployee.Id {
			t.Errorf("GetFullTimeEmployeeById was incorrect, got: %v, want: %v.", employee.Id, test.expectedEmployee.Id)
		}

		if employee.Position != test.expectedEmployee.Position {
			t.Errorf("GetFullTimeEmployeeById was incorrect, got: %v, want: %v.", employee.Position, test.expectedEmployee.Position)
		}

		if employee.BirthYear != test.expectedEmployee.BirthYear {
			t.Errorf("GetFullTimeEmployeeById was incorrect, got: %v, want: %v.", employee.BirthYear, test.expectedEmployee.BirthYear)
		}

		if employee.Code != test.expectedEmployee.Code {
			t.Errorf("GetFullTimeEmployeeById was incorrect, got: %v, want: %v.", employee.Code, test.expectedEmployee.Code)
		}

		if employee.Name != test.expectedEmployee.Name {
			t.Errorf("GetFullTimeEmployeeById was incorrect, got: %v, want: %v.", employee.Name, test.expectedEmployee.Name)
		}

		GetEmployeeById = orginalGetEmployeeById
		GetPersonByCode = originalGetPersonByCode
	}
}
