package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)


type Pizza struct{
	Id uint64  `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Size uint `json:"size"`
}


type ErrorMSg struct{
	Msg string `json:"msg"`
}

var allPizzas []Pizza

func init() {
	Pizza1 := Pizza{
		Id: 1,
		Name: "Pizza1",
		Price: 9.90,
		Size: 30,
	}

	Pizza2 := Pizza{
		Id: 1,
		Name: "Pizza2",
		Price: 7.90,
		Size: 23,
	}

	allPizzas = append(allPizzas, Pizza1, Pizza2)
}


func findPizzaById(Id uint64) (Pizza, bool){
	for _ , pizza := range allPizzas{
		if pizza.Id == Id{
			return pizza, true
		}
	}
	return Pizza{}, false
}

func getAllPizzas(w http.ResponseWriter, r *http.Request){
	log.Println("Get all pizzas.")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(allPizzas)
	if err != nil{
		log.Println("Error encode")
	}
}



func getAllPizzaById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	pizzaId, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil{
		log.Printf("Error convert id %s to number.", vars["id"])
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMSg{Msg: fmt.Sprintf("Recive bad pizza id = %s", vars["id"])})
		return
	}

	pizza, finded := findPizzaById(pizzaId)
	if finded{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(pizza)

	} else{
		log.Printf("Pizza by id %s not found", vars["id"])
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorMSg{Msg: fmt.Sprintf("Pizza by id %s not found", vars["id"])})
	}
}

func main(){
	router:= mux.NewRouter()
	router.HandleFunc("/pizzas", getAllPizzas).Methods("GET")
	router.HandleFunc("/pizza/{id}", getAllPizzaById).Methods("GET")

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}







