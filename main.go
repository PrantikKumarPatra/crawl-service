package main

import (
	"crawl-service/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/crawl", handlers.HandleCrawl)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
