package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	// Prepare data for POST request
	data := map[string]string{
		"apiKey":      "YOUR_API_KEY",
		"phoneNumber": "+336xxxxxxxx",
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshaling data: %v", err)
	}

	// Create HTTP client
	client := &http.Client{Timeout: 10 * time.Second}

	// Create POST request
	req, err := http.NewRequest("POST", "https://api.smspartner.fr/v1/stop-sms/add", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Send POST request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Process your response here
	log.Printf("Response status code: %d", resp.StatusCode)
}
