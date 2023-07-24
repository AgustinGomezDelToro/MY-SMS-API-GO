package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// Prepare data for POST request
	data := map[string]interface{}{
		"apiKey":                "YOUR API KEY",
		"phoneNumbers":          "+336xxxxxxxx",
		"sondageIdent":          "SONDAGE_IDENT",
		"scheduledDeliveryDate": "21/10/2024",
		"time":                  9,
		"minute":                0,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error preparing data: %v", err)
	}

	// Create POST request
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", "https://api.smspartner.fr/v1/sondage/to/send", bytes.NewBuffer(payload))
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

	// Get response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Process your response here
	log.Printf("Response: %s", body)
}
