// Sum Calculator with Channels

/*
Goal: To sum together two halves of a list with the use of channels & go routines
1. Implement the sum Function:
	- Define a function named sum that accepts a slice of integers and a channel. (make own list)
	- Within the function, calculate the sum of the integers in the slice and send the result through the channel.
2. Construct the main Function:
	- Create a slice of integers (nums) containing various numbers.
	- Make two channels (ch1 and ch2) to receive the sums from the calculated halves.
3. Divide the Numbers into Halves:
	-Compute the midpoint of the nums slice to divide it into two halves (left and right).
4. Utilize Goroutines:
	- Print the left and right halves.
	- Implement goroutines to concurrently calculate the sum of each half using the sum function.
	- Use goroutines to calculate the sums of left and right halves and pass the results to their respective channels.
5. Retrieve and Calculate the Total Sum:
	- Receive the calculated sums from channels (ch1 and ch2) in the main function.
	- Calculate the total sum by adding both halves' sums.
6. Display the Results:
	- Print the sums of the first and second halves.
	- Print the total sum of all numbers.
7. Explore and Experiment:
	- Modify the numbers in the nums slice to observe changes in calculations.
	- Experiment with different functionalities within the sum function or main.
	- Write your own tests
*/

package main
