package main

import (
	"diveEvolution/routes"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/", routes.IndexHandler).Methods("GET")
	router.HandleFunc("/updateIndex", routes.UpdateIndexHandler).Methods("Get")
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, router)
}