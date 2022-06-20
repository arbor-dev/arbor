package products

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
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
	err := a.Srv.ListenAndServe()
	if err != nil {
		if err.Error() == "http: Server closed" {
			return
		}
		log.Fatal(err.Error())
	}
	fmt.Println("Spinning down Example Product Service")
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
	a.Router.HandleFunc("/ws/{id:[0-9]+}", a.PriceUpdater).Methods("GET")
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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (a *App) PriceUpdater(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		fmt.Println(err.Error())
		return
	}

	p := product{ID: id}
	p, err = a.Model.getProduct(p)

	if err != nil {
		respondWithError(w, http.StatusNotFound, "Product not found")
		fmt.Println(err.Error())
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	conn.WriteMessage(websocket.TextMessage, []byte(strconv.FormatFloat((float64)(p.Price), 'f', 2, 32)))

	// This function basically would add the connection and its associated item id to a map (unique id
	// for each connection) atomically and would spawn a goroutine in order to listen for incoming messages
	// from the client

	// Given this example, the server should broadcast to all the connections in the array whenever
	// a price is updated
	// The goroutine would wait until the client sends back a close message, on which it will remove the
	// connection from the array

	// Currently in this unimplemented version, the socket will get closed by this server and behaves
	// similarly to a normal HTTP request
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
