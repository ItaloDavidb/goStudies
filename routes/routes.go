package routes

import (
	"net/http"

	"github.com/goStudies/controllers"
	"github.com/gorilla/mux"
)

func LoadRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Index)
	router.HandleFunc("/api/produtos", controllers.GetAllProducts).Methods("GET")
	router.HandleFunc("/api/produtos", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/produtos/{id}", controllers.UpdateProduct).Methods("PATCH")
	router.HandleFunc("/api/produtos/{id}", controllers.RemoveProduct).Methods("DELETE")
	http.Handle("/", router)
}
