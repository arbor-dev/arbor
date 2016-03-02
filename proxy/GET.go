package proxy

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GET(w http.ResponseWriter, url string, r *http.Request) {
	res, err := http.Get(url)
	if err != nil {
		InvalidGET(w, err)
		log.Println(err)
		return
	}

	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)

	switch {
		case contains(JSONHeader, res.Header["Content-Type"]):
			jsonGET(w, url, r, contents)
			break
		case contains(TEXTHeader, res.Header["Content-Type"]):
			textGET(w, url, r, contents)
			break

		case contains(XMLHeader, res.Header["Content-Type"]):
			xmlGET(w, url, r, contents)
			break
		default:
			if err != nil {
				InvalidGET(w, err)
				log.Println(err)
				return
			}
			fmt.Fprintf(w, "%s\n", string(contents))
			break
	}
}

func jsonGET(w http.ResponseWriter, url string, r *http.Request, contents []byte) {
	var data interface{}
	err := json.Unmarshal(contents, &data)
	if err != nil {
		InvalidGET(w, err)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", JSONHeader)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		InvalidGET(w, err)
		log.Println(err)
		return
	}
}

func xmlGET(w http.ResponseWriter, url string, r *http.Request, contents []byte) {
	var data interface{}
	err := xml.Unmarshal(contents, &data)
	if err != nil {
		InvalidGET(w, err)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", XMLHeader)
	w.WriteHeader(http.StatusOK)
	if err := xml.NewEncoder(w).Encode(data); err != nil {
		InvalidGET(w, err)
		log.Println(err)
		return
	}
}

func textGET(w http.ResponseWriter, url string, r *http.Request, contents []byte) {
	var err error
	if err != nil {
		InvalidGET(w, err)
		log.Println(err)
		return
	}

	fmt.Fprintf(w, "%s\n", string(contents))
}

func InvalidGET(w http.ResponseWriter, err error) {
	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	data := map[string]interface{}{"Code": http.StatusNotFound, "Text": "Not Found", "Specfically": err}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
