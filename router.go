package main

import (
	"github.com/anil-appface/go-utility/resthandlers"
	"github.com/gorilla/mux"
)

var AppRouter = mux.NewRouter()

//InitRoot initialises routing to the go app
func InitRoot() {

	exercise1 := resthandlers.NewExercise1Handler(AppRouter)
	//Exercise -1
	AppRouter.Path("/webscrape").HandlerFunc(exercise1.GetTextAndCountForWeb).Methods("GET")

	exercise4 := resthandlers.NewExercise4Handler(AppRouter)
	//Exercise - 4
	AppRouter.Path("/getLastDayofMonth").HandlerFunc(exercise4.GetLastDayOfMonth).Methods("POST")

}
