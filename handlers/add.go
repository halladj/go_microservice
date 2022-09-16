package handlers

import (
	"net/http"
	"practice/data"
)

func (p *Products) AddProducts(rw http.ResponseWriter,
	r *http.Request,
) {

	p.l.Println("Handle Post Products")
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&prod)
	p.l.Printf("Prod: %#v", &prod)

}
