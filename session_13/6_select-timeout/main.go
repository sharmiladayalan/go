package main

import (
	"fmt"
	"time"
)

func main() {

	location := make(chan string)
	status := make(chan string)

	// 🚴 Rider location updates (slow)
	go func() {
		time.Sleep(3 * time.Second)
		location <- "Rider is 2 km away"
	}()

	// 🍔 Order status update (faster)
	go func() {
		time.Sleep(1 * time.Second)
		status <- "Order is being prepared"
	}()

	// ⏱️ wait using select
	select {
	case msg := <-location:
		fmt.Println("LOCATION UPDATE:", msg)

	case msg := <-status:
		fmt.Println("STATUS UPDATE:", msg)

	case <-time.After(2 * time.Second):
		fmt.Println("TIMEOUT: No update received")
	}
}