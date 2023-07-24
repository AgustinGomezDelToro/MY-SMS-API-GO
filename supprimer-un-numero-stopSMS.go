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
	id := "666"

	// Create HTTP client with timeout
	client := &http.Client{Timeout: 10 * time.Second}

	// Create GET request URL
	url := "https://api.smspartner.fr/v1/stop-sms/delete?apiKey=" + apiKey + "&id=" + id

	// Send GET request
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Process your response here
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	log.Printf("Response: %s", body)
}
