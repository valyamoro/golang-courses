package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

var productRepo = repos.NewProductRepo()

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product entities.Product
	json.NewDecoder(r.Body).Decode(&product)
	product = productRepo.Create(product)
	json.NewEncoder(w).Encode(product)
	w.WriteHeader(http.StatusCreated)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productIdLong, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Do not understand {id}")
		return
	}

	product, err := productRepo.GetOne(uint(productIdLong))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product not found!")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
	w.WriteHeader(http.StatusOK)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productIdLong, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Do not understand {id}")
		return
	}

	var products entities.Product
	json.NewDecoder(r.Body).Decode(&product)
	_, err = productRepo.Update(uint(productIdLong), product)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	productIdLong, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Do not understand {id}")
		return
	}
	_, err = productRepo.DeleteOne(uint(productIdLong))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product not found")
		return
	}

	json.NewEncoder(w).Encode("Product deleted successfully!")
	w.WriteHeader(http.StatusNoContent)
}
