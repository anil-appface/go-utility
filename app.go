package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/anil-appface/go-utility/resthandlers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var AppRouter = mux.NewRouter()

//InitRoot initialises routing to the go app
func InitRoot(db *sql.DB) {

	//Exercise -1
	exercise1 := resthandlers.NewExercise1Handler(AppRouter)
	AppRouter.Path("/webscrape").HandlerFunc(exercise1.GetTextAndCountForWeb).Methods("GET")

	//Exercise - 4
	exercise4 := resthandlers.NewExercise4Handler(AppRouter)
	AppRouter.Path("/getLastDayofMonth").HandlerFunc(exercise4.GetLastDayOfMonth).Methods("POST")

	//Exercise - 5
	exercise5 := resthandlers.NewExercise5Handler(AppRouter)
	AppRouter.Path("/getPrimeNumbers/{number}").HandlerFunc(exercise5.GetPrimeNumbers).Methods("GET")

	//Exercise - 3
	exercise3 := resthandlers.NewExercise3Handler(AppRouter, db)
	AppRouter.Path("/users").HandlerFunc(exercise3.GetAll).Methods("GET")
	AppRouter.Path("/user").HandlerFunc(exercise3.UpsertUser).Methods("POST")
	AppRouter.Path("/user/{id}").HandlerFunc(exercise3.DeleteUser).Methods("DELETE")

}

func openDB() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "tester"
	dbPass := "retset"
	dbName := "talentpro"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		fmt.Printf("error while sql open %#v", err)
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	return db, nil
}

//initialise tables
func initialiseTables(db *sql.DB) error {
	userTable := `
	CREATE TABLE IF NOT EXISTS user (
		id char(32) NOT NULL, 
		loginid varchar(200),
		fullname varchar(200), 
		isenabled TINYINT(1),
		created timestamp not null default current_timestamp, 
		updated timestamp not null default current_timestamp on update current_timestamp,
		PRIMARY KEY(id)
	)`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	_, err := db.ExecContext(ctx, userTable)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return err
	}

	return nil
}
