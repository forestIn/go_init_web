package main

import (
	"fmt"
	"go-contacts/app"
	"go-contacts/controllers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user", controllers.GetUser).Methods("GET")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET") //  user/2/contacts

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" //localhost
	}

	fmt.Println(port)

	// err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(headers, methods, origins)(router)))
	// if err != nil {
	// 	fmt.Print(err)
	// }
}
