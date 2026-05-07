package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	// simulate slow network
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Photo uploaded"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "Text message received"
	}()

	// we want 2 messages total
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("From c1:", msg1)

		case msg2 := <-c2:
			fmt.Println("From c2:", msg2)
		}
	}
}