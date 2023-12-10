// Swap Function with Pointers

//Instructions:

/*
1. Write a function that takes two integer pointers and swaps the values they point to
2. May need to think about using a placeholder value
3. Test the function by passing two integer variables printing their values before and after the swap.
*/

package main

import "fmt"

// func with 2 integer pointers, set a temp value for placeholder
func swapPointers(xPtr *int, yPtr *int) {
	temp := *xPtr
	*xPtr = *yPtr
	*yPtr = temp
}

func main() {
	// start with base values to work with
	x := 10
	y := 20

	fmt.Println("x:", x)
	fmt.Println("y:", y)

	// now to swap the values
	xPtr := &x
	yPtr := &y

	swapPointers(xPtr, yPtr)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}
