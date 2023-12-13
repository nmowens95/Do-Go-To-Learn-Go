// Psuedo code:

/*
- need a crawl function that will keep track of the visited
- crawlUrl function
	- check if visited, if visited exit
	- get the url, handle errors
	- handle status code
	- crate a tokenizer html.NewTokenizer(resp.Body)
		- tokenizer next to check each new token, if errortoken break

	- recursively call on itself
*/

package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

func crawl(url string, depth, maxDepth int) {
	visited := make(map[string]bool) // track visited URLs (visited[url] = true/false)

	crawlURL(url, visited, depth, maxDepth)
}

func crawlURL(url string, visited map[string]bool, depth, maxDepth int) {
	if depth > maxDepth || visited[url] {
		return // URL already visited stop crawling or max depth is exceeded
	}

	visited[url] = true
	fmt.Println("Crawling:", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching", url, ":", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:", resp.Status)
		return
	}

	tokenizer := html.NewTokenizer(resp.Body)
	// tokenizer parses the HTML
	// a token is the specific access to the tags name and its attributes
	for {
		tokenType := tokenizer.Next() // will go onto the next token parsing it returning its type or an error
		if tokenType == html.ErrorToken {
			break // at end of doc
		}

		// token.Data contains the name of the tag
		// token.Attr is a slice containing the tag's attributes
		token := tokenizer.Token()                                 //retrieve current token being processed
		if token.Type == html.StartTagToken && token.Data == "a" { // retrieving start tag and data <a>
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					link := attr.Val
					if !strings.HasPrefix(link, "http") { // has a relative url, may look like: /page2.html, ../another-page
						link = url + link // convert to absolute URL
					}
					if !visited[link] {
						crawlURL(link, visited, depth+1, maxDepth) // recursively crawl the URL
					}
				}
			}
		}
	}
}

func main() {
	baseURLs := []string{
		"https://example.com",
		"https://example.org",
		"https://example.net",
		// Add more URLs to crawl concurrently
	}

	var wg sync.WaitGroup
	wg.Add(len(baseURLs))

	for _, URL := range baseURLs {
		go func(url string) {
			defer wg.Done()
			crawl(url, 1, 3)
		}(URL)
	}

	wg.Wait()
	fmt.Println("Crawling is done")
}
