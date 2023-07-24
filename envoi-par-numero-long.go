package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	apiKey := "YOUR API KEY"
	to := "336xxxxxxxx"
	from := "336xxxxxxxx"
	message := "This is your message"

	// Construire les données JSON pour la requête POST
	data := map[string]interface{}{
		"apiKey":  apiKey,
		"to":      to,
		"from":    from,
		"message": message,
	}

	// Convertir les données en JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Erreur lors de la conversion en JSON:", err)
		return
	}

	// Effectuer la requête POST
	url := "https://api.smspartner.fr/v1/vn/send"
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
