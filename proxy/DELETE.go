package proxy

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func DELETE(w http.ResponseWriter, url string, r *http.Request) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		InvalidDELETE(w, err)
		log.Println(err)
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		InvalidDELETE(w, err)
		log.Println(err)
		return
	}

	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)

	switch {

	case contains(JSONHeader, res.Header["Content-Type"]):
		var data interface{}
		err = json.Unmarshal(contents, &data)
		if err != nil {
			InvalidDELETE(w, err)
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", JSONHeader)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(data); err != nil {
			InvalidDELETE(w, err)
			log.Println(err)
			return
		}
		break
	case contains(TEXTHeader, res.Header["Content-Type"]):
		if err != nil {
			InvalidDELETE(w, err)
			log.Println(err)
			return
		}

		fmt.Fprintf(w, "%s\n", string(contents))
		break

	case contains(XMLHeader, res.Header["Content-Type"]):
		var data interface{}
		err = xml.Unmarshal(contents, &data)
		if err != nil {
			InvalidDELETE(w, err)
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", XMLHeader)
		w.WriteHeader(http.StatusOK)
		if err := xml.NewEncoder(w).Encode(data); err != nil {
			InvalidDELETE(w, err)
			log.Println(err)
			return
		}
		break
	default:
		if err != nil {
			InvalidDELETE(w, err)
			log.Println(err)
			return
		}

		fmt.Fprintf(w, "%s\n", string(contents))
		break

	}
}

func InvalidDELETE(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(422) // unprocessable entity
	data := map[string]interface{}{"Code": 422, "Text": "Unprocessable Entity", "Specfically": err}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
