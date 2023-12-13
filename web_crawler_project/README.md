A web crawler, also known as a spider or web spider, is an automated bot or program designed to systematically browse the World Wide Web, typically for the purpose of gathering information from web pages. The primary goal of a web crawler is to index or catalog web content, usually for search engines, but it can serve various other purposes as well.

Here's a breakdown of what a web crawler does:

1. **Fetching Web Pages**: A web crawler starts by fetching web pages (HTML, XML, etc.) from a list of starting URLs or seed/base URLs. These base URLs are typically provided to the crawler as a starting point.

2. **Parsing Content**: Once a web page is fetched, the crawler parses its content. It extracts various pieces of information from the page, such as links, text, images, metadata, etc. Parsing involves analyzing the HTML structure to identify different elements and their attributes.

3. **Following Links**: The crawler identifies and follows hyperlinks (usually found within `<a>` tags) to other web pages within the same domain or even across different domains. It adds these newly discovered URLs to a list for future crawling.

4. **Storing Information**: As the crawler navigates through web pages, it collects and stores relevant information obtained from the pages. This could include text content, titles, URLs, meta data, etc., which may later be used for indexing or other purposes.

5. **Avoiding Duplication**: To avoid endless loops and excessive resource usage, crawlers usually maintain a record of visited URLs to prevent revisiting the same pages repeatedly.

6. **Recursion or Iteration**: Crawlers often use recursive or iterative processes to continuously navigate from one page to another, following links and building a map of the web in the process.

7. **Respecting Robots.txt and Robots Meta Tags**: Web crawlers typically follow rules specified in a website's `robots.txt` file or `robots` meta tags on web pages. These rules instruct crawlers which pages can or cannot be accessed, how frequently they can crawl, etc.

8. **Crawling Ethics and Politeness**: A well-behaved web crawler follows ethical guidelines and practices politeness by not overwhelming servers with too many requests in a short time (rate limiting) and respecting a site's terms of service.

In essence, a web crawler systematically navigates through the web, exploring pages, extracting information, and indexing content for various purposes, such as search engine indexing, data mining, content analysis, link validation, and more.

Resources:
https://blog.devgenius.io/creating-an-efficient-web-crawler-in-go-e4eec36bbf8c
https://go.dev/tour/concurrency/10
https://pkg.go.dev/golang.org/x/net/html // Use for tokenization of web page
