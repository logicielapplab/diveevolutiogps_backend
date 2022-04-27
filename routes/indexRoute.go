package routes

import (
	"diveEvolution/db"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)
var IndexCollection *mongo.Collection
func init()  {
	client := db.ConnectDB()
	IndexCollection = client.Database("DiveEvolution").Collection("Index")
}
func IndexHandler (w http.ResponseWriter, r *http.Request){
	data := db.GetDocument(IndexCollection)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
	fmt.Println("EJECUTADO")
}
func UpdateIndexHandler (w http.ResponseWriter, r *http.Request)  {
	db.UpdateDocumment(IndexCollection)
}
