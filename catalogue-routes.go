package routes

import (
	"github.com/gorilla/mux"
	"simple-REST/pkg/controllers"
)

var RegisterProductStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/product/", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/product/", controllers.GetProduct).Methods("GET")
	router.HandleFunc("/product/{productId}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/product/{productId}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/{productId}", controllers.DeleteProduct).Methods("DELETE")
}