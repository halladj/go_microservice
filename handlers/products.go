package handlers

import (
	"context"
	"log"
	"net/http"
	"practice/data"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter,
	r *http.Request,
) {
	p.l.Println("Handle Get Products")
	//get the data from source
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal data", http.StatusInternalServerError)
	}
}

func (p *Products) AddProducts(rw http.ResponseWriter,
	r *http.Request,
) {

	p.l.Println("Handle Post Products")
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&prod)
	p.l.Printf("Prod: %#v", &prod)

}

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

type KeyProduct struct{}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(r.Body)

		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)

	})
}
