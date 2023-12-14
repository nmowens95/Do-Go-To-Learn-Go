// Concurrent Counter with Mutex 2

/*
1. Initialize Counter: Create a shared counter variable.
2. Import Required Packages: Import the sync package for synchronization utilities.
3. Initialize Mutex: Declare a mutex (sync.Mutex) to control access to the shared counter.
4. Create Goroutines: Use a loop to create multiple goroutines that will concurrently increment the counter.
5. Increment with Mutex: Inside each goroutine, use a mutex to lock before incrementing the counter and unlock afterward.
	This ensures exclusive access to the shared resource during modification.
6. Use sync.WaitGroup: Employ a sync.WaitGroup to wait for all goroutines to finish their increments before proceeding.
7. Wait for Goroutines to Finish: Use wg.Add() to add the number of goroutines, wg.Done() to signal completion from each goroutine,
	and wg.Wait() to wait for all goroutines to finish.
8. Print Final Counter Value: Display the final value of the shared counter after all goroutines have completed their increments.

This approach ensures safe concurrent access to the counter using a mutex to prevent race conditions and synchronize the modifications.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// declare the variables we are going to need
	var (
		counter     = 0
		numRoutines = 10
		wg          sync.WaitGroup
		mu          sync.Mutex
	)

	// add number of goroutines
	wg.Add(numRoutines)

	// use an anonymous go routine function to envoke our increment
	for i := 0; i < numRoutines; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()

			// using sleep to demonstrate work
			time.Sleep(time.Millisecond * 100)

			counter++
		}()
	}

	// wait for all go routines to finish
	wg.Wait()

	fmt.Println(counter)
}
