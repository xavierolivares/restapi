package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

// Book Struct (Model)...kind of like a class
type Book struct {
	//this will be the json field we fetch
	ID		string	`json:"id"`
	Isbn	string	`json:"isbn"`
	Title	string	`json:"title"`
	Author	*Author	`json:"author"`
}

//Author Struct
type Author struct {
	Firstname	string	`json: "firstname"`
	Lastname	string	`json: "lastname"`

}

//Get all books, sort of like app.get()
func getBooks(w http.ResponseWriter, router *http.Request) {

}

//Get single book
func getBook(w http.ResponseWriter, r *http.Request) {

}

//Create a new book
func createBook(w http.ResponseWriter, r *http.Request) {

}

//Update the book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

//Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}
func main() {
	// Init Router
	r := mux.NewRouter()

	// := is typing inference
	// var age int = 35
	// age := 35

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	//RUNNING SERVER, similar to app.listen()
	log.Fatal(http.ListenAndServe(":8000", r))

}
