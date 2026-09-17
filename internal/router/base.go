package router

import (
"github.com/gorilla/mux"
"messenger/internal/usecase/message"
)

func New() *mux.Router {
	router := mux.NewRouter()
	
	u := message.Usecase{}

	_ = newA(router, &u)
	_ = newM(router, &u)

	



	return router

}



