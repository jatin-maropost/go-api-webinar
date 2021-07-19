package main

import (
	"fmt" // Formatter print commands

	"github.com/jatin-maropost/go-api-webinar/config"
	"github.com/jatin-maropost/go-api-webinar/controllers"
	"github.com/jatin-maropost/go-api-webinar/helpers"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Printf("Server is running..")

	helpers.LoadEnvFile(".env")

	// Create router
	router := mux.NewRouter()
	
	// Configure firestore client
	booksHandler := controllers.BooksHandler{Firestore: config.ConfigureFirestore()}
 
	// Api endpoints
	router.HandleFunc("/api/books", booksHandler.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", booksHandler.GetBook).Methods("GET")
	router.HandleFunc("/api/books", booksHandler.CreateBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", booksHandler.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", booksHandler.DeleteBook).Methods("DELETE")

	port, ok := helpers.GetEnvVariable("PORT")
	if !ok {
		port = ":8000"
	}
	// Listen to the request at 8000 post
	log.Fatal(http.ListenAndServe(port, router))
}
