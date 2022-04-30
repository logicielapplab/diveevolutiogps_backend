package routes

import (
	"diveEvolution/db"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

var Header *mongo.Collection
var Footer *mongo.Collection
var HeaderImg *mongo.Collection
var FooterImg *mongo.Collection

func HeaderHandler (w http.ResponseWriter, r *http.Request){
	Header = Client.Database("DiveEvolution").Collection("Header")
	data := db.GetDocument(Header,"76ff0c4a-ca1c-4d62-9304-6e3a71565ff4")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
func FooterHandler (w http.ResponseWriter, r *http.Request){
	Footer = Client.Database("DiveEvolution").Collection("Footer")
	data := db.GetDocument(Footer,"f93746d6-8b27-481f-ad1f-e888f7ef6d0f")
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
	data := db.GetDocument(FooterImg,"52f6cbe4-652e-4e33-a3b7-2aeb5bdae074")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}