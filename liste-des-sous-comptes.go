package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	apiKey := "YOUR_API_KEY"
	url := "https://api.smspartner.fr/v1/subaccount/list?apiKey=" + apiKey

	// Create GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	// Print the response body
	log.Printf("Response body: %s", string(body))
}
