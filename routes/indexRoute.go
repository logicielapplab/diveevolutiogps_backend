package routes

import (
	"diveEvolution/db"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)
var IndexBody *mongo.Collection

func init()  {
	client := db.ConnectDB()
	IndexBody = client.Database("DiveEvolution").Collection("Index")
}
func IndexHandler (w http.ResponseWriter, r *http.Request){
	data := db.GetDocument(IndexBody,"4164da5e-cacd-4827-8891-26945019a5be")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
func UpdateIndexHandler (w http.ResponseWriter, r *http.Request)  {
	db.UpdateDocumment([]*mongo.Collection{IndexBody, Header, Footer})
}
