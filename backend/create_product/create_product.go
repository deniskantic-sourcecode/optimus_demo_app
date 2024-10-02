package create_product

import "net/http"

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Methods", "POST")
}
