package resthandlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anil-appface/go-utility/store"
	"github.com/gorilla/mux"
)

type exercise3Handler struct {
	router *mux.Router
	db     *sql.DB
}

//NewExercise3Handler instatiates the exercise3 handler
func NewExercise3Handler(_router *mux.Router, _db *sql.DB) *exercise3Handler {
	return &exercise3Handler{
		router: _router,
		db:     _db,
	}
}

func (me *exercise3Handler) UpsertUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req *store.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("error while encoding json: %#v", err)
		return
	}

	urs := store.NewUserRepoSQL(me.db)
	user, err := urs.UpsertUser(req)
	if err != nil {
		fmt.Printf("error while insert/update user: %#v", err)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		fmt.Printf("error while decoding json: %#v", err)
		return
	}
}

func (me *exercise3Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	urs := store.NewUserRepoSQL(me.db)
	err := urs.Delete(mux.Vars(r)["id"])
	if err != nil {
		fmt.Printf("error while getting all users: %#v", err)
		return
	}
}

func (me *exercise3Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	urs := store.NewUserRepoSQL(me.db)
	allUsers, err := urs.GetAll()
	if err != nil {
		fmt.Printf("error while getting all users: %#v", err)
		return
	}

	err = json.NewEncoder(w).Encode(allUsers)
	if err != nil {
		fmt.Printf("error while decoding json: %#v", err)
		return
	}
}
