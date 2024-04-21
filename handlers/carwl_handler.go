package handlers

import (
	"crawl-service/dao"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleCrawl(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "Missing URL parameter", http.StatusBadRequest)
		return
	}

	// Start crawling the URL
	foundUrls, err := dao.CrawlURL(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error crawling URL: %v", err), http.StatusInternalServerError)
		return
	}

	// Write the list of found URLs to the response body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foundUrls)
}
