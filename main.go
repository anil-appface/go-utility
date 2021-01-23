package main

import (
	"net/http"
)

func main() {

	//to initialise root
	InitRoot()

	//serve http, to start the server
	http.ListenAndServe(":8000", AppRouter)
}
