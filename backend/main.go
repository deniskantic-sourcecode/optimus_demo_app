package main

import (
	"backend/create_product"
	"fmt"
	"log"
	"net/http"
)

func setupRoutes() {

	mux := http.NewServeMux()
	mux.HandleFunc("/createProduct", create_product.CreateProduct)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func main() {

	fmt.Println("Server is running on port 8080")
	setupRoutes()
}
