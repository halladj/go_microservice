package handlers

import (
	"net/http"
	"practice/data"
	"strconv"

	"github.com/gorilla/mux"
)

func (p Products) UpdateProducts(rw http.ResponseWriter,
	r *http.Request,
) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable To Convert id", http.StatusBadRequest)
	}
	p.l.Println("Handle PUT Products")
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProfuct(id, &prod)
	if err == data.ErrorProdcutNotFound {
		http.Error(rw, "Product Not Found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product Not Found", http.StatusBadRequest)
		return
	}
}
