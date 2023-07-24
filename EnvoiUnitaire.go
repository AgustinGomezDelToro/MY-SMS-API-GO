package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	apiKey := "VOTRE CLE API"
	phoneNumbers := "+336xxxxxxxx"
	message := "Ceci est votre message"
	sender := "monentreprise"
	scheduledDeliveryDate := "21/10/2014"
	time := 9
	minute := 0

	// Construire le corps JSON pour la requête POST
	jsonData := fmt.Sprintf(`{
		"apiKey": "%s",
		"phoneNumbers": "%s",
		"message": "%s",
		"sender": "%s",
		"scheduledDeliveryDate": "%s",
		"time": %d,
		"minute": %d
	}`, apiKey, phoneNumbers, message, sender, scheduledDeliveryDate, time, minute)

	// Effectuer la requête POST
	url := "https://api.smspartner.fr/v1/send"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		fmt.Println("Erreur lors de la création de la requête:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi de la requête:", err)
		return
	}
	defer resp.Body.Close()

	// Traiter la réponse
	if resp.StatusCode == http.StatusOK {
		body := new(bytes.Buffer)
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			fmt.Println("Erreur lors de la lecture de la réponse:", err)
			return
		}
		fmt.Println(body.String())
	} else {
		fmt.Println("La requête POST a échoué. Code de réponse:", resp.StatusCode)
	}
}
