package main

// Import necessary packages
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Prompt the user to enter the URL to crawl
	var url string
	fmt.Print("Enter URL to crawl: ")
	fmt.Scanln(&url)

	// Send a HTTP request to the server for crawling
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/crawl?url=%s", url))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected response status:", resp.StatusCode)
		return
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Display the crawling results received from the server
	fmt.Println("Crawling results:")
	fmt.Println(string(body))
}
