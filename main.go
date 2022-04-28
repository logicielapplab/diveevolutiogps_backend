package main

import (
	"diveEvolution/routes"
	"github.com/gorilla/mux"
	"net/http"
)

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/getIndex", routes.IndexHandler).Methods("GET")
	router.HandleFunc("/getHeader", routes.HeaderHandler).Methods("GET")
	router.HandleFunc("/getFooter", routes.FooterHandler).Methods("GET")
	router.HandleFunc("/updateIndex", routes.UpdateIndexHandler).Methods("Get")
	//port := os.Getenv("PORT")
	http.ListenAndServe(":8080", router)
}