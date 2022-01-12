package handlers

import (
	"SemiTrashApi/DataBase"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Message struct {
	Msg string `json:"msg"`
}

func setContentType(r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Get all books")
	setContentType(r)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(DataBase.AllBooks)
	if err != nil {
		json.NewEncoder(w).Encode(Message{Msg: fmt.Sprintf("Errro return all books %s", err)})
		log.Fatal(fmt.Sprintf("Errro return all books %s", err))
	}
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	setContentType(r)
	bookId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println(fmt.Sprintf("Error get book by id - %s\n", err))
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{Msg: fmt.Sprintf("Error get book by id - %s", err)})
	}

	log.Println(fmt.Printf("Get book by id = %d\n", bookId))
	book, hasBook := DataBase.GetBookById(bookId)
	if hasBook {
		log.Println(fmt.Sprintf("Return book by id - %d\n", book))
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(*book)
	} else {
		log.Println(fmt.Sprintf("Book by id - %d not found\n", bookId))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Message{Msg: fmt.Sprintf("Book by id - %d not found\n", bookId)})
	}

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	setContentType(r)
	log.Println("Try add book")
	decoder := json.NewDecoder(r.Body)
	var book DataBase.Book
	err := decoder.Decode(&book)
	if err != nil {
		log.Println(fmt.Sprintf("Error parse book - %s", err))
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{Msg: fmt.Sprintf("Error parse book - %s", err)})
		return
	}
	DataBase.AddBook(&book)
	w.WriteHeader(http.StatusCreated)
	log.Println(fmt.Sprintf("Added book - %s", book))
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	setContentType(r)
	log.Println("Try update book")
	var book DataBase.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Println(fmt.Sprintf("Error decode body %s"), err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{Msg: fmt.Sprintf("Error decode body %s", err)})
		return
	}
	updatedBook, hasBook := DataBase.GetBookById(book.Id)
	if hasBook {
		log.Println(fmt.Sprintf("Update book - %s", *updatedBook))
		DataBase.UpdateBook(&book)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)

	} else {
		log.Println(fmt.Sprintf("Book %s does not exist ", book))
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{Msg: fmt.Sprintf("Book %s does not exist ", book)})
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	setContentType(r)
	log.Println("Try delete book")
	bookId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println(fmt.Sprintf("Error delete book by id - %s\n", err))
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{Msg: fmt.Sprintf("Error delete book by id - %s", err)})
		return
	}

	log.Println(fmt.Printf("Delete book by id = %d\n", bookId))
	isDeleted := DataBase.DeleteBook(bookId)
	if isDeleted {
		log.Println(fmt.Sprintf("Book by id %d deleted\n", bookId))
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Message{Msg: fmt.Sprintf("Book by id %d deleted", bookId)})
	} else {
		log.Println(fmt.Sprintf("Book by id %d not exist\n", bookId))
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{Msg: fmt.Sprintf("Book by id %d not exist", bookId)})
	}

}
