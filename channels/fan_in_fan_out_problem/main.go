// Fan-out, Fan-in with Channels

/*
Goal: Create a function that generates random numbers concurrently
2. To achieve this we need to perform mulitple go routine operations
3. The generated random numbers should be processed by separate goroutines
4. For the final results, combine each using another goroutine.

1. Define generator Function:
	Declare a function that generates random numbers and sends them through a channel.

2. Write processor Logic:
	Implement a function that receives numbers from a channel, processes them, and sends the processed result to another channel.

3. Create aggregator Function:
	Construct a function that aggregates processed results from multiple channels into a single channel.

4. In the main function:
	- Create channels for communication between generators, processors, and aggregator.
	- Start generators to produce random numbers.
	- Initiate processors to process generated numbers.
	- Aggregate processed results.
	- Use Go routines for the generators and processors when called.

5. Go further:
	- Modify the processing logic or aggregation method in the respective functions.
	- Observe how changes affect the final aggregated results.
	- Build your own tests.
	- Modify constants like numGenerators and numProcessors to change the number of generators and processors.
*/

package main
