package handlers

import (
	"net/http"
	"strconv"

	"github.com/Sahas001/go-micro/data"
	"github.com/gorilla/mux"
)


func (p *Products) Delete(w http.ResponseWriter, r *http.Request) {
	vars:= mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handle DELETE products", id)
	
	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(w, "Product Not Found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product Not Found", http.StatusInternalServerError)
		return
	}
}
