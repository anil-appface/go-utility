package main

import (
	"github.com/anil-appface/go-utility/resthandlers"
	"github.com/gorilla/mux"
)

var AppRouter = mux.NewRouter()

//InitRoot initialises routing to the go app
func InitRoot() {

	//Exercise -1
	exercise1 := resthandlers.NewExercise1Handler(AppRouter)
	AppRouter.Path("/webscrape").HandlerFunc(exercise1.GetTextAndCountForWeb).Methods("GET")

	//Exercise - 4
	exercise4 := resthandlers.NewExercise4Handler(AppRouter)
	AppRouter.Path("/getLastDayofMonth").HandlerFunc(exercise4.GetLastDayOfMonth).Methods("POST")

	//Exercise - 5
	exercise5 := resthandlers.NewExercise5Handler(AppRouter)
	AppRouter.Path("/getPrimeNumbers/{number}").HandlerFunc(exercise5.GetPrimeNumbers).Methods("GET")

}
