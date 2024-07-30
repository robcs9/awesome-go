package main

import (
	"database/sql"
	"log"
	"github.com/mattn/go-sqlite3"
)

type ToDo struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Status string `json:"status"`
}

var DB *sql.DB

func InitDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./awesome.db")
	if err != nil {
		log.Fatal(err)
	}
	_, err = DB.Exec(
		
	)
}