package main

import (
	"net/http"

	"github.com/acm-uiuc/arbor"
)

//URL of the Product Service API
const productServiceURL string = "http://localhost:5000"

//Data format of the API
const productServiceFormat string = "JSON"

//Collection of routes to expose
var productServiceRoutes = arbor.RouteCollection{
	//Route definition
	arbor.Route{
		Name:    "GetProducts",
		Method:  "GET",
		Pattern: "/products",
		Handler: getProducts,
	},
	arbor.Route{
		Name:    "CreateProduct",
		Method:  "POST",
		Pattern: "/product",
		Handler: createProduct,
	},
	arbor.Route{
		Name:    "GetProduct",
		Method:  "GET",
		Pattern: "/products/{id:[0-9]+}",
		Handler: getProduct,
	},
	arbor.Route{
		Name:    "UpdateProduct",
		Method:  "PUT",
		Pattern: "/products/{id:[0-9]+}",
		Handler: updateProduct,
	},
	arbor.Route{
		Name:    "DeleteProduct",
		Method:  "DELETE",
		Pattern: "/products/{id:[0-9]+}",
		Handler: deleteProduct,
	},
}

//Handlers
func getProducts(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, productServiceURL+r.URL.String(), productServiceFormat, "", r)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, productServiceURL+r.URL.String(), productServiceFormat, "", r)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, productServiceURL+r.URL.String(), productServiceFormat, "", r)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, productServiceURL+r.URL.String(), productServiceFormat, "", r)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, productServiceURL+r.URL.String(), productServiceFormat, "", r)
}
