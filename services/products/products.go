package products

import (
	"encoding/json"
	"net/http"
	"strconv"
	"webapp-course-be/db"
	"webapp-course-be/db/dbmodel"
	"webapp-course-be/model"

	"github.com/gorilla/mux"
)

// Create add a new product
func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var product dbmodel.Product
	var response model.Response
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		response.Code = "-1"
		response.Msg = "Error decoding body: " + err.Error()
		payload, _ := json.Marshal(response)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(string(payload)))
		return
	}

	conn := db.GetConnection()
	conn.Product.Insert(product)
	response.Code = "0"
	response.Msg = "Success"
	payload, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(payload)))
	return
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	conn := db.GetConnection()
	products := conn.Product.GetAll()

	payload, _ := json.Marshal(products)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(payload)))
	return
}

func GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var response model.Response
	vars := mux.Vars(r)
	productId, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Code = "-1"
		response.Msg = "Error parsing product id: " + err.Error()
		payload, _ := json.Marshal(response)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(string(payload)))
		return
	}

	conn := db.GetConnection()
	product := conn.Product.GetById(&productId)
	payload, _ := json.Marshal(product)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(string(payload)))
	return
}

func Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var product dbmodel.Product
	var response model.Response

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		response.Code = "-1"
		response.Msg = "Error decoding body: " + err.Error()
		payload, _ := json.Marshal(response)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(string(payload)))
		return
	}

	conn := db.GetConnection()
	idUpdated := conn.Product.Update(&product)
	if idUpdated == nil {
		response.Code = "-1"
		response.Msg = "Product not found"
	} else {
		response.Code = "0"
		response.Msg = "Success"
	}
	payload, _ := json.Marshal(response)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(string(payload)))
	return
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var response model.Response

	vars := mux.Vars(r)
	productId, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Code = "-1"
		response.Msg = "Error parsing product id: " + err.Error()
		payload, _ := json.Marshal(response)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(string(payload)))
		return
	}

	conn := db.GetConnection()
	idDeleted := conn.Product.Delete(&productId)
	if idDeleted == nil {
		response.Code = "-1"
		response.Msg = "Product not found"
	} else {
		response.Code = "0"
		response.Msg = "Success"
	}
	payload, _ := json.Marshal(response)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(string(payload)))
	return
}
