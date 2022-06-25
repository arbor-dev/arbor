package gateway

import (
	"net/http"

	"github.com/arbor-dev/arbor"
)

//URL of the Product Service API
const productServiceHTTPURL string = "http://0.0.0.0:5000"
const prodcutServiceWSURL string = "ws://0.0.0.0:5000"

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
	arbor.Route{
		Name:    "PriceUpdater",
		Method:  "GET",
		Pattern: "/ws/{id:[0-9]+}",
		Handler: websocketPriceUpdater,
	},
}

//Handlers
func getProducts(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, productServiceHTTPURL+r.URL.String(), productServiceFormat, "", r)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, productServiceHTTPURL+r.URL.String(), productServiceFormat, "", r)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, productServiceHTTPURL+r.URL.String(), productServiceFormat, "", r)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, productServiceHTTPURL+r.URL.String(), productServiceFormat, "", r)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, productServiceHTTPURL+r.URL.String(), productServiceFormat, "", r)
}

func websocketPriceUpdater(w http.ResponseWriter, r *http.Request) {
	arbor.ProxyWebsocket(w, prodcutServiceWSURL+r.URL.String(), productServiceFormat, "", r)
}
