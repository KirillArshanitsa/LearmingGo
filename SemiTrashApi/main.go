package main

import (
	"SemiTrashApi/routing"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const (
	ApiVer      = "/ApiV1"
	AllBooksUrl = ApiVer + "/allbooks"
	BookIdUrl   = ApiVer + "/book/{id}"
	BookUrl     = ApiVer + "/book"
)

var (
	port string
)

func init() {
	err := godotenv.Load("config.ini")
	if err != nil {
		log.Fatal(fmt.Sprintf("Error load config file %s\n", err))
	}
	port = os.Getenv("port")
}

func main() {
	router := mux.NewRouter()
	log.Println("Configure router.")
	routing.CreateRoutingForAllBooks(router, AllBooksUrl)
	routing.CreateRoutingForBook(router, BookUrl)
	routing.CreateRoutingForBookId(router, BookIdUrl)
	log.Println("Start on port " + port)
	log.Fatal(http.ListenAndServe("localhost:"+port, router))

}
