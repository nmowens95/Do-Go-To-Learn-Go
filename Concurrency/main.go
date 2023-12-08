// Instructions:

/*
1. The program should accept a list of file paths as input.
2. For each file, create a Go routine that reads the file, counts the words, and sends the count to a shared channel.
3. Use a separate Go routine to receive word counts from the channel and calculate the total count.
4. Implement a function to read a file, split its content into words, and count the number of words.
5. Ensure proper synchronization using channels and wait groups.
6. Print the total word count from all the files at the end.

Example Output:

File 1.txt: 150 words
File 2.txt: 200 words
Total word count: 350 words
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(file string) int {
	doc, err := os.Open(file)
	if err != nil {
		log.Fatalf("Unable read file %v", err) //immediately exits program
	}

	defer doc.Close() // want to wait for the function to execute before closing

	scanner := bufio.NewScanner(doc)
	lineCount := 0

	for scanner.Scan() { // .Scan() will check with a boolean if there is another line
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: % v", err)
	}

	return lineCount
}

func main() {
	file1 := readFile("file1.txt")
	fmt.Println(file1)
	file2 := readFile("file2.txt")
	fmt.Println(file2)
}
