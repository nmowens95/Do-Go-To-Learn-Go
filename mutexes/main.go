// Concurrent Counter with Mutex

/*
1. Implement a concurrent counter using a mutex to synchronize access.
2. Create multiple goroutines that increment the counter simultaneously.
3. Ensure that the counter is safely incremented and print its final value.
*/

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) increment(wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()         // lock to ensure exclusive access
	defer c.mu.Unlock() // ensures unlock
	c.value++
}

func main() {
	const numRoutines = 5

	counter := Counter{}  // initialize counter struct
	var wg sync.WaitGroup // Used to wait for all go routines to finish: Add(), defer/Done(), Wait()

	// add number of go routines
	wg.Add(numRoutines)

	// create go routines to increment counter concurrently
	for i := 0; i < numRoutines; i++ {
		go counter.increment(&wg) // looking at wg address
	}

	wg.Wait() // wait for all go routines to finish

	fmt.Println("Final Counter Value:", counter.value)
}
