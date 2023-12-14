// Create a Function to Fetch URLs

/*
1. Design a function that takes a URL as input, makes an HTTP request, and measures the response time.
	- Utilize net/http package for making HTTP requests.
2. Use Goroutines to concurrently fetch URLs.
	-Each Goroutine should call the function designed in step 1.
3. Create a channel to receive response times or error signals from Goroutines.
4. Use a loop to launch Goroutines for fetching each URL concurrently.
5. Each Goroutine should send the response time or an error signal through the channel.
6. Receive response times or error signals from the channel.
7. Display the response times or error messages for each URL.
8. Implement error handling in case there are issues with fetching the URLs.
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// chan<- T: Send-only channel that can only send values of type T
// <-chan T: Receive-only channel that can only receive values of type T
// chan T: Bidirectional channel that can both send and receive values of type T

func fetchURL(url string, resultChan chan<- time.Duration) { // can only send response times, not read from them
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to fetch address", url, ":", err)
		resultChan <- 0 // send 0 if error
		return
	}
	defer resp.Body.Close()

	_, err = io.Copy(io.Discard, resp.Body) // not storing in memory, allows us to read then discard
	if err != nil {
		fmt.Println("Error reading response", resp, ":", err)
		resultChan <- 0
		return
	}

	duration := time.Since(start)
	resultChan <- duration
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.example.com",
		"https://www.openai.com",
		"https://www.github.com",
	}

	resultChan := make(chan time.Duration, len(urls)) // len(urls) is the buffer size

	for _, url := range urls {
		go fetchURL(url, resultChan)
		duration := <-resultChan // storing the time recieved from channel as var
		if duration != 0 {
			fmt.Printf("URL: %s | Response time: %s\n", url, duration)
		}
		fmt.Println("It didn't work :)")
	}
}
