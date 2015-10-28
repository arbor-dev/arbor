package proxy

import (
	"encoding/json"
	"errors"
)

func contains(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func extract(a interface{}, b string) (val string, err error) {

	j, err := json.Marshal(a)
	if err != nil {
		return "", errors.New("key does not exist")
	}
	// a map container to decode the JSON structure into
	c := make(map[string]interface{})

	// unmarschal JSON
	e := json.Unmarshal(j, &c)

	// panic on error
	if e != nil {
		panic(e)
	}

	// a string slice to hold the keys
	k := make([]string, len(c))

	// iteration counter
	i := 0

	// copy c's keys into k
	for s, _ := range c {
		k[i] = s
		i++
	}

	if contains(b, k) {
		return c[b].(string), nil
	} else {
		return "", errors.New("key does not exist")
	}

}
