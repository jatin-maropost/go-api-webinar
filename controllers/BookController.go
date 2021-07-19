package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
	"github.com/jatin-maropost/go-api-webinar/models"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/api/iterator"
)

// Struct to hold the firestore ref
type BooksHandler struct{
	Firestore *firestore.Client
}

// getBooks function to handle the request at api/books endpoint GET
// Every endpoint handler function needs to have 2 params same like node js
// w is variable of http.ResponseWriter type and r is variable of *http.Request type
func  (bh BooksHandler ) GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Println(&bh.Firestore)
	iter := bh.Firestore.Collection("books").Documents(context.Background())
	var books_list []models.Book
	
	
		
	for {
        doc, err := iter.Next()
        if err == iterator.Done {
                break
        }
		// Creating a book variable 
		var book models.Book

		// Binding data to book pointer variable
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
func (bh BooksHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	
	params := mux.Vars(r) // Get request params
	book_id := params["id"]

	iter :=	bh.Firestore.Collection("books").Where("ID","==",book_id).Limit(1).Documents(context.Background())

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
func (bh BooksHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	

	// Generated a unique id of new book 
	book.ID = uuid.NewV4().String() 
	
	// Created new doc in books collection with auto generated doc id
	_, _, err:=	bh.Firestore.Collection("books").Add(context.Background(),book)

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
func (bh BooksHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	
	// Creating a book variable to hold the book data fetched from the firestore
	var book models.Book

	// Id of book which is to be updated
	book_id := mux.Vars(r)["id"]
	
	// Decode the request payload in book_payload struct variable
	var book_payload models.Book 
	_ = json.NewDecoder(r.Body).Decode(&book_payload)

	iter :=	bh.Firestore.Collection("books").Where("ID","==",book_id).Limit(1).Documents(context.Background())

		book_ref, err := iter.Next()
		if err !=nil {
			log.Printf("Body read error, %v", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		book_ref.DataTo(&book)
		book_payload.ID = book.ID
		bh.Firestore.Collection("books").Doc(book_ref.Ref.ID).Set(context.Background(),book_payload)
	
		// Sending response with update book record
		json.NewEncoder(w).Encode(book_payload)	
}


// Delete book
func (bh BooksHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	
	ctx := context.Background()
	// Id of book which is to be deleted
	book_id := mux.Vars(r)["id"]

	iter :=	bh.Firestore.Collection("books").Where("ID","==",book_id).Limit(1).Documents(context.Background())
		
	// Fetches the book form firestore
		book_ref, err := iter.Next()
		if err !=nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error":"No book found."}`))
			return ;
		}
		
		// Deletes a book 
		_, error := bh.Firestore.Collection("books").Doc(book_ref.Ref.ID).Delete(ctx)
		 
		if error != nil {
        	// Handle any errors in an appropriate way, such as returning them.
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(error.Error()))
			return ;
        
		}
}