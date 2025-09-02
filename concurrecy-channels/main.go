package main

import (
	"fmt"
	"time"
)

// worker is a function that will run in a separate goroutine.
// It accepts a channel to send its result back to the main function.
func worker(done chan string, num int) {
	fmt.Println("👷 Worker: Starting a slow task...")
	time.Sleep(time.Second * 2) // Simulate a task that takes 2 seconds
	fmt.Println("✅ Worker: Task finished.")

	// Send the result back through the channel.
	done <- "worker " + fmt.Sprint(num) + " done"
}

func main() {
	messages := make(chan string, 3)
	t1 := time.Now()
	go worker(messages, 1)
	go worker(messages, 2)
	go worker(messages, 3)

	fmt.Println("⏳ Main: Waiting for the workers to finish...")

	for i := 0; i < 3; i++ {
		result := <-messages
		fmt.Println("Message from worker:", result)
	}
	t2 := time.Now()
	fmt.Println("Total processing time:", t2.Sub(t1))
}
