package routing

import (
	"SemiTrashApi/handlers"
	"github.com/gorilla/mux"
)

func CreateRoutingForAllBooks(router *mux.Router, url string) {
	router.HandleFunc(url, handlers.GetAllBooks).Methods("GET")
}

func CreateRoutingForBookId(router *mux.Router, url string) {
	router.HandleFunc(url, handlers.GetBookById).Methods("GET")
	router.HandleFunc(url, handlers.DeleteBook).Methods("DELETE")
}

func CreateRoutingForBook(router *mux.Router, url string) {
	router.HandleFunc(url, handlers.CreateBook).Methods("POST")
	router.HandleFunc(url, handlers.UpdateBook).Methods("PUT")
	router.HandleFunc(url, handlers.DeleteBook).Methods("DELETE")
}
