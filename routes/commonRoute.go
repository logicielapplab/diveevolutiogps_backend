package routes

import (
	"diveEvolution/db"
	"diveEvolution/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

var Header *mongo.Collection
var Footer *mongo.Collection
var HeaderImg *mongo.Collection
var FooterImg *mongo.Collection

func HeaderHandler (w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	language := params["lang"]
	Header = Client.Database("DiveEvolution").Collection("Header")
	data := db.GetDocument(Header,utils.GetLang(language,"header"))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
func FooterHandler (w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	language := params["lang"]
	Footer = Client.Database("DiveEvolution").Collection("Footer")
	data := db.GetDocument(Footer,utils.GetLang(language,"footer"))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
func HeaderImgHandler (w http.ResponseWriter, r *http.Request){
	HeaderImg = Client.Database("DiveEvolution").Collection("HeaderImg")
	data := db.GetDocument(HeaderImg,"7ec330b9-b663-4725-abb2-0786e28ca6ba")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
func FooterImgHandler (w http.ResponseWriter, r *http.Request){
	FooterImg = Client.Database("DiveEvolution").Collection("FooterImg")
	data := db.GetDocument(FooterImg,"9756d287-94ae-47de-aacc-c454b9ff0ce8")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}