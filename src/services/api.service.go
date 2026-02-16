package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const PotterDBBaseURL = "https://api.potterdb.com/v1"

var httpClient = &http.Client{Timeout: 10 * time.Second}

func callAPI(apiURL string, target interface{}) (int, error) {
	resp, err := httpClient.Get(apiURL)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("erreur HTTP: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, fmt.Errorf("API a retourné le code %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("erreur de décodage JSON: %w", err)
	}

	return http.StatusOK, nil
}
