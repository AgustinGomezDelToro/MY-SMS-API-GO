package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Fields struct {
	APIKey     string      `json:"apiKey"`
	Type       string      `json:"type"`
	Parameters Parameters `json:"parameters"`
}

type Parameters struct {
	Email            string `json:"email"`
	CreditToAttribute int    `json:"creditToAttribute"`
	IsBuyer          int    `json:"isBuyer"`
	Firstname        string `json:"firstname"`
	Lastname         string `json:"lastname"`
}

func main() {
	// Prepare data for POST request
	data := Fields{
		APIKey: "YOUR_API_KEY",
		Type:   "advanced",
		Parameters: Parameters{
			Email:            "aaaa@bbb.ccc",
			CreditToAttribute: 10,
			IsBuyer:          0,
			Firstname:        "prenom",
			Lastname:         "nom",
		},
	}

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(data)

	// Create POST request
	req, err := http.NewRequest("POST", "https://api.smspartner.fr/v1/subaccount/create", payloadBuf)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	// Create HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	// Print the response status and body
	log.Printf("Response status: %s", resp.Status)
	log.Printf("Response body: %s", string(body))
}
