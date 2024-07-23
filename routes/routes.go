package routes

import (
	"net/http"

	"github.com/goStudies/controllers"
	"github.com/gorilla/mux"
)

func LoadRoutes() {
	router := mux.NewRouter()
	http.HandleFunc("/", controllers.Index)
	// router.HandleFunc("/api/produtos", controllers.GetProdutos).Methods("GET")

	// Rota para criar um produto (POST)
	router.HandleFunc("/api/produtos", controllers.CreateProduct).Methods("POST")
}
