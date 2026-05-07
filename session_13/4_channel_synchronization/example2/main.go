package main

import (
	"fmt"
	"time"
)

// Writer Worker
func writer(notes chan string, reviewQueue chan string) {

	for note := range notes {

		fmt.Println("Writing Note:", note)

		time.Sleep(2 * time.Second)

		fmt.Println("Completed Writing:", note)

		// Send to review queue
		reviewQueue <- note
	}

	close(reviewQueue)
}

// Reviewer Worker
func reviewer(reviewQueue chan string, done chan bool) {

	for note := range reviewQueue {

		fmt.Println("*****Reviewing Note*******:", note)

		time.Sleep(1 * time.Second)

		fmt.Println("Correction Completed:", note)
	}

	done <- true
}

func main() {

	// Channels
	notes := make(chan string, 3)

	reviewQueue := make(chan string, 3)

	done := make(chan bool)

	// Start specialized workers
	go writer(notes, reviewQueue)

	go reviewer(reviewQueue, done)

	// Add notes
	notes <- "Math Notes"
	notes <- "Science Notes"
	notes <- "English Notes"

	close(notes)

	// Wait for review completion
	<-done

	fmt.Println("All Notes Processed")
}