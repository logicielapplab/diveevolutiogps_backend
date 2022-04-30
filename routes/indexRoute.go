package routes

import (
	"diveEvolution/db"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)
var Index *mongo.Collection
var IndexImg *mongo.Collection

func IndexHandler (w http.ResponseWriter, r *http.Request){
	Index = Client.Database("DiveEvolution").Collection("Index")
	data := db.GetDocument(Index,"4164da5e-cacd-4827-8891-26945019a5be")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
func IndexImgHandler (w http.ResponseWriter, r *http.Request){
	IndexImg = Client.Database("DiveEvolution").Collection("IndexImg")
	data := db.GetDocument(IndexImg,"15388a1e-faaf-4fcb-9244-1ffa2813be59")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
func UpdateIndexHandler (w http.ResponseWriter, r *http.Request)  {
	db.UpdateDocumment([]*mongo.Collection{Index, Header, Footer})
}
