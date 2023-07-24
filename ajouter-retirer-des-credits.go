package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type SubaccountData struct {
	APIKey         string `json:"apiKey"`
	Credit         string `json:"credit"`
	TokenSubaccount string `json:"tokenSubaccount"`
}

func main() {
	apiKey := "YOUR_API_KEY"
	credit := "100"
	tokenSubaccount := "identifiant du sous-compte"

	// Create POST request data
	data := SubaccountData{
		APIKey:         apiKey,
		Credit:         credit,
		TokenSubaccount: tokenSubaccount,
	}
	jsonData, _ := json.Marshal(data)

	// Create POST request
	req, err := http.NewRequest("POST", "https://api.smspartner.fr/v1/subaccount/credit/add", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
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

	// Print the response body
	log.Printf("Response body: %s", string(body))
}
