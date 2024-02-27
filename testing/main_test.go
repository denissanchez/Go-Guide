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
		{50, 12586269025},
	}

	for _, table := range tables {
		total := Fibonacci(table.a)
		if total != table.n {
			t.Errorf("Fibonacci was incorrect, got: %d, want: %d.", total, table.n)
		}
	}
}
