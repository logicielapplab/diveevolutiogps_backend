package models


type About struct {
	History `bson:"history"`
}
type History struct {
	Introduction string`bson:"introduction"`
	Brief string `bson:"brief"`
	Question string `bson:"question"`
	Button string `bson:"button"`
}
type MVission struct {
	Description string `bson:"description"`
}
type Value struct {
	Description string `bson:"description"`
	Values []string `bson:"values"`
}