package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/danisbagus/golang-hexagon-architecture/internal/core/service"
	v1 "github.com/danisbagus/golang-hexagon-architecture/internal/handler/v1"
	"github.com/danisbagus/golang-hexagon-architecture/internal/repo"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
	appPort := fmt.Sprintf("localhost:%v", os.Getenv("APP_PORT"))
	fmt.Println("Starting the application :", appPort)
	log.Fatal(http.ListenAndServe(appPort, router))
}
