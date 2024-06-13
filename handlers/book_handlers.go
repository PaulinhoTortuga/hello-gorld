package handlers

import (
	"encoding/json"
	"fmt"
	"hello-world/crud/models"
	"hello-world/crud/store"
	"hello-world/crud/utils"
	"net/http"

	"github.com/gorilla/mux"
)


func CreateBook (w http.ResponseWriter, r *http.Request) {
    var book models.Book
    err := json.NewDecoder(r.Body).Decode(&book)
    if(err != nil){
        http.Error(w, err.Error(), http.StatusBadRequest)
        return 
    }
    book.ID, _ = utils.GenerateId()

    store.BookStore[book.ID] = book
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(book)
}

func UpdateBook (w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    book, exist := store.BookStore[id]
    if !exist {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }
    
    var updates map[string]interface{}

    err := json.NewDecoder(r.Body).Decode(&updates)
    if(err != nil){
        http.Error(w, err.Error(), http.StatusBadRequest)
        return 
    }

    err = utils.UpdateStructFields(&book, updates)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    store.BookStore[id] = book

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(book)
}

func GetBooks (w http.ResponseWriter, r *http.Request) {
    var res []models.Book
    for _, v := range store.BookStore {
        res = append(res, v)
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(res)
}

func GetBook (w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    book, exist := store.BookStore[id]
    if !exist {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    book, exists := store.BookStore[id]
    if !exists {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }

    fmt.Printf("Before deletion: %+v\n", book)

    // Delete the book from the store
    delete(store.BookStore, id)

    // Verify deletion
    _, stillExists := store.BookStore[id]
    if stillExists {
        fmt.Printf("Failed to delete book with ID %s\n", id)
    } else {
        fmt.Println("Book deleted successfully")
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(book)
}
