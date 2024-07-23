package routes

import (
	"log"
	"net/http"

	"github.com/goStudies/controllers"
	"github.com/gorilla/mux"
)

func LoadRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Index)
	router.HandleFunc("/api/produtos", controllers.GetAllProducts).Methods("GET")

	// Rota para criar um produto (POST)
	router.HandleFunc("/api/produtos", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api", test).Methods("get")
	http.Handle("/", router)
}
func test(w http.ResponseWriter, r *http.Request) {
	log.Printf("Rota Ok")
}
