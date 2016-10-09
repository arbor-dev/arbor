package services

import (
	"net/http"
	"github.com/acm-uiuc/groot/proxy"
)

//Location
const QuotesURL string = "http://localhost:9494"

//Service Data Type
const QuoteFormat string = "JSON"

//API Interface
var QuotesRoutes = RouteCollection{
	Route{
		"GetAllQuotes",
		"GET",
		"/quotes",
		GetAllQuotes,
	},
    Route{
		"DeleteQuote",
		"DELETE",
		"/quotes/{id}",
		DeleteQuote,
	},
    Route{
		"GetQuote",
		"GET",
		"/quotes/{id}",
		GetQuote,
	},
    Route{
		"CreateQuote",
		"POST",
		"/quotes",
		CreateQuote,
	},
    Route{
		"UpdateQuote",
		"PUT",
		"/quotes/{id}",
		UpdateQuote,
	},
}

//Route handler
func GetAllQuotes(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, QuotesURL+r.URL.String(), QuoteFormat, r)
}

func DeleteQuote(w http.ResponseWriter, r *http.Request) {
	proxy.DELETE(w, QuotesURL+r.URL.String(), QuoteFormat, r)
}

func GetQuote(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, QuotesURL+r.URL.String(), QuoteFormat, r)
}

func CreateQuote(w http.ResponseWriter, r *http.Request) {
	proxy.POST(w, QuotesURL+r.URL.String(), QuoteFormat, r)
}

func UpdateQuote(w http.ResponseWriter, r *http.Request) {
	proxy.PUT(w, QuotesURL+r.URL.String(), QuoteFormat, r)
}
