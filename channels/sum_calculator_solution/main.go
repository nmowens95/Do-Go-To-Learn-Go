// Sum Calculator with Channels

/*
1. Create a function that calculates the sum of a list of integers concurrently using channels.
2. Divide the list into two halves
3. Sum each half concurrently
4. Then, sum the results to find the total sum
5. Print the final sum
*/

package main

import (
	"fmt"
)

// could use this to find median:
/*
func findMedian(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	i := len(nums) / 2
	if len(nums)%2 == 1 {
		return nums[i]
	}
	return (nums[i-1] + nums[i]/2)
}
*/

func sum(nums []int, ch chan int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	ch <- sum // send sum to channel
}

func main() {
	nums := []int{1, 5, 6, 10, 2, 4, 12, 16, 25} // sum total should be 81

	// create channels
	ch1 := make(chan int)
	ch2 := make(chan int)

	median := len(nums) / 2

	// split into to halves
	left := nums[:median]
	right := nums[median:]

	// checking both sides of the list
	fmt.Println(left)
	fmt.Println(right)

	go sum(left, ch1)  // 22
	go sum(right, ch2) // 59

	// recieves sums from channels
	sum1, sum2 := <-ch1, <-ch2

	totalSum := sum1 + sum2
	fmt.Println("1st half:", sum1)
	fmt.Println("2nd half:", sum2)
	fmt.Println("Total sum:", totalSum)
}
