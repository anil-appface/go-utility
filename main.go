package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"gitub.com/anil-appface/resthandlers"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/exercise1", resthandlers.GetTextAndCountForWeb).Methods("GET")

	http.ListenAndServe(":8000", router)
}
