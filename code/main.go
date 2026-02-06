package main

import (
	"database/sql"
	"net/http"

	"app/code/conn"
	"app/code/router"
)

var db *sql.DB

func main() {
	db = conn.ConnectDb()
	r := router.New(db)

	http.ListenAndServe("0.0.0.0:8080", r)
}
