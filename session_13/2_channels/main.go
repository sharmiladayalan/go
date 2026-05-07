package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Create a channel that can store string values
	// Think of it like a communication pipe
	messages := make(chan string)

	// Print the type of the channel
	fmt.Println(reflect.TypeOf(messages))

	// Start a new goroutine (lightweight thread)
	go func() {

		// Send "ping" into the channel
		// <- means send data into channel
		messages <- "ping"

	}()

	// Receive value from the channel
	// Program waits here until data arrives
	msg := <-messages

	// Print received message
	fmt.Println(msg)
}