package routes

import (
	"diveEvolution/db"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

var Header *mongo.Collection
var Footer *mongo.Collection
func init()  {
	client := db.ConnectDB()
	Header = client.Database("DiveEvolution").Collection("Header")
	Footer = client.Database("DiveEvolution").Collection("Footer")
}
func HeaderHandler (w http.ResponseWriter, r *http.Request){
	data := db.GetDocument(Header,"76ff0c4a-ca1c-4d62-9304-6e3a71565ff4")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
func FooterHandler (w http.ResponseWriter, r *http.Request){
	data := db.GetDocument(Footer,"f93746d6-8b27-481f-ad1f-e888f7ef6d0f")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
