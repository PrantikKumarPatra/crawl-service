package dao

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/jackdanger/collectlinks"
)

var mu = &sync.Mutex{}

func CrawlURL(uri string) ([]string, error) {
	start := time.Now()

	// Avoiding and Ignoring URLs which require SSL certificates and other validations.
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	// Opening Client connection
	client := http.Client{Transport: transport}
	resp, err := client.Get(uri)
	if err != nil {
		return nil, fmt.Errorf("error fetching URL: %v", err)
	}
	// Closing the connection
	defer resp.Body.Close()

	// Extracting all links on given page using a custom public module package from Github
	links := collectlinks.All(resp.Body)
	foundUrls := make([]string, 0, len(links))
	for _, link := range links {
		// Fixing the URLs with relative or broken links before adding to the queue.
		absolute := CleanURL(link, uri)
		foundUrls = append(foundUrls, absolute)
	}

	stop := time.Now()
	Display(uri, foundUrls, start, stop)

	return foundUrls, nil
}

// Display all the URLs found on a page and print them on console together.
func Display(uri string, found []string, start time.Time, stop time.Time) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Println("Start time of crawl of this URL:", start)
	fmt.Println("Stop time of crawl of this URL :", stop)
	fmt.Println(uri)

	// Filtering Phone numbers and Email ids which are missed out by url.Parse() and only printing valid URLs.
	for _, str := range found {
		str, err := url.Parse(str)
		if err == nil {
			if str.Scheme == "http" || str.Scheme == "https" {
				fmt.Println("\t", str)
			}
		}
	}
}

func CleanURL(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseURL, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseURL.ResolveReference(uri)
	return uri.String()
}
