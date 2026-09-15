package main

import (
	 "fmt"
	// "messenger/internal/mapping"
	//"messenger/internal/usecase/message"
	"net/http"
	// "github.com/gorilla/mux"
	// "github.com/google/uuid"

	// //"encoding/json"
	// "time"
)

func main() {
	
	//u := message.Usecase{}
	

	

	//http.Handle("/",router)
	

	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}

}
