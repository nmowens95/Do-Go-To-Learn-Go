// Modify Value with Pointer

/*
1. Write a function that takes an integer pointer
2. This function should add 10 to the value it points to
3. Test the function by passing an integer variable
4. Last modify its value, and print the updated value
*/

package main

import "fmt"

// looking to update the value so we need to first set a reference to num
// de-reference the num
// change value, in this case add 10
// update the value that num points to

func addTen(num *int) {
	newVal := *num
	newVal += 10
	*num = newVal
}

func main() {
	num := 10
	// check original value
	fmt.Println("Original value:", num)

	// add 10 to value
	newNum := &num

	addTen(newNum)
	fmt.Println("New value:", num)
}
