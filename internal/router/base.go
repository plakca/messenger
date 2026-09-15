package router

import "github.com/gorilla/mux"

func New() *mux.Router {
	router := mux.NewRouter()
	
	return router
}

