package main

import (
	"fmt"
	"time"
)

// Worker function
func kitchenWorker(orders chan string) {

	for order := range orders {

		fmt.Println("Started Preparing:", order)

		// Simulate cooking time
		time.Sleep(2 * time.Second)

		fmt.Println("Completed:", order)
	}
}

func main() {

	// Buffered channel with capacity 3
	orders := make(chan string, 3)

	// Start worker goroutine
	go kitchenWorker(orders)

	// Sending orders to kitchen
	orders <- "Burger"
	fmt.Println("Order Added: Burger")

	orders <- "Pizza"
	fmt.Println("Order Added: Pizza")

	orders <- "Pasta"
	fmt.Println("Order Added: Pasta")

	// Buffer becomes full here

	// This waits if worker has not consumed any order
	orders <- "Sandwich"
	fmt.Println("Order Added: Sandwich")

	// Close channel after sending all orders
	close(orders)

	// Prevent main from exiting immediately
	time.Sleep(10 * time.Second)
}