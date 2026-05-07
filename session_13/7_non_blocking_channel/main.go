package main

import (
	"fmt"
	"time"
)

func main() {

	messages := make(chan string)
	alerts := make(chan string)

	// Simulate message arriving later
	go func() {
		time.Sleep(3 * time.Second)
		messages <- "New chat message received"
	}()

	// Simulate alert arriving later
	go func() {
		time.Sleep(5 * time.Second)
		alerts <- "CPU usage high"
	}()

	// Check system every second WITHOUT blocking
	for i := 0; i < 7; i++ {

		select {
		case msg := <-messages:
			fmt.Println("📩 MESSAGE:", msg)

		case alert := <-alerts:
			fmt.Println("⚠️ ALERT:", alert)

		default:
			fmt.Println("🔄 No activity... system running")
		}

		time.Sleep(1 * time.Second)
	}
}