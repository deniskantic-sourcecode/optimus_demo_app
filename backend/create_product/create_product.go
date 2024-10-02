package create_product

import (
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10 MB max request per form

	if err != nil {
		fmt.Println("Error has occured with reading form", err.Error()) // message displayed only on backend for debugging
		http.Error(w, "Problem with reading form", http.StatusBadRequest)
		return
	}

}
