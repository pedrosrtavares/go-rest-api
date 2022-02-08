package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pedrosrtavares/go-rest-api/controllers"
	"github.com/pedrosrtavares/go-rest-api/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/api/personalidades", controllers.GetAllPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidade/{id}", controllers.GetOnePersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.PostNewPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidade/{id}", controllers.DeletePersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidade/{id}", controllers.EditPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
