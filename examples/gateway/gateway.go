package main

import (
	"fmt"
	"net/http"

	"github.com/acm-uiuc/arbor"
	"github.com/acm-uiuc/arbor/proxy"
	"github.com/acm-uiuc/arbor/security"
)

var Routes = arbor.RouteCollection{
	arbor.Route{
		"GetProducts",
		"GET",
		"/products",
		getProducts,
	},
	arbor.Route{
		"CreateProduct",
		"POST",
		"/products/{id:[0-9]+}",
		createProduct,
	},
	arbor.Route{
		"GetProduct",
		"GET",
		"/product",
		getProduct,
	},
	arbor.Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	arbor.Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	arbor.Route{
		"Index",
		"GET",
		"/",
		Index,
	},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

//Arbor configurations
func configArbor() {
	security.AccessLogLocation = "tmp/log/access.log"
	security.ClientRegistryLocation = "tmp/clients.db"
	proxy.AccessControlPolicy = "*"
}

func main() {
	configArbor()
	arbor.Boot(Routes, 8000)
}
