package main

import (
	"fmt"
	"time"
)

// Worker function
func worker(id int, jobs chan string, done chan bool) {

	for job := range jobs {

		fmt.Printf("Worker %d started %s\n", id, job)

		time.Sleep(2 * time.Second)

		fmt.Printf("Worker %d completed %s\n", id, job)
	}

	// Notify completion
	done <- true
}

func main() {

	// Job queue
	jobs := make(chan string, 5)

	// Completion signal channel
	done := make(chan bool)

	// Start 3 workers
	go worker(1, jobs, done)
	go worker(2, jobs, done)
	go worker(3, jobs, done)

	// Add jobs
	jobs <- "Order-1"
	jobs <- "Order-2"
	jobs <- "Order-3"
	jobs <- "Order-4"
	jobs <- "Order-5"

	// No more jobs
	close(jobs)

	// Wait for all workers
	<-done
	<-done
	<-done

	fmt.Println("All workers completed")
}