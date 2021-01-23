package resthandlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anil-appface/go-utility/util"
	"github.com/gorilla/mux"
)

type exercise4Handler struct {
	router *mux.Router
}

type dateResponse struct {
	Date           string `json:"Date"`
	LastDayofMonth int    `json:"LastDayofMonth"`
}

//NewExercise4Handler instatiates the exercise4 handler
func NewExercise4Handler(_router *mux.Router) *exercise4Handler {
	return &exercise4Handler{
		router: _router,
	}
}

func (me *exercise4Handler) GetLastDayOfMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type request struct {
		Date string `json:"Date"`
	}

	var req *request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("error while encoding json: %#v", err)
		return
	}

	t, err := util.ConvertStringToDate(req.Date)
	if err != nil {
		fmt.Printf("error while converting date string to time: %#v", err)
		return
	}

	result := &dateResponse{
		Date:           req.Date,
		LastDayofMonth: util.GetLastDayOfMonth(t),
	}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		fmt.Printf("error while decoding json: %#v", err)
		return
	}
}
