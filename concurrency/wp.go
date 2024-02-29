package main

import "fmt"

func main() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
	nWorkers := 3
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}

	for _, task := range tasks {
		jobs <- task
	}

	close(jobs)

	for i := 0; i < len(tasks); i++ {
		<-results
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		fib := Fibonacci(j)
		fmt.Println("Worker", id, "finished job", j, "result", fib)
		results <- fib
	}
}
