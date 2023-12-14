// Concurrent Counter With Mutex 2

/*
Goal: Increment a counter using go routines
1. Define Necessary Variables:
	- Declare variables like counter (initialized to 0), numRoutines (set to the desired number of goroutines), and synchronization primitives (sync.WaitGroup and sync.Mutex).
2 Increment Counter Concurrently:
	- Use the sync.WaitGroup to manage the number of goroutines to be executed.
	- Employ an anonymous goroutine function to increment the counter variable while ensuring synchronization using a sync.Mutex.
3. Parallel Execution:
	- Use a loop to create and launch multiple goroutines concurrently.
	- Each goroutine should increment the counter after acquiring and releasing the mutex lock.
	- Demonstrate some "work" within the goroutine (e.g., using time.Sleep to simulate activity).
4. Wait for Goroutines to Finish:
	- Utilize wg.Wait() to ensure the main function waits until all goroutines finish their execution.
5. Print the Final Counter Value:
	- Output the final value of the counter variable after all goroutines have finished execution.
6. Exploration and Experimentation:
	- Modify the number of goroutines or the simulated work duration within the goroutines to observe different outcomes.
	- Experiment with different synchronization methods (sync.Mutex, etc.) or additional functionalities within the goroutines.
	- Write your own tests.
*/

package main
