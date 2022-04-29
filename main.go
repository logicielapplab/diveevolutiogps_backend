package main

import (
	"diveEvolution/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/api/getIndex", routes.IndexHandler).Methods("GET")
	router.HandleFunc("/api/getHeader", routes.HeaderHandler).Methods("GET")
	router.HandleFunc("/api/getFooter", routes.FooterHandler).Methods("GET")
	router.HandleFunc("/api/updateIndex", routes.UpdateIndexHandler).Methods("Get")
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origin := handlers.AllowedOrigins([]string{"*"})
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, handlers.CORS(methods, origin)(router))
}
func Home (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("https://dive-evolution.herokuapp.com/api/getIndex\nhttps://dive-evolution.herokuapp.com/api/getHeader\nhttps://dive-evolution.herokuapp.com/api/getFooter"))
}
