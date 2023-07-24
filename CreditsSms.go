package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	apiKey := "your_api_key"
	url := "https://api.smspartner.fr/v1/me?apiKey=" + apiKey

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("GET request failed:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Failed to read response body:", err)
			return
		}
		fmt.Println(string(bodyBytes))
	} else {
		fmt.Println("GET request failed. Response Code:", response.StatusCode)
	}
}
