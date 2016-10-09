package proxy

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GET(w http.ResponseWriter, url string, format string, token string, r *http.Request) {
    req, err := http.NewRequest("GET", url, nil)
    if format == "JSON" {
	    req.Header.Set("Content-Type", JSONHeader)
	    req.Header.Set("Accept", "application/json")
    }
    if token != "" {
	     req.Header.Set("Authorization", "Basic " + token)
    }
    client := &http.Client{}
	res, err := client.Do(req)

	if err != nil  || res.StatusCode != http.StatusOK {
		InvalidGET(w, err)
		log.Println(err)
		return
	}

	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)

	switch {
		case contains(JSONHeader, res.Header["Content-Type"]):
			jsonGET(w, url, contents)
			break
		case contains(TEXTHeader, res.Header["Content-Type"]):
			textGET(w, url, contents)
			break
		case contains(HTMLHeader, res.Header["Content-Type"]):
			textGET(w, url, contents)
			break
		case contains(XMLHeader, res.Header["Content-Type"]):
			xmlGET(w, url, contents)
			break
		default:
			if err != nil {
				InvalidGET(w, err)
				log.Println(err)
				return
			}
			textGET(w, url, contents)
			break
	}
}

func jsonGET(w http.ResponseWriter, url string, contents []byte) {
	var data interface{}
	err := json.Unmarshal(contents, &data)
	if err != nil {
		InvalidGET(w, err)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", JSONHeader)
	w.Header().Set("Access-Control-Allow-Origin", AccessControlPolicy)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		InvalidGET(w, err)
		log.Println(err)
		return
	}
}

func xmlGET(w http.ResponseWriter, url string, contents []byte) {
	var data interface{}
	err := xml.Unmarshal(contents, &data)
	if err != nil {
		InvalidGET(w, err)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", XMLHeader)
	w.Header().Set("Access-Control-Allow-Origin", AccessControlPolicy)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		InvalidGET(w, err)
		log.Println(err)
		return
	}
}

func textGET(w http.ResponseWriter, url string, contents []byte) {
	var err error
	if err != nil {
		InvalidGET(w, err)
		log.Println(err)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", AccessControlPolicy)
	fmt.Fprintf(w, "%s\n", string(contents))
}

func InvalidGET(w http.ResponseWriter, err error) {
	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", AccessControlPolicy)
	w.WriteHeader(http.StatusNotFound)
	data := map[string]interface{}{"Code": http.StatusNotFound, "Text": "Not Found", "server-err": err}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
