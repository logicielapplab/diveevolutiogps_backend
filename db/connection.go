package db

import (
	"context"
	"diveEvolution/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var client *mongo.Client
var IndexCollection *mongo.Collection

func ConnectDB () *mongo.Client{
	opt := options.Client().ApplyURI("mongodb+srv://doadmin:Z3d87ni4E91g05aX@logiciel-applab-dab57134.mongo.ondigitalocean.com/admin?authSource=admin&replicaSet=logiciel-applab&tls=true&tlsCAFile=db/ca-certificate.crt")
	client, err := mongo.Connect(context.TODO(), opt)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
func GetDocument(collection *mongo.Collection)  bson.M{
	cur, err := collection.Find(context.TODO(), bson.D{{"_id", "4164da5e-cacd-4827-8891-26945019a5be"}})
	if err != nil {
		log.Fatal(err)
	}
	var result bson.M
	for cur.Next(context.TODO()) {
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Result: ", result)
	}
	return result
}
func UpdateDocumment(collection *mongo.Collection){
	data := models.Index{
		Id: "4164da5e-cacd-4827-8891-26945019a5be",
		Body: models.Body{
			Section1: models.Section1{
				Calidad: []string{"Lorem Ipsum","Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s"},
				Precio: []string{"Lorem Ipsum","Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s"},
			},
			Section2: models.Section2{
				Items: []models.ItemsInfo{
					models.ItemsInfo{
						Title: "Lorem Ipsum",
						Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s",
						Button: "Lorem Ipsum",
					},
					models.ItemsInfo{
						Title: "Lorem Ipsum",
						Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s",
						Button: "Lorem Ipsum",
					},
					models.ItemsInfo{
						Title: "Lorem Ipsum",
						Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s",
						Button: "Lorem Ipsum",
					},
				},
			},
		},
	}
	collection.UpdateOne(context.TODO(), bson.D{{"_id", "4164da5e-cacd-4827-8891-26945019a5be"}}, bson.D{{"$set", data}})
}
