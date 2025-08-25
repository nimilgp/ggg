package main

import (
	"fmt"
	"sync"
	"time"
)

type Truck struct {
	ID int
}

func processTruck(truck Truck) {
	time.Sleep(2 * time.Second) // Simulate processing time
	println("Processing truck ID:", truck.ID)
}

func main() {
	trucks := []Truck{
		{ID: 1},
		{ID: 2},
		{ID: 3},
		{ID: 4},
		{ID: 5},
	}

	t1 := time.Now()
	var wg sync.WaitGroup
	for _, truck := range trucks {
		wg.Add(1)
		go func(t Truck) {
			defer wg.Done()
			processTruck(t)
		}(truck)
	}
	wg.Wait()
	t2 := time.Now()
	fmt.Println("Total processing time:", t2.Sub(t1))
}
