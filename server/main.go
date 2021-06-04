package main

import (
	"encoding/json"
	"fmt" // Formatter print commands
	"jatin/restapi/server/models"

	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/gorilla/mux"

	"context"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"

	"google.golang.org/api/option"
)

// Firestore client global variable
var firestore_client *firestore.Client


// getBooks function to handle the request at api/books endpoint GET
// Every endpoint handler function needs to have 2 params same like node js
// w is variable of http.ResponseWriter type and r is variable of *http.Request type
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	iter := firestore_client.Collection("books").Documents(context.Background())
	var books_list []models.Book
	
		// Creating a book variable 
		var book models.Book
		
	for {
        doc, err := iter.Next()
        if err == iterator.Done {
                break
        }

		// Binding data to book variable pointer
		doc.DataTo(&book)
		books_list = append(books_list,book)

		
		
	}

	// If no books are found then returns (404) not found response code 
	if len(books_list) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":"No books found."}`))	
	} else {
		// Return an array of all the books
		json.NewEncoder(w).Encode(books_list)
	}
	
}

// Get book a on the basis of id in the request params GET
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	
	params := mux.Vars(r) // Get request params
	book_id := params["id"]

	iter :=	firestore_client.Collection("books").Where("ID","==",book_id).Limit(1).Documents(context.Background())

	// Creating a book variable 
	var book models.Book
	// @todo clean it 
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
				break
		}
		// Binding data to book variable pointer
		doc.DataTo(&book)
		json.NewEncoder(w).Encode(book)	
		return
	}
	
	// Return empty book records of no book is found
	json.NewEncoder(w).Encode(book)
}


/* Create Book POST
@todo Add model validations
Add a check for the duplicate uuid in the firestore data
*/
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	

	// Generated a unique id of new book 
	book.ID = uuid.NewV4().String() 
	
	// Created new doc in books collection with auto generated doc id
	_, _, err:=	firestore_client.Collection("books").Add(context.Background(),book)

	if err !=nil {
		
			log.Printf("Body read error, %v", err)
			w.WriteHeader(500) // Return 500 Internal Server Error.
			return
		
	}
	
	// Send response with new book record
	w.WriteHeader(http.StatusCreated) // Return 200 success code
	json.NewEncoder(w).Encode(book)
}

// Update book (Put)
func updateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	
	// Creating a book variable to hold the book data fetched from the firestore
	var book models.Book

	// Id of book which is to be updated
	book_id := mux.Vars(r)["id"]
	
	// Decode the request payload in book_payload struct variable
	var book_payload models.Book 
	_ = json.NewDecoder(r.Body).Decode(&book_payload)

	iter :=	firestore_client.Collection("books").Where("ID","==",book_id).Limit(1).Documents(context.Background())

		book_ref, err := iter.Next()
		if err !=nil {
			log.Printf("Body read error, %v", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		book_ref.DataTo(&book)
		book_payload.ID = book.ID
		firestore_client.Collection("books").Doc(book_ref.Ref.ID).Set(context.Background(),book_payload)
	
		// Sending response with update book record
		json.NewEncoder(w).Encode(book_payload)	
}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	
	ctx := context.Background()
	// Id of book which is to be deleted
	book_id := mux.Vars(r)["id"]

	iter :=	firestore_client.Collection("books").Where("ID","==",book_id).Limit(1).Documents(context.Background())
		
	// Fetches the book form firestore
		book_ref, err := iter.Next()
		if err !=nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error":"No book found."}`))
			return ;
		}
		
		// Deletes a book 
		_, error := firestore_client.Collection("books").Doc(book_ref.Ref.ID).Delete(ctx)
		 
		if error != nil {
        	// Handle any errors in an appropriate way, such as returning them.
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(error.Error()))
			return ;
        
		}
}

// Configure and initialise the firebase admin sdk
func configureFirestore()   {

	// Load firestore configuration keys
	opt := option.WithCredentialsFile("./firestore-config.json")
	
	// Initialise app
	app, err := firebase.NewApp(context.Background(), nil, opt)

	// Log error if error occurs
	if err != nil {
	  log.Fatal(fmt.Errorf("error initializing app: %v", err))
	}

	// Get a Firestore client.
	client,err := app.Firestore(context.Background())
	if err != nil {
		log.Fatal(fmt.Errorf("error initializing app: %v", err))
	  }

	  // Intialise the firestore client in global variable
	  firestore_client = client
}

func main() {
	fmt.Printf("Server is running..")

	// Create router
	router := mux.NewRouter()
	
	// Configure firestore client
	configureFirestore()
 
	// Api endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	port := ":8000"
	// Listen to the request at 8000 post
	log.Fatal(http.ListenAndServe(port, router))

	
}
