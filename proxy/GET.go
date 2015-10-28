package proxy

import (
	"encoding/json"
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
		var data interface{}
		err = json.Unmarshal(contents, &data)
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
		break
	case contains(TEXTHeader, res.Header["Content-Type"]):
		if err != nil {
			InvalidGET(w, err)
			log.Println(err)
			return
		}

		fmt.Fprintf(w, "%s\n", string(contents))
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

func InvalidGET(w http.ResponseWriter, err error) {
	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	data := map[string]interface{}{"Code": http.StatusNotFound, "Text": "Not Found", "Specfically": err}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
