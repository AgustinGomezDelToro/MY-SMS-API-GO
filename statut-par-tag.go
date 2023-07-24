package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// Prepare data for GET request
	apiKey := "YOUR_API_KEY"
	tag := "montag"

	// Create GET request URL
	url := "https://api.smspartner.fr/v1/bulk-status-by-tag?" +
		"apiKey=" + apiKey + "&tag=" + tag

	// Create HTTP client
	client := &http.Client{Timeout: 10 * time.Second}

	// Send GET request
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Get response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Process your response here
	log.Printf("Response: %s", body)
}
