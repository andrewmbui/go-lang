package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//Structs
//Structs are like classes, can have properties and methods, works similar to ES6/Java Class

//Book Struct (Model)
type book struct {
	ID    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
}

//Init books var as a slice Book struct
type books struct {
	bookCollection []book
}

//Functions

//Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	bookList := books{}
	err := allBooksQuery(&bookList)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	response, err := json.Marshal(&bookList)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println("Response", string(response))
	fmt.Fprintf(w, "Test print 1")
	fmt.Fprintf(w, string(response))
}

//Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	/*
		params := mux.Vars(r) // Get params
		Loop through books and find with id
		for _, item := range books {
			if item.ID == params["id"] {
				json.NewEncoder(w).Encode(item)
				return
			}
		}
		json.NewEncoder(w).Encode(&Book{})
	*/
}

//Create a New Book
func createBook(w http.ResponseWriter, r *http.Request) {
	/*
		w.Header().Set("Content-Type", "application/json")
		var book Book
		_ = json.NewDecoder(r.Body).Decode(&book)
		book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID. Randomized ID? cmon
		books = append(books, book)
		json.NewEncoder(w).Encode(book)
	*/
}

//Update Single Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	/*
		params := mux.Vars(r)
		for index, item := range books {
			if item.ID == params["id"] {
				books = append(books[:index], books[:index+1]...)
				var book Book
				_ = json.NewDecoder(r.Body).Decode(&book)
				book.ID = params["id"]
				books = append(books, book)
				json.NewEncoder(w).Encode(book)
				return
			}
		}
		json.NewEncoder(w).Encode(books)
	*/
}

//Delete Single Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	/*
		params := mux.Vars(r)
		for index, item := range books {
			if item.ID == params["id"] {
				books = append(books[:index], books[:index+1]...)
				break
			}
		}
		json.NewEncoder(w).Encode(books)
	*/
}

//Queries
func allBooksQuery(bookList *books) error {
	rows, err := db.Query(`
	SELECT 	uniqueid,
			isbn,
			title
	FROM	books`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		bk := book{}
		err = rows.Scan(
			&bk.ID,
			&bk.Isbn,
			&bk.Title,
		)
		if err != nil {
			return err
		}
		bookList.bookCollection = append(bookList.bookCollection, bk)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

var db *sql.DB

func main() {
	// Init Router
	r := mux.NewRouter()

	/* Mock Data - @todo -implement DB
	books = append(books, Book{ID: "1", Isbn: " 448743", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: " 847564", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})
	*/
	initDb()
	defer db.Close()
	// Route Handlers/ Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func initDb() {
	var err error
	connStr := "user=postgres dbname=socratica password=Magic123 host=localhost sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}
