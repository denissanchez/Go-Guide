package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go DoSomething(i, &wg)
	}

	wg.Wait()
}

func DoSomething(i int, wg *sync.WaitGroup) int {
	defer wg.Done()

	fmt.Printf("Doing something (%v)...\n", i)
	time.Sleep(3 * time.Second)
	fmt.Printf("Done (%v)!", i)

	return i
}
