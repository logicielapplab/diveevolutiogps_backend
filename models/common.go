package models

type Footer struct {
	NavBar `bson:"navBar"`
	Phone string `bson:"phone"`
	SocialMedia `bson:"social_media"`
}
type SocialMedia struct {
	Facebook string `bson:"facebook"`
	Instagram string `bson:"instagram"`
	Twitter string `bson:"twitter"`
}
type NavBar struct {
		NavBarItems []NavBarItem `bson:"nav_bar_items"`
}
type NavBarItem struct {
	Title string `bson:"title"`
	Accordion []string 	`bson:"accordion"`
}

