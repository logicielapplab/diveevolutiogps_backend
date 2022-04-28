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
	router.HandleFunc("/getIndex", routes.IndexHandler).Methods("GET")
	router.HandleFunc("/getHeader", routes.HeaderHandler).Methods("GET")
	router.HandleFunc("/getFooter", routes.FooterHandler).Methods("GET")
	router.HandleFunc("/updateIndex", routes.UpdateIndexHandler).Methods("Get")
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origin := handlers.AllowedOrigins([]string{"*"})
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, handlers.CORS(methods, origin)(router))
}
func Home (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("https://dive-evolution.herokuapp.com/getIndex\nhttps://dive-evolution.herokuapp.com/getHeader\nhttps://dive-evolution.herokuapp.com/getFooter"))
}
