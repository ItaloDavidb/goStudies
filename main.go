package main

import (
	"net/http"

	"github.com/goStudies/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":3000", nil)
}
