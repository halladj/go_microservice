package handlers

import (
	"log"
	"net/http"
	"practice/data"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter,
	r *http.Request,
) {

	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProducts(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		//Expect the id in the URL
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Invalid URI Atoi", http.StatusBadRequest)
			return
		}

		p.updateProducts(id, rw, r)
		return
	}

	//Catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter,
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

func (p *Products) addProducts(rw http.ResponseWriter,
	r *http.Request,
) {

	p.l.Println("Handle Post Products")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
	p.l.Printf("Prod: %#v", prod)

}

func (p *Products) updateProducts(id int, rw http.ResponseWriter,
	r *http.Request,
) {
	p.l.Println("Handle PUT Products")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProfuct(id, prod)
	if err == data.ErrorProdcutNotFound {
		http.Error(rw, "Product Not Found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product Not Found", http.StatusBadRequest)
		return
	}
}
