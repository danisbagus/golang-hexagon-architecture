package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/danisbagus/golang-hexagon-architecture/internal/core/service"
	v1 "github.com/danisbagus/golang-hexagon-architecture/internal/handler/v1"
	"github.com/danisbagus/golang-hexagon-architecture/internal/repo"

	"github.com/gorilla/mux"
)

func main() {
	// multiplexer
	router := mux.NewRouter()

	// repo
	productRepo := repo.NewProductRepo()

	// service
	productService := service.NewProductService(productRepo)

	// handler v1
	productHandlerV1 := v1.ProductHandler{Service: productService}

	// routing v1
	v1Router := router.PathPrefix("/v1").Subrouter()
	v1Router.HandleFunc("/products", productHandlerV1.GetProductListV1).Methods(http.MethodGet).Name("GetProductList")
	v1Router.HandleFunc("/products/{product_id:[0-9]+}", productHandlerV1.GetProductDetailV1).Methods(http.MethodGet).Name("GetProductDetail")

	// starting server
	fmt.Println("Starting the application")
	log.Fatal(http.ListenAndServe("localhost:9000", router))
}
