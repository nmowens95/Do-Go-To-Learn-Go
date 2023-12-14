// Fan-out, Fan-in with Channels

/*
1. Create a function that generates random numbers concurrently
2. To achieve this we need to perform mulitple go routine operations
3. The generated random numbers should be processed by separate goroutines
4. For the final results, combine each using another goroutine.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator(nums int, ch chan int) {
	rand.NewSource(time.Now().UnixNano())

	for i := 0; i < nums; i++ {
		ch <- rand.Intn(100) // random integer in range 0, n and send random num to ch
	}
	close(ch) // close channel after sending all the nums
}

func processor(id int, in, out chan int) {
	for num := range in {
		// Perform some processing on the received number (e.g., multiply by 2)
		result := num * 2
		fmt.Printf("Processor %d: Received %d, Result %d\n", id, num, result)
		out <- result // Send the processed result to the output channel
	}
	close(out) // Close the output channel after processing is done
}

func aggregator(channels ...chan int) chan int {
	resultCh := make(chan int)

	go func() {
		for _, ch := range channels {
			for num := range ch {
				// Aggregate the processed results (e.g., sum them up)
				// Adjust the aggregation logic as needed
				// Here, we're summing up the results
				// Modify this part for different aggregation operations
				// (e.g., find the maximum, minimum, or average)
				// You might need to track the number of results processed
				// to calculate an accurate average, for instance.
				resultCh <- num
			}
		}
		close(resultCh) // Close the result channel when done aggregating
	}()

	return resultCh
}

func main() {
	const numGenerators = 3
	const numProcessors = 5

	// Create channels for communication between generators, processors, and aggregator
	generatorChs := make([]chan int, numGenerators)
	processorChs := make([]chan int, numProcessors)

	// Start generators
	for i := 0; i < numGenerators; i++ {
		generatorChs[i] = make(chan int)
		go generator(5, generatorChs[i]) // Generate 5 random numbers per generator
	}

	// Start processors
	for i := 0; i < numProcessors; i++ {
		processorChs[i] = make(chan int)
		go processor(i+1, generatorChs[i%numGenerators], processorChs[i])
	}

	// Aggregate results
	resultCh := aggregator(processorChs...)

	// Receive and process aggregated results
	for num := range resultCh {
		// Process the aggregated results here, if needed
		// For example, print them or perform further computations
		fmt.Println("Aggregated Result:", num)
	}
}
