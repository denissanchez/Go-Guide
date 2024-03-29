package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	duration1 := 4 * time.Second
	duration2 := 2 * time.Second

	go DoSomething(duration1, c1, 1)
	go DoSomething(duration2, c2, 2)

	for i := 0; i < 2; i++ {
		select {
		case result := <-c1:
			fmt.Println(result)
		case result := <-c2:
			fmt.Println(result)
		}
	}
}

func DoSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}
