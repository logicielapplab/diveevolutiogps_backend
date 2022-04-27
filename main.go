package main

import (
	"diveEvolution/routes"
	"github.com/gorilla/mux"
	"net/http"
)

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/", routes.IndexHandler).Methods("GET")
		router.HandleFunc("/updateIndex", routes.UpdateIndexHandler).Methods("Get")
	http.ListenAndServe(":8080", router)
}