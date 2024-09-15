package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	LoadAppConfig()

	router := mux.NewRouter().StrictSlash(true)

	RegisterProductRoutes(router)

	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func RegisterProductRoutes(router *mux.Router) {
	var muxBase = "/api/products"

	router.HandleFunc(muxBase, conrollers.GetProducts).Methods("GET")
	router.HandleFunc(fmt.Sprintf("%s/{id}", muxBase), controller.GetProductById).Methods("GET")
	router.HandleFunc(muxBase, controllers.CreateProduct).Methods("POST")
	router.HandleFunc(fmt.Sprintf("%s/{id}", muxBase), controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc(fmt.Sprintf("%s/{id}", muxBase), controllers.DeleteProduct).Methods("DELETE")
}
