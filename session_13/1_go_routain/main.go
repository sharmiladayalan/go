package main

import (
	"fmt"
	"time"
)

func demo(from string) {
	fmt.Println("From:", from)
}

func main() {

	// Normal function call
	// Main function waits until demo() xcompletes
	demo("Direct function call")

	// Goroutine call
	// Runs concurrently (main will NOT wait)
	go demo("Goroutine call")

	// Anonymous goroutine function
	// func(data string) -> anonymous function with parameter
	// no function name is given
	// used mostly for one-time logic
	// () at the end executes immediately
	go func(data string) {

		// Local variable inside anonymous function
		message := "Inside anonymous function"

		// Accessing parameter
		fmt.Println("From:", data)

		// Printing local variable
		fmt.Println(message)

		// Loop inside anonymous function
		for i := 1; i <= 3; i++ {
			fmt.Println("Counter:", i)
		}

		// Anonymous functions can access outer variables
		// This behavior is called closure

	}("Call from goroutine func")

	fmt.Println("Message from main function")

	duration := time.Second

	fmt.Println("Sleep with duration", duration)

	// Sleep is used here to wait for goroutines
	// Otherwise main function may exit early
	time.Sleep(duration)

	fmt.Println("All function executed")

	// Deferred anonymous function
	// defer executes LAST before main function exits
	defer func() {
		fmt.Println("Defer function")
	}()
}

