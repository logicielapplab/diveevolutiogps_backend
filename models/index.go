package models

type Index struct {
	Id string `bson:"_id"`
	Header	`bson:"header"`
	Body	`bson:"body"`
	Footer	`bson:"footer"`
}
type Header struct {
	NavBar `bson:"navBar"`
	Title string 	`bson:"title"`
	Button	string `bson:"button"`
}
type Body struct {
	Section1	`bson:"section1"`
	Section2	`bson:"section2"`
}

type Section1 struct {
	Calidad []string	`bson:"calidad"`
	Precio []string	`bson:"precio"`
}
type Section2 struct {
	Items []ItemsInfo	`bson:"items"`
}
type ItemsInfo struct {
	Title string	`bson:"title"`
	Description string	`bson:"description"`
	Button string	`bson:"button"`
}
