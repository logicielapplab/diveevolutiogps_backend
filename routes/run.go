package routes

import (
	"context"
	"diveEvolution/db"
	"diveEvolution/models"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

var Client *mongo.Client
func init() {
	Client = db.ConnectDB()
}
func RunApp(){
	router := mux.NewRouter()
	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/api/getIndex/{lang}", IndexHandler).Methods("GET")
	router.HandleFunc("/api/getContact/{lang}", ContactHandler).Methods("GET")
	router.HandleFunc("/api/getFooter/{lang}", FooterHandler).Methods("GET")
	router.HandleFunc("/api/getIndexImg", IndexImgHandler).Methods("GET")
	router.HandleFunc("/api/getContactImg", ContactImgHandler).Methods("GET")
	router.HandleFunc("/api/getHeaderImg", HeaderImgHandler).Methods("GET")
	router.HandleFunc("/api/getFooterImg", FooterImgHandler).Methods("GET")

	router.HandleFunc("/api/updateIndex", UpdateIndexHandler).Methods("GET")
	router.HandleFunc("/api/updateHeader", writeHeader).Methods("GET")
	router.HandleFunc("/api/updateFooter", writeFooter).Methods("GET")
	router.HandleFunc("/api/updateIndexBody", writeIndexBody).Methods("GET")
	router.HandleFunc("/api/updateContactBody", writeContactBody).Methods("GET")
	router.HandleFunc("/api/updateContactImg", writeContactImg).Methods("GET")
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origin := handlers.AllowedOrigins([]string{"*"})
	//port := os.Getenv("PORT")
	http.ListenAndServe(":8080", handlers.CORS(methods, origin)(router))
}
func Home (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("/api/getIndex/es\n/api/getHeader/es\n/api/getFooter/es\n/api/getContact/es\n/api/getIndexImg\n/api/getHeaderImg\n/api/getFooterImg"))
}
func writeHeader(w http.ResponseWriter, r *http.Request){
	data := models.HeaderImg{
		Id: uuid.NewV4().String(),
		Logo: "https://via.placeholder.com/150/0000FF/808080",
		LangItems: []models.LangItem{
			{
				Lang: "es",
				Img:  "https://res.cloudinary.com/logicielapplab/image/upload/v1651212913/DiveEvolution/Header/es_e6qn2t.svg",
			},
			{
				Lang: "en",
				Img:  "https://res.cloudinary.com/logicielapplab/image/upload/v1651212904/DiveEvolution/Header/gb_gxyw57.svg",
			},
			{
				Lang: "fr",
				Img:  "https://res.cloudinary.com/logicielapplab/image/upload/v1651212904/DiveEvolution/Header/fr_bta7cj.svg",
			},
		},
	}

	HeaderImg := Client.Database("DiveEvolution").Collection("HeaderImg")
	HeaderImg.InsertOne(context.TODO(),data)
}
func writeFooter(w http.ResponseWriter, r *http.Request){
	data := models.FooterImg{
		Id: uuid.NewV4().String(),
		Logo:"https://via.placeholder.com/150/0000FF/808080",
		Background: "https://via.placeholder.com/2000x800/000000/808080",
		WhatsApp: "https://res.cloudinary.com/logicielapplab/image/upload/v1651350046/DiveEvolution/Footer/3225179_app_logo_media_popular_social_icon_juvvzd.svg",
		SocialMediaIcons: models.SocialMediaIcons{
			Facebook: "https://res.cloudinary.com/logicielapplab/image/upload/v1651296519/DiveEvolution/Footer/3225194_app_facebook_logo_media_popular_icon_wrmery.svg",
			Instagram: "https://res.cloudinary.com/logicielapplab/image/upload/v1651296520/DiveEvolution/Footer/3225191_app_instagram_logo_media_popular_icon_hokrzm.svg",
			Twitter: "https://res.cloudinary.com/logicielapplab/image/upload/v1651308103/DiveEvolution/Footer/3225183_app_logo_media_popular_social_icon_hjp5gw.svg",
		},
	}

	HeaderImg := Client.Database("DiveEvolution").Collection("FooterImg")
	HeaderImg.InsertOne(context.TODO(),data)
}
func writeIndexBody(w http.ResponseWriter, r *http.Request){
	data := models.IndexImg{
		Id: uuid.NewV4().String(),
		BodyImg: models.BodyImg{
			Background: "https://via.placeholder.com/2000x900/6a5acd/fffffff",
			Section1: models.Section1{
				Calidad: []string{
					"https://via.placeholder.com/1000x900/d1d1e0/000000",
					"https://res.cloudinary.com/logicielapplab/image/upload/v1651347812/DiveEvolution/IndexBody/award_pifluk.svg",
				},
				Precio: []string{
					"https://via.placeholder.com/1000x900/d1d1e0/000000",
					"https://res.cloudinary.com/logicielapplab/image/upload/v1651347812/DiveEvolution/IndexBody/currency-dollar_koyukh.svg",
				},
			},
			Section2Img: models.Section2Img{
				ItemsImg: []string{
					"https://via.placeholder.com/2000x900/ff6699/ffffff",
					"https://via.placeholder.com/2000x900/ffad33/ffffff",
				},
			},
		},
	}

	IndexBodyImg := Client.Database("DiveEvolution").Collection("IndexImg")
	IndexBodyImg.InsertOne(context.TODO(),data)
}
func writeContactBody(w http.ResponseWriter, r *http.Request){
	data := models.Contact{
		Id: uuid.NewV4().String(),
		Title: "Lorem Ipsum",
		Introduction: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
		ContactInfo: models.ContactInfo{
			Address: "AV. Isabela y Guillermo Vega Gallegos",
			Phone: "(+593) 98 229 1894",
			Email: "jhonatan.quihuiri@gmail.com",
		},
		Button: "Lorem Ipsum",
		FormTitle: "Lorem Ipsum",
	}
	ContactBody := Client.Database("DiveEvolution").Collection("Contact")
	ContactBody.InsertOne(context.TODO(),data)
}
func writeContactImg(w http.ResponseWriter, r *http.Request){
	data := models.ContactImg{
		Id: uuid.NewV4().String(),
		Background: "https://via.placeholder.com/2000x900/6a5acd/fffffff",
		Foreground: "https://via.placeholder.com/350x500/ff0066/fffffff",
		Form: "https://via.placeholder.com/2000x900/00ff00/fffffff",
	}
	ContactImg := Client.Database("DiveEvolution").Collection("ContactImg")
	ContactImg.InsertOne(context.TODO(),data)
}