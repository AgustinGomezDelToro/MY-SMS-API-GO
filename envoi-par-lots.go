package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	apiKey := "VOTRE CLE API"
	sender := "monentreprise"
	scheduledDeliveryDate := "21/10/2014"
	time := 9
	minute := 0

	// Construire les données JSON pour la requête POST
	data := map[string]interface{}{
		"apiKey":               apiKey,
		"sender":               sender,
		"scheduledDeliveryDate": scheduledDeliveryDate,
		"time":                 time,
		"minute":               minute,
		"SMSList": []map[string]string{
			{"phoneNumber": "06xxxxxxx1", "message": "msg 0"},
			{"phoneNumber": "06xxxxxxx2", "message": "msg 1"},
		},
	}

	// Convertir les données en JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Erreur lors de la conversion en JSON:", err)
		return
	}

	// Effectuer la requête POST
	url := "https://api.smspartner.fr/v1/bulk-send"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
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
