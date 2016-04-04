package proxy

import (
	"encoding/json"
	"log"
	"net/http"
)

func DELETE(w http.ResponseWriter, url string, format string, r *http.Request) {
	req, err := http.NewRequest("DELETE", url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)

	if resp.StatusCode != http.StatusOK {
		InvalidDELETE(w, err)
		log.Printf("Failed server %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", AccessControlPolicy)
}

func InvalidDELETE(w http.ResponseWriter, err error) {
	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", AccessControlPolicy)
	w.WriteHeader(http.StatusNotFound)
	data := map[string]interface{}{"Code": http.StatusNotFound, "Text": "Not Found", "Specfically": err}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
