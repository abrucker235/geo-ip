package main

import (
	"github.com/gorilla/mux"
)

//Router is the main configuration for routes.
func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.method).Path(route.pattern).Name(route.name).Handler(route.handlerFunc)
	}

	return router
}
