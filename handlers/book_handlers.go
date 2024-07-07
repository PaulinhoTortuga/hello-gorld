package handlers

import (
	"encoding/json"
	"hello-gorld/crud/models"
	"hello-gorld/crud/utils"
	"io"
	"net/http"
	"os"
	"slices"

	"github.com/gorilla/mux"
)

func CreateBook (w http.ResponseWriter, r *http.Request) {
    var store models.Books
    var book models.Book

    err := json.NewDecoder(r.Body).Decode(&book)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    jsonFile, err := os.Open("./store/books.json")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer jsonFile.Close()

    byteValue, err := io.ReadAll(jsonFile)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }    

    err = json.Unmarshal(byteValue, &store)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    book.ID, _ = utils.GenerateId()
    store.Books = append(store.Books, book)

    content, err := json.Marshal(store)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = os.WriteFile("./store/books.json", content, os.ModePerm)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(book)
}

func UpdateBook (w http.ResponseWriter, r *http.Request) {
    var store models.Books
    var updates map[string]interface{}
    vars := mux.Vars(r)
    id := vars["id"]

    jsonFile, err := os.Open("./store/books.json")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer jsonFile.Close()

    byteValue, err := io.ReadAll(jsonFile)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }    

    err = json.Unmarshal(byteValue, &store)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    idx := slices.IndexFunc(store.Books, func(b models.Book) bool { return b.ID == id })
    if idx == -1 {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }

    book := store.Books[idx]

    err = json.NewDecoder(r.Body).Decode(&updates)
    if(err != nil){
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = utils.UpdateStructFields(&book, updates)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    store.Books[idx] = book

    content, err := json.Marshal(store)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = os.WriteFile("./store/books.json", content, os.ModePerm)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(book)
}

func GetBooks (w http.ResponseWriter, r *http.Request) {
    var res models.Books

    jsonFile, err := os.Open("./store/books.json")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer jsonFile.Close()

    byteValue, err := io.ReadAll(jsonFile)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
            
    err = json.Unmarshal(byteValue, &res)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(res.Books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
    var store models.Books
    vars := mux.Vars(r)
    id := vars["id"]

    jsonFile, err := os.Open("./store/books.json")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer jsonFile.Close()

    byteValue, err := io.ReadAll(jsonFile)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }    

    err = json.Unmarshal(byteValue, &store)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    idx := slices.IndexFunc(store.Books, func(b models.Book) bool { return b.ID == id })
    if idx == -1 {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }

    book := store.Books[idx]

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
    var store models.Books
    vars := mux.Vars(r)
    id := vars["id"]

    jsonFile, err := os.Open("./store/books.json")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer jsonFile.Close()

    byteValue, err := io.ReadAll(jsonFile)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }    

    err = json.Unmarshal(byteValue, &store)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    idx := slices.IndexFunc(store.Books, func(b models.Book) bool { return b.ID == id })
    if idx == -1 {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }

    store.Books = append(store.Books[:idx], store.Books[idx+1:]...)

    content, err := json.Marshal(store)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = os.WriteFile("./store/books.json", content, os.ModePerm)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusOK)
}
