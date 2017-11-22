package products

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	Model  *productModel
	Srv    *http.Server
}

func NewApp() *App {
	a := new(App)
	a.Router = mux.NewRouter()
	a.Model = newProductModel()
	a.initializeRoutes()
	a.Srv = &http.Server{Addr: "0.0.0.0:5000", Handler: a.Router}
	return a
}

func (a *App) Run() {
	fmt.Println("Starting Example Product Service on Port 5000")

	go func() {
		err := a.Srv.ListenAndServe()
		if err != nil {
			if err.Error() == "http: Server closed" {
				return
			}
			log.Fatal(err.Error())
		}
	}()
}

func (a *App) Kill() {
	fmt.Println("Killing example service")
	a.Srv.Shutdown(context.Background())
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/products", a.getProducts).Methods("GET")
	a.Router.HandleFunc("/product", a.createProduct).Methods("POST")
	a.Router.HandleFunc("/products/{id:[0-9]+}", a.getProduct).Methods("GET")
	a.Router.HandleFunc("/products/{id:[0-9]+}", a.updateProduct).Methods("PUT")
	a.Router.HandleFunc("/products/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
}

func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := a.Model.getProducts()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (a *App) createProduct(w http.ResponseWriter, r *http.Request) {
	var p product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		fmt.Println(err.Error())
		return
	}
	defer r.Body.Close()

	a.Model.createProduct(p)
	respondWithJSON(w, http.StatusCreated, p)
}

func (a *App) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		fmt.Println(err.Error())
		return
	}

	p := product{ID: id}
	p, err = a.Model.getProduct(p)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Product not found")
		fmt.Println(err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, p)
}

func (a *App) updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		fmt.Println(err.Error())
		return
	}

	var p product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		fmt.Println(err.Error())
		return
	}
	defer r.Body.Close()
	p.ID = id

	if err := a.Model.updateProduct(p); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		fmt.Println(err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func (a *App) deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		fmt.Println(err.Error())
		return
	}

	if err := a.Model.deleteProduct(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
