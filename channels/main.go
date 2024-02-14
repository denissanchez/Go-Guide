package main

import (
	"fmt"
	"time"
)

func say(msg string, c chan<- string) {
	time.Sleep(time.Second * 2)
	c <- msg
}

func first() {
	c := make(chan string, 2)

	fmt.Println(time.Now())

	go say("Hello", c)
	go say("World", c)

	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println(time.Now())
}

func message(msg string, c chan<- string) {
	c <- msg
}

func second() {
	c := make(chan string, 2)

	c <- "Hello"
	c <- "World"

	fmt.Println(len(c), cap(c))

	close(c)

	for message := range c {
		fmt.Println(message)
	}

	username := make(chan string, 3)
	email := make(chan string, 3)

	go message("Hello @username", username)
	go message("Hello @email", email)

	for i := 0; i < 2; i++ {
		select {
		case msg := <-username:
			fmt.Println(fmt.Sprintf("Username: %s", msg))
		case msg := <-email:
			fmt.Println(fmt.Sprintf("Email: %s", msg))
		}
	}
}

func main() {
	// first()
	second()
}
