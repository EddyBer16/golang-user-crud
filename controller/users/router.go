package controller

import (
	"net/http"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/users/", routesHandler())
	return router
}
